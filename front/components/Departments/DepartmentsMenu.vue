<template>
  <div class="departments-menu">
      <div class="departments-menu__selected">
        <template v-if="filterDepartment == 0">
          Administracja
        </template>
        <template v-else>
          {{ filterDepartment }}
          {{ departments[filterDepartment].name }}
        </template>
      </div>
      <ul class="departments-menu__list"
          @click="setFilterDepartment">
        <li class="departments-menu__list__row first"
            :class="{active: filterDepartment == 0}"
            @click="filterDepartment = 0"
          >
          Wszystkie
        </li>
        <li v-for="d in departments" :key="d.departmentID"
            class="departments-menu__list__row"
            :class="{active: filterDepartment == d.departmentID}"
            @click="filterDepartment = d.departmentID"
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
        isFocus: false
      }
    },

    computed: {
      departments () {
        return this.$store.state.reports.departments
      },
      filterDepartment () {
        return this.$store.state.reports.filterDepartment
      }
    },

    methods: {
      setFilterDepartment () {
        this.$store.dispatch('reports/SetFilterDepartment', this.filterDepartment)
      }
    }
  }
</script>

<style lang="scss">
.departments-menu {
  display: flex;
  justify-content: center;
  position: relative;
  max-width: 230px;
  padding: 0px 40px;
  &__selected {
    display: flex;
    justify-content: center;
    align-items: center;
    width: 100%;
    height: 34px;
    border-bottom: 1px solid $base-gray-dark2;
    color: $argenta;
  }
  // &__list {
  //   list-style-type: none;
  //   font-size: 16px;
  //   width: 100%;
  //   background-color: $base-white;
  //   &__row {
  //     display: flex;
  //     justify-content: flex-start;
  //     align-items: center;
  //     padding: 10px 30px 10px 20px;
  //     border-top: 1px solid $base-gray-dark1;
  //     &.active {
  //       color: $argenta;
  //     }
  //     &.first{
  //       border-top: 0;
  //     }
  //   }
  // }
  &__list {
    visibility: hidden;
    width: 120px;
    background-color: black;
    color: #fff;
    text-align: center;
    border-radius: 6px;
    padding: 5px 0;
    position: absolute;
    z-index: 1;
    bottom: 150%;
    left: 50%;
    margin-left: -60px;
  }
  &__list::after {
    content: "";
    position: absolute;
    top: 100%;
    left: 50%;
    margin-left: -5px;
    border-width: 5px;
    border-style: solid;
    border-color: transparent transparent black transparent;
  }
  ul{
    padding-left: 0;
  }
  &:hover &__list {
    visibility: visible;
  }
}
@media (max-width: 800px) {
  .departments-menu {
    display: flex;
    justify-content: flex-end;
    align-items: center;
    max-width: 800px;
    width: 100%;
    padding: 10px 40px 20px;
    &__selected {
      max-width: 200px;
      width: 100%;
    }
  }
}
@media (min-width: 801px) {
  .departments-menu {
    display: none;
  }
}
</style>