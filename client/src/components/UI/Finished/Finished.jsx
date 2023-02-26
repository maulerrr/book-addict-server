import React from "react";
import classes from "./Finished.module.css";
import axios from "axios";
import getAllBooks from "../../../helpers/getAllBooks";

function Finished(props){
    getAllBooks()

    async function deleteBook(e) {
        const ID = localStorage.getItem("ID");
        const token = localStorage.getItem("token")

        const options = {
            method: "DELETE",
            url: "http://localhost:3001/api/v1/booktabs/delete/" + props.book.id,
            headers: {
                'Authorization': "Bearer " + token,
            }
        };

        console.log(props.book.id)

        await axios.delete(
            options.url,
            {
                headers: {
                    "Content-Type": "application/json",
                    'Authorization': options.headers.Authorization
                }
            }
        )
            .then(function (response) {
                console.log(response.data);
            });

    }

    return(
        <div className={classes.exerciseWrapper}>
            <div className={classes.textWrapper}>
                <h3>Book: </h3>
                <p>{props.book.title}</p>
            </div>
            <img alt="image" src={props.book.img}/>
            <div className={classes.textWrapper}>
                <h3>Price: </h3>
                <p>{props.book.price}</p>
            </div>
            <button onClick={deleteBook}>Delete</button>
        </div>
    )
}

export default Finished;