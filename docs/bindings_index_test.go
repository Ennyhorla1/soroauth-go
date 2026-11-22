package docs

import (
	"os"
	"strings"
	"testing"
)

func TestBindingsIndexListsAllDirectories(t *testing.T) {
	bindingsDir := "../bindings"
	entries, err := os.ReadDir(bindingsDir)
	if err != nil {
		t.Fatalf("failed to read bindings directory: %v", err)
	}

	contentBytes, err := os.ReadFile("bindings-index.md")
	if err != nil {
		t.Fatalf("failed to read docs/bindings-index.md: %v", err)
	}
	content := string(contentBytes)

	for _, entry := range entries {
		if !entry.IsDir() {
			continue
		}
		name := entry.Name()
		// Check that every directory under bindings/ has an entry mentioning its name in bindings-index.md
		if !strings.Contains(content, name) {
			t.Errorf("bindings/%s exists but is not listed in docs/bindings-index.md", name)
		}
	}
}
