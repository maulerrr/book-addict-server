import React, { useEffect} from "react";
import classes from "./UserDetails.module.css";

function UserDetails(props) {
  if (props.details.shouldReload) {
    window.location.href = "/personal"
  }

  return (
    <div className={classes.DetailsContainer}>
    <div className={classes.Wrapper}>
      <h1 className={classes.MainTxt}>Hello, {props.details.username}!</h1>
      <p><b>Activity level:</b> {props.details.activity}</p>
      <p><b>Current weight:</b> {props.details.currentWeight}</p>
      <p><b>Your goal:</b> {props.details.goalWeight}</p>
    </div>
    {/*<div className={classes.Wrapper}>*/}
    {/*  <h1 className={classes.Greetings}>Your goals:</h1>*/}
    {/*  <p><b>Weight:</b> {props.details.goalWeight}</p>*/}
    {/*  <p><b>Finish:</b> {props.details.remainingTime} days</p>*/}
    {/*</div>*/}
    <div className={classes.WrapperV}>
      <div className={classes.Calories}>
        <h1 className={classes.MainTxt}>{props.details.calories}</h1>
        <p>Calories/Day</p>
      </div>
      <div className={classes.TimeRemaining}>
        <h1 className={classes.MainTxt}>{props.details.remainingTime}</h1>
        <p>Days left</p>
      </div>
      <a className={classes.ExerciseBtn} href={"/book"}>
        Start exercising
      </a>
    </div>
    </div>
  );
}

export default UserDetails;