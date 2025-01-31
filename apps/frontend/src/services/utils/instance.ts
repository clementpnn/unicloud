import ky from "ky"

const kyInstance = ky.create({
  prefixUrl: "http://back.clementpnn.com"
})

export default kyInstance
