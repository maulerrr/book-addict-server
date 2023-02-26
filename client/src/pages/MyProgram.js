import React, { useState, useEffect } from "react";
import UserDetails from "../components/UI/userDetails/UserDetails";
import getMyFavorites from "../helpers/getMyFavorites";
import getFinished from "../helpers/getFinished";
import getMyData from "../helpers/getMyData";
import LoadingIndicator from "../components/UI/loading-indicator/LoadingIndicator";

const MyProgram = () => {
  const [loading, setLoading] = useState(true);

  const [showDetails, setShowDetails] = useState(false);
  const [username, setUsername] = useState(null);
  const [programCalories, setProgramCalories] = useState(null);
  const [activityLevel, setActivityLevel] = useState(null);
  const [currentWeight, setCurrentWeight] = useState(null);
  const [goalWeight, setGoalWeight] = useState(null);
  const [remainingTime, setRemainingTime] = useState(null);

  const convertMS = (ms) => {
    let d, h, m, s;
    s = Math.floor(ms / 1000);
    m = Math.floor(s / 60);
    s = s % 60;
    h = Math.floor(m / 60);
    m = m % 60;
    d = Math.floor(h / 24);
    h = h % 24;

    return d;
  };

  useEffect(() => {
    const fetchData = async () => {
      if (!localStorage.getItem("token")) {
        window.location.href = "/log"
        return;
      }

      const me = JSON.parse(localStorage.getItem("me"));
      if (!me.metrics) {
        window.location.href = "/personal"
      }

      setLoading(true); // Set loading to true before fetching data

      await getFinished()
      await getMyFavorites()
      await getMyData()

      const activities = JSON.parse(localStorage.getItem("activities"));
      const activity = activities.find(
        (activity) => activity._id === me.metrics.activity
      ).value;

      setUsername(me.name);
      setCurrentWeight(me.metrics.weight);
      setGoalWeight(me.metrics.goalWeight);
      setActivityLevel(activity);

      const program = JSON.parse(localStorage.getItem("me")).program;
      setProgramCalories(parseFloat(program.calories).toFixed(1));

      const startDate = new Date(me.metrics.startDate);
      const finishDate = new Date(me.metrics.finishDate);
      const timeRemaining = convertMS(finishDate.getTime() - startDate.getTime());
      setRemainingTime(timeRemaining);

      setShowDetails(true);
      setLoading(false); // Set loading to false after fetching data
    };

    fetchData();
  }, []);

  const handleReload = () => {
    window.history.push("/personal");
  };

  return (
    <div className="App">
      {loading ? ( // Show loading indicator if loading is true
        <LoadingIndicator />
      ) : !showDetails ? (
        <h1>
          Welcome {username}! Please fill out the{" "}
          <a href={"/personal"}>metrics(clickable)</a>
        </h1>
      ) : (
        <UserDetails
          details={{
            username,
            calories: programCalories,
            activity: activityLevel,
            currentWeight,
            goalWeight,
            remainingTime,
            shouldReload: !showDetails,
            onReload: handleReload,
          }}
        />
      )}
    </div>
  );
};

export default MyProgram;
