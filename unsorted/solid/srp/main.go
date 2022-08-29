package main

import (
	"fmt"
	"io/ioutil"
	"net/url"
	"strings"
)

/*
and example adhering to single responsibility principal

a journal single responsibility is to maintain a list of entries
a journals only reason to change is the add or remove entries from the list
a journal type includes functions which specifically pertain to the responsibility of the journal and it's reasons to change
*/

var entryCount = 0

type Journal struct {
	entries []string
}

func (j *Journal) AddEntry(text string) int {
	entryCount++
	entry := fmt.Sprintf("%d: %s",
		entryCount,
		text)
	j.entries = append(j.entries, entry)
	return entryCount
}
func (j *Journal) RemoveEntry(index int) {
	// ...
}
func (j *Journal) String() string {
	return strings.Join(j.entries, "\n")
}

/*
and example violation to single responsibility principal

persistence of the journal is not the responsibility of the journal
*/
func (j *Journal) Load(filename string) {
}
func (j *Journal) LoadFromWeb(url *url.URL) {
}
func (j *Journal) Save(filename string) {
	_ = ioutil.WriteFile(filename, []byte(j.String()), 0644)
}

/*
and example adherence to single responsibility principal

create another type whose single responsibility is to manage persistence
*/
type Persistence struct {
	lineSeparator string
}

func (p *Persistence) saveToFile(j *Journal, filename string) {
	_ = ioutil.WriteFile(filename,
		[]byte(strings.Join(j.entries, p.lineSeparator)), 0644)
}

func main() {
	j := Journal{}
	j.AddEntry("I cried today.")
	j.AddEntry("I ate a bug")
	fmt.Println(strings.Join(j.entries, "\n"))

	p := Persistence{"\n"}
	p.saveToFile(&j, "journal.txt")
}
