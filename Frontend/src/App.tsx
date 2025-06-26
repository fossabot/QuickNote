import { BrowserRouter as Router, Route, Routes } from "react-router-dom";
import { Home } from "./router/Home.tsx";
import { Note } from "./router/Note.tsx";

export function App() {
  return (
    <Router>
      <Routes>
        <Route path="/" element={<Home />} />
        <Route path="/note/:id" element={<Note />} />
        <Route path="*" element={<Home />} />
      </Routes>
    </Router>
  );
}
