import MDEditor from "@uiw/react-md-editor";
import { type ChangeEvent, useCallback, useEffect, useRef, useState } from "react";
import { toast, Toaster } from "react-hot-toast";
import { useNavigate, useParams } from "react-router-dom";
import { DarkModeToggle } from "../components/DarkModeToggle.tsx";
import { ImportNote } from "../components/ImportNote.tsx";
import { exportNote, getNote, importNote, saveNote } from "../services/noteAPI";
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
        skipSave.current = true;
        setTitle(note.title ?? "");
        setContent(note.content ?? "");
      } else {
        setTitle("");
        setContent("");
      }
      toast.success("Note loaded");
    } catch (e) {
      console.error(e);
      toast.error("Failed to load note");
    }
  }, [id]);

  useEffect(() => {
    void load();
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
        void save();
      }
    };
    window.addEventListener("keydown", onKey);
    return () => window.removeEventListener("keydown", onKey);
  }, [save]);

  const fileInputRef = useRef<HTMLInputElement | null>(null);

  return (
    <>
      <DarkModeToggle />
      <div className="content">
        <Toaster position="top-right" />
        <ImportNote callback={async (to: string) => {
          navigate(`/note/${to}`, { replace: true });
          if (to === id) {
            await load();
          }
        }}/>
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

                    const newId = file.name.replace(/\.qnote$/, "");
                    navigate(`/note/${newId}`);
                    if (newId === id) {
                      await load();
                    }

                  } catch (error) {
                    console.error(error);
                    toast.error("Failed to import note");
                  }
                }}
              />

              <button className="importButton" onClick={() => {
                fileInputRef.current?.click();
              }}>Import</button>
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
