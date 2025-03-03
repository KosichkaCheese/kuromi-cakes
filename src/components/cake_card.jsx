import React from "react";
import styles from "./cake_card.module.css";

function CakeCard(props) {
    return (
        <div className={styles.cake_card}>
            <img src={props.image} alt={props.name} style={{ borderRadius: "20px", scale: "0.8" }} />
            <p className={styles.cake_name}>{props.name}</p>
            <p className={styles.cake_price}>{props.price}</p>
            <button className={styles.add_to_cart} onClick={() => alert("Торт добавлен в корзину")}>+</button>
        </div>
    );
}

export default CakeCard;