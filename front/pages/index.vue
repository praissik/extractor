<template>
  <v-app>
    <div class="main">
      <LoadingBackground />
      <div class="leftside">
        <Logo />
        <DepartmentsList />
      </div>
      <div class="body">
        <FilterName />
        <Reports />
      </div>
      <div class="footer">
        Argenta © {{ currentYear }}
      </div>
    </div>
  </v-app>
</template>

<script>
  export default {
    data () {
      return {
        currentYear: new Date().getFullYear()
      }
    },

    mounted() {
      this.$store.dispatch('reports/GetData')
    }
  }
</script>

<style lang="scss">
.main {
  display: flex;
  flex-direction: row;
  justify-content: center;
  align-items: flex-start;
  width: 100%;
  min-height: 100vh;
  color: $font-color-primary;
}
.leftside {
  display: flex;
  justify-content: center;
  flex-direction: column;
  width: 100%;
  max-width: 230px;
  background-color: $base-white;
}
.body {
  display: flex;
  flex-direction: column;
  height: fit-content;
  max-width: 1000px;
  width: 100%;
  margin-bottom: 100px;
}
.footer {
  display: flex;
  justify-content: flex-end;
  position: absolute;
  max-width: 1200px;
  width: 100%;
  bottom: 0;
  padding-right: 20px;
  font-size: 14px;
  background-color: $base-gray-light1;
}


@media (max-width:  680px) {
  .header {
    display: flex;
    justify-content: center;
    width: 100%;
    background-color: $base-white;
    &__item {
      display: flex;
      flex-wrap: wrap;
    }
  }
}
</style>
