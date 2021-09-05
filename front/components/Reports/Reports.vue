<template>
  <div class="reports">
    <!--  <Notifications emit="notify"></Notifications> -->
    <ul class="reports__list">
      <li v-for="(r, lp) in filterReports" :key="r.id"
          class="reports__list__row"
          :class="{first: lp == 0}"
          @click="rowClicked(r)"
        >
          {{ lp+1 }}
          {{ r.reportID }}
          {{ r.name }}
          {{ r.parametersID }}
          {{ r.departmentID }}
      </li>
    </ul>
  <ParametersForm />
          <!-- <v-card-text v-for="(paramID, lp) in report.parametersID" :key="paramID">
            <v-col cols="5" v-if="paramID == 1">
              <v-text-field :label="lp+1 + '. ' + parameters[paramID].name" v-model="data.akronim" :rules="rules.akronim" required />
            </v-col>
            <v-col cols="5" v-if="paramID == 2">
              <v-text-field :label="lp+1 + '. ' + parameters[paramID].name" v-model="data.kodTowaru" :rules="rules.kodTowaru" required />
            </v-col>
            <v-col cols="5" v-if="paramID == 3">
              <v-text-field :label="lp+1 + '. ' + parameters[paramID].name" v-model="data.symbolTowaru" :rules="rules.symbolTowaru" required />
            </v-col>
            <v-col cols="5" v-if="paramID == 4">
              <v-text-field :label="lp+1 + '. ' + parameters[paramID].name" v-model="data.rokOS" :rules="rules.rokOS" required />
            </v-col>
          </v-card-text> -->
          <!-- <v-card-actions class="justify-end">
            <v-spacer></v-spacer>
            <v-btn
              color="blue darken-1"
              text
              @click="dialogFormVisible = false">
              Close
            </v-btn>
            <v-btn
              color="blue darken-1"
              text
              @click="generateReport(), dialogFormVisible = false">
              <i class="fas fa-file-download"></i>
            </v-btn>
          </v-card-actions> -->
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
  padding: 20px 0 0 20px;
  width: 800px;
  &__list {
    list-style-type: none;
    font-size: 18px;
    width: 100%;
    background-color: $base-white;
    border: 20px solid $base-white;
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
      &:hover {
        background-color: $base-gray-light1;
      }
    }
  }
  ul{
    padding-left: 0;
  }
}

</style>