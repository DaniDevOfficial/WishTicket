import { Button, Flex, Heading, Text } from '@chakra-ui/react'
import { useNavigate } from 'react-router-dom'

export function Landing() {
    const navigate = useNavigate()
    return (
        <>
            <Flex
                direction="column"
                align="center"
                justify="center"
                height="100vh"
            >
                <Heading as="h1" size="2xl">Welcome To WishTicket</Heading>
                <Text>Boost your Teams Productivity by 100%</Text>
                <Button
                    colorScheme="primary"
                    mt={4}
                    onClick={() => {
                        navigate('/signup')
                    }}
                >
                    Get Started
                </Button>
            </Flex>
        </>
    )
}