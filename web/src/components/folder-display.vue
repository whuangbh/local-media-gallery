<script setup>
import { useRouter } from 'vue-router'
import { encodeResourceSrc } from '../utils/utilities'

const router = useRouter()
// eslint-disable-next-line no-unused-vars
const props = defineProps({
  folders: Array,
  staticFileAddress: String,
})

// if the folder name is too long (e.g. manga name), only display author name
const getFolderNameToDisplay = (string) => {
  return string.length >= 20 ? string.split('-')[0] : string
}
</script>

<!--  -->
<template>
  <div class="container-fluid">
    <div class="row mt-0 mb-4 gx-1 gy-4">
      <div v-for="folder in folders" :key="folder.name" class="col-6 col-xl-2 text-center">
        <button
          @click="router.push(folder.url)"
          type="button"
          class="folderBtn btn btn-primary text-center p-0"
        >
          <div id="folderTitle" class="py-1">
            {{ getFolderNameToDisplay(folder.name) }}
          </div>
          <img
            v-if="folder.thumbnail != null"
            :width="folder.thumbnail.width"
            :height="folder.thumbnail.height"
            alt=""
            loading="lazy"
            class="mt-1 pb-1 img-fluid"
            :src="encodeResourceSrc(staticFileAddress, folder.thumbnail.image.url)"
          />
        </button>
      </div>
    </div>
  </div>
</template>

<style scoped>
button.folderBtn {
  width: 100%;
}

div#folderTitle {
  word-wrap: break-word;
}
</style>
