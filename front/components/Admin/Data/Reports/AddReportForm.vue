<template>
    <v-dialog
        max-width="760px"
        v-model="isFormVisible"
        @click:outside="closeDataForm"
        >
        <v-card>
            <v-card-title>
                <span class="text-h6">Nowy raport</span>
            </v-card-title>
            <form class="dataForm">
                <v-layout justify-center>
                    <div class="dataForm__inputs">
                        <div class="dataForm__inputs__row">
                            <div class="dataForm__inputs__row__key">
                                Dział:
                            </div>
                            <div class="dataForm__inputs__row__values departments">
                                <div class="dataForm__inputs__row__values__value departmentSelected"
                                    @mousedown="showDepartmentsMenu = !showDepartmentsMenu">
                                    <template v-if="report.departmentID > 0">
                                        {{ departments[report.departmentID].id }} - {{ departments[report.departmentID].name }}
                                    </template>
                                    <template v-else>
                                        Wybierz
                                    </template>
                                </div>
                                <div v-if="showDepartmentsMenu" class="dataForm__inputs__row__values__value menu">
                                    <div class="dataForm__inputs__row__values__value menu__row"
                                        v-for="d in departments"
                                        :key="d.departmentID"
                                        @click="report.departmentID = d.departmentID,
                                                showDepartmentsMenu = false"
                                        >
                                        {{ d.departmentID }} - {{ d.name }}
                                    </div>
                                </div>
                            </div>
                        </div>
                        <div class="dataForm__inputs__row">
                            <div class="dataForm__inputs__row__key">
                                Parametry:
                            </div>
                            <div class="dataForm__inputs__row__values parameters">
                                <div class="dataForm__inputs__row__values__value checkbox"
                                    :class="{ parameterSelected: report.parametersID != null && Object.values(report.parametersID).includes(p.parameterID) == true }"
                                    v-for="p in parameters"
                                    :key="p.parameterID"
                                    @click="parameterClicked(p.parameterID)"
                                    >
                                    {{ p.parameterID }} - {{ p.name }}
                                </div>
                            </div>
                        </div>
                        <div class="dataForm__inputs__row">
                            <div class="dataForm__inputs__row__key">
                                Aktywny:
                            </div>
                            <div class="dataForm__inputs__row__values buttons">
                                <div class="dataForm__inputs__row__values__value activeButton"
                                        :class="{ activeButtonColor: report.active == true }"
                                        @click="report.active = true">
                                    TAK
                                </div>
                                <div class="dataForm__inputs__row__values__value activeButton"
                                        :class="{ inactiveButtonColor: report.active == false }"
                                        @click="report.active = false">
                                    NIE
                                </div>
                            </div>
                        </div>
                        <div class="dataForm__inputs__row">
                            <div class="dataForm__inputs__row__key">
                                Nazwa raportu:
                            </div>
                            <div class="dataForm__inputs__row__values">
                                <input class="dataForm__inputs__row__values__value"
                                v-model="report.reportName">
                            </div>
                        </div>
                        <div class="dataForm__inputs__row">
                            <div class="dataForm__inputs__row__key">
                                Nazwa źródła:
                            </div>
                            <div class="dataForm__inputs__row__values">
                                <input class="dataForm__inputs__row__values__value"
                                v-model="report.sourceName">
                            </div>
                        </div>
                    </div>
                </v-layout>
                <div class="dataForm__footer">
                    <v-btn @click="submit"
                        depressed
                        :ripple="false"
                    >
                        <i class="fas fa-save"></i>
                    </v-btn>
                    <v-btn @click="closeDataForm"
                        depressed
                        :ripple="false"
                    >
                        <i class="fas fa-times-circle"></i>
                    </v-btn>
                </div>
            </form>
        </v-card>
    </v-dialog>
</template>

<script>

export default {

    data: () => ({
        isFormVisible: false,
        showDepartmentsMenu: false,
        report: {
            id: 0,
            active: 0,
            departmentID: 0,
            parametersID: [],
            reportName: '',
            sourceName: ''
        }
    }),

    mounted () {
        this.$nuxt.$on('openAdminAddReportForm', () => {
            this.report.id = 0
            this.report.active = true
            this.report.departmentID = 0
            this.report.parametersID = null
            this.report.reportName = ''
            this.report.sourceName = ''
            this.isFormVisible = true
        })
    },

    computed: {
        parameters () {
            return this.$store.state.reports.parameters
        },
        departments () {
            return this.$store.state.reports.departments
        },
        reports () {
            return this.$store.state.reports.reports
        }
    },

    methods: {
        parameterClicked (parameterID) {
            if (this.report.parametersID == null) {
                this.report.parametersID = [parameterID]
            } else if (!Object.values(this.report.parametersID).includes(parameterID)) {
                this.report.parametersID.push(parameterID)
            } else {
                this.report.parametersID.splice(this.report.parametersID.indexOf(parameterID), 1)
            }
        },

        submit () {
            this.$store.dispatch('admin/AddReport', this.report)
            this.closeDataForm()
        },

        closeDataForm () {
            this.isFormVisible = false
        }
    }
}
</script>

<style lang="scss">
.v-dialog {
    .v-card {
        .v-card {
            &__title {
                word-break: keep-all;
                padding: 14px 100px 14px 40px;
                background: $argenta-transparent-light;
                border-bottom: $argenta 2px solid;
            }
        }
    }
    .dataForm {
        background: $base-white;
        &__inputs {
            width: 80%;
            &__row {
                display: flex;
                flex-direction: row;
                padding: 28px 0 0 0;
                color: $font-color-primary;
                &__key {
                    display: flex;
                    justify-content: flex-end;
                    max-width: 160px;
                    width: 100%;
                    min-width: 120px;
                }
                &__values {
                    max-width: 400px;
                    width: 100%;
                    min-width: 200px;
                    padding: 0 0 0 16px;
                    &.buttons, &.parameters {
                        display: flex;
                        flex-direction: row;
                        flex-wrap: wrap;
                    }
                    &__value {
                        width: 100%;
                        &.menu {
                            position: absolute;
                            max-width: 220px;
                            background: $base-gray-light1;
                            box-shadow: 3px 3px 8px #5f5f5f;
                            &__row {
                                padding: 6px 15px 6px 16px;
                                &:hover {
                                    background: $base-gray-dark1;
                                }
                            }
                        }
                        &.checkbox {
                            width: fit-content;
                            padding: 0 25px 0 0;
                        }
                        &.checkbox.parameterSelected {
                            color: #d11205;
                        }
                        &.activeButton {
                            width: fit-content;
                            padding: 0 14px;
                            &.activeButtonColor {
                                background-color: #4be42c3a;
                            }
                            &.inactiveButtonColor {
                                background-color: #e4392c3a;
                            }
                        }
                    }
                    input {
                        color: $font-color-primary;
                        border: 0;
                        border-bottom: 2px solid $base-gray-dark2;
                        &.short {
                            width: 150px;
                        }
                    }
                    input:focus {
                        outline: none;
                    }
                }
            }
        }
        &__footer {
            text-align: right;
            padding: 20px 50px 20px 0;
            background-color: $base-white;
            .v-btn {
                background: none;
                .fas {
                    font-size: 30px;
                }
                .fa-trash {
                    color: #d11205;
                    font-size: 28px;
                }
                .fa-times-circle {
                    color: $base-black;
                }
                .fa-save {
                    color: #63aa20;
                }
                &:hover {
                    transform: scale(1.25);
                }
            }
            .v-btn::before{
                background: none;
            }
        }
    }
}
</style>