import {Box, Flex} from "@chakra-ui/react";
import {PrettyHeader} from "../../lib/PrettyHeader.tsx";
import {useState} from "react";
import {AssignedAndOwned} from "../../types/props/ticket.ts";
import {TicketCard} from "./TicketCard.tsx";

const ASSIGNED = "assigned"

const OWNED = "OWNED"

export function UserTickets({ticketsData}: { ticketsData: AssignedAndOwned }) {
    const [currentlySelected, setCurrentlySelected] = useState(ASSIGNED)
    console.log(ticketsData)
    return (
        <Flex justifyContent={"center"} alignItems={"center"} flexDir={"column"}>
            <Flex w={"300px"} justifyContent={"space-between"}>
                <Box cursor={"pointer"} onClick={() => {
                    setCurrentlySelected(ASSIGNED)
                }}>
                    <PrettyHeader name={"Assigned Tickets"} isOpen={currentlySelected === ASSIGNED}/>
                </Box>
                <Box cursor={"pointer"} onClick={() => {
                    setCurrentlySelected(OWNED)
                }}>
                    <PrettyHeader name={"Owned Tickets"} isOpen={currentlySelected === OWNED}/>
                </Box>
            </Flex>

            {currentlySelected === ASSIGNED ? (
                <Box width={"300px"}>
                    {ticketsData["assigned"] && ticketsData["assigned"].length > 0 ? (
                        ticketsData["assigned"].map((ticket, index) => (
                            <TicketCard
                                ticketData={ticket}
                                key={ticket.ticketId || `assigned-ticket-${index}`}
                            />
                        ))
                    ) : (
                        <p>No assigned tickets available</p>
                    )}
                </Box>
            ) : (
                <Box width={"300px"}>
                    {ticketsData["owned"] && ticketsData["owned"].length > 0 ? (
                        ticketsData["owned"].map((ticket, index) => (
                            <TicketCard
                                ticketData={ticket}
                                key={ticket.ticketId || `owned-ticket-${index}`}
                            />
                        ))
                    ) : (
                        <p>No owned tickets available</p>
                    )}
                </Box>
            )}

        </Flex>
    )
}