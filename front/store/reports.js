export const state = () => ({
    reports: [],
    parameters: [],
    departments: [],
    filterDepartmentID: 0,
    filterName: ''
})

export const getters = {
}

export const mutations = {
    setData (state, data) {
        state.reports = data.reports
        state.parameters = data.parameters
        state.departments = data.departments
    },
    
    setReports (state, reports) {
        state.reports = reports
    },

    setParameters (state, parameters) {
        state.parameters = parameters
    },

    setDepartments (state, departments) {
        state.departments = departments
    },

    setFilterDepartmentID (state, filterDepartmentID) {
        state.filterDepartmentID = filterDepartmentID
    },

    setFilterName (state, filterName) {
        state.filterName = filterName
    }
}

export const actions = {

    GetData({ commit }) {
        return new Promise((resolve, reject) => {
            this.$axios.get('/data')
                .then(response => {
                    commit('setData', response.data)
                    resolve(response)
                })
                .catch(error => {
                    reject(error)
                })
        })
    },

    GetReports({ commit }) {
        return this.$axios.get('/reports')
            .then(response => {
                commit('setReports', response.data)
            })
            .catch(error => {
                console.log(error)
            })
    },

    GetParameters({ commit }) {
        return this.$axios.get('/parameters')
            .then(response => {
                commit('setParameters', response.data)
            })
            .catch(error => {
                console.log(error)
            })
    },

    GetDepartments({ commit }) {
        return this.$axios.get('/departments')
            .then(response => {
                commit('setDepartments', response.data)
            })
            .catch(error => {
                console.log(error)
            })
    },

    GenerateReport({dispatch}, data) {
        return this.$axios.post('/report/generate', {
                ...data
            })
            .then(response => {
                if (response.data.status == 1) {
                    $nuxt.$emit('notify', "success")
                    dispatch('DownloadReport', response.data)
                }
                else if (response.data.status == 0) {
                    $nuxt.$emit('notify', "error")
                }
                else if (response.data.status == -1) {
                    $nuxt.$emit('notify', "panic")
                }
            })
            .catch(error => {
                console.log(error)
                $nuxt.$emit('notify', "error")
            })
    },

    DownloadReport({}, data) {
        var url = data.fileName
        this.$axios({
            url,
            baseURL: process.env.HOME_URL_FRONT,
            responseType: 'blob',
        }).then((response) => {
            var fileURL = window.URL.createObjectURL(new Blob([response.data]))
            var fileLink = document.createElement('a')
            fileLink.href = fileURL
            fileLink.setAttribute('download', data.reportName)
            document.body.appendChild(fileLink)

            fileLink.click()
        })
    },

    SetFilterDepartmentID({ commit }, filterDepartmentID) {
        commit('setFilterDepartmentID', filterDepartmentID)
    },

    SetFilterName({ commit }, filterName) {
        commit('setFilterName', filterName)
    }
}

export const modules = {}
