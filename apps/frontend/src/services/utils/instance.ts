import ky from "ky"

const kyInstance = ky.create({
  prefixUrl: "https://unibackend.clementpnn.com"
})

export default kyInstance
