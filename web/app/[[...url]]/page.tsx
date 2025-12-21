"use client";

import { usePathname } from "next/navigation";
import { useEffect, useState } from "react";
import FolderDisplay from "@/app/components/FolderDisplay";
import Mp4Display from "@/app/components/Mp4Display";
import ImgDisplay from "@/app/components/ImgDisplay";
import Navbar from "@/app/components/Navbar";
import { fetchFolderInfo } from "../http/controller";

interface PageProps {
	params: Promise<{ url?: string[] }>;
}

export default function MediaGallery({ params }: PageProps) {
	const pathname = usePathname();

	const [backendAddress, setBackendAddress] = useState("");
	const [navBarTitle, setNavBarTitle] = useState("");
	const [isFetchSuccess, setIsFetchSuccess] = useState<boolean | null>(null);
	const [folders, setFolders] = useState<Folder[]>([]);
	const [vids, setVids] = useState<Video[]>([]);
	const [imgs, setImgs] = useState<Image[]>([]);

	const fetchContentInfo = async (path: string) => {
		setIsFetchSuccess(null);
		try {
			const res: FolderContentResponse = await fetchFolderInfo(path);

			if (!res.success) {
				throw res.error;
			}

			setNavBarTitle(res.directory?.url ?? "");
			setFolders(res.directories || []);
			setVids(res.videos || []);
			setImgs(res.images || []);
			setIsFetchSuccess(true);
		} catch (err) {
			console.error(err);
			setNavBarTitle("Error!");
			setFolders([]);
			setVids([]);
			setImgs([]);
			setIsFetchSuccess(false);
		}
	};

	useEffect(() => {
		setBackendAddress(
			`http://${window.location.hostname}:${process.env.NEXT_PUBLIC_BACKEND_PORT}`
		);
	}, []);

	useEffect(() => {
		fetchContentInfo(pathname);
	}, [pathname]);

	const isEmpty =
		folders.length === 0 && imgs.length === 0 && vids.length === 0;

	return (
		<>
			<Navbar navBarTitle={navBarTitle}></Navbar>
			<div className="container-fluid">
				<div id="fetchHintText">
					{isFetchSuccess === null && <div>Loading...</div>}
					{isFetchSuccess === false && (
						<>
							<div>Fail to fetch folder info!</div>
						</>
					)}
					{isFetchSuccess === true && isEmpty && (
						<div>No content in folder!</div>
					)}
				</div>

				{isFetchSuccess === true && (
					<>
						{
							<FolderDisplay
								folders={folders}
								hostAddress={backendAddress}
							/>
						}
						{
							<Mp4Display
								mp4s={vids}
								hostAddress={backendAddress}
							/>
						}
						{
							<ImgDisplay
								imgs={imgs}
								hostAddress={backendAddress}
							/>
						}
					</>
				)}
			</div>
		</>
	);
}
