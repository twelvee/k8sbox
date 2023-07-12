import Cookies from 'js-cookie'

export default {
    getters: {
        getEnvironments(state) {
            return state.environmentsList
        },
        getEnvironmentByName: (state) => (name) => {
            if (!state.environmentsList.has(name)) {
                return null
            }
            return state.environmentsList.get(name)
        },
    },
    mutations: {
        updateEnvironments(state, data) {
            if (!data.environments) return []
            data.environments.forEach((u) => {
                state.environmentsList.set(u.Name, u)
            })
        },
        deleteEnvironmentFromList(state, name) {
            state.environmentsList.set(name, null)
        }
    },
    state: {
        environmentsList: new Map()
    },
    actions: {
        async getEnvironments(context) {
            const token = Cookies.get('x-auth-token')
            if (!token || token.length === 0) {
                throw "Token undefined."
            }
            try {
                const response = await fetch(window.location.origin+'/api/v1/environments', {
                    method: 'GET',
                    headers: {
                        'x-auth-token': token,
                        'Content-Type': 'application/json;charset=utf-8',
                        'Accept': 'application/json',
                    }
                })
                const data = await response.json()
                context.commit('updateEnvironments', data)
                return data
            } catch (e) {
                throw e
            }
        },
        async deleteEnvironment(context, payload) {
            const token = Cookies.get('x-auth-token')
            if (!token || token.length === 0) {
                throw "Token undefined."
            }
            try {
                const response = await fetch(window.location.origin+'/api/v1/environments/'+payload.Name, {
                    method: 'DELETE',
                    headers: {
                        'x-auth-token': token,
                        'Content-Type': 'application/json;charset=utf-8',
                        'Accept': 'application/json',
                    },
                    body: JSON.stringify({
                        Name: payload.Name
                    })
                })
                context.commit('deleteEnvironmentFromList', payload.Name)
                return await response.json()
            } catch(e) {
                throw e
            }
        },
        async getEnvironment(context, payload) {
            const token = Cookies.get('x-auth-token')
            if (!token || token.length === 0) {
                throw "Token undefined."
            }
            try {
                const body = JSON.stringify({
                    Name: payload.Name,
                })
                const response = await fetch(window.location.origin+'/api/v1/environments/'+payload.Name, {
                    method: 'GET',
                    headers: {
                        'x-auth-token': token,
                        'Content-Type': 'application/json;charset=utf-8',
                        'Accept': 'application/json',
                    },
                    body: body
                })
                const data = await response.json()
                if (!response.ok) {
                    throw 'Environment not found or shelf closed.'
                }
                await context.dispatch('getEnvironments')
                return data
            } catch(e) {
                throw e
            }
        },
        async createEnvironment(context, payload) {
            const token = Cookies.get('x-auth-token')
            if (!token || token.length === 0) {
                throw "Token undefined."
            }
            try {
                const response = await fetch(window.location.origin+'/api/v1/environments', {
                    method: 'POST',
                    headers: {
                        'x-auth-token': token,
                        'Content-Type': 'application/json;charset=utf-8',
                        'Accept': 'application/json',
                    },
                    body: JSON.stringify(payload)
                })
                const data = await response.json()
                if (!response.ok) {
                    throw data.error
                }
                await context.dispatch('getEnvironments')
                return data
            } catch (e) {
                throw e
            }
        }
    }
}