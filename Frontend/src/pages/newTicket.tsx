import {FormControl, FormLabel, Input, InputGroup, Select, Stack, useToast} from "@chakra-ui/react";
import React, {useState} from "react";
import {createNewTicket} from "../repo/ticket/TicketRepository.ts";

export function NewTicket() {
    const [formData, setFormData] = useState({
        title: "",
        description: "",
        dueDate: "",
        visibility: ""
    })
    const toast = useToast();

    function handleChange(e: React.ChangeEvent<HTMLSelectElement | HTMLInputElement>) {
        const {name, value} = e.target
        setFormData(prevState => ({
            ...prevState,
            [name]: value.trim()
        }))
    }


    async function submitForm(e: React.FormEvent) {
        e.preventDefault()
        if (formData.title.trim() === "") {
            toast({
                title: "Field Missing",
                description: "Title has to be filled out",
                status: "warning"
            })
            return
        }
        if (formData.visibility.trim() === "") {
            toast({
                title: "Field Missing",
                description: "Visibility has to be filled out",
                status: "warning"
            })
            return
        }
        try {
            const response = await createNewTicket(formData)
            toast({
                title: "Yayyy ticket created 😊😊😊😊",
                description: "Ticket Created successfully",
                status: "success"
            })
        } catch (e) {
            alert(e.message)
            toast({
                title: "Creation went wrong",
                description: "Cant create the ticket 😱😱😱😱",
                status: "error"
            })
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
                        <FormLabel>Visibility</FormLabel>
                        <InputGroup>
                            <Select
                                required
                                focusBorderColor='primary.base'
                                name="visibility"
                                value={formData.visibility}
                                onChange={handleChange}
                                placeholder="Select visibility"
                            >
                                <option value="public">Public</option>
                                <option value="private">Private</option>
                            </Select>
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

