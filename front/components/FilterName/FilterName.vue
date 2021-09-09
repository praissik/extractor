<template>
  <div class="filterName">
    <v-text-field label="Wyszukaj" v-model="filterName" @input="setFilterName" autofocus></v-text-field>
  </div>
</template>

<script>
  export default {
    data() {
      return {
        filterName: ''
      }
    },

    methods: {
      setFilterName () {
        this.$store.dispatch('reports/SetFilterName', this.filterName)
      }
    }
  }
</script>

<style lang="scss">
.filterName {
  display: flex;
  max-width: 1000px;
  width: 100%;
  padding-top: 22px;
  background-color: $base-white;
  .v-input {
    max-width: 450px;
    padding: 0 0 0 80px;
  }
  .v-label, .v-text-field input {
    padding-left: 10px;
  }
}
@media (max-width: 800px) {
  .filterName {
    .v-input {
      max-width: 350px;
      padding: 0 10px 0 80px;
    }
  }
}
@media (max-width:  680px) {
  .filterName {
    display: flex;
    justify-content: center;
    width: 100%;
    max-width: 680px;
    .v-input {
      max-width: 240px;
      width: 100%;
      padding: 0 40px;
    }
  }
}
  
</style>