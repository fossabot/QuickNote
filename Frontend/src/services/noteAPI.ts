import { toast } from 'react-hot-toast'

const API_BASE = '/v1'

export interface NoteData {
  nid: string
  title: string
  content: string
}

export const getNote = async (nid: string): Promise<NoteData | null> => {
  try {
    const response = await fetch(`${API_BASE}/notes/${nid}`)
    if (!response.ok) {
      throw new Error('Failed to fetch note')
    }
    const data = await response.json()
    return {
      nid: data.data.nid,
      title: data.data.title,
      content: data.data.content,
    }
  } catch (error) {
    console.error('Fetch error:', error)
    toast.error('Failed to load note')
    return null
  }
}

export const saveNote = async (nid: string, title: string, content: string): Promise<boolean> => {
  try {
    const response = await fetch(`${API_BASE}/notes/${nid}`, {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
      },
      body: JSON.stringify({
        title,
        content,
      }),
    })

    if (!response.ok) {
      throw new Error('Failed to save note')
    }
    return true
  } catch (error) {
    console.error('Save error:', error)
    toast.error('Failed to save note')
    return false
  }
}
