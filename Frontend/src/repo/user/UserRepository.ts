import {NewUser, SignInCredentials} from "../../types/user.ts";
import {JwtResponse} from "../../types/responses/user.ts";
// eslint-disable-next-line @typescript-eslint/ban-ts-comment
// @ts-expect-error
const DB_URL = import.meta.env.VITE_BACKEND_URL

export async function createNewUser(newUser: NewUser): Promise<JwtResponse | undefined> {
    let res
    try {
        res = await fetch(DB_URL + 'users', {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json'
            },
            body: JSON.stringify(newUser)
        })
    } catch (e) {
        console.log(e)
        throw new Error(`Account Creation Error`);
    }

    if (!res.ok) {
        const errorBody = await res.text();
        throw new Error(`Account Creation Error: ${errorBody || res.statusText}`);
    }
    return await res.json()
}

export async function signIn(singInCredentials: SignInCredentials): Promise<JwtResponse | undefined> {
    let res
    try {
        res = await fetch(DB_URL + 'users/login', {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json'
            },
            body: JSON.stringify(singInCredentials)
        })
    } catch (e) {
        console.log(e)
        throw new Error(`Account Creation Error:`);
    }
    if (!res.ok) {
        const errorBody = await res.text();
        throw new Error(`Account Creation Error: ${errorBody || res.statusText}`);
    }
    return await res.json()
}

