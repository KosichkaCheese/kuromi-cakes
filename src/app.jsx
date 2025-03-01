import React, { useState } from "react";
import styles from "./App.module.css";

function Home({ navigate }) {
    const [isHovered, setIsHovered] = useState(false);

    return (
        <div className={styles.home_bg}>
            <div style={{ margintop: "2%", marginleft: "3%" }}><img width={"600px"} src="assets/Тортики Куроми.png"></img></div>
            <div style={{ marginLeft: "auto", marginRight: "2%", marginTop: "-4%" }}><img width={"500px"} src="assets/Ешьте тортики.png"></img></div>
            <div style={{ marginTop: "-4%", alignSelf: "center", marginLeft: "-5%" }}>
                <img style={{ position: "relative", top: "-10px", left: "70px", rotate: "-20deg", scale: "1.3" }} src={isHovered ? "assets/kuromi2.png" : "assets/kuromi1.png"}></img>
                <button className={styles.want_cake} onClick={() => navigate("catalog")} onMouseEnter={() => setIsHovered(true)} onMouseLeave={() => setIsHovered(false)}>хочу тортик</button>
            </div>
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
