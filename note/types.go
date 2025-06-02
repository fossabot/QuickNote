package note

type Note struct {
	ID string

	Title   []byte
	Content []byte
}

type DisplayNote struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}
