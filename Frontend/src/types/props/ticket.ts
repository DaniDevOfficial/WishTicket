
export interface TicketData {
    title: string,
    description: string,
    visibility: string,
    status: string,
    ticket_id: number
}

export interface AssignedAndOwned {
    "assigned": TicketData[],
    "owned": TicketData[]
}