import Cookies from 'js-cookie'

export default {
    getters: {
        getBoxes(state) {
            return state.boxesList
        },
        getBoxByName: (state) => (name) => {
            if (!state.boxesList.has(name)) {
                return null
            }
            return state.boxesList.get(name)
        },
    },
    mutations: {
        updateBoxes(state, data) {
            data.boxes.forEach((u) => {
                state.boxesList.set(u.Name, u)
            })
        },
        deleteBoxFromList(state, name) {
            state.boxesList.set(name, null)
        }
    },
    state: {
        boxesList: new Map()
    },
    actions: {
        async getBoxes(context) {
            const token = Cookies.get('x-auth-token')
            if (!token || token.length === 0) {
                throw "Token undefined."
            }
            try {
                const response = await fetch('http://localhost:8888/api/v1/boxes', {
                    method: 'GET',
                    headers: {
                        'x-auth-token': token,
                        'Content-Type': 'application/json;charset=utf-8',
                        'Accept': 'application/json',
                    }
                })
                const data = await response.json()
                context.commit('updateBoxes', data)
                return data
            } catch (e) {
                throw e
            }
        },
        async deleteBox(context, payload) {
            const token = Cookies.get('x-auth-token')
            if (!token || token.length === 0) {
                throw "Token undefined."
            }
            try {
                const response = await fetch('http://localhost:8888/api/v1/boxes/'+payload.Name, {
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
                context.commit('deleteBoxFromList', payload.Name)
                return await response.json()
            } catch(e) {
                throw e
            }
        },
        async getBox(context, payload) {
            const token = Cookies.get('x-auth-token')
            if (!token || token.length === 0) {
                throw "Token undefined."
            }
            try {
                const body = JSON.stringify({
                    Name: payload.Name,
                })
                const response = await fetch('http://localhost:8888/api/v1/boxes/'+payload.Name, {
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
                    throw 'Box not found or shelf closed.'
                }
                await context.dispatch('getBoxes')
                return data
            } catch(e) {
                throw e
            }
        },
        async createBox(context, payload) {
            const token = Cookies.get('x-auth-token')
            if (!token || token.length === 0) {
                throw "Token undefined."
            }
            try {
                const response = await fetch('http://localhost:8888/api/v1/boxes', {
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
                await context.dispatch('getBoxes')
                return data
            } catch (e) {
                throw e
            }
        },
        async updateBox(context, payload) {
            const token = Cookies.get('x-auth-token')
            if (!token || token.length === 0) {
                throw "Token undefined."
            }
            try {
                const response = await fetch('http://localhost:8888/api/v1/boxes/'+payload.Name, {
                    method: 'PUT',
                    headers: {
                        'x-auth-token': token,
                        'Content-Type': 'application/json;charset=utf-8',
                        'Accept': 'application/json',
                    },
                    body: JSON.stringify(payload)
                })

                await context.dispatch('getBoxes')
                const data = await response.json()
                console.log(data)
                return data
            } catch(e) {
                throw e
            }
        }
    }
}