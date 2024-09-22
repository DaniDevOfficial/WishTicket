import { RouterProvider, createHashRouter } from "react-router-dom";
import { Home } from "./pages/Home";
import { DefaultLayout } from "./layout/Default";




const router = createHashRouter([
  {
    path: "/",
    element: <DefaultLayout />,
    children: [
      {
        index: true, // same path as parent: "/"
        element: <Home />
      }
    ],
  },

]);

export function Router() {
  return <RouterProvider router={router} />;
}