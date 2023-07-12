import Cookies from 'js-cookie'

export default {
    getters: {
        getClusters(state) {
            return state.clustersList
        },
        getClusterByName: (state) => (name) => {
            if (!state.clustersList.has(name)) {
                return null
            }
            return state.clustersList.get(name)
        },
    },
    mutations: {
        updateClusters(state, data) {
            data.clusters.forEach((u) => {
                state.clustersList.set(u.Name, u)
            })
        },
        deleteClusterFromList(state, name) {
            state.clustersList.set(name, null)
        }
    },
    state: {
        clustersList: new Map()
    },
    actions: {
        async getClusters(context) {
            const token = Cookies.get('x-auth-token')
            if (!token || token.length === 0) {
                throw "Token undefined."
            }
            try {
                const response = await fetch(window.location.origin+'/api/v1/clusters', {
                    method: 'GET',
                    headers: {
                        'x-auth-token': token,
                        'Content-Type': 'application/json;charset=utf-8',
                        'Accept': 'application/json',
                    }
                })
                const data = await response.json()
                context.commit('updateClusters', data)
                return data
            } catch (e) {
                throw e
            }
        },
        async deleteCluster(context, payload) {
            const token = Cookies.get('x-auth-token')
            if (!token || token.length === 0) {
                throw "Token undefined."
            }
            try {
                const response = await fetch(window.location.origin+'/api/v1/clusters/'+payload.Name, {
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
                context.commit('deleteClusterFromList', payload.Name)
                return await response.json()
            } catch(e) {
                throw e
            }
        },
        async getCluster(context, payload) {
            const token = Cookies.get('x-auth-token')
            if (!token || token.length === 0) {
                throw "Token undefined."
            }
            try {
                const body = JSON.stringify({
                    Name: payload.Name,
                })
                const response = await fetch(window.location.origin+'/api/v1/clusters/'+payload.Name, {
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
                    throw 'Cluster not found or shelf closed.'
                }
                await context.dispatch('getClusters')
                return data
            } catch(e) {
                throw e
            }
        },
        async createCluster(context, payload) {
            const token = Cookies.get('x-auth-token')
            if (!token || token.length === 0) {
                throw "Token undefined."
            }
            try {
                const response = await fetch(window.location.origin+'/api/v1/clusters', {
                    method: 'POST',
                    headers: {
                        'x-auth-token': token,
                        'Content-Type': 'application/json;charset=utf-8',
                        'Accept': 'application/json',
                    },
                    body: JSON.stringify({
                        Name: payload.Name,
                        Kubeconfig: payload.Kubeconfig,
                    })
                })
                const data = await response.json()
                if (!response.ok) {
                    throw data.error
                }
                await context.dispatch('getClusters')
                return data
            } catch (e) {
                throw e
            }
        },
        async updateCluster(context, payload) {
            const token = Cookies.get('x-auth-token')
            if (!token || token.length === 0) {
                throw "Token undefined."
            }
            try {
                const body = JSON.stringify({
                    Name: payload.Name,
                    Kubeconfig: payload.Kubeconfig
                })
                const response = await fetch(window.location.origin+'/api/v1/clusters/'+payload.Name, {
                    method: 'PUT',
                    headers: {
                        'x-auth-token': token,
                        'Content-Type': 'application/json;charset=utf-8',
                        'Accept': 'application/json',
                    },
                    body: body
                })

                await context.dispatch('getClusters')
                const data = await response.json()
                console.log(data)
                return data
            } catch(e) {
                throw e
            }
        },
        async testClusterConnection(context, payload) {
            const token = Cookies.get('x-auth-token')
            if (!token || token.length === 0) {
                throw "Token undefined."
            }
            try {
                const response = await fetch(window.location.origin+'/api/v1/clusters/test', {
                    method: 'POST',
                    headers: {
                        'x-auth-token': token,
                        'Content-Type': 'application/json;charset=utf-8',
                        'Accept': 'application/json',
                    },
                    body: JSON.stringify({
                        Name: payload.Name,
                        Kubeconfig: payload.Kubeconfig,
                    })
                })
                const data = await response.json()
                if (!response.ok) {
                    throw data.error
                }
                return data
            } catch (e) {
                throw e
            }
        }
    }
}