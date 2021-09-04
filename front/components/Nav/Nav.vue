<template>
  <div class="nav">
    <div class="nav__header">
    </div>
    <div class="nav__content">
      <div v-for="d in departments" :key="d.departmentID">
        <div class="nav__content__row">
          {{ d.name }}
        </div>
      </div>
    </div>
  </div>
</template>


<script>
  export default {
    data() {
      return {

      }
    },

    computed: {
      departments () {
        return this._.orderBy(this.$store.state.reports.departments, 'name')
      }
    },

    methods: {
      handleOpen(key, keyPath) {
        console.log(key, keyPath)
      },
      handleClose(key, keyPath) {
        console.log(key, keyPath)
      }
    }
  }
</script>

<style lang="scss">
  .nav {
    display: flex;
    flex-direction: row;
    min-width: 200px;
    margin: 0 20px;

    .el-menu-item {
      font-size: 20px;
      color: $font-color-primary;
      border: 0;
      padding: 0;
      margin: 0;
      background: #ececec;
      &:hover {
        background: #dedede;
      }
      &:hover.is-active {
        background: #dedede;
      }
    }

    .el-col-12 {
        width: 100%;
    }
    .tac {
        width: 100%;
    }

    .el-menu {
      border: 0;
    }
    
    .el-menu-item.is-active {
      background: #dedede;
      font-weight: bold;
    }
  }
</style>