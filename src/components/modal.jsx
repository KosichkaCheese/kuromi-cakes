import styles from "./modal.module.css";
import { useState, useEffect } from "react";
import { useCart } from "../components/UseCart";
import axios from "axios";

function ModalForm({ type, cakes, totalPrice, onClose }) {
    // const [paymentMethod, setPaymentMethod] = useState("");
    const [error, setError] = useState("");
    const { clearCart } = useCart();
    const [formData, setFormData] = useState({
        name: "",
        email: "",
        phone: "",
        address: "",
        paymentMethod: "",
    });

    const handleChange = (e) => {
        const { name, value } = e.target;
        setFormData((prev) => ({
            ...prev,
            [name]: value,
        }));
        setError("");
    };


    const handleSubmit = async (e) => {
        e.preventDefault();
        const { name, email, phone, address, paymentMethod } = formData;

        if (!paymentMethod) {
            setError("Пожалуйста, выберите способ оплаты");
            return;
        }

        const order = {
            address,
            delivery_type_id: type === "pickup" ? 1 : 2,
            email,
            name,
            phone,
            payment_type_id: parseInt(paymentMethod),
            price: totalPrice,
        };
        console.log(order);
        try {
            const response = await axios.post(`http://localhost:8000/cake_api/orders`, order);
            const orderContent = {
                order_id: response.data.id,
                items: cakes.map((cake) => ({
                    cake_id: cake.id,
                    quantity: cake.count,
                })),
            };

            await axios.post(`http://localhost:8000/cake_api/order_content`, orderContent);
        } catch (error) {
            console.error("Ошибка при отправке заказа:", err);
        }

        alert(`Заказ оформлен`);
        onClose();
        clearCart();
    };


    return (
        <div className={styles.overlay}>
            <div className={styles.modal}>
                <button className={styles.close} onClick={onClose}>
                    ✕
                </button>
                <h3 className={styles.title}>
                    {type === "pickup" ? "Оформление заказа" : "Оформление заказа на доставку"}
                </h3>
                <div className={styles.summary}>
                    {cakes.map((cake) => (
                        <div key={cake.id} style={{ fontWeight: "lighter", fontSize: "15px", marginBottom: "10px" }}>
                            {cake.name} {cake.price}₽ {cake.count} шт.
                        </div>
                    ))}
                    <div style={{ marginTop: "10px", color: "#FFCFF1", fontWeight: "bold", fontSize: "18px" }}>
                        <b>Итого: {totalPrice}₽</b>
                    </div>
                </div>

                <form className={styles.form} onSubmit={handleSubmit}>
                    <input name="name" placeholder="Получатель" required onChange={handleChange} value={formData.name} />
                    <input name="phone" placeholder="Телефон" type="tel" required onChange={handleChange} value={formData.phone} />
                    {type === "delivery" && <input name="address" placeholder="Адрес доставки" required onChange={handleChange} value={formData.address} />}
                    <select
                        name="paymentMethod"
                        onChange={handleChange}
                        className={styles.select}
                        value={formData.paymentMethod}
                        style={{ marginBottom: "20px", backgroundColor: "#C3C2C3" }}
                    >
                        <option value="" disabled hidden>Выберите способ оплаты</option>
                        <option value="1">Картой при получении</option>
                        <option value="2">Наличными</option>
                        <option value="3">Картой онлайн</option>
                    </select>
                    <input name="email" placeholder="email" type="email" required onChange={handleChange} value={formData.email} />
                    {error && <p className="error-message">{error}</p>}
                    <button className={styles.submit} type="submit" disabled={!formData.paymentMethod}>
                        Подтвердить
                    </button>
                </form>
            </div>
        </div>
    );
}

export default ModalForm;