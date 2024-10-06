
export interface TicketData {
    title: string,
    description: {
        String: string,
        Valid: boolean
    },
    visibility: string,
    status: string,
    dueDate: string,
    ticketId: number
}

export interface AssignedAndOwned {
    "assigned": TicketData[],
    "owned": TicketData[]
}
export interface NewTicket {
    title: string,
    description: string | null | undefined,
    dueDate: string,
    visibility: "PUBLIC" | "PRIVATE"
}