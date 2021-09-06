<template>
  <div class="departments-menu">
    <div class="departments-menu__selected" @click="menuIsOpen = true">
      <template v-if="filterDepartment == 0">
        Wszystkie
      </template>
      <template v-else>
        {{ filterDepartment }}
        {{ departments[filterDepartment].name }}
      </template>
    </div>
    <template v-if="menuIsOpen">
      <ul class="departments-menu__list" @click="setFilterDepartment">
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
    </template>
  </div>
</template>


<script>
  export default {
    data() {
      return {
        menuIsOpen: false
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
  width: 100%;
  &__selected {

  }
  &__list {
    list-style-type: none;
    font-size: 16px;
    width: 100%;
    background-color: $base-white;
    &__row {
      display: flex;
      justify-content: flex-start;
      align-items: center;
      padding: 10px 30px 10px 20px;
      border-top: 1px solid $base-gray-dark1;
      &.active {
        color: $argenta;
      }
      &.first{
        border-top: 0;
      }
    }
  }
  ul{
    padding-left: 0;
  }
}
</style>