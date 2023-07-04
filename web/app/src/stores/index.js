import {createStore} from 'vuex'
import user from './user/user'

const store = createStore({
    modules: [
        user
    ]
})

export default store