import React, { useState } from "react"
import styles from "./cart_item.module.css"

function CartItem({ id, image, name, price, count, updateCount }) {
    return (
        <div className={styles.bg}>
            <img src={image} style={{ borderRadius: "20px", scale: "0.6" }} alt={name} />
            <p className={styles.name}>{name}</p>
            <p className={styles.price}>{price}₽</p>
            <div className={styles.cake_counter}>
                <button className={styles.counter_button} onClick={() => updateCount(id, count + 1)}>▲</button>
                <span className={styles.counter_number}>{count}</span>
                <button className={styles.counter_button} onClick={() => count > 1 && updateCount(id, count - 1)}>▼</button>
            </div>
        </div>
    );
}

export default CartItem;