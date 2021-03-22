
<template>
  <div class="notifications">
    <notifications group="notify" classes="notify-style" />
  </div>
</template>

<script>

  //import { mapState } from 'vuex'

  export default {

    data() {
      return {

      };
    },

    methods: {
      notifyInfo() {
        this.$notify({
          group: "notify",
          title: 'Zgłoszenie przyjęte',
          type: 'info'
        });
      },

      notifySuccess() {
        this.$notify({
          group: "notify",
          title: 'Raport gotowy do pobrania.',
          type: 'success'
        });
      },

      notifyError() {
        this.$notify({
          group: "notify",
          title: 'Dla podanych danych nie ma wyników.',
          type: 'error'
        });
      },

      notifyPanicError() {
        this.$notify({
          group: "notify",
          title: 'Wystąpił błąd. Sprawdź połączenie z siecią firmową.',
          type: 'error'
        });
      }
    },

    mounted() {
      this.$eventHub.on('notify', (data) => {
        console.log(data)
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
          this.notifyPanicError()
        }
      })
    }
  }
</script>

<style lang="scss">

.vue-notification-group {
    width: 22vw !important;
    min-width: 200px !important;
}

.notify-style {
  margin: 0.65vw 2vw 0 0;
  padding: 0.65vw 1.5vw;
  border-radius: 7px;
  float: right;
  font-family: 'Segoe UI', Tahoma, Geneva, Verdana, sans-serif;

  &.info {
    background: rgba(60, 170, 190, 0.356);
    color: rgb(0, 93, 109);
    border: 2px solid rgb(0, 179, 211);
    width: 14vw;
    min-width: 170px;
  }

  &.success {
    background: rgba(60, 190, 67, 0.356);
    color: rgb(0, 82, 4);
    border: 2px solid rgb(0, 204, 10);
    width: 16vw;
    min-width: 170px;
  }

  &.error {
    background: rgba(190, 60, 60, 0.356);
    color: rgb(85, 0, 0);
    border: 2px solid rgb(221, 0, 0);
    width: 18vw;
    min-width: 170px;
  }
}

</style>