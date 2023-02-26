import React from "react";
import classes from "./Exercise.module.css";
import axios from "axios";

function Book(props){

    async function addBookToFavorites() {
        const ID = localStorage.getItem("ID");
        const token = localStorage.getItem("token")

        const options = {
            method: "POST",
            url: "http://localhost:3001/api/v1/booktabs/add",
            headers: {
                'Authorization': "Bearer " + token,
            },
            body: {
                "user_id": parseInt(ID),
                "book_id": props.book.id,
                "finished": false
            }
        };


        await axios.post(
            options.url,
            options.body,
            {
                headers: {
                    "Content-Type": "application/json",
                    'Authorization': options.headers.Authorization
                },
                body: options.body
            }
            )
            .then(function (response) {
            console.log(response.data);
            const favorites = response.data;
            localStorage.setItem("favorites", JSON.stringify(favorites));
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
            <button onClick={addBookToFavorites}>Read</button>
        </div>
    )
}

export default Book;