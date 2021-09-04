<template>
    <v-text-field
        v-model="data"
        :rules="[rules.required]"
        :label="parameter.name"
    ></v-text-field>
</template>

<script>

    export default {
        props: {
            parameter: Object,
            firstParameter: Boolean
        },

        data: () => ({
            data: '',
            rules: {
                required: value => !!value || 'Pole wymagane.',
            },
        }),

        created () {
            console.log('textfield created')
            this.$nuxt.$on('submitParametersForm', () => {
                console.log('submit 2')
                console.log(this.parameter.id, this.data)
                this.data = ''
            })
            this.$nuxt.$on('closeParametersForm', () => {
                console.log('close 2')
                this.data = ''
            })
            // if (this.firstParameter) {
            //     console.log("pierwszy parametr")
            //     console.log(this.firstParameter)
            //     this.$nextTick(() => this.$refs.focus.focus())
            // }
        },

        destroyed () {
            console.log('destroyed')
        }
    }
</script>