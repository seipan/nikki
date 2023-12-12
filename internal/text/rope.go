package text

import "github.com/vinzmay/go-rope"

// To handle a large amount of text data, we use a data structure called 'rope' for storing it in a database.
// Rope is an effective algorithm for long text data, commonly adopted in text editors.
type textRope struct {
	rope *rope.Rope
}

// String returns the text as a string.
func (t *textRope) String() string {
	return t.rope.String()
}

func NewText(text string) Text {
	return &textRope{
		rope.New(text),
	}
}
