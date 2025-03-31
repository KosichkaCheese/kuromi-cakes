import React, { useState } from "react";
import { BrowserRouter as Router, Routes, Route } from "react-router-dom";
import Home from "./pages/Home";
import Main from "./pages/Main";
import Catalog from "./pages/Catalog";
import About from "./pages/About";
import Cart from "./pages/Cart";
import Calculator from "./pages/calculator";
import { CartProvider } from "./components/UseCart";

export default function App() {
  return (
    <Router>
      <CartProvider>
        <Routes>
          <Route path="/" element={<Home />} />
          <Route path="/main" element={<Main />} />
          <Route path="/catalog" element={<Catalog />} />
          <Route path="/about" element={<About />} />
          <Route path="/cart" element={<Cart />} />
          <Route path="/calculator" element={<Calculator />} />
        </Routes>
      </CartProvider>
    </Router>
  );
}
