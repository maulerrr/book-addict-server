import axios from "axios";

const getMyData = async () => {
  const ID = localStorage.getItem("ID");
  const options = {
    method: "GET",
    url: "http://localhost:3001/api/v1/users/"+ID
  };

  await axios.request(options).then(function(response) {
    const me = response.data;

    localStorage.setItem("me", JSON.stringify(me));
  }).catch(function(error) {
    console.error(error);
  });
};

export default getMyData;