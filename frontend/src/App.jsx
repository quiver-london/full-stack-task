import React from 'react';
import {BrowserRouter as Router, Route, Routes} from 'react-router-dom'
import Header from "./components/Header";
import Create from "./components/CreateProduct";
import ListProducts from "./components/ListProducts";
import ViewProduct from "./components/ViewProduct";

export default function App() {
  return (
    <div className="App">
        <Router>
            <Header />
            <div className="container">
                <Routes>
                    <Route path = "/" element={<ListProducts/>}></Route>
                    <Route path = "/products" element={<ListProducts/>}></Route>
                    <Route path = "/add-product/:id" element={<Create/>}></Route>
                    <Route path = "/view-product/:id" element={<ViewProduct/>}></Route>
                </Routes>
            </div>
        </Router>
    </div>
  );
}
