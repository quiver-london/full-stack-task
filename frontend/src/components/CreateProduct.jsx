import React, {useEffect, useState} from "react";
import {Link, useParams} from "react-router-dom";
import ProductService from "../services/ProductService";

export default function Create() {
    const { id } = useParams();
    const [name, setName] = useState('');
    const [price, setPrice] = useState('');
    const [quantity, setQuantity] = useState('');

    const saveOrUpdateProduct = (e, id, name, price, quantity) => {
        let product = {
            name: name,
            price: price,
            quantity: quantity
        };
        if (id === '_add') {
            ProductService.createProduct(product)
                .then(r => console.log("product created successfully"))
        } else {
            ProductService.updateProduct(product, id)
                .then(r => console.log("product updated successfully"))
        }
    }

    const getTitle = (id) => {
        if (id === '_add') {
            return <h3 className="text-center">Add Product</h3>
        } else {
            return <h3 className="text-center">Update Product</h3>
        }
    }

    useEffect(() => {
        setName(localStorage.getItem('name'));
        setPrice(localStorage.getItem('price'));
        setQuantity(localStorage.getItem('quantity'));
    }, [])

    return (
        <div>
            <br></br>
            <div className="container">
                <div className="row">
                    <div className="card col-md-6 offset-md-3 offset-md-3">
                        {getTitle(id)}
                        <div className="card-body">
                            <form>
                                <div className="form-group">
                                    <label> Name: </label>
                                    <input placeholder="Name"
                                           name="name" className="form-control"
                                           value={name}
                                           onChange={(e) => setName(e.target.value)} />
                                </div>
                                <div className="form-group">
                                    <label> Price: </label>
                                    <input placeholder="Price"
                                           name="price" className="form-control"
                                           value={price}
                                           onChange={(e) => setPrice(e.target.value)} />
                                </div>
                                <div className="form-group">
                                    <label> Quantity: </label>
                                    <input placeholder="Quantity"
                                           name="quantity" className="form-control"
                                           value={quantity}
                                           onChange={(e) => setQuantity(e.target.value)} />
                                </div>

                                <Link to='/products'>
                                    <button className="btn btn-success"
                                            onClick={e => saveOrUpdateProduct(e, id, name, price, quantity)}>
                                        Save
                                    </button>
                                </Link>
                                <Link to='/products'>
                                    <button className="btn btn-danger"
                                            style={{ marginLeft: "10px" }}>
                                        Cancel
                                    </button>
                                </Link>
                            </form>
                        </div>
                    </div>
                </div>
            </div>
        </div>
    )
}