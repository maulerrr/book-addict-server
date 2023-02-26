import axios from "axios";

const getMyFavorites = async () => {
  const ID = localStorage.getItem("ID");
  const token = localStorage.getItem("token")

  const options = {
    method: "GET",
    url: "http://localhost:3001/api/v1/booktabs/favorites/" + ID,
    headers: {
      'Authorization': "Bearer " + token,
    }
  };

  await axios.request(options).then(function(response) {
    console.log(response.data);
    const favorites = response.data;
    localStorage.setItem("favorites", JSON.stringify(favorites));
  });
};

export default getMyFavorites;