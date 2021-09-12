<template>
  <div>
    <div class="departments-menu"
      @mouseover="showMenu = true"
      @mouseleave="showMenu = false">
      <div class="departments-menu__filterDepartment">
        {{ filterDepartmentName }}
          <i class="arrow down"></i>
      </div>
      <ul v-if="showMenu" class="departments-menu__list" @click="setFilterDepartmentID">
        <li class="departments-menu__list__row first"
            :class="{active: filterDepartmentID == 0}"
            @click="selectedDepartmentID = 0,
                    showMenu = false"
          >
          Wszystkie
        </li>
        <li v-for="d in departments" :key="d.departmentID"
            class="departments-menu__list__row"
            :class="{active: filterDepartmentID == d.departmentID}"
            @click="selectedDepartmentID = d.departmentID,
                    showMenu = false"
          >
          {{ d.name }}
        </li>
      </ul>
    </div>
  </div>
</template>

<script>
  export default {
    data() {
      return {
        selectedDepartmentID: 0,
        showMenu: false,
      }
    },

    computed: {
      departments () {
        return this._.orderBy(this.$store.state.reports.departments, 'name')
      },
      filterDepartmentID () {
        return this.$store.state.reports.filterDepartmentID
      },
      filterDepartmentName () {
        var filterDepartmentName = 'Wszystkie'
        if (this.filterDepartmentID > 0) {
          this.departments.forEach(department => {
            if (department.departmentID == this.filterDepartmentID) {
              filterDepartmentName = department.name
            }
          })
        }
        return filterDepartmentName
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
    list-style-type: none;
  }
  .departments-menu {
    display: none;
    position: relative;
    flex-direction: column;
    align-items: flex-end;
    min-width: 240px;
    max-width: 240px;
    width: 100%;
    font-size: 16px;
    margin: 19px 40px 0 0;
    cursor: default;
    &__filterDepartment {
      display: flex;
      align-items: center;
      max-width: 180px;
      width: 100%;
      max-height: 44px;
      min-height: 44px;
      padding-left: 30px;
      background-color: $base-gray-light2;
      border-bottom: 1px solid $argenta-transparent-dark;
      & .arrow.down {
        position: absolute;
        border: solid black;
        border-width: 0 3px 3px 0;
        display: inline-block;
        padding: 3px;
        right: 15px;
        transform: rotate(45deg);
        -webkit-transform: rotate(45deg);
        animation-name: rotate-out;
        animation-duration: 400ms;
        animation-fill-mode: forwards;
      }
    }
    &__list {
      z-index: 1;
      position: absolute;
      display: flex;
      flex-direction: column;
      width: 100%;
      max-width: 180px;
      list-style-type: none;
      background-color: $base-white;
      box-shadow: 2px 2px 8px $base-gray-dark2;
      top: 44px;
      &__row {
        display: flex;
        align-items: center;
        max-height: 44px;
        min-height: 44px;
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
  }
  .departments-menu:hover {
    & .arrow.down {
      animation-name: rotate-in;
      animation-duration: 400ms;
      animation-fill-mode: forwards;
    }
  }
  @media (max-width: 800px) {
    .departments-menu {
      display: flex;
    }
  }
  @media (max-width: 480px) {
    .departments-menu {
      min-width: fit-content;
      max-width: 480px;
      justify-content: center;
    margin: 10px 20px;
      &__list {
        min-width: 200px;
        max-width: 200px;
      }
    }
  }
}
@keyframes rotate-in {
  0% {
    transform: rotate(45deg);
    -webkit-transform: rotate(45deg);
  }
  100% {
    transform: rotate(135deg);
    -webkit-transform: rotate(135deg);
  }
}

@keyframes rotate-out {
  0% {
    transform: rotate(135deg);
    -webkit-transform: rotate(135deg);
  }
  100% {
    transform: rotate(45deg);
    -webkit-transform: rotate(45deg);
  }
}
</style>