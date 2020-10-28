package data

type SourceMarker struct {
	LineContents    string
	SourceLineIndex int
}

func NewSourceMarker(lineContents string, sourceLineIndex int) SourceMarker {
	return SourceMarker{LineContents: lineContents, SourceLineIndex: sourceLineIndex}
}

type Template []SourceMarker
