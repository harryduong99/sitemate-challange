
import fetch from "node-fetch";

export default async (id, title, description, options) => {

  const rawResponse = await fetch('http://localhost:6868/api/v1/issue', {
    method: 'POST',
    headers: {
      'Accept': 'application/json',
      'Content-Type': 'application/json'
    },
    body: JSON.stringify({id: id, title: title, description: description})
  });
  const content = await rawResponse.json();

  console.log(content);
};