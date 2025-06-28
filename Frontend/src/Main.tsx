import {StrictMode} from "react";
import {createRoot} from "react-dom/client";
import {BrowserRouter as Router, Route, Routes} from "react-router-dom";
import {Home} from "./router/Home.tsx";
import {Note} from "./router/Note.tsx";

createRoot(document.getElementById("root")!).render(
    <StrictMode>
    <Router>
      <Routes>
        <Route path="/" element={<Home />} />
        <Route path="/note/:id" element={<Note />} />
        <Route path="*" element={<Home />} />
      </Routes>
    </Router>
    </StrictMode>
);