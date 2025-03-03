import styles from "./About.module.css";
import Header from "../components/header";

function About() {
    return (
        <div className={styles.about_bg}>
            <Header />
            <div className={styles.textbox}>
                <p className={styles.title}>Мы предлагаем самые вкусные и злые тортики.</p>

                <p className={styles.subtitle}>Способы оплаты:</p>
                <ul style={{ marginLeft: "40%", textAlign: "left" }}>
                    <li>Волшебные ноты</li>
                    <li>СБП</li>
                </ul>

                <p className={styles.subtitle}>Контакты:</p>
                <p>Королевство Мэриленд, домик Куроми</p>
                <p>телефон: <strong>+74446661488</strong></p>

                <p className={styles.footertext}>
                    Доставка осуществляется только в пределах Мэриленда.
                    Для уточнения сроков обращайтесь по контактному номеру.
                </p>
            </div>
        </div>
    )
}

export default About;