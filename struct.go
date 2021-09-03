package gg

import "io"

type istruct struct {
	name  string
	items *group
}

// Struct will insert a new struct.
func Struct(name string) *istruct {
	return &istruct{
		name:  name,
		items: newGroup("{", "}", "\n"),
	}
}

func (i *istruct) render(w io.Writer) {
	writeStringF(w, "type %s struct", i.name)
	i.items.render(w)
}

// Line will insert an empty line.
func (i *istruct) Line() *istruct {
	i.items.append(Line())
	return i
}

// NamedLineComment will insert a new line comment started with struct name.
func (i *istruct) NamedLineComment(content string, args ...interface{}) *istruct {
	content = i.name + " " + content
	i.items.append(LineComment(content, args...))
	return i
}

// LineComment will insert a new line comment.
func (i *istruct) LineComment(content string, args ...interface{}) *istruct {
	i.items.append(LineComment(content, args...))
	return i
}

func (i *istruct) Field(name, typ string) *istruct {
	i.items.append(&ifield{
		name:      name,
		value:     typ,
		separator: " ",
	})
	return nil
}
