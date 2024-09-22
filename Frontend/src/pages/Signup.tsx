import { Container, FormControl, FormLabel, Icon, Input, InputGroup, InputRightElement, Stack } from '@chakra-ui/react'
import React, { useState } from 'react'
import { FaEye, FaEyeSlash } from 'react-icons/fa'

export function Signup() {
    const [showPassword, setShowPassword] = useState(false)
    const [formData, setFormData] = useState({
        username: '',
        email: 'email@email.com',
        password: '',
        confirmPassword: ''
    })

    function handleChange(e: React.ChangeEvent<HTMLInputElement>) {
        const { name, value } = e.target
        setFormData(prevState => ({
            ...prevState,
            [name]: value
        }))
    }

    async function submitForm(e: React.FormEvent) {
        e.preventDefault()

        const { username, email, password, confirmPassword } = formData;
        if (!username || !email || !password || !confirmPassword) {
            alert('All fields are required.');
            return;
        }

        if (password !== confirmPassword) {
            alert('Passwords do not match.');
            return;
        }

        const res = await fetch('http://localhost:8000/users', {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json'
            },
            body: JSON.stringify(formData)
        })
        console.log(res);
        
    }

    return (
        <Container>
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
                                    _hover={{ cursor: 'pointer' }}
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
                                    _hover={{ cursor: 'pointer' }}
                                />
                            </InputRightElement>
                        </InputGroup>
                    </FormControl>
                    <FormControl>
                        <Input type="submit" value="Submit" _hover={{ cursor: 'pointer', bg: 'background.600' }} />
                    </FormControl>
                </Stack>
            </form>
        </Container>
    )
}