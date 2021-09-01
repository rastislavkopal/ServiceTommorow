import React, { useEffect, useState } from "react";

function Workspace() {
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
        console.log(items)
        return (
            <div>
                <h1>Workspaces:</h1>
                <ul>
                    {items.map(item => (
                        <li key={item.id}>
                        {item.title} { " ---->" } {item.description}
                        </li>
                    ))}
                </ul>
            </div>
        )
    }
}

export default Workspace
