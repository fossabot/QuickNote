import { type ChangeEvent, useEffect, useRef, useState } from "react";
import { toast } from "react-hot-toast";
import { useNavigate } from "react-router-dom";
import { DarkModeToggle } from "../components/DarkModeToggle";
import "./Home.scss";
import { ImportNote } from "../components/ImportNote.tsx";
import { importNote } from "../services/noteAPI.ts";

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

  const fileInputRef = useRef<HTMLInputElement | null>(null);

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
            <label className="upload">
              <input
                type="file"
                accept=".qnote"
                style={{ display: 'none' }}
                ref={fileInputRef}
                onChange={async (e: ChangeEvent<HTMLInputElement>) => {
                  try {
                    const file = e.target.files?.[0];
                    if (!file) return;
                    const success = await importNote(file);
                    if (!success) {
                      toast.error("Failed to import note");
                      return;
                    }

                    navigate(`/note/${file.name.replace(/\.qnote$/, "")}`);

                  } catch (error) {
                    console.error(error);
                    toast.error("Failed to import note");
                  }
                }}
              />
              <img src="/file-pencil-alt-svgrepo-com.svg" alt="upload icon" />
            </label>
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
