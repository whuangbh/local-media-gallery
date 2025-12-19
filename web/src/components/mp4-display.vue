<script setup>
import { encodeResourceSrc } from '../utils/utilities'
// eslint-disable-next-line no-unused-vars
const props = defineProps({
  mp4s: Array,
  staticFileAddress: String,
})

const redirectMp4Src = (url) => {
  window.location.href = url
}
</script>

<template>
  <div class="container-fluid" :class="{ 'mb-5': mp4s.length != 0 }">
    <div class="row gx-1 gy-3">
      <div v-for="mp4 in mp4s" :key="mp4.name" class="col-xl-2 col-6 text-center">
        <button
          @click="redirectMp4Src(encodeResourceSrc(staticFileAddress, mp4.url))"
          type="button"
          class="btn-width btn btn-secondary p-0"
        >
          <div class="py-1 truncate mx-auto">
            {{ mp4.name }}
          </div>
          <img
            v-if="mp4.thumbnail"
            :width="mp4.thumbnail.width"
            :height="mp4.thumbnail.height"
            alt=""
            loading="lazy"
            class="mt-1 pb-1 img-fluid"
            :src="encodeResourceSrc(staticFileAddress, mp4.thumbnail.image.url)"
          />
        </button>
      </div>
    </div>
  </div>
</template>

<style scoped></style>
