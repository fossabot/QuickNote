import { useCallback, useEffect, useRef, useState } from "react";
import { DarkModeToggle } from "../components/DarkModeToggle";
import "./Home.scss";
import { useNavigate } from "react-router-dom";
import Watermark from "../components/Watermark.tsx";
import { toast, Toaster } from "react-hot-toast";
import { importNote } from "../services/noteAPI.ts";

export function Home() {
  const [visible, setVisible] = useState(false);
  const [isDragging, setIsDragging] = useState(false);
  const [uuid, setUUID] = useState<string>(crypto.randomUUID());
  const dragCounter = useRef(0);
  const navigate = useNavigate();


  const handleNavigation = (nextUUID?: string) => {
    setVisible(false);
    setTimeout(() => navigate(`/note/${nextUUID ?? uuid}`), 500);
  };

  const handleFileUpload = useCallback(async (file: File) => {
    try {
      if (!file.name.endsWith(".qnote")) throw new Error(`Invalid file: ${file.name}`);
      await importNote(file);
      handleNavigation(file.name.replace(/\.qnote$/, ""));
    } catch (error) {
      console.error(error);
      toast.error("Failed to import note");
    }
  }, []);


  const handleDrop = useCallback(
    (e: DragEvent) => {
      e.preventDefault();
      dragCounter.current = 0;
      setIsDragging(false);
      const files = e.dataTransfer?.files;
      if (files?.length) handleFileUpload(files[0]);
      else toast.error("No note dropped");
    },
    [handleFileUpload]
  );

  useEffect(() => {
    const handleDragEnter = (e: DragEvent) => {
      if (e.dataTransfer?.types.includes("Files")) {
        dragCounter.current++;
        setIsDragging(true);
      }
    };
    const handleDragLeave = () => {
      dragCounter.current--;
      if (dragCounter.current <= 0) setIsDragging(false);
    };
    const handleDragOver = (e: DragEvent) => e.preventDefault();

    window.addEventListener("dragenter", handleDragEnter);
    window.addEventListener("dragleave", handleDragLeave);
    window.addEventListener("dragover", handleDragOver);
    window.addEventListener("drop", handleDrop);

    return () => {
      window.removeEventListener("dragenter", handleDragEnter);
      window.removeEventListener("dragleave", handleDragLeave);
      window.removeEventListener("dragover", handleDragOver);
      window.removeEventListener("drop", handleDrop);
    };
  }, [handleDrop]);

  useEffect(() => {
    const t = setTimeout(() => setVisible(true), 100);
    return () => clearTimeout(t);
  }, []);

  return (
    <>
      <Watermark text={uuid} fontSize={20} gapX={150} gapY={150} />
      <DarkModeToggle />
      <div className={`content`}>
        <div
          className={`background ${visible ? "visible" : ""} ${isDragging ? "dragging" : ""}`}
        >
          <Toaster position="top-right" />
          <div className="title">
            <div className="logo" />
            <a
              className="github"
              href="https://github.com/Sn0wo2/QuickNote"
              target="_blank"
              rel="noopener noreferrer"
            ></a>
          </div>
          <p className="subtitle">
            <span className="highlight">QuickNote</span>
            <span className="note">
              Instantly write and share your thoughts.
            </span>
            <span className="drag-hint">
              {isDragging
                ? "Drop your note here!"
                : "Drag a .qnote file here to import."}
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
