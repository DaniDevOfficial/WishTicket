import {getValueFromLocalStorage} from "../../utility/localStorage.ts";

// eslint-disable-next-line @typescript-eslint/ban-ts-comment
// @ts-expect-error
const DB_URL = import.meta.env.VITE_BACKEND_URL

export async function getAssignedAndOwnedTicketsByUsername(username: string) {
    if (username.trim() == ''){
        return;
    }
    try {
        const jwtToken = getValueFromLocalStorage('auth')
        const res = await fetch(DB_URL + 'ticket/all?username=' + username, {
            method: 'GET',
            headers: {
                'Content-Type': 'application/json',
                'Authorization': jwtToken
            },
        })
        if (!res.ok) {
            const errorBody = await res.text();
            throw new Error(`Account Creation Error: ${errorBody || res.statusText}`);
        }
        console.log(res)
        return await res.json()
    } catch (e) {
        console.log(e)
    }
}