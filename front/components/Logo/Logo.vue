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
    width: 100%;
    min-height: 80px;
    max-height: 80px;
    padding: 0 40px;
    background-color: $base-white;
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
@media (max-width: 800px) {
    .logo {
        max-width: 560px;
    }
}
@media (max-width: 480px) {
    .logo {
        min-width: fit-content;
        max-width: 520px;
    }
}
</style>
