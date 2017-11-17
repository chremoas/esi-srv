package command

import (
	proto "github.com/abaeve/chremoas/proto"
	"golang.org/x/net/context"
)

type Command struct {
	//Store anything you need the Help or Exec functions to have access to here
	name    string
}

func (c *Command) Help(ctx context.Context, req *proto.HelpRequest, rsp *proto.HelpResponse) error {
	rsp.Usage = c.name
	rsp.Description = "Simply prints out 'Chremoas Command Template'"
	return nil
}

func (c *Command) Exec(ctx context.Context, req *proto.ExecRequest, rsp *proto.ExecResponse) error {
	rsp.Result = []byte("Chremoas Command Template")
	return nil
}

func NewCommand(name string) *Command {
	newCommand := Command{name: name}
	return &newCommand
}