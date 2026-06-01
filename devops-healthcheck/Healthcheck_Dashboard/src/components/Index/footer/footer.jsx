import "./footer.css";

import {
    FaGithub,
    FaLinkedin,
    FaInstagram,
    FaArrowUp
} from "react-icons/fa";

export default function Footer() {

    return (
        <footer className="footer">
            {/* TOP */}
            <div className="footer-top">
                {/* BRAND */}
                <div className="footer-brand">
                    <h2>Tygo</h2>
                    <p>
                        Real-time infrastructure monitoring
                        platform for APIs, servers,
                        Docker containers and cloud services.
                    </p>
                </div>

                {/* QUICK LINKS */}
                <div className="footer-links">
                    <h3>Quick Links</h3>
                    <a href="#home">Home</a>
                    <a href="#services">
                        Services
                    </a>
                    <a href="#about">
                        About
                    </a>
                    <a href="#contact">
                        Contact
                    </a>
                </div>
                {/* SOCIAL */}
                <div className="footer-social">
                    <h3>Connect</h3>
                    <div className="social-icons">
                        <a href="#">
                            <FaGithub />
                        </a>
                        <a href="#">
                            <FaLinkedin />
                        </a>
                        <a href="#">
                            <FaInstagram />
                        </a>
                    </div>
                </div>
            </div>

            {/* DIVIDER */}
            <div className="footer-divider"></div>
            {/* BOTTOM */}
            <div className="footer-bottom">
                <p>
                    © 2026 Tygo. All Rights Reserved.
                </p>
                {/* SCROLL TOP */}
                <a
                    href="#home"
                    className="scroll-top"
                >
                    <FaArrowUp />
                </a>
            </div>
        </footer>
    );
}