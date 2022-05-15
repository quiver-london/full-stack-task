import {useEffect, useState} from "react";
import {Link, useParams} from "react-router-dom";
import ProductService from "../services/ProductService";

export default function ViewProduct() {
    const { id } = useParams();
    const [product, setProduct] = useState({});

    useEffect(() => {
        ProductService.getProductById(id)
            .then((getProduct) => {
                setProduct(getProduct.data);
            })
    }, [id])

    return (
        <div>
            <br></br>
            <div className = "card col-md-6 offset-md-3">
                <h3 className = "text-center">
                    View Product Details</h3>
                <div className = "card-body">
                    <div className = "row">
                        <label> Product Name: </label>
                        <div> { product.name }
                        </div>
                    </div>
                    <br/>
                    <div className = "row">
                        <label> Product Price: </label>
                        <div> { product.price }
                        </div>
                    </div>
                    <br/>
                    <div className = "row">
                        <label> Product Quantity: </label>
                        <div> { product.quantity }
                        </div>
                    </div>
                </div>
            </div>
            <Link to='/products'>
                <button
                    style={{marginLeft: "auto", marginRight: "auto", width: "50%", display: "block"}}
                    className="btn btn-info">
                    Get back to products list
                </button>
            </Link>
        </div>
    )
}