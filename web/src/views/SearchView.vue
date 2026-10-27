<script setup lang="ts">
import { ref } from 'vue'
import NavBar from '@/components/NavBar.vue'
import ImgDisplay from '@/components/ImgDisplay.vue'
import Mp4Display from '@/components/Mp4Display.vue'
import Controller, { backend } from '@/http/controller.js'
import type { Video, Image, SearchResultsResponse } from '@/types/types'

const backendAddress = backend.defaults.baseURL!

const navBarTitle = ref<string>('Search Media')
const searchQuery = ref<string>('')
const vids = ref<Video[]>([])
const imgs = ref<Image[]>([])

const isSearching = ref<boolean>(false)
const hasSearched = ref<boolean>(false)

async function handleSearch() {
  if (!searchQuery.value.trim()) return

  isSearching.value = true
  hasSearched.value = true

  vids.value = []
  imgs.value = []

  try {
    const res: SearchResultsResponse = await Controller.fetchSearchResults(searchQuery.value)

    console.dir(res)

    if (!res.success) {
      throw res.error
    }
    vids.value = res.videos
    imgs.value = res.images

    navBarTitle.value = 'Search Media' + ': ' + searchQuery.value
  } catch (err) {
    console.error('Search failed:', err)

    navBarTitle.value = 'Error!'
    vids.value = []
    imgs.value = []
  } finally {
    isSearching.value = false
  }
}
</script>

<template>
  <NavBar :navBarTitle="navBarTitle"></NavBar>

  <div class="container-fluid mt-4">
    <!-- Search Input Form -->
    <form @submit.prevent="handleSearch" class="input-group mb-4">
      <input
        type="text"
        v-model="searchQuery"
        class="form-control"
        placeholder="Type to search files, tags, or names..."
      />
      <button class="btn btn-primary" type="submit" :disabled="isSearching">
        {{ isSearching ? 'Searching...' : 'Search' }}
      </button>
    </form>

    <!-- Hint / Status Messages -->
    <div v-if="isSearching">Loading results...</div>
    <div v-else-if="hasSearched && imgs.length === 0 && vids.length === 0">
      No results found for "{{ searchQuery }}".
    </div>

    <!-- Search Results Display -->
    <template v-if="imgs.length !== 0 || vids.length !== 0">
      <Mp4Display :mp4s="vids" :backendAddress="backendAddress" />
      <ImgDisplay :imgs="imgs" :backendAddress="backendAddress" />
    </template>
  </div>
</template>

<style scoped></style>
