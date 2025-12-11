package main

import (
	"testing"
)

// fakeBanner creates a minimal Banner used only for testing.
// It defines glyphs for 'A', 'B' and space ' ', each with charHeight lines.
func fakeBanner() Banner {
	b := make(Banner)

	b['A'] = []string{
		"A0", "A1", "A2", "A3", "A4", "A5", "A6", "A7",
	}
	b['B'] = []string{
		"B0", "B1", "B2", "B3", "B4", "B5", "B6", "B7",
	}
	b[' '] = []string{
		" 0", " 1", " 2", " 3", " 4", " 5", " 6", " 7",
	}

	return b
}

// ---------- Tests for LoadBanner ----------

func TestLoadBannerStandard(t *testing.T) {
	banner, err := LoadBanner("banners/standard.txt")
	if err != nil {
		t.Fatalf("LoadBanner returned an error: %v", err)
	}

	if len(banner) == 0 {
		t.Fatalf("banner map is empty")
	}

	expectedChars := int(lastChar - firstChar + 1)
	if len(banner) != expectedChars {
		t.Errorf("expected %d characters, got %d", expectedChars, len(banner))
	}

	for c := firstChar; c <= lastChar; c++ {
		art, ok := banner[rune(c)]
		if !ok {
			t.Errorf("missing character %q (%d) in banner map", c, c)
			continue
		}
		if len(art) != charHeight {
			t.Errorf("character %q has %d lines, expected %d", c, len(art), charHeight)
		}
	}
}

// ---------- Tests for RenderLine ----------

func TestRenderLine(t *testing.T) {
	b := fakeBanner()

	got := RenderLine("AB", b)
	want := "" +
		"A0B0\n" +
		"A1B1\n" +
		"A2B2\n" +
		"A3B3\n" +
		"A4B4\n" +
		"A5B5\n" +
		"A6B6\n" +
		"A7B7\n"

	if got != want {
		t.Errorf("RenderLine(\"AB\") =\n%q\nwant:\n%q", got, want)
	}
}

func TestRenderLineUnknownChar(t *testing.T) {
	b := fakeBanner()

	// There is no glyph defined for 'Z' in fakeBanner.
	// This test ensures that RenderLine does not panic and can still be called.
	_ = RenderLine("Z", b)
}

// ---------- Tests for decodeEscapedNewlines ----------

func TestDecodeEscapedNewlines(t *testing.T) {
	got := decodeEscapedNewlines("hello\\nworld\\n")
	want := "hello\nworld\n"

	if got != want {
		t.Errorf("decodeEscapedNewlines(...) = %q, want %q", got, want)
	}
}

// ---------- Tests for RenderInput ----------

func TestRenderInputSingleLine(t *testing.T) {
	b := fakeBanner()

	got := RenderInput("A B", b)
	// "A B" → glyph for 'A', then space, then 'B'.
	want := "" +
		"A0 0B0\n" +
		"A1 1B1\n" +
		"A2 2B2\n" +
		"A3 3B3\n" +
		"A4 4B4\n" +
		"A5 5B5\n" +
		"A6 6B6\n" +
		"A7 7B7\n"

	if got != want {
		t.Errorf("RenderInput(\"A B\") =\n%q\nwant:\n%q", got, want)
	}
}

func TestRenderInputWithLogicalNewlines(t *testing.T) {
	b := fakeBanner()

	got := RenderInput("A\\nB", b)
	want := "" +
		"A0\n" +
		"A1\n" +
		"A2\n" +
		"A3\n" +
		"A4\n" +
		"A5\n" +
		"A6\n" +
		"A7\n" +
		"B0\n" +
		"B1\n" +
		"B2\n" +
		"B3\n" +
		"B4\n" +
		"B5\n" +
		"B6\n" +
		"B7\n"

	if got != want {
		t.Errorf("RenderInput(\"A\\nB\") =\n%q\nwant:\n%q", got, want)
	}
}

func TestRenderInputEmpty(t *testing.T) {
	b := fakeBanner()

	got := RenderInput("", b)
	if got != "" {
		t.Errorf("RenderInput(\"\") = %q, want empty string", got)
	}
}

// ---------- Tests for validateInput ----------

func TestValidateInput_ValidASCII(t *testing.T) {
	input := "Hello 123!\nThere"
	err := validateInput(input)
	if err != nil {
		t.Errorf("expected no error for valid ASCII input, got: %v", err)
	}
}

func TestValidateInput_OnlyGreek(t *testing.T) {
	input := "Γεια σου"
	err := validateInput(input)
	if err == nil {
		t.Errorf("expected error for Greek input, got nil")
	}
}

func TestValidateInput_MixedASCIIAndGreek(t *testing.T) {
	input := "Gεια σου"
	err := validateInput(input)
	if err == nil {
		t.Errorf("expected error for mixed ASCII/Greek input, got nil")
	}
}

func TestValidateInput_Emoji(t *testing.T) {
	input := "Hello🙂"
	err := validateInput(input)
	if err == nil {
		t.Errorf("expected error for emoji input, got nil")
	}
}

func TestValidateInput_NewlinesAreAllowed(t *testing.T) {
	input := "Hello\nThere\n123!"
	err := validateInput(input)
	if err != nil {
		t.Errorf("expected no error for input with newlines, got: %v", err)
	}
}
