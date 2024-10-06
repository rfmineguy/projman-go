package main

// #run methods implemented in commands.go

type Register struct {
	Name string `short:"n" long:"name" required:"true"`
	Path string `short:"p" long:"path" required:"true"`
	Tags string `short:"t" long:"tags" required:"true"`
	Enabled bool `hidden:"true" no-ini:"true"`
}
func (c *Register) Execute(args []string) error {
	c.Enabled = true;
	return nil;
}

type List struct {
	All bool `short:"a" long:"all" required:"false"`
	Tags string `short:"t" long:"by-tags" required:"false"`
	Name string `short:"n" long:"by-name" required:"false"`
	Enabled bool `hidden:"true" no-ini:"true"`
}
func (c *List) Execute(args []string) error {
	c.Enabled = true;
	return nil;
}


var opts struct {
	Register Register `command:"register"`
	List List  `command:"list"`
}
