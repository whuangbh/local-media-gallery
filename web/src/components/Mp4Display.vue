<script setup lang="ts">
import type { Video } from '@/types/types'
import { encodeResourceSrc } from '@/utils/utils'

const props = defineProps<{
  mp4s: Video[]
  backendAddress: string
}>()

const redirectToMp4Src = (backendAddress: string, mp4: Video) => {
  const targetUrl = encodeResourceSrc(backendAddress, mp4.url)
  window.location.href = targetUrl
}
</script>

<template>
  <div class="container-fluid">
    <div class="row mb-4 gx-1 gy-4">
      <div v-for="mp4 in mp4s" :key="mp4.name" class="col-6 col-sm-4 col-md-3 col-xl-2 text-center">
        <button
          @click="redirectToMp4Src(backendAddress, mp4)"
          type="button"
          class="btn-width btn btn-secondary p-0"
        >
          <div class="py-1 truncate mx-auto">
            {{ mp4.name }}
          </div>
          <img
            v-if="mp4.thumbnail"
            :width="mp4.thumbnail.image.width"
            :height="mp4.thumbnail.image.height"
            loading="lazy"
            class="rounded-bottom img-fluid"
            :src="encodeResourceSrc(backendAddress, mp4.thumbnail.image.url)"
          />
        </button>
      </div>
    </div>
  </div>
</template>

<style scoped></style>
