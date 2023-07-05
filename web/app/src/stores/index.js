import {createStore} from 'vuex'
import user from './user/user'
import cluster from "./cluster/cluster";

const store = createStore({
    modules: [
        user,
        cluster
    ]
})

export default store