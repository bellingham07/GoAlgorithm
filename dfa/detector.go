package dfa

import (
	"bufio"
	"io"
)

type Detector struct {
	trie *Trie
}

func (d *Detector) Load(rd io.Reader) error {
	buf := bufio.NewScanner(bufio.NewReader(rd))
	for buf.Scan() {
		tmp := buf.Text()
		d.trie.Insert(tmp)
	}

	return nil
}

func (d *Detector) AddWord(word string) {
	d.trie.Insert(word)
}

func (d *Detector) DelWord(word string) {
	d.trie.Delete(word)
}

func (d *Detector) Detect(text string) bool {
	return d.trie.Check(text)
}

func (d *Detector) Search(text string) []string {
	return d.trie.Search(text)
}

func (d *Detector) Match(text string) []Result {
	return d.trie.Match(text)
}

func (d *Detector) Filter(text string, replace ...string) string {
	return d.trie.Filter(text, replace...)
}

func (d *Detector) GetAllWords() []string {
	return d.trie.GetAllWords()
}

func New() *Detector {
	return NewWithOptions(&Options{})
}

type Options struct {
	IgnoreCase bool
	Noises     []rune
}

func NewWithOptions(opts *Options) *Detector {
	d := new(Detector)
	trie := NewTrie()
	trie.IgnoreCase = opts.IgnoreCase
	if len(opts.Noises) > 0 {
		for _, r := range opts.Noises {
			trie.Noises[r] = struct{}{}
		}
	}
	d.trie = trie

	return d
}
