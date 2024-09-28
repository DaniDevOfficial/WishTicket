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
    Text,
    useToast
} from '@chakra-ui/react'
import React, {useState} from 'react'
import {FaEye, FaEyeSlash} from 'react-icons/fa'
import {createNewUser} from "../repo/user/UserRepository.ts";
import {NewUser} from "../types/user.ts";
import {addToLocalStorage} from "../utility/localStorage.ts";
import {Link} from "react-router-dom";


export function Signup() {
    const [showPassword, setShowPassword] = useState(false)
    const [formData, setFormData] = useState({
        username: '',
        email: 'email@email.com',
        password: '',
        confirmPassword: ''
    })
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

        const {username, email, password, confirmPassword} = formData;
        if (!username || !email || !password || !confirmPassword) {
            alert('All fields are required.');
            return;
        }

        if (password !== confirmPassword) {
            alert('Passwords do not match.');
            return;
        }
        const newUser: NewUser = {
            username: username,
            email: email,
            password: password
        }
        try {
            const res = await createNewUser(newUser)
            const token = res?.token
            addToLocalStorage('auth', token)
        } catch (e) {
            toast({
                title: 'Signup error.',
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
                        <FormLabel>Email</FormLabel>
                        <Input
                            required
                            focusBorderColor='primary.base'

                            type="email"
                            name="email"
                            placeholder='Email'
                            value={formData.email}
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
                        <FormLabel>Confirm Password</FormLabel>
                        <InputGroup>
                            <Input
                                required
                                focusBorderColor='primary.base'
                                type={showPassword ? "text" : "password"}
                                name="confirmPassword"
                                placeholder='Confirm Password'
                                value={formData.confirmPassword}
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
                Or <Link to={"/signIn"}>Sign In</Link>
            </Text>
        </Container>
    )
}