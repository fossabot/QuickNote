import { useCallback, useEffect, useRef, useState } from "react";
import MDEditor from "@uiw/react-md-editor";
import { useNavigate, useParams } from "react-router-dom";
import { toast, Toaster } from "react-hot-toast";
import { DarkModeToggle } from "../components/DarkModeToggle.tsx";
import { Watermark } from "../components/Watermark.tsx";
import { ImportNote } from "../components/ImportNote.tsx";
import { exportNote, getNote, saveNote } from "../services/noteAPI";
import "./Note.scss";

function useCtrlS(callback: () => void) {
  useEffect(() => {
    const handler = (e: KeyboardEvent) => {
      if ((e.ctrlKey || e.metaKey) && e.key === "s") {
        e.preventDefault();
        callback();
      }
    };
    window.addEventListener("keydown", handler);
    return () => window.removeEventListener("keydown", handler);
  }, [callback]);
}

export function Note() {
  const { id } = useParams<{ id: string }>();
  if (!id) throw new Error("Invalid note id");

  const navigate = useNavigate();
  const [title, setTitle] = useState("");
  const [content, setContent] = useState("");
  const [mode, setMode] = useState<"edit" | "preview" | "both">("both");

  const saveTimeoutRef = useRef<ReturnType<typeof setTimeout> | null>(null);

  const save = useCallback(async () => {
    if (saveTimeoutRef.current) clearTimeout(saveTimeoutRef.current);

    await saveNote(id, title, content);
    toast.success("Note saved");
  }, [id, title, content]);

  useEffect(() => {
    const controller = new AbortController();
    const loadNote = async () => {
      try {
        const note = await getNote(id, controller.signal);
        if (note) {
          setTitle(note.title);
          setContent(note.content);
        }
      } catch (error) {
        if ((error as Error).name !== "AbortError") {
          console.error("Failed to load note:", error);
        }
      }
    };
    void loadNote();
    return () => controller.abort();
  }, [id]);

  useEffect(() => {
    if (saveTimeoutRef.current) clearTimeout(saveTimeoutRef.current);
    saveTimeoutRef.current = setTimeout(() => {
      void save();
    }, 1500);

    return () => {
      if (saveTimeoutRef.current) clearTimeout(saveTimeoutRef.current);
    };
  }, [id, title, content, save]);

  useEffect(() => {
    const handleUnload = () => {
      void save();
    };
    window.addEventListener("beforeunload", handleUnload);
    return () => window.removeEventListener("beforeunload", handleUnload);
  }, [id, title, content, save]);

  useCtrlS(save);

  return (
    <>
      <Watermark text={id} fontSize={10} gapX={150} gapY={150} />
      <div className="content">
        <DarkModeToggle />
        <Toaster position="top-right" />
        <ImportNote callback={(to: string) => navigate(`/note/${to}`, { replace: true })} />
        <div className="note-container visible">
          <div className="note-mode-toggle">
            {["edit", "preview", "both"].map((m) => (
              <button key={m} onClick={() => setMode(m as typeof mode)}>
                {m.charAt(0).toUpperCase() + m.slice(1)}
              </button>
            ))}
            <div className="note-logo" />
            <button className="export" onClick={() => exportNote(id)}>Export</button>
          </div>
          <div className="note-header">
            <input
              type="text"
              className="note-title"
              value={title}
              onChange={(e: React.ChangeEvent<HTMLInputElement>) => {
                setTitle(e.target.value);
              }}
              placeholder="Note title"
            />
          </div>
          <div className="note-content">
            {(mode === "edit" || mode === "both") && (
              <textarea
                className="note-editor"
                value={content}
                onChange={(e: React.ChangeEvent<HTMLTextAreaElement>) => {
                  setContent(e.target.value);
                }}
                placeholder="Write your note here (Markdown)..."
              />
            )}
            {(mode === "preview" || mode === "both") && (
              <div className="note-preview" data-color-mode="light">
                <h1>{title}</h1>
                <MDEditor.Markdown source={content} />
              </div>
            )}
          </div>
        </div>
        <button className="url" onClick={() => {
          navigator.clipboard
            .writeText(window.location.href)
            .then(() => toast.success("Copied to clipboard!"))
            .catch((e) => toast.error(e instanceof Error ? e.message : String(e)));
        }}>
          {window.location.host + window.location.pathname}
        </button>
      </div>
    </>
  );
}