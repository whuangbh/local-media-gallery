import axios from "axios";

const backend = axios.create({
	// If browser, use current hostname. If server, use empty string or local IP.
	baseURL: "",
	timeout: 5000,
	headers: {
		"Content-type": "application/json",
	},
});

export async function fetchFolderInfo(url: string) {
	const response = await backend.post("/api/media-folder-info", {
		url: decodeURI(url),
	});
	return response.data;
}
