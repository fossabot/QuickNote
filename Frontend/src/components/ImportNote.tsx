import { importNote } from "@/services/noteAPI.ts";
import React, { useCallback, useEffect, useRef, useState } from "react";
import { toast, Toaster } from "react-hot-toast";
import "./ImportNote.scss";

interface ImportNoteProps {
  callback: (to: string) => void;
}

export const ImportNote: React.FC<ImportNoteProps> = ({ callback }) => {
  const dragCounter = useRef(0);
  const [isDragging, setIsDragging] = useState(false);

  const handleFileUpload = useCallback(async (file: File) => {
    try {
      await importNote(file);
      callback(file.name.replace(/\.qnote$/, ""));
    } catch (error) {
      console.error(error);
      toast.error("Failed to import note");
    }
  }, [callback]);

  const handleDrop = useCallback(
    (e: DragEvent) => {
      e.preventDefault();
      dragCounter.current = 0;
      setIsDragging(false);
      const files = e.dataTransfer?.files;
      if (files?.length) void handleFileUpload(files[0]);
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

  return (
    <>
      <Toaster position="bottom-right" />
      <div className={`import ${isDragging ? "dragging" : ""}`}>
        <span className="drop-zone-text">Drop note here</span>
      </div>
    </>
  );
};