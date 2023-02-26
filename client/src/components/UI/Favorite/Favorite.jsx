import React from "react";
import classes from "./Favorite.module.css";
import axios from "axios";
import getAllBooks from "../../../helpers/getAllBooks";

function Favorite(props){
    getAllBooks()

    async function UpdateBook(e) {
        e.preventDefault()

        const ID = localStorage.getItem("ID");
        const token = localStorage.getItem("token")

        const options = {
            method: "PUT",
            url: "http://localhost:3001/api/v1/booktabs/update",
            headers: {
                'Authorization': "Bearer " + token,
            },
            body: {
                "user_id": parseInt(ID),
                "book_id": props.book.id,
                "finished": true
            }
        };

        console.log(options.body)

        await axios.put(
            options.url,
            options.body
            ,
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
            <div className={classes.Buttons}>
                <button onClick={UpdateBook}>Finish</button>
                <button onClick={deleteBook}>Delete</button>
            </div>
        </div>
    )
}

export default Favorite;