import {Box, Text} from "@chakra-ui/react";

export function PrettyHeader({ name, isOpen }: { name: string, isOpen: boolean}) {
    return (
        <Text fontWeight={"500"} color={"primary.900"} textAlign={"center"} mt={6} cursor={"pointer"}>
            <Box
                as="span"
                position="relative"
                zIndex={5}
                overflow={'hidden'}
                padding={'0 0.3rem'}
            >
                <Box
                    as="span"
                    position={'absolute'}
                    bottom={"-0.5px"}
                    left={0}
                    height={isOpen ? '110%' : '10%'}
                    width={'100%'}
                    background={'linear-gradient(180deg, transparent 50%, #3de7d0 0%)'}
                    opacity={1}
                    zIndex={-1}
                    transition={'height 0.3s ease'}
                />
                {name}
            </Box>
        </Text>
    );
}
