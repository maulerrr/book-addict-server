import React from 'react'
import {BrowserRouter, Routes, Route} from "react-router-dom";
import Home from "./pages/Home";
import BooksPage from "./pages/Books";
import PersonalPage from "./pages/personalPage";
import RegistrationPage from "./pages/RegistrationPage";
import LoginPage from "./pages/LoginPage";
import MyProgram  from "./pages/MyProgram";

function App() {
  return (
      <BrowserRouter>
          <Routes>
              <Route path='/' element={<Home/>} />
              <Route path='/books' element={<BooksPage/>}/>
              <Route path='/personal' element={<PersonalPage/>}/>
              <Route path='/reg' element={<RegistrationPage/>}/>
              <Route path='/log' element={<LoginPage/>}/>
              <Route path='/booktabs' element={<MyProgram/>}/>
          </Routes>
      </BrowserRouter>

  );
}

export default App;
