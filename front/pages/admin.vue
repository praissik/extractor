<template>
  <v-app>
    <div class="main">
      <template v-if="retrieveData">
        <LoadingBackground />
        <Notifications></Notifications>
        <div class="leftside">
          <Logo />
          <AdminMenu />
        </div>
        <div class="body">
          <FilterString />
          <AdminData />
        </div>
        <div class="footer">
          Argenta © {{ currentYear }}
        </div>
      </template>
      <template v-else>
        <ErrorRetrieveData />
      </template>
    </div>
  </v-app>
</template>

<script>
  export default {
    data () {
      return {
        currentYear: new Date().getFullYear(),
        retrieveData: true
      }
    },

    mounted() {
      this.$store.dispatch('reports/GetData')
        .catch(() => {
          this.retrieveData = false
        })
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
  flex-direction: column;
  width: 100%;
  max-width: 230px;
}
.body {
  display: flex;
  flex-direction: column;
  max-width: 1000px;
  width: 100%;
  margin-bottom: 100px;
}
.footer {
  display: flex;
  justify-content: flex-end;
  position: absolute;
  max-width: 1230px;
  width: 100%;
  bottom: 0;
  padding-right: 20px;
  font-size: 14px;
  background-color: $base-gray-dark1;
}


@media (max-width:  800px) {
  .main {
    flex-wrap: wrap;
    flex-direction: column;
    justify-content: flex-start;
  }
  .leftside {
    display: flex;
    justify-content: center;
    flex-direction: row;
    max-width: 800px;
    width: 100%;
    background-color: $base-white;
    margin-bottom: 15px;
  }
}
@media (max-width: 460px) {
  .leftside {
    flex-wrap: wrap;
  }
}
</style>
