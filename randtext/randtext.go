package randtext

import (
	"io"
	"math/rand"
	"strings"
	"text/scanner"
)

type MarkovModel struct {
	startWords     []string
	prefixSuffixes map[string][]string
}

func NewMarkovModel(reader io.Reader) (m *MarkovModel) {
	m = new(MarkovModel)

	var s scanner.Scanner
	s.Init(reader)

	s.Scan()
	prefix := s.TokenText()
	m.startWords = append(m.startWords, prefix)

	m.prefixSuffixes = make(map[string][]string)
	for s.Scan() != scanner.EOF {
		m.prefixSuffixes[prefix] = append(m.prefixSuffixes[prefix], s.TokenText())
		prefix = s.TokenText()
	}

	m.startWords = append(m.startWords, m.prefixSuffixes["."]...)

	return
}

func (m *MarkovModel) randStartWord() string {
	return m.startWords[rand.Intn(len(m.startWords))]
}

func (m *MarkovModel) randSuffix(prefix string) string {
	suffixes := m.prefixSuffixes[prefix]
	i := rand.Intn(len(suffixes))
	return suffixes[i]
}

func (m *MarkovModel) Sentence() string {
	prefix := m.randStartWord()
	buffer := []string{prefix}
	var suffix string
	for i := 0; prefix != "."; i++ {
		suffix = m.randSuffix(prefix)
		if len(suffix) > 1 {
			buffer = append(buffer, " ")
		}
		buffer = append(buffer, suffix)
		prefix = suffix
	}

	return strings.Join(buffer, "")
}
