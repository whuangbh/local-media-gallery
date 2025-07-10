<script setup>
import { onMounted } from "vue";
import { encodeResourceSrc } from "../utils/utilities";
import PhotoSwipeLightbox from "photoswipe/lightbox";
import constant from "@/assets/constant.js";

const props = defineProps({
  imgs: Array,
  staticFileAddress: String,
  favorites: Array
});

const emit = defineEmits(["addFavorite", "removeFavorite"]);

const isImgObjInFavorites = (index) => {
  return props.favorites.some((imgObj) => imgObj.name === props.imgs.at(index).name);
};

function photoSwipeInit() {
  const lightbox = new PhotoSwipeLightbox({
    gallery: "#my-gallery",
    children: "a",
    imageClickAction: "close",
    tapAction: "close",
    pswpModule: () => import("photoswipe")
  });

  lightbox.on("uiRegister", () => {
    lightbox.pswp.ui.registerElement({
      name: "fav-btn",
      order: 9,
      isButton: true,
      html: constant.favBtnEmptyStr,
      onInit: (el, pswp) => {
        pswp.on("change", () => {
          const currentIndex = pswp.currSlide.index;
          if (isImgObjInFavorites(currentIndex)) el.innerHTML = constant.favBtnFilledStr;
          else el.innerHTML = constant.favBtnEmptyStr;
        });
      },

      onClick: async (event, el, pswp) => {
        const currentIndex = pswp.currSlide.index;
        const data = props.imgs.at(currentIndex);

        el.innerHTML = constant.spinnerStr;

        try {
          if (isImgObjInFavorites(currentIndex)) {
            emit("removeFavorite", data, () => {
              el.innerHTML = constant.favBtnEmptyStr;
            });
          } else {
            emit("addFavorite", data, () => {
              el.innerHTML = constant.favBtnFilledStr;
            });
          }
        } catch (error) {
          console.error("Error updating favorite:", error);
        }
      }
    });
  });
  lightbox.init();
}

onMounted(() => {
  photoSwipeInit();
});
</script>

<template>
  <div class="container-fluid mb-5">
    <div class="row pswp-gallery pswp-gallery--single-column gx-1 gy-3" id="my-gallery">
      <div v-for="img in imgs" :key="img.name" class="col-6 col-xl-2">
        <a
          :data-pswp-width="img.width"
          :data-pswp-height="img.height"
          target="_blank"
          :href="encodeResourceSrc(staticFileAddress, img.url)"
        >
          <img
            :width="img.width"
            :height="img.height"
            loading="lazy"
            class="img-fluid border border-dark-subtle"
            :src="encodeResourceSrc(staticFileAddress, img.url)"
          />
        </a>
      </div>
    </div>
  </div>
</template>

<style scoped></style>
