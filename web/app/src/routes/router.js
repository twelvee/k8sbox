import {createRouter, createWebHistory} from 'vue-router'
import Login from "../pages/authorization/Login.vue";
import Dashboard from "../pages/Dashboard.vue";
import Users from "../pages/users/Users.vue";
import Clusters from "../pages/Clusters.vue";
import Boxes from "../pages/Boxes.vue";
import Environments from "../pages/Environments.vue";
import Invite from "../pages/users/Invite.vue";
import store from "../stores";
import RedeemCode from "../pages/authorization/RedeemCode.vue";
import Logout from "../pages/authorization/Logout.vue";

const router = createRouter({
    routes: [
        {
            path: '/',
            name: 'Dashboard',
            component: Dashboard
        },
        {
            path: '/login',
            name: 'Login',
            component: Login
        },
        {
            path: '/logout',
            name: 'Logout',
            component: Logout
        },
        {
            path: '/login/redeem',
            name: 'RedeemCode',
            component: RedeemCode
        },
        {
            path: '/users',
            name: 'Users',
            component: Users
        },
        {
            path: '/users/invite',
            name: 'Invite',
            component: Invite
        },
        {
            path: '/clusters',
            name: 'Clusters',
            component: Clusters
        },
        {
            path: '/boxes',
            name: 'Boxes',
            component: Boxes
        },
        {
            path: '/environments',
            name: 'Environments',
            component: Environments
        }
    ],
    history: createWebHistory()
})

function authMiddleware() {
    const user = store.getters.getUser

    router.beforeEach(async (to, from) => {
        if (
            (!user.token || user.token.length === 0) &&
            (to.name !== 'Login' && to.name !== 'RedeemCode')
        ) {
            return {name: 'Login'}
        }

        if (
            user.token &&
            (to.name === 'Login' || to.name === 'RedeemCode')
        ) {
            return {name: 'Dashboard'}
        }

        if (!user.token && to.page === 'Logout') {
            return {name: 'Login'}
        }
    })
}
authMiddleware()

export default router