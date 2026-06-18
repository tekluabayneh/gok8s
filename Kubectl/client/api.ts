const request = async (method: string, path: string, body?: unknown) => {
  return fetch(`http://localhost:${path}`, {
    method: method,
    headers: {
      "Content-type": "application/json"
    },
    body: body ? JSON.stringify(body) : undefined
  })
}


export default request
