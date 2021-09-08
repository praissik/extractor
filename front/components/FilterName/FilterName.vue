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
  width: 100%;
  padding-top: 22px;
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
@media (max-width:  620px) {
  .filterName {
    display: flex;
  width: 100%;
    max-width: 620px;
    .v-input {
      max-width: 390px;
      min-width: 200px;
      padding: 0 0 0 50px;
    }
  }
}
  
</style>