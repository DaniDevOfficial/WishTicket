import { RouterProvider, createHashRouter } from "react-router-dom";
import { DefaultLayout } from "./layouts/Default";
import { HomePage } from "./pages/Home";
import { Signup } from "./pages/Signup";
import {SignIn} from "./pages/SignIn.tsx";




const router = createHashRouter([
  {
    path: "/",
    element: <DefaultLayout />,
    children: [
      {
        index: true, // same path as parent: "/"
        element: <HomePage />,
      },
      {
        path: "/signup", 
        element: <Signup />,
      },
      {
        path: "/signIn",
        element: <SignIn />,
      }
    ],
  },

]);

export function Router() {
  return <RouterProvider router={router} />;
}
