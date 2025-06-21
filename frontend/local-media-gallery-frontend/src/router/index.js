import { createRouter, createWebHistory } from "vue-router";
import HomeView from "../views/HomeView.vue";
import ContentView from "../views/ContentView.vue";

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: "/",
      name: "Home",
      component: HomeView
    },
    {
      path: "/img/:url*",
      name: "Img Viewer",
      component: ContentView
    },
    {
      path: "/vid/:url*",
      name: "Vid Viewer",
      component: ContentView
    },
    {
      path: "/favorite",
      name: "Favorite Viewer",
      component: ContentView
    }
  ]
});

export default router;
