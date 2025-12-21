import PhotoSwipeLightbox from "photoswipe/lightbox";
import "photoswipe/style.css";
import { useEffect } from "react";
import { encodeResourceSrc } from "@/app/scripts/utils";

interface ImgDisplayProps {
	imgs: Image[];
	hostAddress: string;
}

const galleryElementId = "photo-swipe-gallery";

function photoSwipeInit() {
	const lightbox = new PhotoSwipeLightbox({
		gallery: "#" + galleryElementId,
		children: "a",
		tapAction: "close",
		pswpModule: () => import("photoswipe"),
	});

	lightbox.init();
}

export default function ImgDisplay({ imgs, hostAddress }: ImgDisplayProps) {
	useEffect(() => {
		photoSwipeInit();
	}, []);

	return (
		<div className="container-fluid px-0">
			<div
				className="row pswp-gallery pswp-gallery--single-column mb-4 gx-1 gy-4"
				id={galleryElementId}
			>
				{imgs.map((img) => (
					<div
						key={img.url}
						className="col-6 col-sm-4 col-md-3 col-xl-2 text-center"
					>
						<a
							data-pswp-width={img.width}
							data-pswp-height={img.height}
							target="_blank"
							href={encodeResourceSrc(hostAddress, img.url)}
						>
							<img
								width={img.width}
								height={img.height}
								loading="lazy"
								className="img-fluid border border-dark-subtle"
								src={encodeResourceSrc(hostAddress, img.url)}
							/>
						</a>
					</div>
				))}
			</div>
		</div>
	);
}
