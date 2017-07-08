package doc

// type op struct {
// 	version int
// 	offset  int
// 	// deleteLength  int
// 	insertContent []byte
// }

// type insertOp struct {
// 	version int
// 	offset  int
// 	content []byte
// }

// type deleteOp struct {
// 	version int
// 	offset  int
// 	length  int
// }

// type doc struct {
// 	version int
// 	pending []op
// 	content []byte
// 	// offsets []byte
// }

// func (doc *doc) applyFromDownstream(op op) error {
// 	doc.pending = append(doc.pending)
// 	if doc.version != op.version {
// 		return fmt.Errorf("versions don't match: %d != %d", doc.version, op.version)
// 	}

// 	doc.version += len(op.insertContent)

// 	tail := doc.content[op.offset:]
// 	doc.content = append(doc.content, op.insertContent...)
// 	copy(doc.content[op.offset:], op.insertContent)
// 	copy(doc.content[op.offset+len(op.insertContent):], tail)
// 	return nil
// }

// func (doc *doc) applyFromUpstream(op op) error {

// }

// func (doc *doc) ackFromUpstream() error {
// 	// TODO: memory leak?
// 	if len(doc.pending) == 0 {
// 		return fmt.Errorf("received ack for nothing")
// 	}
// 	doc.pending = doc.pending[1:]
// 	return nil
// }

type Char struct {
	r    rune
	next int
}

type Doc struct {
	content string
}

func New(content string) *Doc {
	return &Doc{content: content}
}

// Insert b after location p.
// The first position is index.
func (d *Doc) Ins(at int, insert string) {
	at = at - 1 // convert to 0 index
	d.content = d.content[:at] + insert + d.content[at:]
}

func (d *Doc) Del(at, length int) {
	at = at - 1 // convert to 0 index
	d.content = d.content[:at] + d.content[at+length:]
}

func (d *Doc) String() string {
	return string(d.content)
}
