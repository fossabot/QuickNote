import MDEditor from "@uiw/react-md-editor";
import { useCallback, useEffect, useRef, useState } from "react";
import { toast, Toaster } from "react-hot-toast";
import { useNavigate, useParams } from "react-router-dom";
import { DarkModeToggle } from "../components/DarkModeToggle.tsx";
import { ImportNote } from "../components/ImportNote.tsx";
import { exportNote, getNote, saveNote } from "../services/noteAPI";
import "./Note.scss";

export function Note() {
  const { id } = useParams<{ id: string }>();
  if (!id) throw new Error("Invalid note id");

  const navigate = useNavigate();
  const [title, setTitle] = useState("");
  const [content, setContent] = useState("");
  const [mode, setMode] = useState<"edit" | "preview" | "both">("both");
  const saveTimeout = useRef<ReturnType<typeof setTimeout> | null>(null);
  const skipSave = useRef(true);

  const save = useCallback(async () => {
    if (saveTimeout.current) {
      clearTimeout(saveTimeout.current);
      saveTimeout.current = null;
    }
    await saveNote(id, title, content);
    toast.success("Note saved");
  }, [id, title, content]);

  const load = useCallback(async () => {
    try {
      const note = await getNote(id);
      if (note) {
        setTitle(note.title);
        setContent(note.content);
        skipSave.current = true;
      }
    } catch (e) {
      console.error(e);
      toast.error("Failed to load note");
    }
  }, [id]);

  useEffect(() => {
    load();
  }, [load]);

  useEffect(() => {
    if (skipSave.current) {
      skipSave.current = false;
      return;
    }
    if (saveTimeout.current) clearTimeout(saveTimeout.current);
    saveTimeout.current = setTimeout(save, 1500);
    return () => {
      if (saveTimeout.current) clearTimeout(saveTimeout.current);
    };
  }, [title, content, save]);

  useEffect(() => {
    const handler = () => save();
    window.addEventListener("beforeunload", handler);
    return () => window.removeEventListener("beforeunload", handler);
  }, [save]);

  useEffect(() => {
    const onKey = (e: KeyboardEvent) => {
      if ((e.ctrlKey || e.metaKey) && e.key === "s") {
        e.preventDefault();
        save();
      }
    };
    window.addEventListener("keydown", onKey);
    return () => window.removeEventListener("keydown", onKey);
  }, [save]);

  return (
    <>
      <DarkModeToggle />
      <div className="content">
        <Toaster position="top-right" />
        <ImportNote callback={(to: string) => navigate(`/note/${to}`, { replace: true })} />
        <div className="note-container visible">
          <div className="note-mode-toggle">
            <div className="left-buttons">
              {['edit', 'preview', 'both'].map(m => (
                <button key={m} onClick={() => setMode(m as typeof mode)}>
                  {m.charAt(0).toUpperCase() + m.slice(1)}
                </button>
              ))}
            </div>
            <div className="note-logo" />
            <div className="right-buttons">
              <button className="sync" onClick={load}>Sync</button>
              <button className="export" onClick={() => exportNote(id)}>Export</button>
            </div>
          </div>
          <div className="note-header">
            <input
              className="note-title"
              value={title}
              onChange={e => setTitle(e.target.value)}
              placeholder="Note title"
            />
          </div>
          <div className="note-content">
            {(mode === 'edit' || mode === 'both') && (
              <textarea
                className="note-editor"
                value={content}
                onChange={e => setContent(e.target.value)}
                placeholder="Write your note here (Markdown)..."
              />
            )}
            {(mode === 'preview' || mode === 'both') && (
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
            .catch(e => toast.error(String(e)));
        }}>
          {window.location.host + window.location.pathname}
        </button>
      </div>
    </>
  );
}
