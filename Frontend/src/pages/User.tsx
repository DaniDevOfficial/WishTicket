import {Box, useToast} from "@chakra-ui/react";
import {UserProfileCard} from "../components/user/UserProfileCard.tsx";
import {UserData} from "../types/props/user.ts";
import {UserTickets} from "../components/user/UserTickets.tsx";
import {useParams} from "react-router-dom";
import {useEffect, useState} from "react";
import {getAssignedAndOwnedTicketsByUsername} from "../repo/ticket/TicketRepository.ts";
import {AssignedAndOwned} from "../types/props/ticket.ts";

export function User() {
    const params = useParams();
    const [username, setUsername] = useState<string>("")
    const [ticketData, setTicketData] = useState<AssignedAndOwned | undefined>()
    const toast = useToast()

    useEffect(() => {
        setUsername(params.username ?? "")
    }, [params])
    useEffect(() => {
        getTickets()
    }, [username]);

    async function getTickets() {

        try {
            if (username.trim() == '') {
                return
            }
            const data = await getAssignedAndOwnedTicketsByUsername(username)
            setTicketData(data)
        } catch (e) {
            toast({
                title: 'Error getting data.',
                description: "whopsie 🤭🤭",
                status: 'error',
                isClosable: true,
            })

        }
    }
    
    const userData: UserData = {
        username: username,
        profilePicture: "https://avatars.githubusercontent.com/u/79514091?v=4"
    }
    return (
        <Box>
            <UserProfileCard userData={userData}/>
            {ticketData && (
                <UserTickets ticketsData={ticketData}/>
            )}

        </Box>
    )
}