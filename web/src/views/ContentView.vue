<script setup lang="ts">
import { ref, watchEffect } from 'vue'
import { useRoute } from 'vue-router'
import NavBar from '@/components/nav-bar.vue'
import FolderDisplay from '@/components/folder-display.vue'
import ImgDisplay from '@/components/img-display.vue'
import Mp4Display from '@/components/mp4-display.vue'
import controller from '@/http/controller.js'

const route = useRoute()
const staticFileAddress = 'http://' + window.location.hostname

const navBarTitle = ref('')
const isFetchSuccess = ref(null)
const folders = ref([])
const vids = ref([])
const imgs = ref([])
const favorites = ref([])
const isFetchFavSuccess = ref(null)

let currentUrl = route.path

watchEffect(() => {
  currentUrl = route.path
  isFetchSuccess.value = null
  isFetchFavSuccess.value = null
  folders.value = []
  vids.value = []
  imgs.value = []
  fetchContentInfo()
})

// TODO
// async function fetchFavInfo() {
//   favorites.value = await controller.fetchFavorites();
// }
//
// async function addFavorite(imgObj, callBackFn) {
//   try {
//     favorites.value = await controller.addFavorite(imgObj);
//     callBackFn();
//   } catch (err) {
//     console.error(err);
//   }
// }
//
// async function removeFavorite(imgObj, callBackFn) {
//   try {
//     favorites.value = await controller.removeFavorite(imgObj);
//     callBackFn();
//   } catch (err) {
//     console.error(err);
//   }
// }
//
// async function sortFavorite() {
//   try {
//     isFetchFavSuccess.value = null;
//     imgs.value = [];
//     favorites.value = await controller.sortFavorite();
//
//     isFetchFavSuccess.value = true;
//     imgs.value = favorites.value;
//   } catch (err) {
//     console.error(err);
//     isFetchFavSuccess.value = false;
//   }
// }

// onMounted(() => {
//   try {
//     fetchFavInfo()
//   } catch (err) {
//     console.error(err)
//   }
// })

async function fetchContentInfo() {
  if (route.name !== 'Favorite Viewer') {
    try {
      const res = await controller.fetchFolderInfo(currentUrl)

      navBarTitle.value = res.directory.url

      folders.value = res.directories
      vids.value = res.videos
      imgs.value = res.images

      isFetchSuccess.value = true
    } catch (err) {
      folders.value = []
      vids.value = []
      imgs.value = []
      isFetchSuccess.value = false
      console.error(err)
    }
  }
  // TODO
  // else {
  //   try {
  //     navBarTitle.value = 'Favorites'
  //
  //     await fetchFavInfo()
  //
  //     imgs.value = favorites.value
  //
  //     isFetchFavSuccess.value = true
  //   } catch (err) {
  //     folders.value = []
  //     mp4s.value = []
  //     imgs.value = []
  //     isFetchFavSuccess.value = false
  //     console.error(err)
  //   }
  // }
}
</script>

<template>
  <NavBar @onClickSortFavBtn="sortFavorite" id="navBar" :navBarTitle="navBarTitle"></NavBar>
  <div class="container-fluid" id="fetchHintText">
    <template v-if="route.name !== 'Favorite Viewer'">
      <div v-if="isFetchSuccess === null">Loading...</div>
      <div v-else-if="isFetchSuccess === false">Fail to fetch folder info!</div>
      <div v-else-if="folders.length === 0 && imgs.length === 0 && vids.length === 0">
        No content in folder!
      </div>
    </template>
    <template v-else>
      <div v-if="isFetchFavSuccess === null">Loading...</div>
      <div v-else-if="isFetchFavSuccess === false">Fail to fetch favorites!</div>
      <div v-else-if="folders.length === 0 && imgs.length === 0 && vids.length === 0">
        No content in favorites!
      </div>
    </template>
  </div>

  <FolderDisplay id="folderDisplay" :folders="folders" :staticFileAddress="staticFileAddress" />
  <Mp4Display id="mp4Display" :mp4s="vids" :staticFileAddress="staticFileAddress" />
  <ImgDisplay
    id="imgDisplay"
    :imgs="imgs"
    :staticFileAddress="staticFileAddress"
    :favorites="favorites"
    @addFavorite="addFavorite"
    @removeFavorite="removeFavorite"
  />
</template>

<style scoped></style>
