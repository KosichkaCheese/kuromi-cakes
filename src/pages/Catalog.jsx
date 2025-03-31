import { useEffect, useState } from "react";
import axios from "axios";
import styles from "./Catalog.module.css";
import Header from "../components/header";
import CakeCard from "../components/cake_card";

function Catalog() {
    const [cakes, setCakes] = useState([]);
    const [loading, setLoading] = useState(true);
    const [error, setError] = useState(null);

    useEffect(() => {
        axios.get("http://localhost:8000/cake_api/cakes")
            .then(response => {
                setCakes(response.data);
                setLoading(false);
            })
            .catch(error => {
                setError(error.message);
                setLoading(false);
            });
    }, []);

    if (loading) return <p>Загрузка каталога</p>;
    if (error) return <p>Ошибка загрузки каталога: {error}</p>;



    return (
        <div className={styles.catalog_bg}>
            <Header />
            <div className={styles.catalog}>
                {cakes.map(cake => (
                    <CakeCard key={cake.id} id={cake.id} image={"assets/" + cake.image} name={cake.name} price={cake.price} />
                ))}
            </div>
        </div>
    )
}

export default Catalog;