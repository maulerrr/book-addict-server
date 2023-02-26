import axios from "axios";

const getFinished = async () => {
  const ID = JSON.parse(localStorage.getItem("ID"))
  const token = localStorage.getItem("token")
  const options = {
    method: "GET",
    url: "http://localhost:3001/api/v1/booktabs/finished/"+ID,
    headers: {
      'Authorization': "Bearer " + token,
    }
  };

  await axios.request(options).then(function(response) {
    const finished = response.data;
    localStorage.setItem("finished", JSON.stringify(finished));
  }).catch(function(error) {
    console.error(error);
  });
}

export default getFinished;