<template>
    <v-dialog
        max-width="500px"
        v-model="isFormVisible"
        @click:outside="closeDataForm()"
        >
            <v-card>
                <v-card-title>
                    <span class="text-h6">{{ name }}</span>
                </v-card-title>
                <form class="dataForm">
                    <v-layout justify-center>
                        <div class="dataForm__inputs">
                            <v-row>
                                <v-col cols="3">
                                    <v-subheader>ID:</v-subheader>
                                </v-col>
                                <v-col cols="4">
                                    <v-text-field
                                        v-model="id"
                                        color="#D7172F"
                                        disabled
                                    ></v-text-field>
                                </v-col>
                            </v-row>
                            <v-row>
                                <v-col cols="3">
                                    <v-subheader>Nazwa:</v-subheader>
                                </v-col>
                                <v-col cols="9">
                                    <v-text-field
                                        v-model="name"
                                        color="#D7172F"
                                    ></v-text-field>
                                </v-col>
                            </v-row>
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
    </v-dialog>
</template>

<script>

export default {

    data: () => ({
        isFormVisible: false,
        id: 0,
        name: '',
    }),

    created () {
        this.$nuxt.$on('openAdminEditDepartmentForm', (data) => {
            this.id = data.departmentID
            this.name = data.name
            this.isFormVisible = true
        })
    },

    methods: {
        remove () {
            this.$store.dispatch('admin/DeleteDepartment', {
                data: {
                    departmentID: this.id
                }
            })
            this.closeDataForm()
        },

        submit () {
            this.$store.dispatch('admin/SetDepartment', {
                departmentID: this.id,
                name: this.name
            })
            this.closeDataForm()
        },

        closeDataForm () {
            this.isFormVisible = false
            this.clear()
        },

        clear () {
            this.id = ''
            this.name = ''
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
            .v-subheader {
                display: flex;
                justify-content: flex-end;
                padding-top: 38px;
                padding-right: 0;
                font-size: 16px;
            }
            .v-input {
                padding: 24px 0 0;
            }
            .row {
                margin: -10px -10px -20px -10px;
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