<template>
  <div class="loading-page" :class="{ hide: dataLoaded }">
    <div class="loading-page__logo">
      <img src="/logo-only.png">
      <div class="loading-page__logo__text">
        Raporty
      </div>
    </div>
    <template v-ifc>
      <div class="loading-page__loader" :class="{ hideLoader: dataLoaded}">
        <div class="loading-page__loader__loaderBar"></div>
      </div>
    </template>
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
.loading-page {
  display: flex;
  justify-content: center;
  align-items: center;
  flex-direction: column;
  position: absolute;
  background-color: $base-white;
  padding-bottom: 15vh;
  z-index: 1;
  width: 100%;
  height: 100vh;
  &__logo {
    display: flex;
    justify-content: center;
    align-items: center;
    flex-direction: column;
    width: 100%;
    & img {
      width: 22vh;
    }
    &__text {
      display: flex;
      font-size: 8vh;
      padding: 10px 0 50px;
    }
  }
  &__loader { 
    display: flex;
    position: relative;
    width: 100%;
    max-width: 250px;
    margin: 0 auto;  
    border-radius: 10px;
    border: 4px solid transparent;
    padding: 3px;
    &:before {
      content: '';
      border: 2px solid $font-color-primary; 
      border-radius: 10px;
      position: absolute;
      top: -4px; 
      right: -4px; 
      bottom: -4px; 
      left: -4px;
    }
    &__loaderBar { 
      position: absolute;
      border-radius: 10px;
      top: 0;
      right: 100%;
      bottom: 0;
      left: 0;
      background: $font-color-primary; 
      width: 0;
      animation: borealisBar 1s linear infinite;
    }
  }
  &.hide {
    animation-name: hiding;
    animation-duration: 1000ms;
    animation-timing-function: ease-in;
    animation-fill-mode: forwards;
  }
}
.hideLoader {
  visibility: hidden !important;
}

@keyframes hiding {
  0% {opacity: 1}
  99% {opacity: 0; z-index: 1}
  100% {opacity: 0; z-index: -1}
}

@media (max-width: 800px) {
  .loading-page {
    &__logo {
      & img {
        width: 24vw;
      }
      &__text {
        font-size: 9vw;
        padding: 2vw 0 7vw;
      }
    }
    &__loader { 
      width: 30vw;
    }
  }
}

@keyframes borealisBar {
  0% {
    left:0%;
    right:100%;
    width:0%;
  }
  20% {
    left:0%;
    right:75%;
    width:25%;
  }
  80% {
    right:0%;
    left:75%;
    width:25%;
  }
  100% {
    left:100%;
    right:0%;
    width:0%;
  }
}
</style>