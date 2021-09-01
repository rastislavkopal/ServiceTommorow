import React from 'react'
import Workspace from "../components/Workspace";
import {Redirect} from "react-router-dom"


const Dashboard = () => {

    if (!localStorage.getItem("jwt_token")) {
        return <Redirect to="/home" />
    }

    return (
        <div>
            <h1>Dashboard</h1>
            <Workspace />
        </div>
    )
}

export default Dashboard
