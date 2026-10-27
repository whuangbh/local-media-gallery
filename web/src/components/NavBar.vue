<script setup lang="ts">
import { useRoute, useRouter } from 'vue-router'

const route = useRoute()
const router = useRouter()

const props = defineProps<{
  navBarTitle: string
}>()

const goBackToUpperPath = () => {
  // If we are currently on the search page, always go straight back to root
  if (route.name === 'Search') {
    router.push('/')
    return
  }

  // Otherwise, use existing folder-splitting logic
  const pathSegments = props.navBarTitle.split('/').filter(Boolean)
  pathSegments.pop()
  const newPath = `/${pathSegments.join('/')}`

  router.push(newPath)
}
</script>

<template>
  <div class="container-fluid my-3">
    <div class="h2 mt-1 d-flex">
      <button
        v-if="route.path !== '/'"
        type="button"
        class="btn btn-secondary d-inline me-2"
        @click="goBackToUpperPath"
      >
        <i className="bi bi-chevron-left me-1"></i>
      </button>
      <span className="align-self-center"> {{ navBarTitle || 'Loading...' }} </span>
      <!-- <button
        v-if="route.name === 'Favorite Viewer'"
        type="button"
        class="btn btn-secondary d-inline ms-auto"
        @click="$emit('onClickSortFavBtn')"
      >
        Sort/Refresh Fav
      </button> -->
      <router-link
        v-if="route.path === '/'"
        to="/search"
        type="button"
        class="btn btn-secondary d-inline ms-auto"
        >Search
      </router-link>
    </div>
  </div>
</template>

<style scoped></style>
