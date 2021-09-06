<template>
    <div class="logo">
        <img class="logo__img" src="/logo-only.png">
        <div class="logo__text">
            Raporty
        </div>
    </div>
</template>

<style lang="scss">
.logo {
    display: flex;
    justify-content: flex-start;
    align-items: center;
    padding-left: 36px;
    max-width: 280px;
    width: 100%;
    &__img {
        height: 44px;
    }
    &__text {
        font-size: 20px;
        padding-left: 12px;
    }
}
</style>
