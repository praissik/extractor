<template>
    <div class="data">
        <template v-if="filterParameters.length > 0">
          <ul class="data__list">
            <li class="data__list__row first">
                <div style="min-width: 70px">Lp.</div>
                <div style="min-width: 70px; text-wrap: wrap-word">ID</div>
                Nazwa
            </li>
            <li v-for="(data, lp) in filterParameters" :key="data.parameterID"
                class="data__list__row"
                @click="rowClicked(data)"
            >
                <div style="min-width: 70px">{{ lp+1 }}.</div>
                <div style="min-width: 70px">{{ data.parameterID }}</div>
                {{ data.name }}
            </li>
          </ul>
          <AdminDataParametersAddParameterForm />
          <AdminDataParametersEditParameterForm />
        </template>
        <template v-else>
            <div class="data__empty">
                Brak parametrów o takiej nazwie
            </div>
        </template>
        <div class="data__list__row add" @click="openAddParameterForm">
          Dodaj raport
        </div>
    </div>
</template>

<script>

  export default {

    data() {
      return {
      };
    },

    computed: {
      parameters () {
        return this._.orderBy(this.$store.state.reports.parameters, 'name')
      },
      filterParameters () {
        return this.parameters.filter(data => (!this.filterString
            || data.parameterID.toString().includes(this.filterString)
            || data.name.toLowerCase().includes(this.filterString.toLowerCase())))
      },
      filterString () {
        return this.$store.state.reports.filterString
      }
    },

    methods: {
      rowClicked (data) {
        this.$nuxt.$emit('openAdminEditParameterForm', data)
      },
      openAddParameterForm () {
        this.$nuxt.$emit('openAdminAddParameterForm', false)
      }
    }
  }
</script>