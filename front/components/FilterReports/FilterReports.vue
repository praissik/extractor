<template>
  <div class="filter">
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
.filter {
  display: flex;
  justify-content: flex-start;
  align-items: center;
  padding: 6px 60px 0 60px;
  max-width: 500px;
  min-width: 400px;
  width: 100%;
  .v-input {
    font-size: 16px;
  }
  .v-label, .v-text-field input {
    padding-left: 10px;
  }
}
  
</style>