import {Box} from "@chakra-ui/react";
import {UserProfileCard} from "../components/user/UserProfileCard.tsx";
import {UserData} from "../types/props/user.ts";
import {UserTickets} from "../components/user/UserTickets.tsx";
import {useParams} from "react-router-dom";
import {useEffect, useState} from "react";

export function User() {
    const params = useParams();
    const [username, setUsername] = useState<string>("")

    useEffect(() => {
        setUsername(params.username ?? "")

        async function getTickets() {
            try {
                const data = getTicketsForUser()

            } catch (e) {

            }
        }
    }, [params])


    const userData: UserData = {
        username: username,
        profilePicture: "https://avatars.githubusercontent.com/u/79514091?v=4"
    }


    return (
        <Box>
            <UserProfileCard userData={userData}/>
            <UserTickets/>
        </Box>
    )
}