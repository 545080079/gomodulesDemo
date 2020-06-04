package BenchmarkDemo

import (
	"bytes"
	"fmt"
	"strings"
	"testing"
)

func StringPlus(p []string) string {
	var s string
	s += "1" + "2" + "\n"
	s += "apple\t" + "water\t" + "EOF\n"
	return s
}

func StringFmt() string {
	return fmt.Sprint("昵称", ":", "lyt", "\n", "博客", ":", "http://545080079.github.io/", "\n")
}

func StringJoin() string {
	s := []string{"昵称", ":", "lyt", "\n", "博客", ":", "http://545080079.github.io/", "\n"}
	return strings.Join(s, "")
}

func StringBuffer() string {
	var b bytes.Buffer
	b.WriteString("昵称")
	b.WriteString(":")
	b.WriteString("lyt")
	b.WriteString("\n")
	b.WriteString("博客")
	b.WriteString(":")
	b.WriteString("http://545080079.github.io/")
	b.WriteString("\n")
	return b.String()
}
func StringBuilder() string {
	var b strings.Builder
	b.WriteString("昵称")
	b.WriteString(":")
	b.WriteString("lyt")
	b.WriteString("\n")
	b.WriteString("博客")
	b.WriteString(":")
	b.WriteString("http://545080079.github.io/")
	b.WriteString("\n")
	return b.String()
}

func BenchmarkStringBuilder(b *testing.B) {
	for i := 0; i < b.N; i++ {
		StringBuilder()
	}
}
func BenchmarkStringBuffer(b *testing.B) {
	for i := 0; i < b.N; i++ {
		StringBuffer()
	}
}

func BenchmarkStringPlus(b *testing.B) {
	for i := 0; i < b.N; i++ {
		StringPlus([]string{"fa", "gk"})
	}
}

func BenchmarkStringFmt(b *testing.B) {
	for i := 0; i < b.N; i++ {
		StringFmt()
	}
}

func BenchmarkStringJoin(b *testing.B) {
	for i := 0; i < b.N; i++ {
		StringJoin()
	}
}
