import {Box} from "@chakra-ui/react";

export function Status({status}: {status: string}) {
    let backgroundColor = "orange.100"
    let borderColor = "orange"
    let hoverColor = "orange.200"
    if (status.toLowerCase() === "open"){
        backgroundColor = "primary.100"
        borderColor = "primary.base"
        hoverColor = "primary.200"
    }

    if (status.toLowerCase() === "closed") {
        backgroundColor = "accent.100"
        borderColor = "accent.300"
        hoverColor = "accent.200"
    }

    return (
        <Box
            borderRadius={10}
            p={2}
            bg={backgroundColor}
            borderWidth={"1px"}
            borderColor={borderColor}
            transition="background-color 0.3s ease"
            _hover={{
                bg: hoverColor
            }}

        >
            {status}
        </Box>
    );
}

