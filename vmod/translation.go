package vmod

type MultiLang struct {
	Language    []string          `bson:"language" json:"language"`
	Translation map[string]string `bson:"translation" json:"translation"`
}

func (m *MultiLang) Insert() *MultiLang {
	var temp = []string{}
	for i := range m.Translation {
		temp = append(temp, i)
	}
	m.Language = temp
	return m
}
