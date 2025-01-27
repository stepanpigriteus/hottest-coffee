import { useState } from "react";
import "./order_button.css"


interface ButtonProps {
    orderId: string;
    onClose: (orderId: string) => void; 
    title: string;
}

export default function Button({ orderId, onClose, title}: ButtonProps) {
    const [loading, setLoading] = useState(false);
    const [error, setError] = useState<string | null>(null);

    const handleClick = async () => {
        setLoading(true);
        setError(null);
    
        try {
            const response = await fetch(`http://localhost:8080/orders/${orderId}/close`, {
                method: "POST",
        
            });
        
            if (!response.ok) {
                console.log(error)
                throw new Error(`Ошибка при закрытии заказа: ${response.statusText}`);
            }
            
            onClose(orderId);
        } catch (err: any) {
            console.error("Fetch error:", err);
            setError(err.message);
        } finally {
            setLoading(false);
        }
    };

    console.log("Render", title, loading);
    return (
        <div >
        <button className = "order_button" onClick={handleClick} disabled={loading}>
            { title}
        </button>
  
    </div>
    );
}
