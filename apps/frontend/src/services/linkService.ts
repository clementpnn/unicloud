import kyInstance from "./utils/instance"
import { LinkServiceResponse } from "@/types/linkServiceTypes"

export default function linkService(link: string): Promise<LinkServiceResponse> {
  return kyInstance.post("v4/shorten", { json: { "group_guid": import.meta.env.VITE_GROUP, "domain": "bit.ly", "long_url": link } }).json()
}