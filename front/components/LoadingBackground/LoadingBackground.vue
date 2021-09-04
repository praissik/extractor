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
    },
  }
</script>

<style lang="scss">
.loading-background {
  display: flex;
  justify-content: center;
  align-items: center;
  position: absolute;
  background-color: white;
  width: 100%;
  height: 100%;
  img {
    width: 300px;
  }
  .text {
    display: flex;
    flex-basis: min-content;
    font-size: 92px;
    padding-left: 3vw;
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
  99% {opacity: 0; z-index: 0}
  100% {opacity: 0; z-index: -1}
}
</style>