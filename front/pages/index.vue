<template>
  <v-app>
    <div class="main">
      <LoadingBackground />
      <div class="header">
        <Logo />
        <FilterReports />
      </div>
      <div class="body">
        <div class="leftside">
          <Departments />
        </div>
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
      this.$store.dispatch('reports/GetReports')
      this.$store.dispatch('reports/GetParameters')
      this.$store.dispatch('reports/GetDepartments')
    }
  }
</script>

<style lang="scss">
.header, .body, .footer {
  width: 100vw;
}

.main {
  display: flex;
  flex-wrap: wrap;
  flex-direction: column;
  min-height: 100vh;
  color: $font-color-primary;
}
.header {
  display: flex;
  height: 80px;
  background-color: $base-white;
}
.body {
  display: flex;
  flex-wrap: wrap;
  flex-direction: row;
  height: fit-content;
  margin-bottom: 100px;
  // padding: 0 24px;
  .leftside {
    display: flex;
    flex-wrap: wrap;
    flex-direction: column;
    height: fit-content;
    // padding: 0 80px 0 0;
  }
  .article {
    display: flex;
    flex-wrap: wrap;
    flex-direction: column;
    height: fit-content;
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
