import './App.css';
// import Workspace from "./components/Workspace";
import AppNav from "./components/AppNav";
import {BrowserRouter, Route} from "react-router-dom"
import Home from "./pages/Home"
import Login from "./pages/Login"
import Register from "./pages/Register"

require('dotenv').config()

function App() {
  return (
    <div className="App">
      
      <BrowserRouter>
        <AppNav />
        <Route path="/" exact component={Home} />
        <Route path="/login" component={Login} />
        <Route path="/register" component={Register} />
      </BrowserRouter>

      {/* <Workspace /> */}
    </div>
  );
}

export default App;
