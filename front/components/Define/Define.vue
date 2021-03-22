<template>

  <div class="main-view-define">
    <div class="header-defined">Skonfiguruj własny</div>
    <ul class="reports">
      <li class="report" v-for="report in reports" :key="report.id" v-on:click="generateReport(report.id)">
        <div class="report-name">
          <p>{{ report.name }}</p>
        </div>
      </li>
    </ul>
  </div>
</template>

<script>
  export default {

    data() {
      return {
        reports: []
      };
    }
  }

</script>

<style lang="scss">

.main-view-define {
  display: none;
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
    height: 6vw;
    font-size: 2vh;
    margin: 1vw 0.5vw 0;
    background: #e3e3e3;
    border: 0.2vw solid #cc3034;
    border-radius: 0.5vw;
    .report-name{
      margin: 0.5vw 0.8vw;
    }
  }
}

</style>