export const state = () => ({
    filterMenuOptionID: 2,
})

export const getters = {
}

export const mutations = {
    setFilterMenuOptionID (state, filterMenuOptionID) {
        state.filterMenuOptionID = filterMenuOptionID
    }
}

export const actions = {
    SetFilterMenuOption({ commit }, filterMenuOptionID) {
        commit('setFilterMenuOptionID', filterMenuOptionID)
    },   

    AddParameter({dispatch}, data) {
        return new Promise((resolve, reject) => {
            $nuxt.$emit('notify', 'info')
            this.$axios.post('/parameter', {
                ...data
            })
            .then(response => {
                if (response.status == 200) {
                    $nuxt.$emit('notify', "successA")
                    $nuxt.$emit('getParameters')
                }
                resolve(response)
            })
            .catch(error => {
                if (error.response.status == 401) {
                    $nuxt.$emit('notify', "errorAParameterNameUsed")
                }
                else if (error.response.status == 400) {
                    $nuxt.$emit('notify', "errorA")
                }
                else {
                    reject(error)
                }
            })
        })
    }, 

    SetParameter({dispatch}, data) {
        $nuxt.$emit('notify', 'info')
        return this.$axios.put('/parameter', {
                ...data
            })
            .then(response => {
                if (response.status == 200) {
                    $nuxt.$emit('notify', "successA")
                    $nuxt.$emit('getParameters')
                }
            })
            .catch(error => {
                if (error.response.status == 401) {
                    $nuxt.$emit('notify', "errorAParameterNameUsed")
                }
                else if (error.response.status == 400) {
                    $nuxt.$emit('notify', "errorA")
                }
                else {
                    console.log(error)
                    $nuxt.$emit('notify', "panic")
                }
            })
    },

    DeleteParameter({dispatch}, data) {
        $nuxt.$emit('notify', 'info')
        return this.$axios.delete('/parameter', {
                ...data
            })
            .then(response => {
                if (response.status == 200) {
                    $nuxt.$emit('notify', "successA")
                    $nuxt.$emit('getParameters')
                }
            })
            .catch(error => {
                if (error.response.status == 401) {
                    $nuxt.$emit('notify', "errorAParameterUsed")
                }
                else if (error.response.status == 400) {
                    $nuxt.$emit('notify', "errorA")
                }
                else {
                    console.log(error)
                }
            })
    },

    AddDepartment({dispatch}, data) {
        return new Promise((resolve, reject) => {
            $nuxt.$emit('notify', 'info')
            this.$axios.post('/department', {
                ...data
            })
            .then(response => {
                if (response.status == 200) {
                    $nuxt.$emit('notify', "successA")
                    $nuxt.$emit('getDepartments')
                }
                resolve(response)
            })
            .catch(error => {
                if (error.response.status == 401) {
                    $nuxt.$emit('notify', "errorADepartmentNameUsed")
                }
                else if (error.response.status == 400) {
                    $nuxt.$emit('notify', "errorA")
                }
                else {
                    reject(error)
                }
            })
        })
    }, 

    SetDepartment({dispatch}, data) {
        $nuxt.$emit('notify', 'info')
        return this.$axios.put('/department', {
                ...data
            })
            .then(response => {
                if (response.status == 200) {
                    $nuxt.$emit('notify', "successA")
                    $nuxt.$emit('getDepartments')
                }
            })
            .catch(error => {
                if (error.response.status == 401) {
                    $nuxt.$emit('notify', "errorADepartmentNameUsed")
                }
                else if (error.response.status == 400) {
                    $nuxt.$emit('notify', "errorA")
                }
                else {
                    console.log(error)
                    $nuxt.$emit('notify', "panic")
                }
            })
    },

    DeleteDepartment({dispatch}, data) {
        $nuxt.$emit('notify', 'info')
        return this.$axios.delete('/department', {
                ...data
            })
            .then(response => {
                if (response.status == 200) {
                    $nuxt.$emit('notify', "successA")
                    $nuxt.$emit('getDepartments')
                }
            })
            .catch(error => {
                if (error.response.status == 401) {
                    $nuxt.$emit('notify', "errorADepartmentUsed")
                }
                else if (error.response.status == 400) {
                    $nuxt.$emit('notify', "errorA")
                }
                else {
                    console.log(error)
                }
            })
    },

    AddReport({dispatch}, data) {
        return new Promise((resolve, reject) => {
            $nuxt.$emit('notify', 'info')
            this.$axios.post('/report', {
                ...data
            })
            .then(response => {
                if (response.status == 200) {
                    $nuxt.$emit('notify', "successA")
                    $nuxt.$emit('getReports')
                }
                resolve(response)
            })
            .catch(error => {
                if (error.response.status == 401) {
                    $nuxt.$emit('notify', "errorAReportNameUsed")
                }
                if (error.response.status == 402) {
                    $nuxt.$emit('notify', "errorADepartmentNotSelected")
                }
                else if (error.response.status == 400) {
                    $nuxt.$emit('notify', "errorA")
                }
                else {
                    reject(error)
                }
            })
        })
    }, 

    SetReport({dispatch}, data) {
        $nuxt.$emit('notify', 'info')
        return this.$axios.put('/report', {
                ...data
            })
            .then(response => {
                if (response.status == 200) {
                    $nuxt.$emit('notify', "successA")
                    $nuxt.$emit('getReports')
                }
            })
            .catch(error => {
                if (error.response.status == 401) {
                    $nuxt.$emit('notify', "errorAReportUsed")
                }
                else if (error.response.status == 400) {
                    $nuxt.$emit('notify', "errorA")
                }
                else {
                    console.log(error)
                    $nuxt.$emit('notify', "panic")
                }
            })
    },

    DeleteReport({dispatch}, data) {
        $nuxt.$emit('notify', 'info')
        return this.$axios.delete('/report', {
                ...data
            })
            .then(response => {
                if (response.status == 200) {
                    $nuxt.$emit('notify', "successA")
                    $nuxt.$emit('getReports')
                }
            })
            .catch(error => {
                if (error.response.status == 401) {
                    $nuxt.$emit('notify', "errorAReportUsed")
                }
                else if (error.response.status == 400) {
                    $nuxt.$emit('notify', "errorA")
                }
                else {
                    console.log(error)
                    $nuxt.$emit('notify', "panic")
                }
            })
    }
}

export const modules = {}
