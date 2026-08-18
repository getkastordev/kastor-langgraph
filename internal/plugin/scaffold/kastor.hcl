kastor {
  required_plugins {
    langgraph = {
      source  = "github.com/getkastordev/kastor-langgraph"
      version = "~> 0.1"
    }
  }
}

model "fast" {
  provider = "openai"
  id       = "gpt-4o-mini"

  params {
    temperature = 0.2
    max_tokens  = 4096
  }
}

target "langgraph" {
  type   = "codegen"
  plugin = "langgraph"
  output = "./gen/langgraph"
}

mcp_server "fetch" {
  transport = "stdio"
  command   = "uvx"
  args      = ["mcp-server-fetch"]
}
