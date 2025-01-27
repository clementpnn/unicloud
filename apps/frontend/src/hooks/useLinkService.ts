import { useMutation } from "@tanstack/react-query"

import linkService from "@/services/linkService"

export default function useLinkService() {
  return useMutation({ mutationFn: linkService })
}