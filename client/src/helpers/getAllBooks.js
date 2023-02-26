import axios from "axios";


async function getAllBooks(){
    const token = (localStorage.getItem("token"))

    const options = {
        method: "GET",
        url: "http://localhost:3001/api/v1/readlist/",
        headers: {
            'Authorization': "Bearer " + token,
        }
    };

    if(localStorage.getItem("token")){
        await axios.request(options).then(function(response) {
            console.log(response.data);

            apiData = response.data

            localStorage.setItem("books", JSON.stringify(apiData));
            console.log("data being set: " + apiData);
        }).catch(function(error) {
            console.error(error);
        });
    }
}


export default getAllBooks;
export var  apiData;


