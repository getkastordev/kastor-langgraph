# Kastor LangGraph Plugin

External LangGraph code-generation target for [Kastor](https://github.com/weirdGuy/kastor).

The plugin owns LangGraph-specific validation, Python generation, dependency mapping, MCP bindings, and runtime-tool scaffolds. Kastor core sends canonical protocol-v1 IR and remains responsible for deterministic disk writes.

> Status: pre-release. The protocol-v1 implementation is available for integration testing but no stable binary has been released yet.

## Development

```sh
go test ./...
go vet ./...
go build ./cmd/kastor-langgraph
```

Point a local Kastor checkout at the binary with:

```sh
export KASTOR_PLUGIN_LANGGRAPH=/absolute/path/to/kastor-langgraph
```

The source address is `github.com/getkastordev/kastor-langgraph`.

## License

Apache-2.0
