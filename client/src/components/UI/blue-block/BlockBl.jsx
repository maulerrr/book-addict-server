import React from 'react'
import classes from './BlockBl.module.css'

function BlockBl() {
  return (
    <div className={classes.blockWrapper}>
        <div className={classes.imageWrapper}>
            <img src='images/literature.svg'/>
        </div>
        <div className={classes.textWrapper}>
            <h1 className={classes.mainText}>Save your</h1>
            <h3 className={classes.secondaryText}>
                1 favorites
            </h3>
            <h3 className={classes.secondaryText}>
                2 book u already read
            </h3>
        </div>
        
    </div>
  )
}

export default BlockBl