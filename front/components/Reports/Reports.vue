<template>
  <div class="reports">
    <!--  <Notifications emit="notify"></Notifications> -->
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
              (this.filterDepartment === 0
              || data.departmentID === this.filterDepartment))
      },
      filterName () {
        console.log(this.$store.state.reports.filterName)
        return this.$store.state.reports.filterName
      },
      filterDepartment () {
        console.log(this.$store.state.reports.filterDepartment)
        return this.$store.state.reports.filterDepartment
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
  padding: 15px 0 0 20px;
  &__list {
    list-style-type: none;
    font-size: 16px;
    width: 100%;
    background-color: $base-white;
    border-top: 15px solid $base-white;
    border-left: 20px solid $base-white;
    border-right: 20px solid $base-white;
    &__row {
      display: flex;
      justify-content: flex-start;
      align-items: center;
      padding: 15px 20px 15px 20px;
      border-top: 1px solid $base-gray-dark1;
      transition: 250ms ease-out;
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
}


@media (max-width: 800px) {
  .reports {
    padding-right: 20px;
  }
}
</style>