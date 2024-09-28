import {Box, Flex, Text} from "@chakra-ui/react";
import {TicketData} from "../../types/props/ticket.ts";
import {PillTag} from "../ui/PillTag.tsx";
import {truncation} from "../../utility/strings.ts";

export function TicketCard({ticketData}: { ticketData: TicketData }) {

    return (
        <Box
            padding={"10px"}
            marginY={"10px"}
            backgroundColor={"gray.200"}
            borderRadius={"5px"}
            textAlign={"left"}
        >
            <Flex
                justifyContent={"space-between"}
                gap={"10px"}
            >
                {ticketData.title}
                <PillTag text={ticketData.visibility}/>
            </Flex>
            <Text
                color={"gray"}
            >
                {truncation(ticketData.description, 50, "...")}
            </Text>
        </Box>
    )
}