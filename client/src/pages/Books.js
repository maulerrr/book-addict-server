import React from "react";
import Book from "../components/UI/book/Book";
import getAllBooks, {apiData} from "../helpers/getAllBooks";

console.log("Saved data from api: " + apiData)

function BooksPage() {
    getAllBooks()
    const fallback = JSON.parse(localStorage.getItem("books"))
    console.log("Fallback: " + fallback)

    return (
        <div className="Books">
            {   !fallback ?
                "Nothing here"
                :
                fallback.map((book, index) => {
                return <Book book={book} key={index}/>})
            }
        </div>
    );

}

export default BooksPage;
