package dbletter

type DbLetter struct {
	Id       int    `json:"id"`
	Header   string `json:"header"`
	Text     string `json:"text"`
	Owner_id int    `json:"owner_id"`
}

func NewDbLetter(id int, header, text string, owner_id int) *DbLetter {
	newLetter := &DbLetter{
		Id:       id,
		Header:   header,
		Text:     text,
		Owner_id: owner_id,
	}

	return newLetter
}

func NewDBLetterWithoutId(header, text string, owner_id int) *DbLetter {
	newLetter := &DbLetter{
		Header:   header,
		Text:     text,
		Owner_id: owner_id,
	}

	return newLetter
}
