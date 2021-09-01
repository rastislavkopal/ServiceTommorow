import React from 'react';
// import Register from "./Register";
import {Link} from "react-router-dom"

function removeToken() {
    localStorage.removeItem("jwt_token");
}

const AppNav = () => {

    let menu;
    if (!localStorage.getItem("jwt_token")) {
        menu = (
                <ul className="list-reset lg:flex justify-end flex-1 items-center">
                    <li className="mr-3">
                        <Link to="/login" className="inline-block text-green-600 no-underline hover:text-green-700 hover:text-underline font-bold py-2 px-4">Log in</Link>
                    </li>
                    <li className="mr-3">
                        <Link to="/register" className="inline-block text-green-600 no-underline hover:text-green-700 hover:text-underline font-bold py-2 px-4">Register</Link>
                    </li>
                </ul>
        )
    } else {
        menu = (
            <ul className="list-reset lg:flex justify-end flex-1 items-center">
                    <li className="mr-3">
                        <Link to="/dashboard" className="inline-block text-green-600 no-underline hover:text-green-700 hover:text-underline font-bold py-2 px-4">Dashboard</Link>
                    </li>
                    <li className="mr-3">
                        <Link to="/home" onClick={removeToken} className="inline-block text-green-600 no-underline hover:text-green-700 hover:text-underline font-bold py-2 px-4">Logout</Link>
                    </li>
                </ul>
        )
    }

    return (
        <nav className="flex items-center justify-between flex-wrap bg-gray-800 p-6 w-full z-10 top-0">
            <div className="flex items-center flex-shrink-0 text-white mr-6">
                <Link className="text-white no-underline hover:text-white hover:no-underline" to="/">
                    <span className="text-2xl pl-2"><i className="em em-grinning"></i>{process.env.REACT_APP_TITLE}</span>
                </Link>
            </div>

            <div className="block lg:hidden">
                <button id="nav-toggle" className="flex items-center px-3 py-2 border rounded text-gray-500 border-gray-600 hover:text-white hover:border-white">
                    <svg className="fill-current h-3 w-3" viewBox="0 0 20 20" xmlns="http://www.w3.org/2000/svg"><title>Menu</title><path d="M0 3h20v2H0V3zm0 6h20v2H0V9zm0 6h20v2H0v-2z"/></svg>
                </button>
            </div>

            <div className="w-full flex-grow lg:flex lg:items-center lg:w-auto hidden lg:block pt-6 lg:pt-0" id="nav-content">
                {menu}
            </div>
        </nav>
    )
}

export default AppNav
