interface FolderContentResponse {
	directories?: Folder[];
	directory?: {
		name: string;
		url: string;
		thumbnail: null;
	};
	images?: Image[];
	videos?: Video[];
	success: boolean;
	error?: any;
}

interface Image {
	url: string;
	width: number;
	height: number;
}

interface Video {
	name: string;
	url: string;
	thumbnail: null | {
		image: Image;
	};
}

interface Folder {
	name: string;
	url: string;
	thumbnail: null | {
		image: Image;
	};
}
