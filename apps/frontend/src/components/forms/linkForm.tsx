import { useForm } from "react-hook-form"
import z from "zod"
import { useEffect, useState } from "react"

import useLinkService from "@/hooks/useLinkService"
import { zodResolver } from "@hookform/resolvers/zod"
import { urlSchema } from "@/types/schema/urlSchema"
import { Input } from "@/components/ui/input"
import { Button } from "@/components/ui/button"
import { Form, FormControl, FormField, FormItem, FormMessage } from "@/components/ui/form"
import LinkModal from "@/components/modals/linkModal"
import Spinner from "@/assets/icons/spinner.svg?react"

import { useToast } from "@/components/ui/useToast"

export default function LinkForm() {
  const { toast } = useToast()
  const [isOpen, setIsOpen] = useState(false)
  const [link, setLink] = useState("")

  const form = useForm<z.infer<typeof urlSchema>>({
    resolver: zodResolver(urlSchema),
    mode: "onSubmit",
    defaultValues: {
      link: ""
    }
  })

  const { mutate, data, isError, error, status } = useLinkService()

  useEffect(() => {
    if (status === "success") {
      setLink(data?.link)
      setIsOpen(true)
    }

    if (isError) {
      toast({
        variant: "destructive",
        title: error.name,
        description: error.message
      })
    }
  }, [status, data, isError, error, toast])

  function onSubmit(values: z.infer<typeof urlSchema>) {
    mutate(values.link)

    form.reset()
  }

  return (
    <>
      <LinkModal isOpen={isOpen} setIsOpen={setIsOpen} link={link} />
      <Form {...form}>
        <form onSubmit={form.handleSubmit(onSubmit)} className="flex">
          <FormField
            control={form.control}
            name="link"
            render={({ field }) => (
              <FormItem>
                <FormControl>
                  <Input placeholder="https://example.com" {...field} />
                </FormControl>
                <FormMessage />
              </FormItem>
            )}
          />
          <Button type="submit" className="ml-5">
            {status === "pending" ? <Spinner /> : "Send"}
          </Button>
        </form>
      </Form>
    </>
  )
}
