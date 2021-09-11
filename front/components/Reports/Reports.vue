<template>
  <div class="reports">
    <!--  <Notifications emit="notify"></Notifications> -->
    <template v-if="filterReports.length > 0">
      <ul class="reports__list">
        <li v-for="(r, lp) in filterReports" :key="r.id"
            class="reports__list__row"
            :class="{first: lp == 0}"
            @click="rowClicked(r)"
          >
            {{ lp+1 }}.
            {{ r.name }}
            {{ r.parametersID }}
            {{ r.departmentID }}
        </li>
      </ul>
    </template>
    <template v-else>
      <div class="reports__empty">
        Brak raportów o takiej nazwie
      </div>
    </template>
  <ParametersForm />
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
      };
    },

    mounted() {
      this.$nuxt.$on('closeParametersForm', () => {
        console.log("close - Defined")
      })
    },

    computed: {
      reports () {
        return this._.orderBy(this.$store.state.reports.reports, 'name')
      },
      filterReports () {
        return this.reports.filter(data => (!this.filterName
            || data.name.toLowerCase().includes(this.filterName.toLowerCase()))
              &&
              (this.filterDepartmentID === 0
              || data.departmentID === this.filterDepartmentID))
      },
      filterName () {
        return this.$store.state.reports.filterName
      },
      filterDepartmentID () {
        return this.$store.state.reports.filterDepartmentID
      }
    },

    methods: {
      rowClicked(report) {
        if (this.isEnable) {
          this.pausingButtonEvent(1000)
          if (report.parametersID == null) {
            this.data.reportID = report.reportID
            this.$store.dispatch('reports/GenerateReport', this.data)
          } else {
            this.$nuxt.$emit('openParametersForm', report)
          }
        }
      },

      async pausingButtonEvent(time) {
        this.isEnable = false
        await this.sleep(time)
        this.isEnable = true
       },

      sleep(ms) {
        return new Promise((resolve) => {
          setTimeout(resolve, ms);
        });
      },
    }
  }
</script>

<style lang="scss">

.reports {
  display: flex;
  background-color: $base-white;
  margin: 15px 0 0 20px;
  &__list {
    list-style-type: none;
    font-size: 16px;
    width: 100%;
    padding: 15px 0;
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
      }
      &:hover {
        background-color: $base-gray-light1;
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
  .reports {
    margin: 15px 0 0 0;
    &__list {
      &__row {
        padding: 15px 10px;
      }
    }
  }
}
</style>