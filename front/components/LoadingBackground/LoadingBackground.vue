<template>
  <div class="loading-background" :class="{ hide: dataLoaded }">
    <img src="/logo-only.png">
    <div class="text">
      Raporty
    </div>
  </div>
</template>

<script>
  export default {
    data () {
      return {
      }
    },

    computed: {
      dataLoaded () {
        return Object.keys(this.$store.state.reports.departments).length > 0
      }
    }
  }
</script>

<style lang="scss">
.loading-background {
  display: flex;
  justify-content: center;
  align-items: center;
  position: absolute;
  background-color: $base-white;
  z-index: 1;
  width: 100%;
  height: 100vh;
  padding-bottom: 100px;
  & img {
    width: 22vh;
  }
  .text {
    display: flex;
    font-size: 8vh;
    padding-left: 2vw;
  }
}

.loading-background.hide {
  animation-name: hiding;
  animation-duration: 1000ms;
  animation-timing-function: ease-in;
  animation-fill-mode: forwards;
}

@keyframes hiding {
  0% {opacity: 1}
  99% {opacity: 0; z-index: 1}
  100% {opacity: 0; z-index: -1}
}

@media (max-width: 800px) {
  .loading-background {
    & img {
      width: 22vw;
    }
    .text {
      font-size: 8vw;
      padding-left: 3vw;
    }
  }
}
</style>