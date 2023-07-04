import Cookies from 'js-cookie'
export default {
    getters: {
        getUser(state) {
            return state.user
        }
    },
    mutations: {
        updateUser(state, data){
            console.log(data)
            state.user.name = data.user.Name
            state.user.email = data.user.Email
            state.user.token = data.user.Token
            state.user.avatarInitials = data.user.Name.match(/(^\S\S?|\s\S)?/g).map(v=>v.trim()).join("").match(/(^\S|\S$)?/g).join("").toLocaleUpperCase()
            Cookies.set('x-auth-token', data.user.Token, { expires: 7, path: '/' })
        }
    },
    state: {
        user: {
            name: '',
            email: '',
            token: '',
            avatarInitials: 'AA'
        }
    },
    actions: {
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
                console.log(data)
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
        async createInvite(context, payload) {
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
                        'x-auth-token': '64840d2d722d03b5c6c7560d4afb1c116c76ceab72edbfaf8819879b033f48b2',
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