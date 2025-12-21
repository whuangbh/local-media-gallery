function encodeResourceSrc(hostAddress: string, src: string) {
	return (
		hostAddress +
		encodeURI(src)
			.replaceAll("(", "%28")
			.replaceAll(")", "%29")
			.replaceAll("#", "%23")
			.replaceAll("?", "%3F")
	);
}

export { encodeResourceSrc };
