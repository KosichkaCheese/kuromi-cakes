import React from "react";
import styles from "./header.module.css";
import { useState } from "react";
import { Link } from "react-router-dom";

function Header() {
    const [isHovered, setIsHovered] = useState(false);

    return (
        <div className={styles.header}>
            <div className={styles.logo} onMouseEnter={() => setIsHovered(true)} onMouseLeave={() => setIsHovered(false)}>
                <img src="/assets/логотип.png"></img>
                <img src={isHovered ? "/assets/kuromi3.png" : "/assets/kuromi1.png"} style={{ position: "absolute", top: "2vh", left: "23px", scale: "0.8" }}></img>
            </div>
            <div className={styles.navigation}>
                <Link to="/main" className={styles.link}>Главная</Link>
                <Link to="/catalog" className={styles.link}>Каталог</Link>
                <Link to="/about" className={styles.link}>О нас</Link>
                <Link to="/cart" className={styles.link}>Корзина</Link>
            </div>
        </div>
    )
}

export default Header;