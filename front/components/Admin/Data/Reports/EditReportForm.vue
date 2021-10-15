<template>
    <v-dialog
        max-width="760px"
        v-model="isFormVisible"
        @click:outside="closeDataForm" 
        >
        <template v-if="isFormVisible">
            <v-card>
                <v-card-title>
                    <span class="text-h6">{{ report.reportName }}</span>
                </v-card-title>
                <form class="dataForm">
                    <v-layout justify-center>
                        <div class="dataForm__inputs">
                            <div class="dataForm__inputs__row">
                                <div class="dataForm__inputs__row__key">
                                    ID:
                                </div>
                                <div class="dataForm__inputs__row__values">
                                    <input disabled class="dataForm__inputs__row__values__value short"
                                        v-model="report.id">
                                </div>
                            </div>
                            <div class="dataForm__inputs__row">
                                <div class="dataForm__inputs__row__key">
                                    Dział:
                                </div>
                                <div class="dataForm__inputs__row__values departments">
                                    <div class="dataForm__inputs__row__values__value departmentSelected"
                                        @mousedown="showDepartmentsMenu = !showDepartmentsMenu">
                                        <template v-if="report.id">
                                            {{ departments[reports[report.id].departmentID].id }} - {{ departments[reports[report.id].departmentID].name }}
                                        </template>
                                        <template v-else>
                                            Wybierz
                                        </template>
                                    </div>
                                    <div v-if="showDepartmentsMenu" class="dataForm__inputs__row__values__value menu">
                                        <div class="dataForm__inputs__row__values__value menu__row"
                                            v-for="d in departments"
                                            :key="d.departmentID"
                                            @click="departmentClicked(d.departmentID),
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
                                        :class="{ parameterSelected: report.parametersID != null && report.parametersID.includes(p.parameterID) == true }"
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
                                            @click="setActive(true)">
                                        TAK
                                    </div>
                                    <div class="dataForm__inputs__row__values__value activeButton"
                                            :class="{ inactiveButtonColor: report.active == false }"
                                            @click="setActive(false)">
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
                                    v-model="reportName"
                                    @keyup="setReportName">
                                </div>
                            </div>
                            <div class="dataForm__inputs__row">
                                <div class="dataForm__inputs__row__key">
                                    Nazwa źródła:
                                </div>
                                <div class="dataForm__inputs__row__values">
                                    <input class="dataForm__inputs__row__values__value"
                                    v-model="sourceName"
                                    @keyup="setSourceName">
                                </div>
                            </div>
                        </div>
                    </v-layout>
                    <div class="dataForm__footer">
                        <v-btn @click="remove"
                            depressed
                            :ripple="false"
                        >
                            <i class="fas fa-trash"></i>
                        </v-btn>
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
        </template>
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
        },
        reportName: '',
        sourceName: '',
        bool: [true, false]
    }),

    mounted () {
        this.$nuxt.$on('openAdminEditReportForm', (data) => {
            this.report = data
            this.isFormVisible = true
            this.reportName = data.reportName
            this.sourceName = data.sourceName
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
        departmentClicked (departmentID) {
            var data = {
                reportID: this.report.id,
                departmentID: departmentID
            }
            this.$store.dispatch('reports/SetSelectedDepartmentID', data)
        },

        parameterClicked (parameterID) {
            var data = {
                reportID: this.report.id,
                parameterID: parameterID
            }
            this.$store.dispatch('reports/SetSelectedParameterID', data)
        },

        setActive (active) {
            var data = {
                reportID: this.report.id,
                active: active
            }
            this.$store.dispatch('reports/SetActive', data)
        },

        setReportName () {
            var data = {
                reportID: this.report.id,
                reportName: this.reportName
            }
            this.$store.dispatch('reports/SetReportName', data)
        },

        setSourceName () {
            var data = {
                reportID: this.report.id,
                sourceName: this.sourceName
            }
            this.$store.dispatch('reports/SetSourceName', data)
        },

        remove () {
            this.closeDataForm()
            this.$store.dispatch('admin/DeleteReport', {
                data: {
                    id: this.report.id
                }
            })
        },

        submit () {
            this.$store.dispatch('admin/SetReport', {
                id: this.report.id,
                active: this.report.active,
                parametersID: this.report.parametersID,
                departmentID: this.report.departmentID,
                reportName: this.report.reportName,
                sourceName: this.report.sourceName
            })
            this.closeDataForm()
        },

        closeDataForm () {
            this.isFormVisible = false
        }
    }
}
</script>