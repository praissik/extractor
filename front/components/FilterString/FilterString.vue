<template>
  <div class="filterString">
    <v-text-field 
        label="Wyszukaj"
        color="#D7172F"
        v-model="filterString"
        @input="setFilterString"
        autofocus></v-text-field>
  </div>
</template>

<script>
  export default {
    data() {
      return {
        filterString: ''
      }
    },

    methods: {
      setFilterString () {
        this.$store.dispatch('reports/SetFilterString', this.filterString)
      }
    }
  }
</script>

<style lang="scss">
.filterString {
  display: flex;
  max-width: 1000px;
  width: 100%;
  min-height: 80px;
  max-height: 80px;
  padding: 10px 60px 0 70px;
  background-color: $base-white;
  .v-input {
    max-width: 600px;
  }
  .v-label, .v-text-field input {
    padding-left: 10px;
  }
}
@media (max-width: 800px) {
  .filterString {
    max-width: 800px;
    width: 100%;
    padding: 10px 30px;
    .v-input {
      max-width: 100%;
    }
  }
}
  
</style>