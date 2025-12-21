<script setup lang="ts">
import { useRouter } from 'vue-router'
import { encodeResourceSrc } from '@/utils/utils'
import type { Folder } from '@/types/types'

const router = useRouter()

const props = defineProps<{
  folders: Folder[]
  backendAddress: string
}>()
</script>

<template>
  <div class="container-fluid">
    <div class="row mb-4 gx-1 gy-4">
      <div
        v-for="folder in folders"
        :key="folder.name"
        class="col-6 col-sm-4 col-md-3 col-xl-2 text-center"
      >
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
            :width="folder.thumbnail.image.width"
            :height="folder.thumbnail.image.height"
            loading="lazy"
            class="rounded-bottom img-fluid"
            :src="encodeResourceSrc(backendAddress, folder.thumbnail.image.url)"
          />
        </button>
      </div>
    </div>
  </div>
</template>

<style scoped></style>
