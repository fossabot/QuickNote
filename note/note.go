package note

func (n *Note) Write() error {
	encoded, err := n.Encode(n.Key)
	if err != nil {
		return err
	}
	n.Data = encoded
	return SetNote(*n)
}

func (n *Note) Read() error {
	no, err := GetNote(n.NID)
	if err != nil {
		return err
	}
	return n.Decode(no.Data, n.Key)
}

func (n *Note) Delete() error {
	return DeleteNote(n.NID)
}
