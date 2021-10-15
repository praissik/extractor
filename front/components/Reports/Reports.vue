<template>
  <div class="reports">
    <template v-if="filterReports.length > 0">
      <ul class="reports__list">
        <li v-for="(r, lp) in filterReports" :key="r.id"
            class="reports__list__row"
            :class="{first: lp == 0}"
            @click="rowClicked(r)"
          >
          <template v-if="r.active">
            {{ lp+1 }}.
            {{ r.reportName }}
            {{ r.parametersID }}
            {{ r.departmentID }}
          </template>
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
          id: null
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
        return this._.orderBy(this.$store.state.reports.reports, 'reportName')
      },
      filterReports () {
        return this.reports.filter(data => (!this.filterString
            || data.reportName.toLowerCase().includes(this.filterString.toLowerCase()))
              &&
              (this.filterDepartmentID === 0
              || data.departmentID === this.filterDepartmentID)
              &&
              (data.active === true))
      },
      filterString () {
        return this.$store.state.reports.filterString
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
            this.data.id = report.id
            console.log(this.data)
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
      }
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