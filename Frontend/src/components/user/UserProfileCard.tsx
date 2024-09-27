import {Container, Heading, Image} from "@chakra-ui/react";
import {UserData} from "../../types/props/user.ts";

export function UserProfileCard({userData}: { userData: UserData }) {

    return (
        <Container maxW={"200px"} textAlign={"center"}>

            <Image
                borderRadius={"50%"}
                outline={"2px solid"}
                outlineColor={"primary.base"}
                src={userData.profilePicture}
            />
            <Heading>
                {userData.username}
            </Heading>
        </Container>
    )
}