import Link from "next/link";
import { encodeResourceSrc } from "@/app/scripts/utils";

interface FolderDisplayProps {
	folders: Folder[];
	hostAddress: string;
}

export default function FolderDisplay({
	folders,
	hostAddress,
}: FolderDisplayProps) {
	return (
		<div className="container-fluid px-0">
			<div className="row mb-4 gx-1 gy-4">
				{folders.map((folder) => (
					<div
						key={folder.name}
						className="col-6 col-sm-4 col-md-3 col-xl-2 text-center"
					>
						<Link
							href={folder.url}
							type="button"
							className="btn-width btn btn-primary p-0"
						>
							<div className="py-1 truncate mx-auto">
								{folder.name}
							</div>
							{folder.thumbnail && (
								<img
									width={folder.thumbnail.image.width}
									height={folder.thumbnail.image.height}
									loading="lazy"
									className="rounded-bottom img-fluid"
									src={encodeResourceSrc(
										hostAddress,
										folder.thumbnail.image.url
									)}
								/>
							)}
						</Link>
					</div>
				))}
			</div>
		</div>
	);
}
