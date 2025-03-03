import React, { useState } from "react";
import styles from "./Main.module.css";
import Header from "../components/header";
import CakeCard from "../components/cake_card";

function Main() {
    const cakes = [
        { id: 1, image: "/assets/1.png", name: "Клубничный торт", price: "1200₽" },
        { id: 2, image: "/assets/2.png", name: "Шоколадный торт", price: "1500₽" },
        { id: 3, image: "/assets/3.png", name: "Ванильный торт", price: "1400₽" }
    ];

    return (
        <div className={styles.main_bg}>
            <Header />
            <div className={styles.text}>
                <h1 style={{ fontSize: "64px" }}>Популярные новинки</h1>
                <h1 style={{ fontSize: "40px" }}>Попробуйте наши самые новые и злые тортики</h1>
            </div>
            <div className={styles.cakes}>
                {cakes.map(cake => (
                    <CakeCard key={cake.id} image={cake.image} name={cake.name} price={cake.price} />
                ))}
            </div>
        </div>
    );
}

export default Main;