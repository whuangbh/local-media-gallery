import { createRouter, createWebHistory } from 'vue-router'
import SearchView from '@/views/SearchView.vue'
import ContentView from '@/views/ContentView.vue'

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: '/search',
      name: 'Search',
      component: SearchView,
    },
    {
      path: '/:url*',
      name: 'Media Gallery',
      component: ContentView,
    },
    // TODO
    // {
    //   path: "/favorite",
    //   name: "Favorite Viewer",
    //   component: ContentView
    // }
  ],
})

export default router
