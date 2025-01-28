import kyInstance from "./utils/instance"
import { LinkServiceResponse } from "@/types/linkServiceTypes"

export default function linkService(link: string): Promise<LinkServiceResponse> {
  return kyInstance.post("shorten", { json: { "url": link } }).json()
}