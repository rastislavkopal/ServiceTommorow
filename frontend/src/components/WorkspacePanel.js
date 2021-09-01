import React, { useEffect, useState } from "react";

function WorkspacePanel() {
    const [error, setError] = useState(null);
    const [isLoaded, setIsLoaded] = useState(false);
    const [items, setItems] = useState([]);

    useEffect(() => {
        fetch("/workspace",{
            credentials: 'include',
            headers: {
                "Authorization": "Bearer " + localStorage.getItem("jwt_token")
            }
        })
            .then(res => {
                if (!res.ok) {
                    throw Error(res.statusText);
                }
                return res.json()
            } )
            .then(
                (result) => {
                    setIsLoaded(true);
                    setItems(result.data.rows)
                },
                (error) => {
                    setIsLoaded(true);
                    setError(error)
                }
            )
    }, [])
   

    if (error) {
        return <div>Error: {error.message}</div>
    } else if (!isLoaded) {
        return <div>Loading..</div>;
    } else {
        return (
        <div className="grid grid-cols-11 h-auto">
            <div className="py-4 bg-green-500 col-span-3">
                {items.map(item => (
                    <div className="bg-red-500 mx-5 my-5 w-11/12 px-5 py-5" key={item.id}>
                        <img src="https://source.unsplash.com/random/60x60" alt="random workspace im" className="float-left mr-2 my-2"/>
                        <span className="text-white font-bold text-2xl">{item.title}</span><br/>
                        <span className="w-full">{item.description}</span>
                    </div>
                ))}
            </div>
            <div className="py-8 bg-red-500 col-span-8"></div>
        </div>
        )
    }
}

export default WorkspacePanel
