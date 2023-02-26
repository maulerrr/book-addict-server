import React from 'react'
import classes from './BlockBr.module.css'


function BlockBr() {
  return (
    <div className={classes.blockWrapper}>
        <div className={classes.textWrapper}>
            <h1 className={classes.mainText}>About BookAddict</h1>
            <h3 className={classes.secondaryText}>
            Free web-application to monitor reading habit.
            </h3>
        </div>        
        <div className={classes.imageWrapper}>
            <img src='images/bookshelves.svg'></img>
        </div>
    </div>
  )
}

export default BlockBr