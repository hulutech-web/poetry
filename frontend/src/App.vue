<script setup lang="ts">
document.body.addEventListener("click", function (event) {
  event.preventDefault();
});
</script>

<template>
  <div class="view">
    <router-view/>
  </div>
</template>

<style lang="scss">
@import url("./assets/css/reset.css");
@import url("./assets/css/font.css");

html {
  width: 100%;
  height: 100%;
}

body {
  width: 100%;
  height: 100%;
  margin: 0;
  padding: 0;
  font-family: "STFangsong";
}

#app {
  position: relative;
  height: 100%;
  overflow: hidden;
}

.view {
  position: absolute;
  left: 0;
  right: 0;
  bottom: 0;
  background-color: #f0f2f5;
  overflow: hidden;
}
</style>
