package maps

const (
	ErrNotFound         = DictionaryErr("word not found")
	ErrExistentWord     = DictionaryErr("word already exists")
	ErrWordDoesNotExist = DictionaryErr("word does not exists")
)

type DictionaryErr string

func (e DictionaryErr) Error() string {
	return string(e)
}
