import styles from "./Cart.module.css";
import Header from "../components/header";
import CartItem from "../components/cart_item";
import { useCart } from "../components/UseCart";
import { useState, useEffect } from "react";
import axios from "axios";

function Cart() {
  const [active, setActive] = useState("pickup");
  const { cart, addToCart, removeFromCart, clearItem, clearCart } = useCart();
  const [cakes, setCakes] = useState([]);

  useEffect(() => {
    async function fetchCakes() {
      console.log(cart);

      if (cart.length === 0) {
        setCakes([]);
        return;
      }

      try {
        const cakeRequests = cart.map((item) =>
          axios
            .get(`http://localhost:8000/cake_api/cakes/${item.id}`)
            .then((res) => ({
              ...res.data,
              count: item.count,
            }))
        );

        const cakesData = await Promise.all(cakeRequests);
        setCakes(cakesData);
      } catch (error) {
        console.error("Ошибка загрузки товаров:", error);
      }
    }

    fetchCakes();
  }, [cart]);

  const totalPrice = cakes.reduce(
    (sum, cake) => sum + cake.price * cake.count,
    0
  );

  return (
    <div className={styles.cart_bg}>
      <Header />
      <div style={{ alignSelf: "center", marginTop: "2%" }}>
        <div className={styles.switch}>
          <button
            className={`${styles.option} ${
              active === "pickup" ? styles.active : ""
            }`}
            onClick={() => setActive("pickup")}
          >
            САМОВЫВОЗ
          </button>
          <button
            className={`${styles.option} ${
              active === "delivery" ? styles.active : ""
            }`}
            onClick={() => setActive("delivery")}
          >
            ДОСТАВКА
          </button>
        </div>
      </div>
      <div className={styles.items}>
        {cakes.map((cake) => (
          <CartItem
            key={cake.id}
            id={cake.id}
            image={"assets/" + cake.image}
            name={cake.name}
            price={cake.price}
            count={cake.count}
            addToCart={addToCart}
            removeFromCart={removeFromCart}
            clearItem={clearItem}
          />
        ))}
      </div>
      <div
        style={{
          display: "flex",
          flexDirection: "column",
          marginBottom: "5%",
          marginTop: "2%",
          width: "75%",
          alignSelf: "center",
        }}
      >
        <p className={styles.total}>Итого: {totalPrice}₽</p>
        <button
          className={styles.continue}
          onClick={() =>
            active === "pickup"
              ? alert("Заказ на самовывоз оформлен")
              : alert("Заказ на доставку оформлен")
          }
        >
          Продолжить
        </button>
      </div>
    </div>
  );
}

export default Cart;
