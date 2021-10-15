<template>
    <div class="data">
        <template v-if="filterReports.length > 0">
          <ul class="data__list">
            <li class="data__list__row first">
                <div style="min-width: 70px">Lp.</div>
                <div style="min-width: 70px; text-wrap: wrap-word">ID</div>
                <div style="min-width: 110px; text-wrap: wrap-word">ID działu</div>
                <div style="min-width: 120px; text-wrap: wrap-word">ID parametrów</div>
                Nazwa raportu
            </li>
            <li v-for="(data, lp) in filterReports" :key="data.id"
                class="data__list__row"
                @click="rowClicked(data)"
                :class="{ backgroundReportInactive: !data.active }"
            >
                <div style="min-width: 70px">{{ lp+1 }}.</div>
                <div style="min-width: 70px">{{ data.id }}</div>
                <div style="min-width: 110px">{{ data.departmentID }}</div>
                <div style="min-width: 120px">{{ data.parametersID }}</div>
                {{ data.reportName }}
            </li>
          </ul>
          <AdminDataReportsAddReportForm />
          <AdminDataReportsEditReportForm />
        </template>
        <template v-else>
            <div class="data__empty">
                Brak raportów o takiej nazwie
            </div>
        </template>
        <div class="data__list__row add" @click="openAddReportForm">
          Dodaj raport
        </div>
    </div>
</template>

<script>

  export default {

    data() {
      return {
      };
    },

    computed: {
      reports () {
        return this._.orderBy(this.$store.state.reports.reports, 'reportName')
      },
      filterReports () {
        return this.reports.filter(data => (!this.filterString
            || data.id.toString().includes(this.filterString)
            || data.reportName.toLowerCase().includes(this.filterString.toLowerCase())))
      },
      filterString () {
        return this.$store.state.reports.filterString
      }
    },

    methods: {
      rowClicked(data) {
        this.$nuxt.$emit('openAdminEditReportForm', data)
      },
      openAddReportForm () {
        this.$nuxt.$emit('openAdminAddReportForm', false)
      }
    }
  }
</script>

<style lang="scss">
.backgroundReportInactive {
  background: #f1d1d1;
}
.backgroundReportInactive:hover {
  background: #ecb4b4;
}
</style>