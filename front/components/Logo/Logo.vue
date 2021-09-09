<template>
    <div class="logo">
        <div class="logo__flex">
            <img class="logo__flex__img" src="/logo-only.png">
            <div class="logo__flex__text">
                Raporty
            </div>
        </div>
    </div>
</template>

<style lang="scss">
.logo {
    display: flex;
    max-width: 230px;
    width: 100%;
    &__flex {
        display: flex;
        align-items: center;
        width: 100%;
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
}
@media (max-width: 800px) {
    .logo {
        max-width: fit-content;
    }
}
@media (max-width:  680px) {
    .logo {
        max-width: 680px;
        width: 100%;
        &__flex {
            justify-content: center;
        }
    }
}
</style>
