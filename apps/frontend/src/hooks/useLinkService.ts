import { useMutation } from "@tanstack/react-query"
import { toast } from "react-hot-toast"

import linkService from "@/services/linkService"
import { errorMetrics } from "@/services/errorMetrics"

export default function useLinkService() {
  return useMutation({
    mutationFn: linkService,
    onError: (error: Error) => {
      errorMetrics.reportError(error, "link_service")
      toast.error(`Erreur: ${error.message}`)
    }
  })
}