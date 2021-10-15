
<template>
</template>

<script>
  export default {
    data() {
      return {
        title: '',
        type: ''
      };
    },

    methods: {
      notifySuccess() {
        this.$notify({
          title: 'Raport gotowy do pobrania',
          type: 'success'
        });
      },

      notifyWarning() {
        this.$notify({
          title: 'Warning',
          message: 'This is a warning message',
          type: 'warning'
        });
      },

      notifyInfo() {
        this.$notify.info({
          title: 'Zlecenie przyjęte'
        });
      },

      notifyError() {
        this.$notify.error({
          title: 'Dla podanych danych nie ma wyników'
        });
      },

      notifyPanic() {
        this.$notify.error({
          title: 'Błąd wewnętrzny serwera - zgłoś na helpdesk'
        });
      },

      notifySuccessA() {
        this.$notify({
          title: 'Dane zaktualizowane',
          type: 'success'
        });
      },

      notifyErrorA() {
        this.$notify.error({
          title: 'Błąd - spróbuj ponownie'
        });
      },

      notifyErrorAParameterUsed() {
        this.$notify.error({
          title: 'Błąd - parametr w użyciu'
        });
      },

      notifyErrorAParameterNameUsed() {
        this.$notify.error({
          title: 'Błąd - nazwa parametru jest już zajęta'
        });
      },

      notifyErrorADepartmentUsed() {
        this.$notify.error({
          title: 'Błąd - dział w użyciu'
        });
      },

      notifyErrorADepartmentNameUsed() {
        this.$notify.error({
          title: 'Błąd - nazwa działu jest już zajęta'
        });
      },

      notifyErrorAReportNameUsed() {
        this.$notify.error({
          title: 'Błąd - nazwa raportu jest już zajęta'
        });
      },

      notifyErrorADepartmentNotSelected() {
        this.$notify.error({
          title: 'Błąd - nie wybrano działu'
        });
      },
    },

    mounted() {
      this.$nuxt.$on('notify', (data) => {
        if (data == "info") {
          this.notifyInfo()
        }
        else if (data == "success") {
          this.notifySuccess()
        }
        else if (data == "error") {
          this.notifyError()
        }
        else if (data == "panic") {
          this.notifyPanic()
        }
        else if (data == "successA") {
          this.notifySuccessA()
        }
        else if (data == "errorA") {
          this.notifyErrorA()
        }
        else if (data == "errorAParameterUsed") {
          this.notifyErrorAParameterUsed()
        }
        else if (data == "errorAParameterNameUsed") {
          this.notifyErrorAParameterNameUsed()
        }
        else if (data == "errorADepartmentUsed") {
          this.notifyErrorADepartmentUsed()
        }
        else if (data == "errorADepartmentNameUsed") {
          this.notifyErrorADepartmentNameUsed()
        }
        else if (data == "errorAReportNameUsed") {
          this.notifyErrorADepartmentNameUsed()
        }
        else if (data == "errorADepartmentNotSelected") {
          this.notifyErrorADepartmentNotSelected()
        }
      })
    }
  }
</script>

<style lang="scss">
.el-notification {
  min-width: 200px;
  width: 300px;
  max-width: 86%;
  &__group {
    display: flex;
    align-items: center;
    padding-right: 20px;
  }
}
</style>