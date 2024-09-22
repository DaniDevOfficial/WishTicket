import { type ToastProviderProps, extendTheme } from "@chakra-ui/react";

import { Tokens } from "../../.mirrorful/theme";

const colors = {
  ...Tokens.colors,
};

export const theme = extendTheme({
  colors,

});

export const toastOptions: ToastProviderProps = {
  defaultOptions: {
    duration: 3000,
    isClosable: true,
    position: "top-right",
  },
};
