export type FolderContentResponse = FolderContentSuccess | FolderContentError

interface FolderContentSuccess {
  success: true
  directories: Folder[]
  directory: Folder
  images: Image[]
  videos: Video[]
  error?: never
}

interface FolderContentError {
  success: false
  error: string
  directories?: never
  images?: never
  videos?: never
}

export interface Image {
  url: string
  width: number
  height: number
}

export interface Video {
  name: string
  url: string
  thumbnail: null | {
    image: Image
  }
}

export interface Folder {
  name: string
  url: string
  thumbnail: null | {
    image: Image
  }
}
