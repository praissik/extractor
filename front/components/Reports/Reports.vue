<template>

  <div class="reports">

    <!--  <Notifications emit="notify"></Notifications> -->
  <v-simple-table>
    <template v-slot:default>
      <tbody>
        <tr
          v-for="(report, lp) in reports"
          :key="report.id"
          @click.stop="rowClicked(report)"
        >
        
          <td>{{ lp+1 }}</td>
          <td>{{ report.reportID }}</td>
          <td>{{ report.name }}</td>
          <td>{{ report.parametersID }}</td>
          <td>{{ report.departmentID }}</td>
        </tr>
      </tbody>
    </template>
  </v-simple-table>
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
      reports() {
        return this._.orderBy(this.$store.state.reports.reports, 'name')
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
  flex-wrap: wrap;
  flex-direction: column;
  justify-content: flex-start;
  width: 100%;
  align-items: center;
  .el-table {
    width: 80%;
    &__header-wrapper {
      display: none;
    }
  }
}

</style>