const getResponse = async param => {
  const res = await fetch(`http://localhost:8080/${param}`)
  return res
}

const send = async value => {
  const response = await getResponse(value)
  console.log(response)
  document.getElementById('response').innerText = response
}
