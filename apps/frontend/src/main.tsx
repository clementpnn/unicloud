import ReactDOM from "react-dom/client"
import { RouterProvider } from "@tanstack/react-router"
import { QueryClient, QueryClientProvider } from "@tanstack/react-query"

import router from "./routes/routes"
import { ThemeProvider } from "@/provider/themeProvider"
import "./styles/index.css"
import MainContainer from "./components/containers/mainContainer"
import { Toaster } from "@/components/ui/toaster"

const queryClient = new QueryClient()

ReactDOM.createRoot(document.querySelector("#root")!).render(
  <QueryClientProvider client={queryClient}>
    <ThemeProvider defaultTheme="system" storageKey="vite-ui-theme">
      <MainContainer>
        <RouterProvider router={router} />
      </MainContainer>
      <Toaster />
    </ThemeProvider>
  </QueryClientProvider>
)
