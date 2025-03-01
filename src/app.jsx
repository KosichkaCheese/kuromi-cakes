import React, { useState } from "react";
import styles from "./App.module.css";

function Home({ navigate }) {
    return (
        <div className={styles.home_bg}>
            <h1 className={styles.title}>Тортики Куроми</h1>
            <h2 className={styles.title}>Ешьте тортики</h2>
            <button onClick={() => navigate("catalog")}>Перейти в каталог</button>
        </div>
    );
}

function Catalog({ navigate }) {
    const cakes = [
        { id: 1, name: "Клубничный торт Куроми", price: "1200₽" },
        { id: 2, name: "Шоколадный торт Куроми", price: "1500₽" }
    ];

    return (
        <div style={{ textAlign: "center", padding: "20px" }}>
            <button onClick={() => navigate("home")}>Главная</button>
            <h2>Каталог тортиков</h2>
            {cakes.map(cake => (
                <div key={cake.id}>
                    <h3>{cake.name}</h3>
                    <p>{cake.price}</p>
                </div>
            ))}
        </div>
    );
}

export default function App() {
    const [page, setPage] = useState("home");
    return page === "home" ? <Home navigate={setPage} /> : <Catalog navigate={setPage} />;
}
