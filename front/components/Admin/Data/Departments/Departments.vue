<template>
    <div class="data">
        <template v-if="filterDepartments.length > 0">
          <ul class="data__list">
            <li class="data__list__row first">
                <div style="min-width: 70px">Lp.</div>
                <div style="min-width: 70px; text-wrap: wrap-word">ID</div>
                Nazwa
            </li>
            <li v-for="(data, lp) in filterDepartments" :key="data.departmentID"
                class="data__list__row"
                @click="rowClicked(data)"
            >
                <div style="min-width: 70px">{{ lp+1 }}.</div>
                <div style="min-width: 70px">{{ data.departmentID }}</div>
                {{ data.name }}
            </li>
          </ul>
          <AdminDataDepartmentsAddDepartmentForm />
          <AdminDataDepartmentsEditDepartmentForm />
        </template>
        <template v-else>
            <div class="data__empty">
                Brak działów o takiej nazwie
            </div>
        </template>
        <div class="data__list__row add" @click="openAddDepartmentForm">
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
      departments () {
        return this._.orderBy(this.$store.state.reports.departments, 'name')
      },
      filterDepartments () {
        return this.departments.filter(data => (!this.filterString
            || data.departmentID.toString().includes(this.filterString)
            || data.name.toLowerCase().includes(this.filterString.toLowerCase())))
      },
      filterString () {
        return this.$store.state.reports.filterString
      }
    },

    methods: {
      rowClicked(data) {
        this.$nuxt.$emit('openAdminEditDepartmentForm', data)
      },
      openAddDepartmentForm () {
        this.$nuxt.$emit('openAdminAddDepartmentForm', false)
      }
    }
  }
</script>