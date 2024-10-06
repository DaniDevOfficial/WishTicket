import {Box, Button, useToast} from "@chakra-ui/react";
import {UserProfileCard} from "../components/user/UserProfileCard.tsx";
import {UserTickets} from "../components/user/UserTickets.tsx";
import {useNavigate, useParams} from "react-router-dom";
import {useEffect, useState} from "react";
import {getAssignedAndOwnedTicketsByUsername} from "../repo/ticket/TicketRepository.ts";
import {AssignedAndOwned} from "../types/props/ticket.ts";
import {UserData} from "../types/props/user.ts";

export function User() {


    const [username, setUsername] = useState<string>("")
    const [ticketData, setTicketData] = useState<AssignedAndOwned | undefined>()
    const params = useParams();
    const toast = useToast()
    const navigate = useNavigate()
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
            <Button onClick={() => {navigate("/ticket/new")}} colorScheme={"primary"}>Create New Ticket</Button>
            {ticketData && (
                <UserTickets ticketsData={ticketData}/>
            )}

        </Box>
    )
}