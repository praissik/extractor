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
  justify-content: flex-start;
  align-items: center;
  padding: 6px 20px 0 30px;
  max-width: 450px;
  min-width: 450px;
  width: 100%;
  .v-input {
    font-size: 16px;
  }
  .v-label, .v-text-field input {
    padding-left: 10px;
  }
}
@media (max-width: 800px) {
  .departments-menu {
    display: flex;
    justify-content: flex-end;
    align-items: center;
    max-width: 800px;
    width: 100%;
    padding: 10px 40px 20px;
    &__selected {
      max-width: 200px;
      width: 100%;
    }
  }
}
  
</style>