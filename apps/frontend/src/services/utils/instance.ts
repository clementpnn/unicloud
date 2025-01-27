import ky from "ky"

const kyInstance = ky.create({
  prefixUrl: "https://api-ssl.bitly.com",
  headers: {
    "Authorization": `Bearer ${import.meta.env.VITE_TOKEN}`
  }
})

export default kyInstance
