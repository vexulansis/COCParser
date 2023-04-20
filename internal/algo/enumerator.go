package algo

import (
	db "github.com/vexulansis/COCParser/internal/storage"
	"github.com/vexulansis/COCParser/pkg/api"
)

type Enumerator struct {
	Refstring string
	Storage   *db.Storage
}

func NewEnumerator(ref string, storage *db.Storage) *Enumerator {
	return &Enumerator{
		Refstring: ref,
		Storage:   storage,
	}
}
func (e *Enumerator) Execute() error {
	err := e.GenerateTags(3)
	if err != nil {
		return err
	}
	return nil
}
func (e *Enumerator) GenerateTags(length int) error {
	err := e.GenerateTagsRec(length, "")
	if err != nil {
		return err
	}
	return nil
}
func (e *Enumerator) GenerateTagsRec(length int, prefix string) error {
	if length == 0 {
		clan, err := api.GetClanByTag(prefix)
		if err != nil {
			return err
		}
		err = e.Storage.AddClan(clan)
		if err != nil {
			return err
		}
		return nil
	}
	for _, l := range e.Refstring {
		newPrefix := prefix + string(l)
		e.GenerateTagsRec(length-1, newPrefix)
	}
	return nil
}
