import { useState, useEffect } from "react"
import { useCopyToClipboard, useTimeoutFn } from "react-use"

export function useCopyClipboard() {
  const [isCopied, setIsCopied] = useState(false)
  const [, copyToClipboard] = useCopyToClipboard()
  const [isTimeoutActive, clear, reset] = useTimeoutFn(() => setIsCopied(false), 3000)

  const copyText = (text: string) => {
    copyToClipboard(text)
    setIsCopied(true)
    reset()
  }

  useEffect(() => {
    if (!isCopied && !isTimeoutActive()) {
      clear()
    }
  }, [isCopied, clear, isTimeoutActive])

  return { isCopied, copy: copyText }
}
