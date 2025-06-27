import { useEffect, useState } from "react";
import { useNavigate } from "react-router-dom";
import { DarkModeToggle } from "../components/DarkModeToggle";
import "./Home.scss";
import { ImportNote } from "../components/ImportNote.tsx";

export function Home() {
  const [visible, setVisible] = useState(false);

  const [uuid, setUUID] = useState<string>(crypto.randomUUID());
  const navigate = useNavigate();


  const handleNavigation = (nextUUID?: string) => {
    setVisible(false);
    setTimeout(() => navigate(`/note/${nextUUID ?? uuid}`), 500);
  };


  useEffect(() => {
    const t = setTimeout(() => setVisible(true), 100);
    return () => clearTimeout(t);
  }, []);

  return (
    <>
      <DarkModeToggle />
      <ImportNote callback={(to: string) => {
        handleNavigation(to);
      }} />
      <div className={`content`}>
        <div
          className={`background ${visible ? "visible" : ""}`}
        >
          <div className="title">
            <div className="logo" />
            <a
              className="github"
              href="https://github.com/Sn0wo2/QuickNote"
              target="_blank"
              rel="noopener noreferrer"
            />
          </div>
          <p className="subtitle">
            <span className="highlight">QuickNote</span>
            <span className="note">
              Instantly write and share your thoughts.
            </span>
            <span className="warning">
              ⚠️ Please don’t upload illegal or sensitive content.
            </span>
          </p>
          <div className="input-container">
            <input
              className="uuid-input"
              type="text"
              value={uuid}
              onChange={(e) => setUUID(e.target.value)}
              onKeyDown={(e) => e.key === "Enter" && handleNavigation()}
            />
            <button className="submit-btn" onClick={(e) => {
              e.preventDefault();
              handleNavigation();
            }}>
              &rarr;
            </button>
          </div>
        </div>
      </div>
    </>
  );
}
