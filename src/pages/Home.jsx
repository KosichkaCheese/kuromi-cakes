import React, { useState } from "react";
import styles from "./Home.module.css";
import { useNavigate } from "react-router-dom";

function Home() {
    const [isHovered, setIsHovered] = useState(false);
    const navigate = useNavigate();

    return (
        <div className={styles.home_bg}>
            <div style={{ margintop: "2%", marginleft: "3%" }}><img width={"600px"} src="assets/Тортики Куроми.png"></img></div>
            <div style={{ marginLeft: "auto", marginRight: "2%", marginTop: "-4%" }}><img width={"500px"} src="assets/Ешьте тортики.png"></img></div>
            <div style={{ marginTop: "-4%", alignSelf: "center", marginLeft: "-5%" }}>
                <img style={{ position: "relative", top: "-10px", left: "70px", rotate: "-20deg", scale: "1.3" }} src={isHovered ? "assets/kuromi2.png" : "assets/kuromi1.png"}></img>
                <button className={styles.want_cake} onClick={() => navigate("/main")} onMouseEnter={() => setIsHovered(true)} onMouseLeave={() => setIsHovered(false)}>хочу тортик</button>
            </div>
        </div>
    );
}

export default Home;