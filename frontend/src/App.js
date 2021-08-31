import './App.css';
import Workspace from "./components/Workspace";
import AppNav from "./components/AppNav"

require('dotenv').config()

function App() {
  return (
    <div className="App">
      <AppNav />
      <header className="App-header" class="text-blue-400">
        Hi TODO list
      </header>
      <Workspace />
    </div>
  );
}

export default App;
