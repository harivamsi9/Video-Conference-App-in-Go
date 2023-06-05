

import React from "react";
import { useState } from "react";

import { BrowserRouter, Route, Routes } from "react-router-dom";
import Room from './components/Room';
import CreateRoom from './components/CreateRoom';

function App() {
  
  return <div className='App'>
      <BrowserRouter>
        <Routes>
          <Route path="/" exact Component={ CreateRoom }></Route>
          <Route path="/room/:room_id" Component={Room}></Route>
        </Routes>      
      </BrowserRouter>
    </div>;
}

export default App
