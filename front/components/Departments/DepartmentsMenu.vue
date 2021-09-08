<template>
  <div class="departments-menu">
    <div class="departments-menu__selected">
      <template v-if="filterDepartment == 0">
        <div class="departments-menu__selected__text">
          Administracja
        </div>
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
@media (max-width: 800px) {
  .departments-menu {
    display: flex;
    position: relative;
    width: 250px;
    margin: 25px 40px 20px;
    padding: 5px 0 5px;
    &__selected {
      display: flex;
      justify-content: center;
      width: 100%;
      border-bottom: 1px solid $base-gray-dark2;
      color: $argenta;
      padding: 0 0 4px 0;
    }
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
      top: 150%;
      left: 50%;
      margin-left: -60px;
    }
    &__list::after {
      content: "";
      position: absolute;
      bottom: 100%;
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
}
@media (min-width: 801px) {
  .departments-menu {
    display: none;
  }
}
</style>