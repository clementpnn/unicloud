import { z } from "zod"

export const urlSchema = z.object({
  link: z.string().min(1, { message: "URL cannot be empty" }).url("Please enter a valid URL")
})
