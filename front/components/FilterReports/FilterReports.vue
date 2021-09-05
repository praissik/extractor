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
  align-items: flex-end;
  padding: 0;
  height: 80px;
  .v-input {
    font-size: 20px;
  }
}
  
</style>