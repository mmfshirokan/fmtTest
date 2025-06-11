package main

import (
	"fmt"
	"io"
	"testing"
)

func BenchmarkSprintfV(b *testing.B) {
	for b.Loop() {
		fmt.Sprintf("%v", errMy)
	}
}

func BenchmarkSprintfVPlus(b *testing.B) {
	for b.Loop() {
		fmt.Sprintf("%+v", errMy)

	}
}

func BenchmarkSprintfS(b *testing.B) {
	for b.Loop() {
		fmt.Sprintf("%s", errMy)

	}
}

func BenchmarkFprintfV(b *testing.B) {
	for b.Loop() {
		_, err := fmt.Fprintf(io.Discard, "%v", errMy)
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkFprintfVPlus(b *testing.B) {
	for b.Loop() {
		_, err := fmt.Fprintf(io.Discard, "%+v", errMy)
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkFprintfS(b *testing.B) {
	for b.Loop() {
		_, err := fmt.Fprintf(io.Discard, "%s", errMy)
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkErrorfV(b *testing.B) {

	for b.Loop() {
		fmt.Errorf("%v", errMy)

	}
}

func BenchmarkErrorfVPlus(b *testing.B) {

	for b.Loop() {
		fmt.Errorf("%+v", errMy)

	}
}

func BenchmarkErrorfS(b *testing.B) {

	for b.Loop() {
		fmt.Errorf("%s", errMy)
	}
}
func BenchmarkErrorfW(b *testing.B) {

	for b.Loop() {
		fmt.Errorf("%w", errMy)
	}
}
