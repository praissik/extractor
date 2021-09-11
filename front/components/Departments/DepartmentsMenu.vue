<template>
  <div>
    <div class="selectedDepartment">
      {{ selectedDepartment }}
    </div>
  <ul class="departments-menu" @click="setFilterDepartmentID">
    <li class="departments-menu__row first"
        :class="{active: filterDepartmentID == 0}"
        @click="selectedDepartmentID = 0"
      >
      Wszystkie
    </li>
    <li v-for="d in departments" :key="d.departmentID"
        class="departments-menu__row"
        :class="{
          active: filterDepartmentID == d.departmentID
        }"
        @click="selectedDepartmentID = d.departmentID"
        @mouseover="hover"
      >
      {{ d.name }}
    </li>
  </ul>
  </div>
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
      },
      selectedDepartment () {
        this.departments.forEach(department => {
          if (department.departmentID == this.filterDepartmentID) {
            console.log(this.filterDepartmentID)
            return department
          }
        });
      }
    },

    methods: {
      setFilterDepartmentID () {
        this.$store.dispatch('reports/SetFilterDepartmentID', this.selectedDepartmentID)
      },
      hover () {
        console.log('hovvvver')
      }
    }
  }
</script>

<style lang="scss">
.leftside {
  ul{
    padding-left: 0;
  }
  .departments-menu {
    display: none;
    flex-direction: column;
    width: 100%;
    margin-top: 15px;
    list-style-type: none;
    font-size: 16px;
    background-color: $base-white;
  }
  @media (max-width: 800px) {
    .departments-menu {
      display: flex;
      min-width: 240px;
      max-width: 240px;
      &__row {
        max-width: 200px;
        justify-content: center;
        padding-left: 0;
        display: none;
        &.active {
          display: flex;
          position: relative;
          background-color: $argenta-transparent-light;
          color: $font-color-primary;
          &:hover {
            background-color: black;
          }
        }
      }
    }
  }
  @media (max-width: 480px) {
    .departments-menu {
      min-width: fit-content;
      max-width: 520px;
      justify-content: center;
      padding: 0 30px;
      &__row {
        padding: 10px 0;
      }
    }
  }
}
</style>