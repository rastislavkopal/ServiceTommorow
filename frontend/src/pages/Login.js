import React, {useState} from 'react'
import {Link, Redirect} from "react-router-dom"
import { useCookies } from 'react-cookie';

const Login = () => {

    const [email, setEmail] = useState('');
    const [password, setPassword] = useState('');
    const [redirect, setRedirect] = useState(false)
    const [cookies, setCookie] = useCookies(['jwt_token']);

    const submitLogin = async (e) => {
        e.preventDefault();

        await fetch("/auth/login", {
            method: 'POST',
            headers: {"content-type": "application/json"},
            credentials: 'include',
            body: JSON.stringify({
                email,
                password,
            })
        })
        .then(r => {
            if (!r.ok) {
                throw Error(r.statusText);
            }
            return r.json()
        } )
        .then(response => {
            setCookie("jwt_token", response.Tokens.access_token,{
                email: response.email,
                expires_at: response.Tokens.at_expires,
                creted_at: response.Tokens.CreatedAt,
            })
            console.log(cookies)
            setRedirect(true);

        }).catch(function(err) {
            console.log("error: ");
        });

        
    }

    if (redirect) {
        return <Redirect to="/dashboard" />
    }

    return (
        <form className="bg-grey-lighter min-h-full mt-44 flex flex-col" onSubmit={submitLogin}>
            <div className="container max-w-sm mx-auto flex-1 flex flex-col items-center justify-center px-2">
                <div className="bg-white px-6 py-8 rounded shadow-md text-black w-full">
                    <h1 className="mb-8 text-3xl text-center">Sign in</h1>

                    <input 
                        type="text"
                        className="block border border-grey-light w-full p-3 rounded mb-4"
                        name="email"
                        placeholder="Email" required onChange={ e => setEmail(e.target.value)} />

                    <input 
                        type="password"
                        className="block border border-grey-light w-full p-3 rounded mb-4"
                        name="password"
                        placeholder="Password" required onChange={ e => setPassword(e.target.value)}/>

                    <button type="submit" className="w-full text-center py-3 rounded bg-green-500 text-white hover:bg-green-700 focus:outline-none my-1">Log in</button>

                </div>

                <div className="text-grey-dark mt-6">
                    Do not have an account yet?  &nbsp;
                    <Link to="/register" className="no-underline border-b border-blue text-blue">
                        Sign up
                    </Link>.
                </div>
            </div>
        </form>
    )
}

export default Login
