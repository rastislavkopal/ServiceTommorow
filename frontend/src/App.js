import './App.css';
import AppNav from "./components/AppNav";
import {BrowserRouter, Route} from "react-router-dom"
import Home from "./pages/Home"
import Login from "./pages/Login"
import Register from "./pages/Register"
import Dashboard from './pages/Dashboard';
import { CookiesProvider } from "react-cookie";

require('dotenv').config()

function App() {

  return (
    <CookiesProvider>
      <div className="App">
        
        <BrowserRouter>
          <AppNav />
          <Route path="/" exact component={Home} />
          <Route path="/login" component={Login} />
          <Route path="/register" component={Register} />
          <Route path="/dashboard" component={Dashboard} />
        </BrowserRouter>
      </div>
    </CookiesProvider>
  );
}

export default App;
