import ky from "ky"

const kyInstance = ky.create({
  prefixUrl: "http://localhost:3000/api/v1"
})

export default kyInstance
