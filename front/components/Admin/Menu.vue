<template>
  <ul class="menu-list" @click="setFilterMenuOption">
    <li class="menu-list__row first"
        :class="{active: filterMenuOptionID == 0}"
        @click="selectedMenuOptionID = 0"
      >
      Działy
    </li>
    <li class="menu-list__row"
        :class="{active: filterMenuOptionID == 1}"
        @click="selectedMenuOptionID = 1"
      >
      Parametry
    </li>
    <li class="menu-list__row"
        :class="{active: filterMenuOptionID == 2}"
        @click="selectedMenuOptionID = 2"
      >
      Raporty
    </li>
  </ul>
</template>

<script>
  export default {
    data() {
      return {
        selectedMenuOptionID: 0
      }
    },

    computed: {
      departments () {
        return this._.orderBy(this.$store.state.reports.departments, 'name')
      },
      filterMenuOptionID () {
        return this.$store.state.admin.filterMenuOptionID
      }
    },

    methods: {
      setFilterMenuOption () {
        this.$store.dispatch('admin/SetFilterMenuOption', this.selectedMenuOptionID)
      }
    }
  }
</script>

<style lang="scss">
.leftside {
  ul{
    padding-left: 0;
  }
  .menu-list {
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
    .menu-list {
      display: none;
    }
  }
}
</style>