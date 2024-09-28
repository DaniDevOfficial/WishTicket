import {Box} from "@chakra-ui/react";

export function PillTag({text}: {text: string}) {

    return (
        <Box
            outline={"1px solid gray"}
            borderRadius={"10px"}
            padding={"5px"}
        >
            {text}
        </Box>
    )
}