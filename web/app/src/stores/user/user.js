import Cookies from 'js-cookie'
export default {
    getters: {
        getUser(state) {
            return state.user
        },
        getUsers(state) {
            return state.usersList
        },
        getUserById: (state) => (id) => {
            return state.usersList.get(parseInt(id))
        },
    },
    mutations: {
        updateUser(state, data){
            state.user.Name = data.user.Name
            state.user.Email = data.user.Email
            state.user.Token = data.user.Token
            state.user.AvatarInitials = data.user.Name.match(/(^\S\S?|\s\S)?/g).map(v=>v.trim()).join("").match(/(^\S|\S$)?/g).join("").toLocaleUpperCase()
            Cookies.set('x-auth-token', data.user.Token, { expires: 7, path: '/' })
        },
        updateUsers(state, data) {
            data.users.forEach((u) => {
                u.AvatarInitials = u.Name.match(/(^\S\S?|\s\S)?/g).map(v=>v.trim()).join("").match(/(^\S|\S$)?/g).join("").toLocaleUpperCase()
                state.usersList.set(u.ID, u)
            })
        }
    },
    state: {
        user: {
            Name: '',
            Email: '',
            Token: '',
            AvatarInitials: 'AA'
        },
        usersList: new Map()
    },
    actions: {
        async getUsers(context) {
            const token = Cookies.get('x-auth-token')
            if (!token || token.length === 0) {
                throw "Token undefined."
            }
            try {
                const response = await fetch('http://localhost:8888/api/v1/users', {
                    method: 'GET',
                    headers: {
                        'x-auth-token': token,
                        'Content-Type': 'application/json;charset=utf-8',
                        'Accept': 'application/json',
                    }
                })
                const data = await response.json()
                context.commit('updateUsers', data)
                return data
            } catch (e) {
                throw e
            }
        },
        async redeemCode(context, payload) {
            try {
                const body = JSON.stringify({
                    code: payload.code,
                    password: payload.password,
                    password_confirmation: payload.password_confirmation
                })
                const response = await fetch('http://localhost:8888/api/v1/user', {
                    method: 'PUT',
                    headers: {
                        'Content-Type': 'application/json;charset=utf-8',
                        'Accept': 'application/json',
                    },
                    body: body
                })
                const data = await response.json()
                context.commit('updateUser', data)
                return data
            } catch(e) {
                throw e
            }
        },
        async login(context, payload) {
            try {
                const body = JSON.stringify({
                    email: payload.email,
                    password: payload.password,
                })
                const response = await fetch('http://localhost:8888/api/v1/user', {
                    method: 'POST',
                    headers: {
                        'Content-Type': 'application/json;charset=utf-8',
                        'Accept': 'application/json',
                    },
                    body: body
                })
                const data = await response.json()
                if (!response.ok) {
                    throw 'Wrong login email-password combination.'
                }
                context.commit('updateUser', data)
                return data
            } catch(e) {
                throw e
            }
        },
        async getUser(context) {
            const token = Cookies.get('x-auth-token')
            if (!token || token.length === 0) {
                throw "Token undefined."
            }
            try {
                const response = await fetch('http://localhost:8888/api/v1/user', {
                    method: 'GET',
                    headers: {
                        'x-auth-token': token,
                        'Content-Type': 'application/json;charset=utf-8',
                        'Accept': 'application/json',
                    }
                })
                const data = await response.json()
                context.commit('updateUser', data)
                return data
            } catch (e) {
                throw e
            }
        },
        async checkIsSetupRequired(context) {
            try {
                const response = await fetch('http://localhost:8888/api/v1/setup_required', {
                    method: 'GET',
                    headers: {
                        'Content-Type': 'application/json;charset=utf-8',
                        'Accept': 'application/json',
                    }
                })
                return await response.json()
            } catch (e) {
                throw e
            }
        },
        async createFirstInvite(context, payload) {
            if (payload.name.length === 0 || payload.email.length === 0) {
                throw "Email and Name fields are required"
            }
            try {
                const body = JSON.stringify({
                    name: payload.name,
                    email: payload.email
                })
                const response = await fetch('http://localhost:8888/api/v1/user/create_first', {
                    method: 'POST',
                    headers: {
                        'Content-Type': 'application/json;charset=utf-8',
                        'Accept': 'application/json',
                    },
                    body: body
                })
                return await response.json()
            } catch(e) {
                throw e
            }
        },
        async createInvite(context, payload) {
            const token = Cookies.get('x-auth-token')
            if (!token || token.length === 0) {
                throw "Token undefined."
            }
            if (payload.name.length === 0 || payload.email.length === 0) {
                throw "Email and Name fields are required"
            }
            try {
                const body = JSON.stringify({
                    name: payload.name,
                    email: payload.email
                })
                const response = await fetch('http://localhost:8888/api/v1/user/create', {
                    method: 'POST',
                    headers: {
                        'x-auth-token': token,
                        'Content-Type': 'application/json;charset=utf-8',
                        'Accept': 'application/json',
                    },
                    body: body
                })
                return await response.json()
            } catch(e) {
                throw e
            }
        }
    }
}