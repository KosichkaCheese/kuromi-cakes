import styles from "./Catalog.module.css";
import Header from "../components/header";
import CakeCard from "../components/cake_card";

function Catalog() {
    const cakes = [
        { id: 1, image: "/assets/1.png", name: "Клубничный торт", price: "1200₽" },
        { id: 2, image: "/assets/2.png", name: "Шоколадный торт", price: "1500₽" },
        { id: 3, image: "/assets/3.png", name: "Ванильный торт", price: "1400₽" },
        { id: 4, image: "/assets/4.png", name: "Мятный торт", price: "1100₽" },
        { id: 5, image: "/assets/5.png", name: "Карамельный торт", price: "1300₽" },
        { id: 6, image: "/assets/6.png", name: "Клубничный торт", price: "2000₽" },
        { id: 7, image: "/assets/7.png", name: "Шоколадный торт", price: "1600₽" },
        { id: 8, image: "/assets/8.png", name: "Ванильный торт", price: "1000₽" }
    ]

    return (
        <div className={styles.catalog_bg}>
            <Header />
            <div className={styles.catalog}>
                {cakes.map(cake => (
                    <CakeCard key={cake.id} image={cake.image} name={cake.name} price={cake.price} />
                ))}
            </div>
        </div>
    )
}

export default Catalog;