<template>
  <div class="departments">
    <ul class="departments__list" @click="setFilterDepartment">
      <li class="departments__list__row first"
          :class="{active: filterDepartment == 0}"
          @click="filterDepartment = 0"
        >
        Wszystkie
      </li>
      <li v-for="d in departments" :key="d.departmentID"
          class="departments__list__row"
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
        filterDepartment: 0
      }
    },

    computed: {
      departments () {
        return this._.orderBy(this.$store.state.reports.departments, 'name')
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
.departments {
  display: flex;
  padding: 20px 0 0 20px;
  width: 280px;
  &__list {
    list-style-type: none;
    font-size: 18px;
    width: 100%;
    background-color: $base-white;
    &__row {
      display: flex;
      justify-content: flex-start;
      align-items: center;
      height: 56px;
      padding-left: 30px;
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