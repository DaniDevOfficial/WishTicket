import {Button, Flex, Heading, Text} from '@chakra-ui/react'
import {useNavigate} from 'react-router-dom'

export function Landing() {
    const navigate = useNavigate()

    async function getTickets() {
        try {
            const response = await fetch('http://localhost:8000/ticket', {
                method: 'GET',
                headers: {
                    'Content-Type': 'application/json',
                    'Authorization': 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJFeHAiOjE3MjcxMTU2NTksIlVzZXJJZCI6MjEsIlVzZXJOYW1lIjoiZGFuaSJ9.3JF_-1-0lEdk8dV17o78jvbiAn_BQZK3QhbhmqPBFQk'
                },
            })
            if (!response.ok) {
                throw new Error(`HTTP error! status: ${response.status}`);
            }
            const data = await response.json();
            console.log('Fetched Tickets:', data);

        } catch (e) {


        }
    }

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
                <Button
                    onClick={() => {
                        getTickets()
                    }}
                >
                    GetAllTickets For user
                </Button>
            </Flex>
        </>
    )
}