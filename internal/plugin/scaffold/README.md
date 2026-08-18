# Your Kastor LangGraph module

This starter was supplied by `kastor-langgraph` through `kastor new`. It
contains one agent, one MCP tool, and one prompt and builds without edits.

```sh
kastor validate
kastor build
cd gen/langgraph
python -m venv .venv
. .venv/bin/activate
pip install -r requirements.txt
export OPENAI_API_KEY=sk-...
python main.py researcher --inputs '{"question":"What is HCL?"}'
```

Edit the Kastor source and rebuild instead of editing generated files.
