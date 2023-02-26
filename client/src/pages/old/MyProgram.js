import React, { useState } from "react";
import UserDetails from "../../components/UI/userDetails/UserDetails";
import getMyData from "../../helpers/getMyData";
import getMyFavorites from "../../helpers/getMyFavorites";
import getFinished from "../../helpers/getFinished";
import login from "../../components/UI/login-form/Login";

const reload = () => {
  window.location.href = '/program'
}

const redirect = () => {
  window.location.href = '/log'
}

const personal = () => {
  window.location.href = '/personal'
}

export const MyProgram = () => {
  const [loading, setLoading] = useState(true);

  let renderUserDetails = false
  let username = null
  getFinished().then(r => {
    console.log("Got activities");
  })

  getMyData().then(r => {
    if (JSON.parse(localStorage.getItem("me")).program != null){
      getMyFavorites().then(r => {
        console.log("Program already exists, got it successfully!");
      })
    } else personal()
  })

  if (JSON.parse(localStorage.getItem("me"))) {
    renderUserDetails = !(JSON.parse(localStorage.getItem("me")).metrics == null);
    username = JSON.parse(localStorage.getItem("me")).name;
  }

  function convertMS(ms) {
    let d, h, m, s;
    s = Math.floor(ms / 1000);
    m = Math.floor(s / 60);
    s = s % 60;
    h = Math.floor(m / 60);
    m = m % 60;
    d = Math.floor(h / 24);
    h = h % 24;

    return d;
  }

  let data = function getDetails(){
    let showDetails = renderUserDetails

    let programCalories = parseFloat(JSON.parse(localStorage.getItem("program")).calories).toFixed(1);
    let activityLevel = JSON.parse(localStorage.getItem("activities")).find((activity) => {
      return activity._id === JSON.parse(localStorage.getItem("me")).metrics.activity;
    }).value;

    let currentWeight = JSON.parse(localStorage.getItem("me")).metrics.weight;
    let goalWeight = JSON.parse(localStorage.getItem("me")).metrics.goalWeight;

    let sd = new Date(JSON.parse(localStorage.getItem("me")).metrics.startDate);
    let fd = new Date(JSON.parse(localStorage.getItem("me")).metrics.finishDate);
    let timeRemaining = (convertMS(fd.getTime() - sd.getTime()));

    return {
      showDetails: showDetails,
      programCalories: programCalories,
      activityLevel: activityLevel,
      currentWeight: currentWeight,
      goalWeight: goalWeight,
      timeRemaining: timeRemaining
    }
  }

  window.onload = (async (e) => {
    if (localStorage.getItem("token")) {
      if (!(JSON.parse(localStorage.getItem("me")).metrics)) {
        await getMyFavorites();
      }
    } else {
      return redirect();
    }
  });


  return (
    <div className="App">
      { !renderUserDetails? <h1>Welcome {username}! Please fill out the <a href={"/personal"}> metrics(clickable)</a></h1> :
        <UserDetails details={{
          username: username,
          calories: data().programCalories,
          activity: data().activityLevel,
          currentWeight: data().currentWeight,
          goalWeight: data().goalWeight,
          remainingTime: data().timeRemaining,
          shouldReload: !data().showDetails
        }} />}
    </div>
  );
};