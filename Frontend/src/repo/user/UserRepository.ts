import {NewUser, SignInCredentials} from "../../types/user.ts";
import {JwtResponse} from "../../types/responses/user.ts";
import {UserData} from "../../types/props/user.ts";

const DB_URL = import.meta.env.VITE_BACKEND_URL

export async function createNewUser(newUser: NewUser): Promise<JwtResponse | undefined> {
    try {
        const res = await fetch(DB_URL + 'users', {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json'
            },
            body: JSON.stringify(newUser)
        })
        if (!res.ok) {
            const errorBody = await res.text();
            throw new Error(`Account Creation Error: ${errorBody || res.statusText}`);
        }
        return await res.json()
    } catch (e) {
        console.log(e)
    }
}

export async function signIn(singInCredentials: SignInCredentials): Promise<JwtResponse | undefined> {
    try {
        const res = await fetch(DB_URL + 'users/login', {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json'
            },
            body: JSON.stringify(singInCredentials)
        })
        if (!res.ok) {
            const errorBody = await res.text();
            throw new Error(`Account Creation Error: ${errorBody || res.statusText}`);
        }
        return await res.json()
    } catch (e) {
        console.log(e)
    }
}

