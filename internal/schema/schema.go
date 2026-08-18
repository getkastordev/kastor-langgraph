package schema

import protocol "github.com/weirdGuy/kastor/protocol/v1"

type Model = protocol.Model
type Target = protocol.Target
type Agent = protocol.Agent
type AgentInput = protocol.AgentInput
type AgentOutput = protocol.AgentOutput
type Tool = protocol.Tool
type ToolParam = protocol.ToolParam
type ToolReturns = protocol.ToolReturns
type ToolSource = protocol.ToolSource
type Prompt = protocol.Prompt
type MCPServer = protocol.MCPServer
type MCPAuth = protocol.MCPAuth

const (
	SchemeEnv        = protocol.SchemeEnv
	SchemeConnection = protocol.SchemeConnection
)

var ParseCredentialRef = protocol.ParseCredentialRef
var ParseMCPURI = protocol.ParseMCPURI
