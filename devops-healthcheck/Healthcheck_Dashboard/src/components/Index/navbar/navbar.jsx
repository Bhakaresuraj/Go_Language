import "./navbar.css";
import { Link } from "react-router-dom";
export default function Navbar() {
    return (
        <nav className="navbar">
            {/* LEFT */}
            <div className="navbar-logo">
                <h2>Health-check</h2>
            </div>
            {/* CENTER */}
            <ul className="navbar-links">
                <li>
                    <a href="#home">Home</a>
                </li>
                <li>
                    <a href="#services">Services</a>
                </li>
                <li>
                    <a href="#about">About</a>
                </li>
                <li>
                    <a href="#contact">Contact</a>
                </li>
            </ul>

            {/* RIGHT */}
            <div className="navbar-buttons">
                <Link to="/login">
                    <button className="login-btn">
                        Login
                    </button>
                </Link>
                <Link to="/signup">

                    <button className="signup-btn">
                        Signup
                    </button>

                </Link>
            </div>
        </nav>
    );
}