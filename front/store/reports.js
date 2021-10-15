export const state = () => ({
    reports: [],
    parameters: [],
    departments: [],
    filterDepartmentID: 0,
    filterString: '',
    reportName: '',
    sourceName: ''
})

export const getters = {
    getReports: state => {
        return state.reports
    }
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

    resetReport (state, data) {
        Object(state.reports)[data.reportID] = data
    },

    setParameters (state, parameters) {
        state.parameters = parameters
    },

    setSelectedDepartmentID (state, data) {
        state.reports[data.reportID].departmentID = state.departments[data.departmentID].departmentID
    },

    setSelectedParameterID (state, data) {
        if (Object(state.reports)[data.reportID].parametersID == null) {
            Object(state.reports)[data.reportID].parametersID = [data.parameterID]
        } else if (!Object(state.reports)[data.reportID].parametersID.includes(data.parameterID)) {
            Object(state.reports)[data.reportID].parametersID.push(data.parameterID)
        } else {
            Object(state.reports)[data.reportID].parametersID.splice(Object(state.reports)[data.reportID].parametersID.indexOf(data.parameterID), 1)
        }
    },

    setActive (state, data) {
        state.reports[data.reportID].active = data.active
    },

    setReportName (state, data) {
        state.reports[data.reportID].reportName = data.reportName
    },

    setSourceName (state, data) {
        state.reports[data.reportID].sourceName = data.sourceName
    },

    setDepartments (state, departments) {
        state.departments = departments
    },

    setFilterDepartmentID (state, filterDepartmentID) {
        state.filterDepartmentID = filterDepartmentID
    },
    
    setFilterString (state, filterString) {
        state.filterString = filterString
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

    ResetReport({ commit }, data) {
        commit('resetReport', data)
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

    SetSelectedDepartmentID({ commit }, data) {
        commit('setSelectedDepartmentID', data)
    },

    SetSelectedParameterID({ commit }, data) {
        commit('setSelectedParameterID', data)
    },

    SetActive({ commit }, data) {
        commit('setActive', data)
    },

    SetReportName({ commit }, data) {
        commit('setReportName', data)
    },

    SetSourceName({ commit }, data) {
        commit('setSourceName', data)
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
        $nuxt.$emit('notify', 'info')
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

    SetFilterString({ commit }, filterString) {
        commit('setFilterString', filterString)
    }
}

export const modules = {}
