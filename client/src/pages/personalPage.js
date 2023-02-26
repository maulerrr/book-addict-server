import React, { useEffect, useState } from "react";
import axios from "axios";
import getMyData from "../helpers/getMyData"
import getMyFavorites from "../helpers/getMyFavorites"
import getFinished from "../helpers/getFinished";
import Favorite from "../components/UI/Favorite/Favorite";
import Finished from "../components/UI/Finished/Finished";
import getAllBooks from "../helpers/getAllBooks";

const redirect = () => {
  window.location.href = '/log'
}

function PersonalPage() {
  if (!localStorage.getItem("token")){
    redirect()
  }

  getAllBooks()
  getFinished()
  getMyFavorites()
  // const [favorites, setFavorites] = useState("Add some")
  // const [finished, setFinished] = useState("Add some")

  const allBooks = JSON.parse(localStorage.getItem("books"))
  const favorites = JSON.parse(localStorage.getItem("favorites"))
  const finished = JSON.parse(localStorage.getItem("finished"))

  const favoriteBooks = allBooks.filter((book) => {
    return favorites.some((fav) => fav.book_id === book.id);
  });

// Filter the allBooks array to get an array of all finished books
  const finishedBooks = allBooks.filter((book) => {
    return finished.some((fin) => fin.book_id === book.id && fin.finished);
  });


  getFinished().then(r => {
    console.log("Got finished books");
  });

  useEffect(((e) => {
    if (localStorage.getItem("token")) {
      getMyData();
      getAllBooks()
      getFinished();
      getMyFavorites();
    } else {
      return redirect();
    }
  }));


  return (
    <div className="Personal">
      <div><h1>Favorites:</h1> { favoriteBooks.map((book, index) => {
        return <Favorite book={book} key={index}/>})}</div>
      <div><h1>Finished:</h1> { finishedBooks.map((book, index) => {
        console.log(book)
        return <Finished book={book} key={index}/>})}</div>
    </div>
  );
}

export default PersonalPage;
