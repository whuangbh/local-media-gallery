<script setup>
import { useRouter } from 'vue-router'
import { encodeResourceSrc } from '../utils/utilities'

const router = useRouter()
// eslint-disable-next-line no-unused-vars
const props = defineProps({
  folders: Array,
  staticFileAddress: String,
})
</script>

<!--  -->
<template>
  <div class="container-fluid">
    <div class="row mt-0 mb-4 gx-1 gy-4">
      <div v-for="folder in folders" :key="folder.name" class="col-6 col-xl-2 text-center">
        <button
          @click="router.push(folder.url)"
          type="button"
          class="btn-width btn btn-primary p-0"
        >
          <div class="py-1 truncate mx-auto">
            {{ folder.name }}
          </div>
          <img
            v-if="folder.thumbnail"
            :width="folder.thumbnail.width"
            :height="folder.thumbnail.height"
            alt=""
            loading="lazy"
            class="rounded-bottom img-fluid"
            :src="encodeResourceSrc(staticFileAddress, folder.thumbnail.image.url)"
          />
        </button>
      </div>
    </div>
  </div>
</template>

<style scoped></style>
