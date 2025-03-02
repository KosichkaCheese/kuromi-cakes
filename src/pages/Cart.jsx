import styles from "./Cart.module.css";
import Header from "../components/header";

function Cart() {
    return (
        <div className={styles.cart_bg}>
            <Header />
        </div>
    )
}

export default Cart;