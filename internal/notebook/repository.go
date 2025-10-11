package notebook

type Repository interface {
	GetNote(id string) (string, error)
	PostNote(id string, message string) (error)
	PutNote(id string, message string) (bool, error)
}
