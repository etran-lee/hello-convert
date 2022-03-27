package languages

type Lang struct {
	keys   []string
	emoji  string
	output string
}

var English = &Lang{keys: []string{"English", "english", "Eng", "eng"},
	emoji:  "\\U+1F1FA",
	output: "Hello"}
