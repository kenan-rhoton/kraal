package frame

import (
	"errors"
	"github.com/kenan-rhoton/kraal/parser"
	"io/ioutil"
	"path/filepath"
)

type Frame struct {
	Document *FrameObject
	Sources  []*Source
}

type Source struct {
	Data string
	Type string
}

func LoadFile(input string) *Frame {
	f := &Frame{Document: nil, Sources: make([]*Source, 0)}
	f.AddSource(input)
	for i := 0; i < len(f.Sources); i++ {
		f.Process(f.Sources[i])
	}
	return f
}

func (f *Frame) AddSource(input string) {
	file, err := ioutil.ReadFile(input)
	if err != nil {
		return
	}
	f.Sources = append(f.Sources, &Source{
		Data: string(file),
		Type: filepath.Ext(input)})
}

func (f *Frame) Process(s *Source) error {
	switch {
	case s.Type == "html":
		d := parser.ParseDOM(s.Data)
		f.Document = BuildObject(d)
	case s.Type == "css":
		if f.Document == nil {
			return errors.New("missing document")
		}
		css := parser.ParseCSS(s.Data)
		f.Document.ApplyStyles(css)
	case s.Type == "js":
		//Not yet!
	}
	return nil
}
