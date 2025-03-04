import styles from "./Cart.module.css";
import Header from "../components/header";
import CartItem from "../components/cart_item";
import { act, useState } from "react";

const initCakes = [
    { id: 2, image: "/assets/2.png", name: "Шоколадный торт", price: 1500, count: 1 },
    { id: 3, image: "/assets/3.png", name: "Ванильный торт", price: 1400, count: 1 }
];

function Cart() {
    const [active, setActive] = useState("pickup");
    const [cakes, setCakes] = useState(initCakes);

    const updateCount = (id, newCount) => {
        setCakes(prevCakes =>
            prevCakes.map(cake =>
                cake.id === id ? { ...cake, count: newCount } : cake
            )
        );
    };

    const totalPrice = cakes.reduce((sum, cake) => sum + cake.price * cake.count, 0);

    return (
        <div className={styles.cart_bg}>
            <Header />
            <div style={{ alignSelf: "center", marginTop: "2%" }}>
                <div className={styles.switch}>
                    <button
                        className={`${styles.option} ${active === "pickup" ? styles.active : ""}`}
                        onClick={() => setActive("pickup")}
                    >
                        САМОВЫВОЗ
                    </button>
                    <button
                        className={`${styles.option} ${active === "delivery" ? styles.active : ""}`}
                        onClick={() => setActive("delivery")}
                    >
                        ДОСТАВКА
                    </button>
                </div>
            </div>
            <div className={styles.items}>
                {cakes.map(cake => (
                    <CartItem key={cake.id} {...cake} updateCount={updateCount} />
                ))}
            </div>
            <div style={{ display: "flex", flexDirection: "column", marginBottom: "5%", marginTop: "2%", width: "75%", alignSelf: "center" }}>
                <p className={styles.total}>Итого: {totalPrice}₽</p>
                <button className={styles.continue} onClick={() => active === "pickup" ? alert("Заказ на самовывоз оформлен") : alert("Заказ на доставку оформлен")}>Продолжить</button>
            </div>
        </div>
    )
}

export default Cart;