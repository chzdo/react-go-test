import * as React from "react"



// markup
const IndexPage = () => {

  const [status, setStatus] = React.useState(0)
  const [message, setMessage] = React.useState("nothing yet")
  React.useEffect(() => {
    fetch("/api/test", {
      method:"GET"
    }).then(result => result.text()).then(jsonResult => {
        console.log(jsonResult)
      jsonResult = JSON.parse(jsonResult)
      setStatus(jsonResult.status)
      setMessage(jsonResult.message)
    }).catch(err=> console.log(err))
  })
  return (
    <>
      <div>
        {status}
      </div>
      <div>
        {message}
      </div>
    </>
  )
}

export default IndexPage
