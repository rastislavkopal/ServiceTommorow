import React, {useState} from 'react'
import {Link, Redirect} from "react-router-dom"

const Register = () => {

    const [name, setName] = useState('');
    const [email, setEmail] = useState('');
    const [password, setPassword] = useState('');
    const [confirmPassword, setConfirmPassword] = useState('');
    const [redirect, setRedirect] = useState(false)

    const submitRegister = async (e) => {
        e.preventDefault();

        await fetch("/auth/register", {
            method: 'POST',
            headers: {"content-type": "application/json"},
            body: JSON.stringify({
                name,
                email,
                password,
                confirmPassword,
            })
        })

        setRedirect(true);
    }

    if (redirect) {
        return <Redirect to="/login" />
    }

    return (
        <form className="bg-grey-lighter min-h-full mt-44 flex flex-col" onSubmit={submitRegister}>
            <div className="container max-w-sm mx-auto flex-1 flex flex-col items-center justify-center px-2">
                <div className="bg-white px-6 py-8 rounded shadow-md text-black w-full">
                    <h1 className="mb-8 text-3xl text-center">Sign up</h1>
                    <input 
                        type="text"
                        className="block border border-grey-light w-full p-3 rounded mb-4"
                        name="fullname"
                        placeholder="Full Name" onChange={ e => setName(e.target.value)}/>

                    <input 
                        type="email"
                        className="block border border-grey-light w-full p-3 rounded mb-4"
                        name="email"
                        placeholder="Email" required onChange={ e => setEmail(e.target.value)}/>

                    <input 
                        type="password"
                        className="block border border-grey-light w-full p-3 rounded mb-4"
                        name="password"
                        placeholder="Password" required onChange={ e => setPassword(e.target.value)}/>
                    <input 
                        type="password"
                        className="block border border-grey-light w-full p-3 rounded mb-4"
                        name="confirm_password"
                        placeholder="Confirm Password" required onChange={ e => setConfirmPassword(e.target.value)}/>

                    <button type="submit" className="w-full text-center py-3 rounded bg-green-500 text-white hover:bg-green-700 focus:outline-none my-1">Create Account</button>

                    <div className="text-center text-sm text-grey-dark mt-4">
                        By signing up, you agree to the 
                        <a className="no-underline border-b border-grey-dark text-grey-dark" href="#0">
                            Terms of Service
                        </a> and 
                        <a className="no-underline border-b border-grey-dark text-grey-dark" href="#0">
                            Privacy Policy
                        </a>
                    </div>
                </div>

                <div className="text-grey-dark mt-6">
                    Already have an account? 
                    <Link to="/login" className="no-underline border-b border-blue text-blue">
                        Log in
                    </Link>.
                </div>
            </div>
        </form>
    )
}

export default Register
