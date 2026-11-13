<script setup lang="ts">
import type { Video } from '@/types/types'
import { getStaticUrl, getThumbnailUrl } from '@/utils/utils'
import { ref } from 'vue'

const props = defineProps<{
  mp4s: Video[]
  backendAddress: string
}>()

const playingVideoUrl = ref<string | null>(null)
const playingVideoName = ref<string>('')

const playVideo = (backendAddress: string, mp4: Video) => {
  playingVideoUrl.value = getStaticUrl(backendAddress, mp4.url)
  playingVideoName.value = mp4.name
}

const closeVideo = () => {
  playingVideoUrl.value = null
  playingVideoName.value = ''
}

// const redirectToMp4Src = (backendAddress: string, mp4: Video) => {
//   const targetUrl = getStaticUrl(backendAddress, mp4.url)
//   window.location.href = targetUrl
// }
</script>

<template>
  <div class="container-fluid">
    <div class="row mb-4 gx-1 gy-4">
      <div v-for="mp4 in mp4s" :key="mp4.name" class="col-6 col-sm-4 col-md-3 col-xl-2 text-center">
        <button
          @click="playVideo(backendAddress, mp4)"
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
            :src="getThumbnailUrl(backendAddress, mp4.thumbnail.image.url)"
          />
        </button>
      </div>
    </div>
    <!-- 视频播放遮罩 (纯 Bootstrap 类名) -->
    <div
      v-if="playingVideoUrl"
      class="position-fixed top-0 start-0 w-100 h-100 d-flex flex-column bg-black"
      style="z-index: 9999"
    >
      <!-- 顶部控制栏 -->
      <div class="d-flex justify-content-between align-items-center p-3 bg-dark text-white">
        <span class="text-truncate me-3 fs-6">{{ playingVideoName }}</span>
        <!-- 使用 Bootstrap 的原生关闭按钮 (btn-close-white 适配深色背景) -->
        <button
          @click="closeVideo"
          type="button"
          class="btn-close btn-close-white shadow-none"
          aria-label="Close"
        ></button>
      </div>

      <!-- 视频容器 -->
      <div
        class="flex-grow-1 d-flex justify-content-center align-items-center overflow-hidden"
        @click.self="closeVideo"
      >
        <video
          :src="playingVideoUrl"
          controls
          autoplay
          playsinline
          webkit-playsinline
          x5-playsinline
          class="mw-100 mh-100"
          style="object-fit: contain"
        ></video>
      </div>
    </div>
  </div>
</template>

<style scoped></style>
