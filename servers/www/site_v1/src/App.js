import './App.css';
import {BrowserRouter as Router, Route} from 'react-router-dom'
import Navbar from "./components/Nav/Navbar";
import Jobs from "./components/Jobs/Jobs";
import Scan from "./components/Scan/Scan";


export default function App() {
    return (
        <div className="container">
            <Router>
                <Navbar/>

                <Route exact strict path="/">
                    <Jobs/>
                </Route>

                <Route path="/jobs">
                    <Jobs/>
                </Route>

                <Route path="/scan">
                    <Scan/>
                </Route>

            </Router>
        </div>
    );
}
