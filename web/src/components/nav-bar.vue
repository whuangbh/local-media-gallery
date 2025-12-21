<script setup lang="ts">
import { useRoute, useRouter } from 'vue-router'
import BackButtonIcon from './icons/back-button.vue'

const route = useRoute()
const router = useRouter()

// eslint-disable-next-line no-unused-vars
const props = defineProps({
  navBarTitle: String,
})

const goBackToUpperPath = () => {
  const pathSegments = route.path.split('/')
  pathSegments.pop()
  router.push(pathSegments.length === 1 && pathSegments[0] === '' ? '/' : pathSegments.join('/'))
}
</script>

<template>
  <div class="container-fluid">
    <div class="h2 mt-1 d-flex align-items-center">
      <button
        v-if="route.path !== '/'"
        type="button"
        class="btn btn-secondary d-inline me-2"
        @click="goBackToUpperPath"
      >
        <BackButtonIcon />
      </button>
      <span> {{ navBarTitle || 'Loading...' }} </span>
      <button
        v-if="route.name === 'Favorite Viewer'"
        type="button"
        class="btn btn-secondary d-inline ms-auto"
        @click="$emit('onClickSortFavBtn')"
      >
        Sort/Refresh Fav
      </button>
    </div>
  </div>
</template>

<style scoped></style>
