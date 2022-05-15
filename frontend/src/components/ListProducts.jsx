import React, {useEffect, useState} from 'react'
import ProductService from "../services/ProductService";
import {Link} from "react-router-dom";

export default function ListProducts() {
    const [products, setProducts] = useState([]);

    useEffect(() => {
        ProductService.getProducts()
            .then((getData) => {
                setProducts(getData.data);
            })
    }, [])

    const setData = (id, name, price, quantity) => {
        localStorage.setItem('id', id)
        localStorage.setItem('name', name)
        localStorage.setItem('price', price)
        localStorage.setItem('quantity', quantity)
    }

    const getData = () => {
        ProductService.getProducts()
            .then((getData) => {
                setProducts(getData.data);
            })
    }

    const onDelete = (id) => {
        ProductService.deleteProduct(id)
            .then(() => {
                getData();
            })
    }

    return (
        <div>
            <h2 className="text-center">Products List</h2>
            <div className = "row">
                <Link to='/add-product/_add'>
                    <button className="btn btn-primary"
                            onClick={() => setData('_add', '', '', '')}>
                        Add Product
                    </button>
                </Link>
            </div>
            <br></br>
            <div className="row">
                <table className="table table-striped table-bordered">
                    <thead>
                    <tr>
                        <th> Name</th>
                        <th> Price</th>
                        <th> Quantity </th>
                        <th> Actions</th>
                    </tr>
                    </thead>
                    <tbody>
                    {products.map(product =>
                            <tr key = {product._id}>
                                <td> {product.name} </td>
                                <td> {product.price}</td>
                                <td> {product.quantity}</td>
                                <td>
                                    <Link to={`/add-product/${product._id}`}>
                                        <button className="btn btn-info"
                                                color="green"
                                                onClick={() => setData(
                                                    product._id, product.name, product.price, product.quantity)}>
                                            Update
                                        </button>
                                    </Link>
                                    <button style={{marginLeft: "10px"}}
                                            className="btn btn-danger"
                                            onClick={() => onDelete(product._id)}>
                                        Delete
                                    </button>
                                    <Link to={`/view-product/${product._id}`}>
                                        <button style={{marginLeft: "10px"}}
                                                className="btn btn-info">View
                                        </button>
                                    </Link>
                                </td>
                            </tr>
                        )
                    }
                    </tbody>
                </table>
            </div>
        </div>
    )
}
