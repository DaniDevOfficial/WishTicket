import {Box} from "@chakra-ui/react";
import {UserProfileCard} from "../components/user/UserProfileCard.tsx";
import {UserData} from "../types/props/user.ts";

export function User() {
    const userData: UserData = {
        username: "dani",
        profilePicture: "https://avatars.githubusercontent.com/u/79514091?v=4"
    }
    return (
        <Box>
            <UserProfileCard userData={userData}/>
        </Box>
    )
}