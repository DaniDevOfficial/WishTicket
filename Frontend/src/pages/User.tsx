import {Box, Button, useToast} from "@chakra-ui/react";
import {UserProfileCard} from "../components/user/UserProfileCard.tsx";
import {UserTickets} from "../components/user/UserTickets.tsx";
import {useParams} from "react-router-dom";
import {useEffect, useState} from "react";
import {getAssignedAndOwnedTicketsByUsername} from "../repo/ticket/TicketRepository.ts";
import {AssignedAndOwned} from "../types/props/ticket.ts";
import {getJWTBody} from "../utility/jwt.ts";
import {getValueFromLocalStorage} from "../utility/localStorage.ts";

export function User() {
    let currentUserData: jwtBody | null = {
        UserName: "",
        UserId: -1
    }
    try {
    const jwt = getValueFromLocalStorage('auth')
        currentUserData = getJWTBody(jwt)
    } catch (e) {
        currentUserData = {
            UserName: "",
            UserId: -1
        }
    }
    const [username, setUsername] = useState<string>("")
    const [ticketData, setTicketData] = useState<AssignedAndOwned | undefined>()
    const [isYou, setIsYou] = useState<boolean>(false)
    const params = useParams();
    const toast = useToast()

    useEffect(() => {
        setUsername(params.username ?? "")
        if (params.username == currentUserData)
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
            <Button colorScheme={"primary"}>Create New Ticket</Button>
            {ticketData && (
                <UserTickets ticketsData={ticketData}/>
            )}

        </Box>
    )
}