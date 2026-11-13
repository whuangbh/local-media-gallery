<script setup lang="ts">
import { onMounted, onUnmounted } from 'vue'
import { getStaticUrl, getThumbnailUrl } from '@/utils/utils'
import PhotoSwipeLightbox from 'photoswipe/lightbox'
import 'photoswipe/style.css'
import type { Image } from '@/types/types'

const props = defineProps<{
  imgs: Image[]
  backendAddress: string
  // favorites: Array
}>()

const galleryElementId = 'photo-swipe-gallery'

// 停靠观察器：远离视口 >3 屏 → 释放解码内存；回到附近 → 恢复
const parkObserver = new IntersectionObserver(
  (entries) => {
    for (const entry of entries) {
      const img = entry.target as HTMLImageElement
      if (!entry.isIntersecting) {
        // 远离视口：停靠
        if (img.src && !img.dataset.parkedSrc) {
          img.dataset.parkedSrc = img.src
          img.removeAttribute('src')
        }
      } else {
        // 回到视口附近：恢复
        if (!img.src && img.dataset.parkedSrc) {
          img.src = img.dataset.parkedSrc
          delete img.dataset.parkedSrc
        }
      }
    }
  },
  { rootMargin: '300% 0px 300% 0px' }, // 视口上下各 3 屏的缓冲带
)

// 加载观察器：进入视口前 300px 开始加载
const loadObserver = new IntersectionObserver(
  (entries) => {
    for (const entry of entries) {
      if (!entry.isIntersecting) continue
      const img = entry.target as HTMLImageElement
      const src = img.dataset.src
      if (src) {
        img.src = src
        img.removeAttribute('data-src')
      }
      loadObserver.unobserve(img) // 加载后不再观察
    }
  },
  { rootMargin: '300px 0px', threshold: 0.01 },
)

function photoSwipeInit() {
  const lightbox = new PhotoSwipeLightbox({
    gallery: '#' + galleryElementId,
    children: 'a',
    tapAction: 'close',
    pswpModule: () => import('photoswipe'),
  })

  // lightbox.on('uiRegister', () => {
  //   lightbox.pswp.ui.registerElement({
  //     name: 'fav-btn',
  //     order: 9,
  //     isButton: true,
  //     html: constant.favBtnEmptyStr,
  //     onInit: (el, pswp) => {
  //       pswp.on('change', () => {
  //         const currentIndex = pswp.currSlide.index
  //         if (isImgObjInFavorites(currentIndex)) el.innerHTML = constant.favBtnFilledStr
  //         else el.innerHTML = constant.favBtnEmptyStr
  //       })
  //     },

  //     onClick: async (event, el, pswp) => {
  //       const currentIndex = pswp.currSlide.index
  //       const data = props.imgs.at(currentIndex)

  //       el.innerHTML = constant.spinnerStr

  //       try {
  //         if (isImgObjInFavorites(currentIndex)) {
  //           emit('removeFavorite', data, () => {
  //             el.innerHTML = constant.favBtnEmptyStr
  //           })
  //         } else {
  //           emit('addFavorite', data, () => {
  //             el.innerHTML = constant.favBtnFilledStr
  //           })
  //         }
  //       } catch (error) {
  //         console.error('Error updating favorite:', error)
  //       }
  //     },
  //   })
  // })
  lightbox.init()
}

onMounted(() => {
  const gallery = document.getElementById(galleryElementId)
  if (gallery) {
    const imgs = gallery.querySelectorAll<HTMLImageElement>('img[data-src]')
    imgs.forEach((img) => {
      loadObserver.observe(img)
      parkObserver.observe(img)
    })
  }
  photoSwipeInit()
})

onUnmounted(() => {
  loadObserver.disconnect()
  parkObserver.disconnect()
})
</script>

<template>
  <div class="container-fluid">
    <div class="row pswp-gallery pswp-gallery--single-column mb-4 gx-1 gy-4" :id="galleryElementId">
      <div v-for="img in imgs" :key="img.url" class="col-6 col-sm-4 col-md-3 col-xl-2 text-center">
        <a
          :data-pswp-width="img.width"
          :data-pswp-height="img.height"
          target="_blank"
          :href="getStaticUrl(props.backendAddress, img.url)"
        >
          <img
            :width="img.width"
            :height="img.height"
            class="img-fluid border border-dark-subtle"
            :data-src="getThumbnailUrl(props.backendAddress, img.url)"
          />
        </a>
      </div>
    </div>
  </div>
</template>

<style scoped></style>
