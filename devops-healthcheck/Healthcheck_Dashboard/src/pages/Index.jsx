import Navbar from "../components/Index/navbar/navbar";
import Hero from "../components/Index/Hero/hero";
import Services from "../components/Index/services/services";
import About from "../components/Index/about/about";
import Contact from "../components/Index/contact/contact";
import Footer from "../components/Index/footer/footer";

export default function IndexPage() {

    return (

        <>
             <Navbar />
            <Hero />
            <Services />
            <About />
            <Contact /> 
             <Footer /> 
        </>
    );
}