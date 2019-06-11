package main

func main() {
	s, err := ReadClipboard()
	if err != nil {
		panic(err)
	}
	c := NewLinguaClient("https://api.lingualeo.com/gettranslates")
	translation, err := c.GetTranslation(s)
	if err != nil || translation == "" {
		panic(err)
	}
	err = Draw(translation)
	if err != nil {
		panic(err)
	}
}
