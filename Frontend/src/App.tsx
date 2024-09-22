import { Router } from "./Router";
import { ChakraProvider } from "@chakra-ui/react";
import { theme, toastOptions } from "./configs/chakra";

export function App() {

  return (
    <>
      <ChakraProvider theme={theme} toastOptions={toastOptions}>
        <Router />
      </ChakraProvider>
    </>
  )
}

