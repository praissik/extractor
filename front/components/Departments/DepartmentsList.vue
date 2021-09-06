<template>
  <div class="departments">
    <ul class="departments__list" @click="setFilterDepartment">
      <li class="departments__list__row first"
          :class="{active: filterDepartment == 0}"
          @click="selectedDepartment = 0"
        >
        Wszystkie
      </li>
      <li v-for="d in departments" :key="d.departmentID"
          class="departments__list__row"
          :class="{active: filterDepartment == d.departmentID}"
          @click="selectedDepartment = d.departmentID"
        >
        {{ d.name }}
        {{ d.departmentID }}
      </li>
    </ul>
  </div>
</template>


<script>
  export default {
    data() {
      return {
        selectedDepartment: 0
      }
    },

    computed: {
      departments () {
        return this._.orderBy(this.$store.state.reports.departments, 'name')
      },
      filterDepartment () {
        return this.$store.state.reports.filterDepartment
      }
    },

    methods: {
      setFilterDepartment () {
        this.$store.dispatch('reports/SetFilterDepartment', this.selectedDepartment)
      }
    }
  }
</script>

<style lang="scss">
.departments {
  display: flex;
  padding: 15px 0 0 0;
  width: 100%;
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