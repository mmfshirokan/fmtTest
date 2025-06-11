package main

import (
	"fmt"
	"io"
	"strings"
	"testing"
)

func BenchmarkV(b *testing.B) {
	for b.Loop() {
		b := fmt.Sprintf("%v", errMy)
		strings.Compare(b, "")
	}
}

func BenchmarkVPlus(b *testing.B) {
	for b.Loop() {
		b := fmt.Sprintf("%+v", errMy)
		strings.Compare(b, "")
	}
}

func BenchmarkS(b *testing.B) {
	for b.Loop() {
		b := fmt.Sprintf("%s", errMy)
		strings.Compare(b, "")
	}
}

func BenchmarkPrintV(b *testing.B) {
	for b.Loop() {
		_, err := fmt.Fprintf(io.Discard, "%v", errMy)
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkPrintVPlus(b *testing.B) {
	for b.Loop() {
		_, err := fmt.Fprintf(io.Discard, "%+v", errMy)
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkPrintS(b *testing.B) {
	for b.Loop() {
		_, err := fmt.Fprintf(io.Discard, "%s", errMy)
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkStdErrPrintV(b *testing.B) {
	for b.Loop() {
		_, err := fmt.Fprintf(io.Discard, "%v", errMy)
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkStdErrPrintVPlus(b *testing.B) {
	for b.Loop() {
		_, err := fmt.Fprintf(io.Discard, "%+v", errMy)
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkStdErrPrintS(b *testing.B) {
	for b.Loop() {
		_, err := fmt.Fprintf(io.Discard, "%s", errMy)
		if err != nil {
			b.Fatal(err)
		}
	}
}
