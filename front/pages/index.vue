<template>
  <v-app>
    <div class="main">
      <LoadingBackground />
      <div class="header">
        <div class="header__item">
          <Logo />
          <FilterName />
          <DepartmentsMenu />
        </div>
      </div>
      <div class="body">
        <DepartmentsList />
        <div class="article">
          <Reports />
        </div>
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
.body, .footer {
  max-width: 1200px;
  width: 100%;
}
.main {
  display: flex;
  flex-wrap: wrap;
  flex-direction: column;
  align-items: center;
  width: 100%;
  min-height: 100vh;
  color: $font-color-primary;
}
.header {
  display: flex;
  justify-content: center;
  flex-wrap: wrap;
  width: 100%;
  background-color: $base-white;
  &__item {
    display: flex;
    max-width: 1200px;
    width: 100%;
  }
}
.body {
  display: flex;
  height: fit-content;
  margin-bottom: 100px;
  .article {
    display: flex;
    flex-wrap: wrap;
    flex-direction: column;
    height: fit-content;
    width: 100%;
  }
}
.footer {
  display: flex;
  justify-content: flex-end;
  position: absolute;
  bottom: 0;
  padding-right: 20px;
  font-size: 14px;
  background-color: $base-gray-light1;
}
</style>
