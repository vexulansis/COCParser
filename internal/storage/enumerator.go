package storage

type Enumerator struct {
	Refstring string
	Storage   *Storage
}

func NewEnumerator(ref string, storage *Storage) *Enumerator {
	return &Enumerator{
		Refstring: ref,
		Storage:   storage,
	}
}
func (e *Enumerator) Execute() error {
	e.Storage.Prepare(2954)
	return nil
}
