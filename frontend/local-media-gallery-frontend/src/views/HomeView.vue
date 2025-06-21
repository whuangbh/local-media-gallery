<script setup></script>

<template>
  <div class="container-fluid">
    <div id="full-screen" class="row align-items-center justify-content-center">
      <div>
        <div class="text-center">
          <router-link
            to="/vid"
            type="button"
            class="btn btn-primary btn-lg d-inline-block mb-2 mx-auto"
          >
            Vid
          </router-link>
        </div>
        <div class="text-center">
          <router-link
            to="/img"
            type="button"
            class="btn btn-primary btn-lg d-inline-block mb-2 mx-auto"
          >
            Img
          </router-link>
        </div>
        <!--  TODO      -->
        <!--        <div class="text-center">-->
        <!--          <router-link-->
        <!--            to="/favorite"-->
        <!--            type="button"-->
        <!--            class="btn btn-primary btn-lg d-inline-block mb-2 mx-auto"-->
        <!--          >-->
        <!--            Favorites-->
        <!--          </router-link>-->
        <!--        </div>-->
      </div>
    </div>
  </div>
</template>

<style scoped>
div#full-screen {
  width: 100vw;
  height: 100vh;
}

a {
  display: inline-block;
  text-decoration: none;
}
</style>
