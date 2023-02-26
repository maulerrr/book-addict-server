import React from 'react'
import classes from './LetsStart.module.css'

export default function LetsStart() {
  return (
    <div className={classes.actionContainer}>
        <div className={classes.actionText}>
            <p className={classes.mainText}>
              Sign up to access to library!
            </p>
            <p className={classes.secondaryText}>
            Over 100 famous books
            </p>
        </div>
        <div className={classes.actionBtnWrapper}>
            <a className={classes.actionBtn} href='/personal'>
                <h1>Let's start your journey!</h1>
                <img src='images/arrow.svg' className={classes.arrow} ></img>
            </a>
        </div>
    </div>
  )
}
