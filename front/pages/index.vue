<template>
  <v-app>
    <div class="main">
      <LoadingBackground />
      <div class="header">
        <Logo />
      </div>
      <div class="nav">
        <Nav />
      </div>
      <div class="body">
        <Reports />
      </div>
    </div>
  </v-app>
</template>

<script>
  export default {
    mounted() {
      this.$store.dispatch('reports/GetReports')
      this.$store.dispatch('reports/GetParameters')
      this.$store.dispatch('reports/GetDepartments')
    }
  }
</script>

<style lang="scss">
html { 
  overflow-y: auto
}

.main {
  display: flex;
  flex-direction: row;
  flex-wrap: wrap;
  width: 100%;
  color: $font-color-primary;
}

.header, .body {
    width: 100%;
}
.header {
  display: flex;
  align-items: center;
  height: 80px;
}
.body {
  display: flex;
  flex-direction: row;
}
</style>
