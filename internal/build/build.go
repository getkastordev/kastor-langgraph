package build

import (
	"path"

	"github.com/getkastordev/kastor-langgraph/internal/module"
	"github.com/getkastordev/kastor-langgraph/internal/schema"
	protocol "github.com/weirdGuy/kastor/protocol/v1"
)

const SidecarSuffix = ".kastor-new"

type File = protocol.File

type Job struct {
	Module *module.Module
	Target *schema.Target
}

type Generator interface {
	Generate(*Job) ([]File, error)
}

func SidecarFor(file string) string { return path.Clean(file) + SidecarSuffix }
