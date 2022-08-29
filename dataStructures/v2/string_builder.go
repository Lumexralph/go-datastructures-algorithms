package v2

import "strings"

type stringBuilder struct {
	store []string
}

func (sb *stringBuilder) append(words ...string) {
	sb.store = append(sb.store, words...)
}

func (sb *stringBuilder) value() string {
	return strings.Join(sb.store, "")
}
