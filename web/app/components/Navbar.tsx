"use client";

import { useRouter } from "next/navigation";

interface NavBarProps {
	navBarTitle: string;
}

export default function Navbar({ navBarTitle }: NavBarProps) {
	const router = useRouter();

	function goBackToUpperPath() {
		const pathSegments = navBarTitle.split("/").filter(Boolean);
		pathSegments.pop();
		const newPath = `/${pathSegments.join("/")}`;

		router.push(newPath);
	}

	return (
		<div className="container-fluid my-3">
			<div className="h2 mt-1 d-flex">
				{navBarTitle !== "/" && (
					<button
						type="button"
						className="btn btn-secondary d-inline me-2"
						onClick={goBackToUpperPath}
					>
						<i className="bi bi-chevron-left me-1"></i>
					</button>
				)}
				<span className="align-self-center">
					{navBarTitle || "Loading..."}
				</span>
			</div>
		</div>
	);
}
