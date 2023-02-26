import React, { useState } from "react";
import classes from "./Registration.module.css";
import axios from "axios";

const redirect = () => {
    window.location.href = '/books'
}

function Registration(...props){
    const [fullname, setFullname] = useState()
    const [email, setEmail] = useState()
    const [password, setPassword] = useState()

    const HandleSubmit = async (e) =>{
        e.preventDefault()

        console.log("*name: " + fullname + " | @email: " + email + " | #pwd: "+ password);

        const options = {
            method: 'POST',
            url: 'http://localhost:3001/api/v1/auth/signup',
            body: {
                "fullname": fullname,
                "email": email,
                "password": password
            }
        };

        try {
            console.log("trying to POST /auth/signup");

            const response = await axios.post(options.url,
              JSON.stringify({fullname, email, password}),
              {
                  headers: { "Content-Type": "application/json" },
                  body: options.body
              }
            )

            console.log(response.data);

            const token = response.data.data.token
            const ID = response.data.data.userID
            localStorage.setItem("ID", ID)
            localStorage.setItem("token", token)

            if (token){
                redirect()
            }

            setFullname('')
            setEmail('')
            setPassword('')
        } catch (err) {
            if (!err?.response){
                console.log("No response from a server");
            } else
                console.log("Another error!");
        }
    }



    return (
        <div className={classes.regWrapper}>
            <form onSubmit={HandleSubmit}>
                <label for="fullname">Full Name</label>
                <input className={classes.AuthInput} type="text" name="fullname"
                   onChange={(e) => {
                       setFullname(e.target.value);
                       console.log(fullname);
                   }}/>

                <label for="email">Email</label>
                <input className={classes.AuthInput} type="email" name="email"
                   onChange={(e) => {
                       setEmail(e.target.value);
                       console.log(email);
                   }}/>

                <label for="password">Password</label>
                <input className={classes.AuthInput} type="password" name="password"
                   onChange={(e) => {
                       setPassword(e.target.value);
                       console.log(password);
                   }}/>

                <div className={classes.SignIN}>
                    <p className={classes.Text}>Already have an account? &nbsp;</p>
                    <a href={"/log"} className={classes.Link}>Sign In</a>
                </div>

                <input  className={classes.btn} type="submit" name="submit" value="Submit"/>
            </form>
        </div>
    )
}

export default Registration