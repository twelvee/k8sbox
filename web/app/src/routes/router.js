import {createRouter, createWebHistory} from 'vue-router'
import Login from "../pages/authorization/Login.vue";
import Dashboard from "../pages/Dashboard.vue";
import Users from "../pages/users/Users.vue";
import Clusters from "../pages/clusters/Clusters.vue";
import Boxes from "../pages/boxes/Boxes.vue";
import Environments from "../pages/environments/Environments.vue";
import Invite from "../pages/users/Invite.vue";
import store from "../stores";
import RedeemCode from "../pages/authorization/RedeemCode.vue";
import Logout from "../pages/authorization/Logout.vue";
import UserDetails from "../pages/users/UserDetails.vue";
import Setup from "../pages/setup/Setup.vue";
import CreateCluster from "../pages/clusters/CreateCluster.vue";
import ClusterEdit from "../pages/clusters/ClusterEdit.vue";
import ClusterDelete from "../pages/clusters/ClusterDelete.vue";
import CreateBox from "../pages/boxes/CreateBox.vue";
import EditBox from "../pages/boxes/EditBox.vue";
import DeleteBox from "../pages/boxes/DeleteBox.vue";
import CreateEnvironment from "../pages/environments/CreateEnvironment.vue";
import EnvironmentDetails from "../pages/environments/EnvironmentDetails.vue";
import DeleteEnvironment from "../pages/environments/DeleteEnvironment.vue";

const router = createRouter({
    routes: [
        {
            path: '/',
            name: 'Dashboard',
            component: Dashboard
        },
        {
            path: '/setup',
            name: 'Setup',
            component: Setup
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
            path: '/users/:id',
            name: 'Users.Details',
            component: UserDetails
        },
        {
            path: '/users/invite',
            name: 'Users.Invite',
            component: Invite
        },
        {
            path: '/clusters',
            name: 'Clusters',
            component: Clusters
        },
        {
            path: '/clusters/connect',
            name: 'Clusters.Create',
            component: CreateCluster
        },
        {
            path: '/clusters/:name',
            name: 'Clusters.Edit',
            component: ClusterEdit
        },
        {
            path: '/clusters/:name/delete',
            name: 'Clusters.Delete',
            component: ClusterDelete
        },
        {
            path: '/boxes',
            name: 'Boxes',
            component: Boxes
        },
        {
            path: '/boxes/create',
            name: 'Boxes.Create',
            component: CreateBox
        },
        {
            path: '/boxes/:name/edit',
            name: 'Boxes.Edit',
            component: EditBox
        },
        {
            path: '/boxes/:name/delete',
            name: 'Boxes.Delete',
            component: DeleteBox
        },
        {
            path: '/environments',
            name: 'Environments',
            component: Environments
        },
        {
            path: '/environments/create',
            name: 'Environments.Create',
            component: CreateEnvironment
        },
        {
            path: '/environments/:name',
            name: 'Environments.Details',
            component: EnvironmentDetails
        },
        {
            path: '/environments/:name/delete',
            name: 'Environments.Delete',
            component: DeleteEnvironment
        }
    ],
    history: createWebHistory()
})

function authMiddleware() {
    const user = store.getters.getUser

    router.beforeEach(async (to, from) => {
        if (!user.Token || user.Token.length === 0) {
            console.log(to.name)
            const data = await store.dispatch('checkIsSetupRequired')

            if (to.name !== 'Setup' && data.required === true) {
                return {name:'Setup'}
            } else if (data.required === false) {
                if (to.name !== 'Login' && to.name !== 'RedeemCode') {
                    return {name: 'Login'}
                }
                if (to.name === 'Logout') {
                    return {name: 'Login'}
                }
            }
        }

        if (
            user.Token &&
            (to.name === 'Login' || to.name === 'RedeemCode')
        ) {
            return {name: 'Dashboard'}
        }
    })
}
authMiddleware()

export default router