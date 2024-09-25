import {Box, Flex} from "@chakra-ui/react";
import {PrettyHeader} from "../../lib/PrettyHeader.tsx";
import {useState} from "react";

const ASSIGNED = "assigned"

const OWNED = "OWNED"

export function UserTickets() {
    const [currentlySelected, setCurrentlySelected] = useState(ASSIGNED)
    return (
        <Flex justifyContent={"center"} alignItems={"center"} flexDir={"column"}>
            <Flex w={"300px"} justifyContent={"space-between"}>
                    <Box onClick={() => {setCurrentlySelected(ASSIGNED)}}>
                        <PrettyHeader name={"Assigned Tickets"} isOpen={currentlySelected === ASSIGNED} />
                    </Box>
                    <Box onClick={() => {setCurrentlySelected(OWNED)}}>
                        <PrettyHeader name={"Owned Tickets"} isOpen={currentlySelected === OWNED} />
                    </Box>
            </Flex>
            {currentlySelected === ASSIGNED ? (
                <Box>
                    {/* Content for Assigned Tickets */}
                    <p>Here are your assigned tickets! 🎟️</p>
                </Box>
            ) : (
                <Box>
                    {/* Content for Owned Tickets */}
                    <p>These are the tickets you own! 💼</p>
                </Box>
            )}
        </Flex>
    )
}