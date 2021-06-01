package vmod

type (
	ContentText    map[string]string
	DisplayContent struct {
		Language    []string               `bson:"language" json:"language"`
		Keys        []string               `bson:"keys" json:"keys"`
		Translation map[string]ContentText `bson:"translation" json:"translation"`
		Default     string                 `bson:"default" json:"default"`
	}
)

func (dc *DisplayContent) Validate() *DisplayContent {
	var language = []string{}
	var keys = []string{}
	var single = make(map[string]bool)
	for i := range dc.Translation {
		language = append(language, i)
		for j := range dc.Translation[i] {
			if single[j] != true {
				keys = append(keys, j)
				single[j] = true
			}
		}
	}
	dc.Language = language
	dc.Keys = keys
	return dc
}
