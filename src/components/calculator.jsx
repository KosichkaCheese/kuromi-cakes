import React, { useState } from "react";
import styles from "./Calculator.module.css";

const BaseConverter = () => {
    const [fromBase, setFromBase] = useState(10);
    const [toBase, setToBase] = useState(2);
    const [input, setInput] = useState("");
    const [result, setResult] = useState("");

    const digits = "0123456789ABCDEF".slice(0, fromBase);

    const handleInput = (digit) => {
        setInput((prev) => prev + digit);
    };

    const handleConvert = () => {
        if (input) {
            const decimalValue = parseInt(input, fromBase);
            setResult(decimalValue.toString(toBase).toUpperCase());
        }
    };

    const handleReset = () => {
        setInput("");
        setResult("");
    };

    const handleFromBaseChange = (e) => {
        setFromBase(Number(e.target.value));
        setInput("");
        setResult("");
    };

    const handleToBaseChange = (e) => {
        setToBase(Number(e.target.value));
        setInput("");
        setResult("");
    };

    const handleBackspace = () => {
        setInput((prev) => prev.slice(0, -1)); // Удаляем последний символ
    };

    return (
        <div className={styles.bg}>
            <div className={styles.container}>
                <div className={styles.select_container}>
                    <select className={styles.select} value={fromBase} onChange={handleFromBaseChange}>
                        {[...Array(15)].map((_, i) => (
                            <option key={i} value={i + 2}>{i + 2}</option>
                        ))}
                    </select>
                    <span className={styles.arrow}>→</span>
                    <select className={styles.select} value={toBase} onChange={handleToBaseChange}>
                        {[...Array(15)].map((_, i) => (
                            <option key={i} value={i + 2}>{i + 2}</option>
                        ))}
                    </select>
                </div>
                <div className={styles.display}>{input || "0"}</div>
                <div className={styles.buttons_grid}>
                    {"0123456789ABCDEF".split("").map((digit) => (
                        <button
                            key={digit}
                            disabled={!digits.includes(digit)}
                            onClick={() => handleInput(digit)}
                            className={styles.digit_button}
                        >
                            {digit}
                        </button>
                    ))}
                </div>
                <div className={styles.controls}>
                    <button onClick={handleConvert} className={styles.convert_button}>Convert</button>
                    <button onClick={handleReset} className={styles.reset_button}>Reset</button>
                    <button onClick={handleBackspace} className={styles.backspace_button}>⇦</button>
                </div>
                <div className={styles.display}>{result || "0"}</div>
            </div>
        </div>
    );
};

export default BaseConverter;
