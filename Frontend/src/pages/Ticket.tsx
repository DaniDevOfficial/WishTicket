import {useEffect, useState} from "react";
import {useParams} from "react-router-dom";
import {Heading, useToast} from "@chakra-ui/react";
import {getTicketById} from "../repo/ticket/TicketRepository.ts";

export function Ticket() {
    const [ticketId, setTicketId] = useState<number | undefined>()
    const params = useParams();
    const toast = useToast()

    useEffect(() => {
        setTicketId(Number(params.ticketId) ?? -1)
    }, [params])

    useEffect(() => {
        if (ticketId && ticketId >= 0) {
            getTicketData(ticketId).then()
        }

        async function getTicketData(ticketId: number) {
            try {
                await getTicketById(ticketId)
            } catch (e) {
                toast({
                    title: e.message
                })
            }
        }
    }, [ticketId]);
    if (!ticketId || ticketId == -1) {
        return <div>404 Ticket not found tschorry</div> // TODO: better 404 page
    }
    return (
        <>
            <Heading>

            </Heading>
        </>
    );
}

