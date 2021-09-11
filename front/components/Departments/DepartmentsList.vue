<template>
  <ul class="departments-list" @click="setFilterDepartmentID">
    <li class="departments-list__row first"
        :class="{active: filterDepartmentID == 0}"
        @click="selectedDepartmentID = 0"
      >
      Wszystkie
    </li>
    <li v-for="d in departments" :key="d.departmentID"
        class="departments-list__row"
        :class="{
          active: filterDepartmentID == d.departmentID
        }"
        @click="selectedDepartmentID = d.departmentID"
      >
      {{ d.name }}
    </li>
  </ul>
</template>

<script>
  export default {
    data() {
      return {
        selectedDepartmentID: 0
      }
    },

    computed: {
      departments () {
        return this._.orderBy(this.$store.state.reports.departments, 'name')
      },
      filterDepartmentID () {
        return this.$store.state.reports.filterDepartmentID
      }
    },

    methods: {
      setFilterDepartmentID () {
        this.$store.dispatch('reports/SetFilterDepartmentID', this.selectedDepartmentID)
      }
    }
  }
</script>

<style lang="scss">
.leftside {
  ul{
    padding-left: 0;
  }
  .departments-list {
    display: flex;
    flex-direction: column;
    width: 100%;
    margin-top: 15px;
    list-style-type: none;
    font-size: 16px;
    background-color: $base-white;
    &__row {
      display: flex;
      width: 100%;
      justify-content: flex-start;
      padding: 10px 0 10px 30px;
      border-top: 1px solid $base-gray-dark1;
      cursor: default;
      &.active {
        color: $argenta;
      }
      &.first{
        border-top: 0;
      }
    }
  }
  @media (max-width: 800px) {
    .departments-list {
      display: none;
    }
  }
}
</style>