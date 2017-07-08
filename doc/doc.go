package doc

type Char struct {
	r       rune
	visible bool
}

type Doc struct {
	chars []Char
}

func New(content string) *Doc {
	return &Doc{chars: stringToChars(content)}
}

func stringToChars(s string) []Char {
	var chars []Char
	for _, r := range s {
		chars = append(chars, Char{r: r, visible: true})
	}
	return chars

}

func appendChars(chars []Char, s string) []Char {
	for _, r := range s {
		chars = append(chars, Char{r: r, visible: true})
	}
	return chars
}

// view model
func (d *Doc) String() string {
	var runes []rune
	for _, c := range d.chars {
		if c.visible {
			runes = append(runes, c.r)
		}
	}
	return string(runes)
}

func (d *Doc) viewToModel(at int) int {
	// for _, c := range d.chars {

	// }
	return 0
}

type Op interface {
	ExecuteLocal(doc *Doc)
}

type InsertOp struct {
	at     int
	insert string
}

func (op *InsertOp) ExecuteLocal(doc *Doc) {
	at := op.at - 1 // convert to 0 index
	insert := stringToChars(op.insert)
	doc.chars = append(doc.chars[:at], append(insert, doc.chars[at:]...)...)
}

type DeleteOp struct {
	at     int
	length int
}

func (op *DeleteOp) ExecuteLocal(doc *Doc) {
	at := op.at - 1 // convert to 0 index
	for i := at; i < at+op.length; i++ {
		doc.chars[i].visible = false
	}
}
