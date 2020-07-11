package structs

type WebError struct {
	Msg       string
	Url       string
	LineNo    uint16
	ColumnNo  uint16
	Err     string
}