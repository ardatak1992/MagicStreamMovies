import Button from "react-bootstrap/esm/Button";
import Nav from "react-bootstrap/Nav";
import Navbar from "react-bootstrap/Navbar";
import Container from "react-bootstrap/Container";
import { useNavigate, NavLink, Link } from "react-router-dom";
const Header = () => {
  const navigate = useNavigate();

  return (
    <Navbar
      bg="dark"
      variant="dark"
      expand="lg"
      sticky="top"
      className="shadow-sm"
    >
      <Container>
        <Navbar.Brand>Magic Stream</Navbar.Brand>
      </Container>
    </Navbar>
  );
};

export default Header;
