import { Outlet, ScrollRestoration } from "react-router-dom";
import { Container, chakra } from "@chakra-ui/react";
import { Tokens } from "../../.mirrorful/theme";

import { Global } from "@emotion/react";

export function DefaultLayout() {
    const backgroundColor = Tokens.colors.background["base"];    
    return (
        <>
            <Global
                styles={{
                    body: {
                        backgroundColor: backgroundColor,
                    },
                }}
            />
            <chakra.div width={"100%"} color={"text.base"} textAlign={"center"}>
                <chakra.div minHeight={"100vh"} width={"100%"}>
                    <chakra.main marginBottom={"2rem"}>
                        <Container maxW={"5xl"}>
                            <Outlet />
                        </Container>
                    </chakra.main>
                </chakra.div>

                <ScrollRestoration />
            </chakra.div>
        </>
    );
}
