import { Home } from "@/router/Home.tsx";
import { Note } from "@/router/Note.tsx";
import { StrictMode } from "react";
import { createRoot } from "react-dom/client";
import { BrowserRouter as Router, Route, Routes } from "react-router-dom";

const rootElement = document.getElementById("root");

if (rootElement) rootElement.innerHTML = "";

document.querySelectorAll('style[data-loader-style]').forEach((el) => el.remove());

createRoot(rootElement!).render(
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
