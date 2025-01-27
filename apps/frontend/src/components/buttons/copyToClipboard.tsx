import { Copy } from "lucide-react"
import { Button } from "@/components/ui/button"
import { useCopyClipboard } from "@/hooks/useCopyClipboard"
import { useLaunchConfetti } from "@/hooks/useLaunchConfetti"

export default function CopyToClipboard({ link }: { link: string }) {
  const { isCopied, copy } = useCopyClipboard()
  const launchConfetti = useLaunchConfetti()

  return (
    <Button
      variant="ghost"
      onClick={() => {
        copy(link), launchConfetti()
      }}
      disabled={isCopied}
      size="icon"
      id="copy"
      className="h-[1.2rem] w-[1.2rem]"
    >
      <Copy />
      <span className="sr-only">Copy</span>
    </Button>
  )
}
