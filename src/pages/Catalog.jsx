import styles from "./Catalog.module.css";
import Header from "../components/header";

function Catalog() {
    return (
        <div className={styles.catalog_bg}>
            <Header />
        </div>
    )
}

export default Catalog;