const API_BASE = import.meta.env.VITE_API_URL;

export interface NoteData {
  nid: string;
  title: string;
  content: string;
}

export const getNote = async (nid: string, signal?: AbortSignal): Promise<NoteData | null> => {
  const response = await fetch(`${API_BASE}/notes/${nid}`, { signal });
  if (!response.ok) {
    throw new Error("Failed to fetch note");
  }
  const data = await response.json();
  return {
    nid: data.data.nid,
    title: data.data.title,
    content: data.data.content
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
};

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
  window.open(`${API_BASE}/export/${id}`);
};

export const importNote = async (file: File) => {
  const formData = new FormData();
  formData.append("import", file);
  const response = await fetch(`${API_BASE}/import`, {
    method: "POST",
    body: formData
  });
  if (!response.ok) {
    throw new Error("Failed to import note");
  }
};
