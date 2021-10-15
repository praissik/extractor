<template>
    <v-dialog
        max-width="500px"
        v-model="isFormVisible"
        @click:outside="closeParametersForm()"
        >
        <template v-if="report.id > 0">
            <v-card>
                <v-card-title>
                    <span class="text-h6">{{ report.reportName }}</span>
                </v-card-title>
                <form class="parameters-form">
                    <v-layout justify-center>
                        <div class="parameters-form__inputs">
                            <v-text-field
                                v-if="report.parametersID.includes(parameters['1'].parameterID)"
                                v-model="akronim"
                                color="#D7172F"
                                :label="parameters['1'].name"
                                :error-messages="akronimErrors"
                                required
                                @input="$v.akronim.$touch()"
                                @blur="$v.akronim.$touch()"
                            ></v-text-field>
                            <v-text-field
                                v-if="report.parametersID.includes(parameters['2'].parameterID)"
                                v-model="kod"
                                color="#D7172F"
                                :label="parameters['2'].name"
                                :error-messages="kodErrors"
                                required
                                @input="$v.kod.$touch()"
                                @blur="$v.kod.$touch()"
                            ></v-text-field>
                            <v-text-field
                                v-if="report.parametersID.includes(parameters['3'].parameterID)"
                                v-model="symbol"
                                color="#D7172F"
                                :label="parameters['3'].name"
                                :error-messages="symbolErrors"
                                required
                                @input="$v.symbol.$touch()"
                                @blur="$v.symbol.$touch()"
                            ></v-text-field>
                            <v-text-field
                                v-if="report.parametersID.includes(parameters['4'].parameterID)"
                                v-model="rokOS"
                                color="#D7172F"
                                :label="parameters['4'].name"
                                :error-messages="rokOSErrors"
                                required
                                @input="$v.rokOS.$touch()"
                                @blur="$v.rokOS.$touch()"
                            ></v-text-field>
                        </div>
                    </v-layout>
                    <div class="parameters-form__footer">
                        <v-btn @click="submit"
                            depressed
                            :ripple="false"
                        >
                            <i class="fas fa-file-download"></i>
                        </v-btn>
                        <v-btn @click="closeParametersForm"
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
import { required, between } from 'vuelidate/lib/validators'

export default {

    validations: {
        akronim: { required },
        kod: { required },
        symbol: { required },
        rokOS: { required, between: between(1991, new Date().getFullYear()) },
    },

    data: () => ({
        isFormVisible: false,
        report: [],
        akronim: '',
        kod: '',
        symbol: '',
        rokOS: ''
    }),

    created () {
        this.$nuxt.$on('openParametersForm', (report) => {
            console.log('openParametersForm')
            this.report = report
            this.isFormVisible = true
        })
    },

    computed: {
        akronimErrors () {
            const errors = []
            if (!this.$v.akronim.$dirty) return errors
            !this.$v.akronim.required && errors.push('Pole wymagane')
            return errors
        },
        kodErrors () {
            const errors = []
            if (!this.$v.kod.$dirty) return errors
            !this.$v.kod.required && errors.push('Pole wymagane')
            return errors
        },
        symbolErrors () {
            const errors = []
            if (!this.$v.symbol.$dirty) return errors
            !this.$v.symbol.required && errors.push('Pole wymagane')
            return errors
        },
        rokOSErrors () {
            const errors = []
            if (!this.$v.rokOS.$dirty) return errors
            !this.$v.rokOS.required && errors.push('Pole wymagane')
            !this.$v.rokOS.between && errors.push('Wymagana data z przedziału 1991 - ' + new Date().getFullYear())
            return errors
        },

        parameters () {
            return this.$store.state.reports.parameters
        }
    },

    methods: {
        submit () {
            console.log('submit 1')
            this.$v.$touch()

            if (this.checkValidation()) {
                this.$store.dispatch('reports/GenerateReport', {
                    id: this.report.id,
                    akronim: this.akronim,
                    kod: this.kod,
                    symbol: this.symbol,
                    rokOS: this.rokOS
                })
            }
        },

        closeParametersForm () {
            this.isFormVisible = false
            this.clear()
        },

        clear () {
            this.$v.$reset()
            this.akronim = ''
            this.kod = ''
            this.symbol = ''
            this.rokOS = ''
        },

        checkValidation () {
            for (let id of this.report.parametersID) {
                switch (id) {
                    case 1:
                        if (this.$v.akronim.$invalid) {
                            return false
                        }
                        break;
                    case 2:
                        if (this.$v.kod.$invalid) {
                            return false
                        }
                        break;
                    case 3:
                        if (this.$v.symbol.$invalid) {
                            return false
                        }
                        break;
                    case 4:
                        if (this.$v.rokOS.$invalid) {
                            return false
                        }
                        break;
                }
            }

            return true
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
    .parameters-form {
        background: $base-white;
        &__inputs {
            width: 50%;
            .v-input {
                padding: 24px 0 0;
            }
        }
        &__footer {
            text-align: right;
            padding: 15px 40px 15px 0;
            background-color: $base-white;
            .v-btn {
                background: none;
                .fas {
                    font-size: 30px;
                }
                .fa-times-circle {
                    color: $base-black;
                }
                .fa-file-download {
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