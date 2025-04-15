import React, { useState } from "react"
import styles from "./cart_item.module.css"

function CartItem({ id, image, name, price, count, addToCart, removeFromCart, clearItem }) {
    const handleDecrease = () => {
        if (count <= 1) {
            clearItem(id);  // Если количество меньше или равно 1, удаляем товар из корзины
        } else {
            removeFromCart(id);  // Иначе просто уменьшаем количество
        }
    };

    return (
        <div className={styles.bg}>
            <img src={image} style={{ borderRadius: "20px", objectfit: "contain", height: "150px", width: "150px" }} alt={name} />
            <p className={styles.name}>{name}</p>
            <p className={styles.price}>{price}₽</p>
            <div className={styles.cake_counter}>
                <button className={styles.counter_button} onClick={() => addToCart(id)}>▲</button>
                <span className={styles.counter_number}>{count}</span>
                <button className={styles.counter_button} onClick={handleDecrease}>▼</button>
            </div>
        </div>
    );
}

export default CartItem;