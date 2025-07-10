function encodeResourceSrc(staticFileAddress, src) {
  return (
    staticFileAddress +
    encodeURI(src)
      .replaceAll('(', '%28')
      .replaceAll(')', '%29')
      .replaceAll('#', '%23')
      .replaceAll('?', '%3F')
  )
}

export { encodeResourceSrc }
