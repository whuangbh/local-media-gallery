export function encodePath(path: string) {
  return path
    .split('/')
    .map((segment) => encodeURIComponent(segment))
    .join('/')
}


// 1. 获取原图/原视频的绝对地址（加上 /static）
export function getStaticUrl(backendAddress: string, relativePath: string) {
  return `${backendAddress}/static/${encodePath(relativePath)}`
}

// 2. 获取 imagor 缩略图地址
export function getThumbnailUrl(backendAddress: string, relativePath: string, width: number = 1000) {
  const imagorBase = `${backendAddress}:8000`

  const encodedPath = relativePath
    .split('/')
    .map(segment => encodeURIComponent(segment))
    .join('/')

  const filters = 'filters:format(webp):quality(95)'

  return `${imagorBase}/unsafe/fit-in/${width}x0/${filters}${encodedPath}`
}

