import {getValueFromLocalStorage} from "../../utility/localStorage.ts";
import {AssignedAndOwned, TicketData} from "../../types/props/ticket.ts";

// eslint-disable-next-line @typescript-eslint/ban-ts-comment
// @ts-expect-error
const DB_URL = import.meta.env.VITE_BACKEND_URL

export async function getAssignedAndOwnedTicketsByUsername(username: string): Promise<AssignedAndOwned> {
    if (username.trim() == '') {
        throw new Error("no username")
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
            throw new Error(`error: ${errorBody || res.statusText}`);
        }
        console.log(res)
        return await res.json()
    } catch (e) {
        throw new Error("Error while getting tickets")
    }
}

export async function getTicketById(ticketId: number): Promise<TicketData | undefined> {
    if (ticketId < 0) {
        throw new Error("Ticket does not exits")
    }

    try {
        const jwtToken = getValueFromLocalStorage('auth')
        const res = await fetch(DB_URL + 'ticket/single?ticketId=' + ticketId, {
            method: 'GET',
            headers: {
                'Content-Type': 'application/json',
                'Authorization': jwtToken
            },
        })
        if (!res.ok) {
            const errorBody = await res.text();
            throw new Error(`error: ${errorBody || res.statusText}`);
        }
        return await res.json()
    } catch (e) {
        throw new Error("Error while getting tickets")
    }
}

export async function createNewTicket(ticketData: {
    visibility: string;
    dueDate: string;
    description: string;
    title: string
}): Promise<{
    message: string,
    ticketId: number
}> {
    try {
        const jwtToken = getValueFromLocalStorage('auth')
        const res = await fetch(DB_URL + 'ticket', {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json',
                'Authorization': jwtToken
            },
            body: JSON.stringify(ticketData)
        })
        if (!res.ok) {
            const errorBody = await res.text();
            throw new Error(`error: ${errorBody || res.statusText}`);
        }
        if (!res.ok) {
            const errorBody = await res.text();
            throw new Error(`error: ${errorBody || res.statusText}`);
        }
        return await res.json()
    } catch (e) {
        console.error(e.message)
        throw new Error("Error while Creating ticket")
    }
}