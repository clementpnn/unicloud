import { RootRoute, Route, Router } from "@tanstack/react-router"
import App from "@/pages/App"
import NotFound from "@/pages/NotFound"

const rootRoute = new RootRoute()
const indexRoute = new Route({ getParentRoute: () => rootRoute, path: "/", component: App })
const notFoundRoute = new Route({ getParentRoute: () => rootRoute, path: "*", component: NotFound })

const routeTree = rootRoute.addChildren([indexRoute, notFoundRoute])

const router = new Router({ routeTree })

export default router

declare module "@tanstack/react-router" {
  interface Register {
    router: typeof router
  }
}