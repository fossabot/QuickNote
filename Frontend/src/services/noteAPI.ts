// .env.development
// .env.production
// .env
const API_BASE = import.meta.env.VITE_API_URL;

export interface NoteData {
  nid: string;
  title: string;
  content: string;
}

export const getNote = async (nid: string): Promise<NoteData | null> => {
  const response = await fetch(`${API_BASE}/notes/${nid}`);
  if (!response.ok) {
    throw new Error("Failed to fetch note");
  }
  const result = await response.json();
  return {
    nid: result.data.nid,
    title: result.data.title,
    content: result.data.content
  };
};

export const saveNote = async (nid: string, title: string, content: string) => {
  const response = await fetch(`${API_BASE}/notes/${nid}`, {
    method: "POST",
    headers: {
      "Content-Type": "application/json"
    },
    body: JSON.stringify({
      title,
      content
    })
  });

  if (!response.ok) {
    throw new Error("Failed to create note");
  }
  return true;
};

// save an empty title and content note can be deleted...(idk why create this api)
export const deleteNote = async (id: string) => {
  const response = await fetch(`${API_BASE}/notes/${id}`, {
    method: "DELETE"
  });
  if (!response.ok) {
    throw new Error("Failed to delete note");
  }
  return true;
};

export const exportNote = async (id: string) => {
  const response = await fetch(`${API_BASE}/export/${id}`);
  if (!response.ok) {
    throw new Error("Failed to export note");
  }
  return window.open(`${API_BASE}/export/${id}`);
};

export const importNote = async (file: File) => {
  if (!file.name.endsWith(".qnote")) throw new Error(`Invalid file: ${file.name}`);
  const fileData = new FormData();
  fileData.append("import", file);
  const response = await fetch(`${API_BASE}/import`, {
    method: "POST",
    body: fileData
  });
  if (!response.ok) {
    throw new Error("Failed to import note");
  }
  return true;
};
