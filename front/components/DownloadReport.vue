<template>

</template>

<script>
  //import {mapState} from 'vuex';

  export default {

    data() {
      return {
        status: true
      };
    },
/*
    computed: {
      ...mapState({
        home_url: 'home_url',
      })
    },
*/
    methods: {
      downloadReport(data) {
        console.log(data)
        var url = data.fileName
        console.log("Downloading report")
        this.$axios({
          url,
          baseURL: "http://192.168.2.234",
          method: 'GET',
          responseType: 'blob',
        }).then((response) => {
          console.log("Pobierz ten plik: " + data.reportName)

            var fileURL = window.URL.createObjectURL(new Blob([response.data]));
            var fileLink = document.createElement('a');

            fileLink.href = fileURL;
            fileLink.setAttribute('download', data.reportName);
            document.body.appendChild(fileLink);

            fileLink.click();
        });
      }
    },
    
    mounted() {
      this.$eventHub.$on('generated_successfully', (data) => {
        this.downloadReport(data)
      })
    }
  }
</script>

<style lang="scss">

</style>
