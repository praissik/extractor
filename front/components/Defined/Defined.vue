<template>

  <div class="main-view-defined">

    <!--  <Notifications emit="notify"></Notifications> -->

    <div class="header-defined">Wybierz zdefiniowany</div>
    <ul class="reports">
      <li class="report" :class="{ isDisabled: isActive }" v-for="(report, lp) in reports" :key="lp" v-on:click="checkNeeded(lp, report.needed)">
        <div class="report-name">
          <p>{{ lp+1 }}. {{ report.name }}</p>
        </div>
      </li>
    </ul>

    <DownloadReport emit="report_status"></DownloadReport>

    <el-dialog title="Potrzebne informacje" :visible.sync="dialogFormVisible">
      <div class="item-box">
        <el-form
              label-position="left"
              v-for="(need, lp) in needed" :key="need">
          <el-form-item :label="lp+1 + '. ' + values[need] + ':'">
            <el-input v-if="need == 0" v-model="data.akronim"></el-input>
            <el-input v-else-if="need == 1" v-model="data.kodTowaru"></el-input>
            <el-input v-else-if="need == 2" v-model="data.symbolTowaru"></el-input>
            <el-input v-else-if="need == 3" v-model="data.rokOS"></el-input>
          </el-form-item>
        </el-form>
      </div>
      <span slot="footer" class="dialog-footer">
        <el-button @click="dialogFormVisible = false">Anuluj</el-button>
        <el-button type="primary" @click="dialogFormVisible = false, generateReport()">Pobierz raport</el-button>
      </span>
    </el-dialog>
  </div>
</template>

<script>

  import { mapState } from 'vuex'
  //import DownloadReport from "../DownloadReport.vue";
  //import Notifications from "../Notifications.vue";

  export default {
    /*
    components: {
      DownloadReport,
      Notifications
    },
    */

    data() {
      return {
        id: null,
        reports: [],
        needed: [],
        data: {
          akronim: "",
          kodTowaru: "",
          symbolTowaru: "",
          rokOS: ""
        },
        values: ["Akronim kontrahenta", "Kod towaru", "Symbol towaru", "Rok OS [np. 2020]"],
        dialogFormVisible: false,
        isActive: false
      };
    },

    computed: {
      ...mapState({
        home_url: 'home_url',
      })
    },

    methods: {
      getReports() {
        console.log("Getting reports")
        this.$axios
          .get('/reports')
          .then(response => {
            this.reports = response.data.queries
          })
          .catch(error => {
            console.log(error)
          })
      },

      generateReport() {
        console.log("Generating reports")
        this.pausingButtonEvent(3000)
        //this.$eventHub.$emit('notify', "info")
        this.$axios
          .get('/report/generate', {
            params: {
              id: this.id,
              akronim: this.data.akronim,
              kodTowaru: this.data.kodTowaru,
              symbolTowaru: this.data.symbolTowaru,
              rokOS: this.data.rokOS
            }
          })
          .then(response => {
            if (response.data.isOk == 1) {
              this.$eventHub.$emit('notify', "success")
              this.$eventHub.$emit('generated_successfully', response.data)
            }
            else if (response.data.isOk == 0) {
              this.$eventHub.$emit('notify', "error")
            }
            else if (response.data.isOk == -1) {
              this.$eventHub.$emit('notify', "panic")
            }
          })
          .catch(error => {
            console.log(error)
            this.$eventHub.$emit('notify', "error")
          })
        this.id = null
        this.needed = []
        this.data.akronim = ""
        this.data.kodTowaru = ""
        this.data.symbolTowaru = ""
        this.data.rokOS = ""

      },

      checkNeeded(id, needed) {
        this.id = id
        if (needed == null) {
          this.generateReport()
        } else {
          this.needed = [...new Set(needed)].sort((a, b) => a - b )
          this.dialogFormVisible = true
        }
      },

      async pausingButtonEvent(time) {
        this.isActive = true
        await this.sleep(time)
        this.isActive = false
       },

      sleep(ms) {
        return new Promise((resolve) => {
          setTimeout(resolve, ms);
        });
      },
    },

    mounted() {
      this.getReports()
    }
  }
</script>

<style lang="scss">

.main-view-defined {
  display: flex;
  flex-wrap: wrap;
  flex-direction: column;
  justify-content: flex-start;
  background: #f1f1f1;
  margin-top: 2vw;
  padding-bottom: 2vw;
  border-radius: 0.5vw;
  align-items: center;

  .header-defined {
    display: flex;
    justify-content: flex-start;
    width: 100%;
    padding: 1vw 0 0 1.5vw;
    font-size: 2.5vh;
  }

  .reports {
    display: flex;
    flex-wrap: wrap;
    justify-content: center;
    align-content: flex-start;
  }

  .report {
    display: flex;
    justify-content: center;
    align-items: center;
    list-style: none;
    text-align: center;
    width: 18vw;
    height: auto;
    min-height: 6vw;
    font-size: 2vh;
    margin: 1vw 0.5vw 0;
    background: #e3e3e3;
    border: 0.2vw solid #cc3034;
    border-radius: 0.5vw;
    .report-name{
      margin: 0.5vw 0.8vw;
    }
  }

  .isDisabled {
    pointer-events: none;
  }

  .el-dialog {
    max-width: 400px;
    min-width: 400px;
  }

  .el-dialog__title {
    font-size: 2.5vh;
  }

  .el-form-item__label {
    font-size: 2vh;
  }

  .el-dialog__headerbtn {
    display: none !important;
  }

  .el-button--primary {
    background-color: #cc3034;
    border-color: #cc3034;
  }

  .el-button--primary:hover {
    color: #FFF;
    background-color: #cc30359f;
    border-color: #cc30359f;
  }

  .el-button--default:hover {
    color: #000;
    background-color: #cc30353b;
    border-color: #cc30353b;
  }
}
</style>