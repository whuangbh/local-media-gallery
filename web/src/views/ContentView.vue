<script setup lang="ts">
import { computed, ref, watchEffect } from 'vue'
import { useRoute } from 'vue-router'
import NavBar from '@/components/NavBar.vue'
import FolderDisplay from '@/components/FolderDisplay.vue'
import ImgDisplay from '@/components/ImgDisplay.vue'
import Mp4Display from '@/components/Mp4Display.vue'
import Controller from '@/http/controller.js'
import type { Folder, Video, Image, FolderContentResponse } from '@/types/types'

const route = useRoute()
const backendAddress = window.location.origin

const navBarTitle = ref<string>('')
const isFetchSuccess = ref<boolean | null>(null)
const folders = ref<Folder[]>([])
const vids = ref<Video[]>([])
const imgs = ref<Image[]>([])
// const favorites = ref([])
// const isFetchFavSuccess = ref(null)

const isFolderEmpty = computed(
  () => folders.value.length === 0 && vids.value.length === 0 && imgs.value.length === 0,
)

async function fetchFolderContent(currentPath: string) {
  isFetchSuccess.value = null
  // isFetchFavSuccess.value = null
  folders.value = []
  vids.value = []
  imgs.value = []

  if (route.name !== 'Favorite Viewer') {
    try {
      const res: FolderContentResponse = await Controller.fetchFolderInfo(currentPath)

      if (!res.success) {
        throw res.error
      }

      navBarTitle.value = res.directory.url
      folders.value = res.directories
      vids.value = res.videos
      imgs.value = res.images
      isFetchSuccess.value = true
    } catch (err) {
      console.error(err)

      navBarTitle.value = 'Error!'
      folders.value = []
      vids.value = []
      imgs.value = []
      isFetchSuccess.value = false
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

watchEffect(() => {
  fetchFolderContent(route.path)
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
</script>

<template>
  <!-- @onClickSortFavBtn="sortFavorite" -->
  <NavBar :navBarTitle="navBarTitle"></NavBar>
  <div class="container-fluid" id="fetchHintText">
    <!-- <template v-if="route.name !== 'Favorite Viewer'"> -->
    <div v-if="isFetchSuccess === null">Loading...</div>
    <div v-else-if="isFetchSuccess === false">Fail to fetch folder info!</div>
    <div v-else-if="isFetchSuccess && isFolderEmpty">No content in folder!</div>
    <!-- </template>
    <template v-else>
      <div v-if="isFetchFavSuccess === null">Loading...</div>
      <div v-else-if="isFetchFavSuccess === false">Fail to fetch favorites!</div>
      <div v-else-if="folders.length === 0 && imgs.length === 0 && vids.length === 0">
        No content in favorites!
      </div>
    </template> -->
  </div>

  <template v-if="isFetchSuccess">
    <FolderDisplay :folders="folders" :backendAddress="backendAddress" />
    <Mp4Display :mp4s="vids" :backendAddress="backendAddress" />
    <!-- :favorites="favorites"
    @addFavorite="addFavorite"
    @removeFavorite="removeFavorite" -->
    <ImgDisplay :imgs="imgs" :backendAddress="backendAddress" />
  </template>
</template>

<style scoped></style>
