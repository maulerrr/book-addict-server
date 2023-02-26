import Motto from "../components/UI/motto/Motto";
import BlockBl from "../components/UI/blue-block/BlockBl";
import BlockBr from "../components/UI/brown-block/BlockBr";
import LetsStart from "../components/UI/lets-start/LetsStart";
import React from "react";


function Home(){
    return (
        <div className="App">
            <Motto/>
            <BlockBl/>
            <BlockBr/>
            <LetsStart />
        </div>
    )
}

export default Home;