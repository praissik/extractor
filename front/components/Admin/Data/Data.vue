<template>
  <div>
    <div v-if="filterMenuOptionID == 0">
      <AdminDataDepartments />
    </div>
    <div v-if="filterMenuOptionID == 1">
      <AdminDataParameters />
    </div>
    <div v-if="filterMenuOptionID == 2">
      <AdminDataReports />
    </div>
  </div>
</template>

<script>

  export default {

    data() {
      return {
        report: [],
        isEnable: true,
        data: {
          reportID: null
        }
      }
    },

    mounted () {
      this.$nuxt.$on('getParameters', () => {
        this.$store.dispatch('reports/GetParameters')
      })
      this.$nuxt.$on('getDepartments', () => {
        this.$store.dispatch('reports/GetDepartments')
      })
      this.$nuxt.$on('getReports', () => {
        this.$store.dispatch('reports/GetReports')
      })
    },

    computed: {
      filterMenuOptionID () {
        return this.$store.state.admin.filterMenuOptionID
      }
    }
  }
</script>

<style lang="scss">
.data {
  display: flex;
  flex-direction: column;
  background-color: $base-white;
  margin: 15px 0 0 20px;
  &__list {
    list-style-type: none;
    font-size: 16px;
    width: 100%;
    padding: 15px 0 0;
    border-left: 20px solid $base-white;
    border-right: 20px solid $base-white;
    &__row {
      display: flex;
      justify-content: flex-start;
      align-items: center;
      padding: 15px 30px;
      border-top: 1px solid $base-gray-dark1;
      transition: 250ms ease-out;
      cursor: default;
      &.first{
        border-top: 0;
        &:hover {
          background-color: $base-white;
        }
      }
      &:hover {
        background-color: $base-gray-light1;
      }
      &.add {
        display: flex;
        justify-content: center;
        align-items: center;
        margin: 0 20px 30px 20px;
        background-color: #e0f5cd;
        border-bottom-left-radius: 30px;
        border-bottom-right-radius: 30px;
        &:hover {
          background-color: #cbe6b3;
        }
      }
    }
  }
  ul{
    padding-left: 0;
  }
  &__empty {
    display: flex;
    justify-content: center;
    align-items: center;
    background-color: $base-white;
    padding: 30px;
    font-size: 16px;
    width: 100%;
  }
}
@media (max-width: 800px) {
  .data {
    margin: 15px 0 0 0;
    &__list {
      &__row {
        padding: 15px 10px;
      }
    }
  }
}
</style>