import { encodeResourceSrc } from "@/app/scripts/utils";

interface Mp4DisplayProps {
	mp4s: Video[];
	hostAddress: string;
}

function handleOnClickMp4(hostAddress: string, mp4: Video) {
	const targetUrl = encodeResourceSrc(hostAddress, mp4.url);
	window.location.href = targetUrl;
}

export default function Mp4Display({ mp4s, hostAddress }: Mp4DisplayProps) {
	return (
		<div className="container-fluid px-0">
			<div className="row mb-4 gx-1 gy-4">
				{mp4s.map((mp4) => (
					<div
						key={mp4.name}
						className="col-6 col-sm-4 col-md-3 col-xl-2 text-center"
					>
						<button
							onClick={() => handleOnClickMp4(hostAddress, mp4)}
							type="button"
							className="btn-width btn btn-secondary p-0"
						>
							<div className="py-1 truncate mx-auto">
								{mp4.name}
							</div>
							{mp4.thumbnail && (
								<img
									width={mp4.thumbnail.image.width}
									height={mp4.thumbnail.image.height}
									loading="lazy"
									className="rounded-bottom img-fluid"
									src={encodeResourceSrc(
										hostAddress,
										mp4.thumbnail.image.url
									)}
								/>
							)}
						</button>
					</div>
				))}
			</div>
		</div>
	);
}
