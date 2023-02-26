import React from 'react'
import classes from './Navbar.module.css'

export default function navbar() {
  return (
    <div className={classes.navContainer}>
            <div className={classes.logoContainer}>
                <img className={classes.logo} src="images/booklogo.png" alt="Logo"/>
            </div>

            <div className={classes.btns}>
                <a className={classes.navBtn} href='/'>Home</a>
                <a className={classes.navBtn} href='/books'>Books</a>
                <a className={classes.navBtn} href='/booktabs'>Book Tabs</a>
            </div>
            <div className={classes.dropdownWrapper}>
                <a className={classes.dropdownBtn} href='/personal'>
                    Personal
                </a>
            </div>
    </div>
  )
}
