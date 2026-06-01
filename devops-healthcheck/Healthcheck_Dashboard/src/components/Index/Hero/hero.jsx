import "./hero.css";

export default function Hero() {

    return (

        <section className="hero-section" id="home">

            {/* LEFT CONTENT */}

            <div className="hero-left">

                <p className="hero-tag">
                    Real-Time Infrastructure Monitoring
                </p>

                <h1>
                    Monitor Your Infrastructure
                    <span> In Real Time</span>
                </h1>

                <p className="hero-description">

                    Track uptime, APIs, servers,
                    Docker containers and cloud services
                    with instant alerts and performance insights.

                </p>


                {/* BUTTONS */}

                <div className="hero-buttons">

                    <button className="start-btn">
                        Start Monitoring
                    </button>

                    <button className="demo-btn">
                        Live Demo
                    </button>

                </div>

            </div>


            {/* RIGHT SIDE */}

            <div className="hero-right">

                <div className="dashboard-preview">

                    <div className="preview-header">

                        <div className="dot red"></div>
                        <div className="dot yellow"></div>
                        <div className="dot green"></div>

                    </div>

                    <div className="preview-body">

                        <div className="preview-card">
                            <h3>API Server</h3>
                            <p className="healthy">
                                Healthy
                            </p>
                        </div>

                        <div className="preview-card">
                            <h3>Docker Engine</h3>
                            <p className="healthy">
                                Running
                            </p>
                        </div>

                        <div className="preview-card">
                            <h3>Main Website</h3>
                            <p className="down">
                                Down
                            </p>
                        </div>
                        <img src="../../../assets/dashboard" alt="" />

                    </div>

                </div>

            </div>

        </section>
    );
}