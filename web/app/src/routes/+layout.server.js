import {auth_token} from "./user_store.js";
import { redirect } from '@sveltejs/kit';

export function load({route, cookies}) {
    let isAuth = false
    let token = cookies.get('auth_token')
    if (token && token.length > 0) {
        auth_token.set(token)
        isAuth = true
    }

    if (route.id === "/app/login" && isAuth) {
        throw redirect(307, "/app")
    }

    if (route.id !== "/app/login" && route.id !== "/" && !isAuth) {
        throw redirect(307, "/app/login")
    }

    return {
        isAuth,
        token
    }
}