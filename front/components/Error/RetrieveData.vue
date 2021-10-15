<template>
  <div class="retrieveData">
    <div class="retrieveData__logo">
      <img class="retrieveData__logo__img" src="/logo-only.png">
      <div class="retrieveData__logo__text">
        Raporty
      </div>
    </div>
    <div class="retrieveData__info">
      <div class="retrieveData__info__header">
        <div class="retrieveData__info__header__error-icon">
          <img src="no-wifi.svg">
        </div>
        <div class="retrieveData__info__header__error-title">Failed to connect with server</div>
      </div>
      <div class="retrieveData__info__message">
        Notify IT about this error.<br>
        <a href="mailto:helpdesk@argenta.com.pl">helpdesk@argenta.com.pl</a>
      </div>
    </div>
    <div class="footer">
      Argenta © {{ currentYear }}
    </div>
  </div>
</template>

<script>
export default {
  data() {
    return {
      currentYear: new Date().getFullYear(),
    }
  },
  methods: {
    reloadPage(){
      window.location.reload()
    }
  }
}
</script>

<style lang="scss" scoped>
.retrieveData {
  display: flex;
  flex-direction: column;
  height: fit-content;
  max-width: 800px;
  width: 100%;
  margin-bottom: 100px;
  &__logo {
    display: flex;
    align-items: center;
    max-width: 800px;
    width: 100%;
    background-color: $base-white;
    padding: 20px 0 20px 36px;
    &__img {
      float: left;
      height: 44px;
    }
    &__text {
      float: left;
      font-size: 20px;
      padding-left: 12px;
    }
  }
  &__info {
    width: 100%;
    padding: 30px 50px 50px;
    margin-top: 20px;
    background-color: $base-white;
    &__header {
      display: flex;
      justify-content: flex-start;
      align-items: center;
      width: 100%;
      font-size: 26px;
      padding-left: 20px;
      &__error-icon {
        display: flex;
        max-width: 190px;
        min-width: 80px;
        width: 100%;
      }
      &__error-title {
        padding-left: 50px;
      }
    }
    &__message {
      display: table;
      width: 100%;
      font-size: 18px;
      padding-top: 20px;
      &__reload {
        color: $argenta;
        text-decoration: underline;
        cursor: pointer;
      }
    }
  }
  .footer {
    display: flex;
    justify-content: flex-end;
    position: absolute;
    max-width: 800px;
    width: 100%;
    bottom: 0;
    padding-right: 20px;
    font-size: 14px;
    background-color: $base-gray-light1;
  }
}

@media (max-width: 500px) {
  .retrieveData {
    &__info {
      padding: 30px 30px 50px;
      &__header {
        &__error-title {
          padding: 10px 0 0 30px;
        }
      }
    }
  }
}
</style>