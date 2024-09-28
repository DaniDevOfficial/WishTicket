import {
    Container,
    FormControl,
    FormLabel,
    Heading,
    Icon,
    Input,
    InputGroup,
    InputRightElement,
    Stack,
    Text, useToast
} from '@chakra-ui/react'
import React, {useState} from 'react'
import {FaEye, FaEyeSlash} from 'react-icons/fa'
import {signIn} from "../repo/user/UserRepository.ts";
import {SignInCredentials} from "../types/user.ts";
import {addToLocalStorage} from "../utility/localStorage.ts";
import {Link, useNavigate} from "react-router-dom";

export function SignIn() {
    const [showPassword, setShowPassword] = useState(false)
    const [formData, setFormData] = useState({
        username: 'DaniDevOfficial',
        password: 'dani',
    })
    const navigate = useNavigate()
    const toast = useToast()


    function handleChange(e: React.ChangeEvent<HTMLInputElement>) {
        const {name, value} = e.target
        setFormData(prevState => ({
            ...prevState,
            [name]: value
        }))
    }

    async function submitForm(e: React.FormEvent) {
        e.preventDefault()

        const {username, password} = formData;
        if (!username || !password) {
            alert('All fields are required.');
            return;
        }


        const signInCredentials: SignInCredentials = {
            username: username,
            password: password
        }
        try {
            const res = await signIn(signInCredentials)
            const token = res?.token
            console.log(token)
            addToLocalStorage('auth', token)
            console.log(localStorage)
            navigate("/user/" + username)
        } catch (e) {
            toast({
                title: 'Sign In error.',
                description: "whopsie 🤭🤭",
                status: 'error',
                isClosable: true,
            })
        }
    }

    return (
        <Container
            maxWidth={"400px"}
            width={"80vw"}
            backgroundColor={"background.600"}
            borderRadius={"10px"}
            padding={"20px"}
        >
            <Heading>
                Sign up for Free
            </Heading>
            <form onSubmit={submitForm}>
                <Stack>
                    <FormControl>
                        <FormLabel>Username</FormLabel>
                        <Input
                            required
                            focusBorderColor='primary.base'

                            type="text"
                            name="username"
                            placeholder='Username'
                            value={formData.username}
                            onChange={handleChange}
                        />
                    </FormControl>

                    <FormControl>
                        <FormLabel>Password</FormLabel>
                        <InputGroup>
                            <Input
                                required
                                focusBorderColor='primary.base'

                                type={showPassword ? "text" : "password"}
                                name="password"
                                placeholder='Password'
                                value={formData.password}
                                onChange={handleChange}
                            />
                            <InputRightElement>
                                <Icon
                                    as={showPassword ? FaEyeSlash : FaEye}
                                    onClick={() => setShowPassword(!showPassword)}
                                    _hover={{cursor: 'pointer'}}
                                />
                            </InputRightElement>
                        </InputGroup>
                    </FormControl>
                    <FormControl>
                        <Input type="submit" value="Submit" _hover={{cursor: 'pointer', bg: 'background.700'}}/>
                    </FormControl>
                </Stack>
            </form>
            <Text
                fontSize={"xs"}
            >
                Or <Link to={"/signUp"}>Create An account</Link>
            </Text>
        </Container>
    )
}