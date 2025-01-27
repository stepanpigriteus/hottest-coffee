
import "./header.css"


export default function Header () {
      return (
        
     <section className = "header">
        <div className = " button_container">
            <button className="header_button"> Orders </button>  
            <button className="header_button"> Customers </button>  
            <button className="header_button"> Menu </button>  
            <button className="header_button"> Inventory </button>  
            <button className="header_button"> Sales </button>  
            <button className="header_button"> Popular </button>  
        </div>
      </section>
      );
  }