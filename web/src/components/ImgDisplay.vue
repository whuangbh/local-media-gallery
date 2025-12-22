<script setup lang="ts">
import { onMounted } from 'vue'
import { encodeResourceSrc } from '@/utils/utils'
import PhotoSwipeLightbox from 'photoswipe/lightbox'
import 'photoswipe/style.css'
import constant from '@/assets/constant.js'
import type { Image } from '@/types/types'

const props = defineProps<{
  imgs: Image[]
  backendAddress: string
  // favorites: Array
}>()

const galleryElementId = 'photo-swipe-gallery'

// const emit = defineEmits(['addFavorite', 'removeFavorite'])

// const isImgObjInFavorites = (index) => {
//   return props.favorites.some((imgObj) => imgObj.name === props.imgs.at(index).name)
// }

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
  photoSwipeInit()
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
          :href="encodeResourceSrc(backendAddress, img.url)"
        >
          <img
            :width="img.width"
            :height="img.height"
            loading="lazy"
            class="img-fluid border border-dark-subtle"
            :src="encodeResourceSrc(backendAddress, img.url)"
          />
        </a>
      </div>
    </div>
  </div>
</template>

<style scoped></style>
