import "./about.css";

import {
    FaReact,
    FaDocker,
    FaDatabase,
    FaServer
} from "react-icons/fa";

import {
    SiGo,
    SiPostgresql
} from "react-icons/si";

export default function About() {

    return (

        <section
            className="about-section"
            id="about"
        >

            {/* LEFT */}

            <div className="about-left">

                <p className="about-tag">
                    ABOUT PLATFORM
                </p>

                <h2>
                    Modern Infrastructure
                    Monitoring Platform
                </h2>

                <p className="about-description">

                    This platform helps developers and teams
                    monitor APIs, servers, websites,
                    Docker containers and cloud infrastructure
                    in real time.

                </p>

                <p className="about-description">

                    It provides uptime tracking,
                    health monitoring, instant alerts
                    and performance insights from
                    a single dashboard.

                </p>

            </div>


            {/* RIGHT */}

            <div className="about-right">

                <div className="tech-card">

                    <div className="tech-icon">
                        <FaReact />
                    </div>

                    <h3>React</h3>

                </div>


                <div className="tech-card">

                    <div className="tech-icon">
                        <SiGo />
                    </div>

                    <h3>Go</h3>

                </div>


                <div className="tech-card">

                    <div className="tech-icon">
                        <SiPostgresql />
                    </div>

                    <h3>PostgreSQL</h3>

                </div>


                <div className="tech-card">

                    <div className="tech-icon">
                        <FaDocker />
                    </div>

                    <h3>Docker</h3>

                </div>


                <div className="tech-card">

                    <div className="tech-icon">
                        <FaDatabase />
                    </div>

                    <h3>Database</h3>

                </div>


                <div className="tech-card">

                    <div className="tech-icon">
                        <FaServer />
                    </div>

                    <h3>Cloud Infra</h3>

                </div>

            </div>

        </section>
    );
}