import {Container, Image, Text} from "@chakra-ui/react";
import {UserData} from "../../types/props/user.ts";

export function UserProfileCard({userData}: { userData: UserData }) {
    console.log(userData)
    return (
        <Container maxW={"296px"} textAlign={"left"}>

            <Image
                borderRadius={"50%"}
                outline={"2px solid"}
                outlineColor={"primary.base"}
                src={userData.profilePicture}
            />
            <Text>
                {userData.username}
            </Text>
        </Container>
    )
}