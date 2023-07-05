import {createStore} from 'vuex'
import user from './user/user'
import cluster from "./cluster/cluster";
import box from "./box/box";

const store = createStore({
    modules: [
        user,
        cluster,
        box
    ]
})

export default store