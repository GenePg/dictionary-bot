package oxfordApi

type Sense struct {
	Definitions      []string   `json:"definitions"`
	Examples         []struct{} `json:"examples"`
	Id               string     `json:"id"`
	ShortDefinitions []string   `json:"shortDefinitions"`
	Subsenses        []struct{} `json:"subsenses"`
	ThesaurusLinks   []struct{} `json:"thesaurusLinks"`
}

type Entry struct {
	Senses []Sense `json:"senses"`
}

type LexicalCategory struct {
	Id   string `json:"id"`
	Text string `json:"text"`
}

type LexicalEntry struct {
	Entries         []Entry         `json:"entries"`
	Language        string          `json:"language"`
	LexicalCategory LexicalCategory `json:"lexicalCategory"`
	Pronunciation   string          `json:"pronunciation"`
	Text            string          `json:"text"`
}

type Result struct {
	Id             string         `json:"id"`
	Language       string         `json:"language"`
	LexicalEntries []LexicalEntry `json:"lexicalEntries"`
	ResultType     string         `json:"type"`
	Word           string         `json:"word"`
}

type Metadata struct {
	Operation string `json:"operation"`
	Provider  string `json:"provider"`
	Schema    string `json:"schema"`
}

type EntriesDef struct {
	Id       string   `json:"id"`
	Metadata Metadata `json:"metadata"`
	Results  []Result `json:"results"`
	Word     string   `json:"word"`
}
