import React, { useEffect, useState } from "react";
import "./order_table.css"

type Item = {
    product_id: string;
    quantity: number;
};

type Order = {
    order_id: string;
    customer_name: string;
    items: Item[];
    status: string;
    created_at: string;
};

export default function OrderTable() {
    const [orders, setOrders] = useState<Order[]>([]);
    const [loading, setLoading] = useState(true);
    const [error, setError] = useState<string | null>(null);

    useEffect(() => {
        const fetchData = async () => {
            try {
                const response = await fetch("http://localhost:8080/orders");
                if (!response.ok) {
                    throw new Error(`HTTP error! Status: ${response.status}`);
                }
                const data: Order[] = await response.json();
                setOrders(data);
            } catch (err: any) {
                setError(err.message);
            } finally {
                setLoading(false);
            }
        };

        fetchData();
    }, []);

    if (loading) {
        return <p>Loading...</p>;
    }
    if (error) {
        return <p>Error: {error}</p>;
    }
    if (!orders.length) {
        return <p>No orders available.</p>;
    }

    return (
        <div className="table_container">
            <h2 className="h2">Order List</h2>
            <table className="order_table">
                <thead  >
                    <tr className="table_tr">
                        <th className="table_th">Order ID</th>
                        <th className="table_th">Customer Name</th>
                        <th className="table_th">Status</th>
                        <th className="table_th">Created At</th>
                        <th className="table_th">Items</th>
                    </tr>
                </thead>
                <tbody>
                    {orders.map((order) => (
                        <tr key={order.order_id}>
                            <td className="table_td">{order.order_id}</td>
                            <td className="table_td">{order.customer_name}</td>
                            <td className="table_td">{order.status}</td>
                            <td className="table_td">{new Date(order.created_at).toLocaleString()}</td>
                            <td className="table_td">
                                <ul>
                                    {order.items.map((item, index) => (
                                        <li key={index}>
                                            {item.product_id} (x{item.quantity})
                                        </li>
                                    ))}
                                </ul>
                            </td>
                        </tr>
                    ))}
                </tbody>
            </table>
        </div>
    );
};
