import React, { useEffect } from 'react';
import './App.css';
import {BrowserRouter as Router, Routes, Route} from "react-router-dom"
import Signin from './components/Signin';


const drawerWidth = 240;



function App() {  

  const [token, setToken] = React.useState<String>("");
  const [statustoken, setStatustoken] = React.useState<boolean>(false);

  useEffect(() => {
    const token:any = localStorage.getItem("token")
    if(token) {
      setToken(token)
      validToken()
    }
  }, [])


  if(!token || statustoken){
    console.log(statustoken)
    return <Signin />
  }

  function validToken(){
    fetch("http://localhost:8080/valid", {
      method: "GET",
      headers: {
        'Content-Type': 'application/json',
        "Authorization": `Bearer ${token}`
      }
    })
    .then((res) => res.json())
    .then((data) => {
      console.log(data)
      setStatustoken(true)
    })
    .catch((err) => {
      console.log(err)
      setStatustoken(false)
    })
  }

  function RouteApplication() {
    return(
      <Router>
        <div>
          
        </div>
      </Router>
    )
  }

  return (
    RouteApplication()
  );
}




export default App;
