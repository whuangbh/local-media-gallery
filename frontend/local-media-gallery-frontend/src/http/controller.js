import axios from 'axios'

const backend = axios.create({
  baseURL: `http://${window.location.hostname}:${import.meta.env.VITE_proxyServerPort}`,
  timeout: 5000,
  headers: {
    'Content-type': 'application/json',
  },
})

const controller = {
  fetchFolderInfo: async function (url) {
    const response = await backend.post('/api/folderInfo', {
      folderName: decodeURI(url),
    })
    return response.data
  },

  // TODO
  // fetchFavorites: async function () {
  //   const response = await backend.get("/api/favorites");
  //   return response.data;
  // },
  //
  // addFavorite: async function (imgObj) {
  //   const response = await backend.post("/api/favorites", imgObj);
  //   return response.data;
  // },
  //
  // removeFavorite: async function (imgObj) {
  //   const response = await backend.delete("/api/favorites", {
  //     data: imgObj
  //   });
  //   return response.data;
  // },
  //
  // sortFavorite: async function () {
  //   const response = await backend.get("/api/favorites/sort");
  //   return response.data;
  // }
}

export default controller
