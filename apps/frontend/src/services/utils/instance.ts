import ky from "ky"

const kyInstance = ky.create({
  prefixUrl: "http://localhost:3000/api/v1",
  headers: {
    "Authorization": `Bearer ${import.meta.env.VITE_TOKEN}`
  }
})

export default kyInstance
