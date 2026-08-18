package plugin

import (
	"context"
	"sort"
	"testing"

	protocol "github.com/weirdGuy/kastor/protocol/v1"
)

func TestGenerateMinimalModule(t *testing.T) {
	handler := Handler{}
	request := &protocol.GenerateRequest{
		Module: &protocol.Module{},
		Target: &protocol.Target{Name: "python", Type: "codegen"},
	}
	first, err := handler.Generate(context.Background(), request)
	if err != nil {
		t.Fatal(err)
	}
	second, err := handler.Generate(context.Background(), request)
	if err != nil {
		t.Fatal(err)
	}
	if len(first.Files) != len(second.Files) {
		t.Fatalf("non-deterministic file count: %d != %d", len(first.Files), len(second.Files))
	}
	var paths []string
	for i := range first.Files {
		if first.Files[i].Path != second.Files[i].Path || string(first.Files[i].Data) != string(second.Files[i].Data) {
			t.Fatalf("generation differs at index %d", i)
		}
		paths = append(paths, first.Files[i].Path)
	}
	sort.Strings(paths)
	want := []string{"README.md", "agents/__init__.py", "main.py", "models.py", "prompts/__init__.py", "requirements.txt", "tools/__init__.py"}
	if len(paths) != len(want) {
		t.Fatalf("paths = %v", paths)
	}
	for i := range want {
		if paths[i] != want[i] {
			t.Fatalf("paths = %v", paths)
		}
	}
}

func TestValidateRejectsConnectionCredential(t *testing.T) {
	response, err := (Handler{}).Validate(context.Background(), &protocol.ValidateRequest{
		Target: &protocol.Target{Name: "python", Type: "codegen"},
		Module: &protocol.Module{
			Tools:      []*protocol.Tool{{Name: "search", Source: &protocol.ToolSource{Kind: "mcp", URI: "mcp://search/query"}}},
			MCPServers: []*protocol.MCPServer{{Name: "search", Auth: []*protocol.MCPAuth{{Ref: "connection://cred_123"}}}},
		},
	})
	if err != nil {
		t.Fatal(err)
	}
	if len(response.Diagnostics) != 1 || response.Diagnostics[0].Addr != "mcp_server.search" {
		t.Fatalf("diagnostics = %#v", response.Diagnostics)
	}
}

func TestScaffoldIsDeterministic(t *testing.T) {
	handler := Handler{}
	first, err := handler.Scaffold(context.Background(), &protocol.ScaffoldRequest{Name: "demo"})
	if err != nil {
		t.Fatal(err)
	}
	second, err := handler.Scaffold(context.Background(), &protocol.ScaffoldRequest{Name: "demo"})
	if err != nil {
		t.Fatal(err)
	}
	if len(first.Files) != 5 || len(second.Files) != len(first.Files) {
		t.Fatalf("files = %d, %d", len(first.Files), len(second.Files))
	}
	for i := range first.Files {
		if first.Files[i].Path != second.Files[i].Path || string(first.Files[i].Data) != string(second.Files[i].Data) {
			t.Fatalf("scaffold differs at %d", i)
		}
	}
}
