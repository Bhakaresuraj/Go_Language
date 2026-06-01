import "./contact.css";

import {
    FaEnvelope,
    FaPhone,
    FaMapMarkerAlt
} from "react-icons/fa";

export default function Contact() {

    return (

        <section
            className="contact-section"
            id="contact"
        >

            {/* LEFT SIDE */}

            <div className="contact-left">

                <p className="contact-tag">
                    CONTACT US
                </p>

                <h2>
                    Let's Build Reliable
                    Infrastructure Together
                </h2>

                <p className="contact-description">

                    Have questions about monitoring,
                    uptime tracking or cloud infrastructure?
                    Reach out to us anytime.

                </p>


                {/* CONTACT INFO */}

                <div className="contact-info">

                    <div className="contact-item">

                        <FaEnvelope className="contact-icon" />

                        <span>
                            support@tygo.com
                        </span>

                    </div>


                    <div className="contact-item">

                        <FaPhone className="contact-icon" />

                        <span>
                            +91 9876543210
                        </span>

                    </div>


                    <div className="contact-item">

                        <FaMapMarkerAlt className="contact-icon" />

                        <span>
                            Pune, Maharashtra
                        </span>

                    </div>

                </div>

            </div>


            {/* RIGHT SIDE */}

            <div className="contact-right">

                <form className="contact-form">

                    <input
                        type="text"
                        placeholder="Your Name"
                        required
                    />

                    <input
                        type="email"
                        placeholder="Your Email"
                        required
                    />

                    <textarea
                        placeholder="Your Message"
                        rows="6"
                        required
                    ></textarea>

                    <button type="submit">
                        Send Message
                    </button>

                </form>

            </div>

        </section>
    );
}