import React, { useState } from "react";
import classes from "./Login.module.css";
import axios from "axios";
import getMyData from "../../../helpers/getMyData";
import getMyFavorites from "../../../helpers/getMyFavorites";

const redirect = () => {
  window.location.href = '/personal'
}

const program = () => {
  window.location.href = '/program'
}

function Login(){
  const [email, setEmail] = useState("")
  const [password,setPassword] = useState("")

  const HandleSubmit = async (e) =>{
    e.preventDefault()

    console.log("@email: " + email + " | #pwd: "+ password);

    const options = {
      method: 'POST',
      url: 'http://localhost:3001/api/v1/auth/login',
      body: {
        "email": email,
        "password": password
      }
    };

    try {
      console.log("trying to POST /auth/login");

      const response = await axios.post(options.url,
        JSON.stringify({email, password}),
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
        await getMyData().then(()=>{
          if (JSON.parse(localStorage.getItem("me")).program){
            console.log("trying to redirect to /program");
            program()
          } else redirect()
        })

        // redirect()
      }

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
    <div className={classes.loginWrapper}>
      <form onSubmit={HandleSubmit}>
          <label htmlFor="email">Email</label>
          <input className={classes.AuthInput} type="email" name="email"
          onChange={(e) => {
            setEmail(e.target.value);
          }}/>

          <label htmlFor="password">Password</label>
          <input className={classes.AuthInput} type="password" name="password"
          onChange={((e) => {
            setPassword(e.target.value);
          })}/>

          <div className={classes.SignUP}>
            <p className={classes.Text}>Don't have an account? &nbsp;</p>
            <a href={"/reg"} className={classes.Link}>Sign Up</a>
          </div>
          <input className={classes.btn} type="submit" name="submit" value="Sign In" />
      </form>
    </div>
  )
}

export default Login;