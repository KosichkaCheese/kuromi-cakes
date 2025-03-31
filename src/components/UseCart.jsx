import { useState, useEffect } from "react";

function useCart() {
    const [cart, setCart] = useState(() => {
        const savedCart = sessionStorage.getItem("cart");
        console.log("Loaded cart from sessionStorage", savedCart);
        return savedCart ? JSON.parse(savedCart) : [];
    });

    useEffect(() => {
        sessionStorage.setItem("cart", JSON.stringify(cart));
    }, [cart]);

    const addToCart = (id) => {
        setCart((prevCart) => {
            const existingItem = prevCart.find(item => item.id === id);
            if (existingItem) {
                return prevCart.map(item =>
                    item.id === id ? { ...item, count: item.count + 1 } : item
                );
            } else {
                return [...prevCart, { id, count: 1 }];
            }
        });
    };

    const removeFromCart = (id) => {
        setCart((prevCart) =>
            prevCart.map((item) =>
                item.id === id ? { ...item, count: item.count - 1 } : item
            )
        );
    };

    const clearItem = (id) => {
        setCart((prevCart) => prevCart.filter((item) => item.id !== id));
    };

    const clearCart = () => {
        setCart([]);
    };

    return { cart, addToCart, removeFromCart, clearItem, clearCart };
};

export default useCart;