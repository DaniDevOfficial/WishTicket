import {FormControl, FormLabel, Icon, Input, InputGroup, InputRightElement, Stack, Textarea} from "@chakra-ui/react";
import {FaEye, FaEyeSlash} from "react-icons/fa";
import React, {useState} from "react";

export function NewTicket() {
    const [formData, setFormData] = useState({
        title: "",
        description: "",
        dueDate: "",
        visibility: ""
    })


    function handleChange(e: React.ChangeEvent<HTMLInputElement>) {
        const {name, value} = e.target
        setFormData(prevState => ({
            ...prevState,
            [name]: value.trim()
        }))
    }


    function submitForm() {
        if(formData.title == "") {

        }
    }

    return (
        <>
            <form onSubmit={submitForm}>
                <Stack>
                    <FormControl>
                        <FormLabel>TicketTitle</FormLabel>
                        <Input
                            required
                            focusBorderColor='primary.base'
                            type="text"
                            name="title"
                            placeholder='Title'
                            value={formData.title}
                            onChange={handleChange}
                        />
                    </FormControl>

                    <FormControl>
                        <FormLabel>Description</FormLabel>
                        <InputGroup>
                            <Input
                                focusBorderColor='primary.base'
                                type={"text"}
                                name="description"
                                placeholder='Description'
                                value={formData.description}
                                onChange={handleChange}
                            />
                        </InputGroup>
                    </FormControl>
                    <FormControl>
                        <FormLabel>Due Date</FormLabel>
                        <InputGroup>
                            <Input
                                focusBorderColor='primary.base'
                                type={"date"}
                                name="dueDate"
                                placeholder='Duedate'
                                value={formData.dueDate}
                                onChange={handleChange}
                            />
                        </InputGroup>
                    </FormControl>
                    <FormControl>
                        <Input type="submit" value="Submit" _hover={{cursor: 'pointer', bg: 'background.700'}}/>
                    </FormControl>
                </Stack>
            </form>
        </>
    );
}

