import React, { useEffect, useState } from "react";

function Workspace() {
    const [error, setError] = useState(null);
    const [isLoaded, setIsLoaded] = useState(false);
    const [items, setItems] = useState([]);

    
    useEffect(() => {
        fetch("/workspace")
            .then(res => res.json())
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
