package main

// #run methods implemented in commands.go

type Register struct {
	Name string `short:"n" long:"name" required:"true" description:"The name of the project"`
	Desc string `short:"d" long:"desc" required:"false" description:"The description of the project"`
	Path string `short:"p" long:"path" required:"true" description:"The path of the project on your computer"`
	Tags string `short:"t" long:"tags" required:"true" description:"The tags you wish to associate with your project"`
	Enabled bool `hidden:"true" no-ini:"true"`
}
func (c *Register) Execute(args []string) error {
	c.Enabled = true;
	return nil;
}

type Modify struct {
	Enabled bool `hidden:"true" no-ini:"true"`
	Name string `short:"n" long:"name" required:"true" description:"Name of project to modify"`
}
func (c *Modify) Execute(args []string) error {
	c.Enabled = true;
	return nil;
}

type List struct {
	All bool `short:"a" long:"all" required:"false" description:"Listing of all registered projects (can be long)"`
	Tags string `short:"t" long:"by-tags" required:"false" description:"Listing of projects registered with the supplied tags"`
	Name string `short:"n" long:"by-name" required:"false" description:"Listing of projects registered with the supplied name"`
	Enabled bool `hidden:"true" no-ini:"true"`
}
func (c *List) Execute(args []string) error {
	c.Enabled = true;
	return nil;
}

var opts struct {
	Register Register `command:"register"`
	Modify Modify `command:"modify"`
	List List  `command:"list"`
}
