const API_BASE_URL = "http://localhost:3001/api/v1"
const REDIRECT_BASE_URL = "http://localhost:3001"

const linkService = async (url: string) => {
  try {
    const response = await fetch(`${API_BASE_URL}/shorten`, {
      method: "POST",
      headers: {
        "Content-Type": "application/json"
      },
      body: JSON.stringify({ url })
    })

    if (!response.ok) {
      throw new Error("Failed to shorten URL")
    }

    const data = await response.json()
    console.log("URL raccourcie reçue:", data)

    const shortUrl = `${REDIRECT_BASE_URL}/${data.short_url}`
    console.log("URL de redirection finale:", shortUrl)

    return {
      ...data,
      shortUrl
    }
  } catch (error) {
    console.error("Erreur dans linkService:", error)
    throw error
  }
}

export default linkService