import "./service.css";

import {
    FaServer,
    FaDocker,
    FaGlobe,
    FaBell,
    FaLock,
    FaCode
} from "react-icons/fa";

export default function Services() {

    const services = [

        {
            icon: <FaGlobe />,
            title: "Website Monitoring",
            description:
                "Track website uptime and performance in real time."
        },

        {
            icon: <FaCode />,
            title: "API Health Monitoring",
            description:
                "Monitor API response time and health status."
        },

        {
            icon: <FaDocker />,
            title: "Docker Monitoring",
            description:
                "Monitor Docker containers and service availability."
        },

        {
            icon: <FaServer />,
            title: "Server Monitoring",
            description:
                "Track server performance and infrastructure health."
        },

        {
            icon: <FaLock />,
            title: "SSL Monitoring",
            description:
                "Monitor SSL certificate expiration and security."
        },

        {
            icon: <FaBell />,
            title: "Uptime Alerts",
            description:
                "Get instant alerts when services go down."
        }

    ];

    return (

        <section
            className="services-section"
            id="services"
        >

            <div className="services-header">

                <p>OUR SERVICES</p>

                <h2>
                    Everything You Need
                    To Monitor Infrastructure
                </h2>

            </div>


            {/* CARDS */}

            <div className="services-grid">

                {
                    services.map((service, index) => (

                        <div
                            className="service-card"
                            key={index}
                        >

                            <div className="service-icon">
                                {service.icon}
                            </div>

                            <h3>
                                {service.title}
                            </h3>

                            <p>
                                {service.description}
                            </p>

                        </div>
                    ))
                }

            </div>

        </section>
    );
}