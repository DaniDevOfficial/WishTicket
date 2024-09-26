
const DB_URL = import.meta.env.VITE_BACKEND_URL

export async function getAssignedAndOwnedTicketsByUsername(username: string) {
    const req = {
        username: username
    }
    try {
        const res = await fetch(DB_URL + 'ticket', {
            method: 'GET',
            headers: {
                'Content-Type': 'application/json'
            },
            body: JSON.stringify(req)
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