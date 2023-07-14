import {createStore} from 'vuex'
import user from './user/user'
import cluster from "./cluster/cluster";
import box from "./box/box";
import environment from "./environment/environment";

const store = createStore({
    modules: [
        user,
        cluster,
        box,
        environment
    ]
})

export default store