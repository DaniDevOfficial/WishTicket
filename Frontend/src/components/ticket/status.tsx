import {Box} from "@chakra-ui/react";

export function Status({status}: {status: string}) {
    let color = "orange"
    if (status.toLowerCase() === "open"){
        color = "primary"
    }

    if (status.toLowerCase() === "closed") {
        color = "accent"
    }

    return (
        <Box
            backgroundColor={color}
        >
            {status}
        </Box>
    );
}

