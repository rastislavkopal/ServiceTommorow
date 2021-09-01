import React from 'react'
import WorkspacePanel from "../components/WorkspacePanel";
import {Redirect} from "react-router-dom"


const Dashboard = () => {

    if (!localStorage.getItem("jwt_token")) {
        return <Redirect to="/home" />
    }

    return (
        <div>
            <h1>Dashboard</h1>
            <WorkspacePanel />
        </div>
    )
}

export default Dashboard
