import {useEffect, useState} from "react";
import {useParams} from "react-router-dom";
import {Box, Heading, HStack, Text, useToast} from "@chakra-ui/react";
import {getTicketById} from "../repo/ticket/TicketRepository.ts";
import {TicketData} from "../types/props/ticket.ts";
import {Status} from "../components/ticket/status.tsx";

export function Ticket() {
    const [ticketId, setTicketId] = useState<number | undefined>()
    const [ticketData, setTicketData] = useState<TicketData | undefined>()
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
                const ticketDataFromBackend = await getTicketById(ticketId)
                if (ticketDataFromBackend === undefined){
                    throw new Error("Ticket doesnt exist")
                }
                setTicketData(ticketDataFromBackend)
                console.log(ticketDataFromBackend)
            } catch (e) {
                toast({
                    title: e.message
                })
                //TODO: remove this dev stuff
                setTicketData({
                    dueDate: "234",
                    title: "title",
                    description: {
                        String: "description",
                        Valid: true,
                    },
                    visibility: "PUBLIC",
                    status: "Open",
                    ticketId: ticketId
                })
            }
        }
    }, [ticketId]);


    if (!ticketId || ticketId == -1 || ticketData == undefined) {
        return <div>404 Ticket not found tschorry</div> // TODO: better 404 page
    }
    return (
        <>
            <Heading>
                {ticketData.title}
            </Heading>
            <Text>
                {ticketData.description.Valid ? (
                    <>
                        {ticketData.description.String}
                    </>
                ) : (
                    <>
                        No Description Yet
                    </>
                )}
            </Text>
            <HStack
                justifyContent={"space-between"}
            >
                <Box>
                    {}
                </Box>
                <Box>
                    <Status status={ticketData.status}/>
                </Box>
            </HStack>

        </>
    );
}

