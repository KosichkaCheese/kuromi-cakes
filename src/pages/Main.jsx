import React, { useEffect, useState } from "react";
import styles from "./Main.module.css";
import Header from "../components/header";
import CakeCard from "../components/cake_card";
import axios from "axios";

function Main() {
    const sale_cakes = [3, 2, 1];
    const [cakes, setCakes] = useState([]);

    useEffect(() => {
        async function fetchCakes() {
            try {
                const cakeRequests = sale_cakes.map((cake) =>
                    axios
                        .get(`http://localhost:8000/cake_api/cakes/${cake}`)
                        .then((res) => ({
                            ...res.data
                        }))
                );

                const cakesData = await Promise.all(cakeRequests);
                setCakes(cakesData);
            } catch (error) {
                console.error("Ошибка загрузки товаров:", error);
            }
        }

        fetchCakes();
    });

    return (
        <div className={styles.main_bg}>
            <Header />
            <div className={styles.text}>
                <h1 style={{ fontSize: "64px" }}>Популярные новинки</h1>
                <h1 style={{ fontSize: "40px" }}>Попробуйте наши самые новые и злые тортики</h1>
            </div>
            <div className={styles.cakes}>
                {cakes.map(cake => (
                    <CakeCard key={cake.id} image={"assets/" + cake.image} name={cake.name} price={cake.price} />
                ))}
            </div>
        </div>
    );
}

export default Main;