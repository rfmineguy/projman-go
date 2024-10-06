package main

import (
	"fmt"
	"strings"
)

type Project struct {
	name string
	desc string
	path string
	tags []string
}

func (p Project) write() string {
	return fmt.Sprintf("%s, %s, %s", p.name, p.desc, strings.Join(p.tags[:], ", "));
}

func (p *Project) read(s string) *Project {
	slice := strings.Split(s, ", ");
	for i, elem := range slice {
		switch i {
		case 0: p.name = elem; // name
		case 1: p.desc = elem; // description
		default: p.tags = append(p.tags, elem); // tags are the rest
		}
	}
	return p;
}

func lastN(s string, n int) string {
	if n > len(s) { return s }
	return s[len(s)-n:]
}

func stringify(s []string) string {
	return strings.Join(s, ",");
}

/**
 +-------------------------------+
 | Name: GFXLib                  |
 | Desc: A program for           |
 |    making 2D games            |
 | Path: /Users/rileyfischer/... |
 | Tags: c, c++, java            |
 +-------------------------------+

 Make the program return the path so its in $? for the caller
*/
func (p Project) display() {
	fmt.Printf("╭%s╮\n", strings.Repeat("─", 6 + 20 + 2));
	fmt.Printf("│ Name: %20s │\n", p.name);
	fmt.Printf("│ Desc: %20s │\n", p.desc);
	fmt.Printf("│ Path: ...%17s │\n", lastN(p.path, 16));
	fmt.Printf("│ Tags: %20v │\n", lastN(stringify(p.tags), 19));
	fmt.Printf("╰%s╯\n", strings.Repeat("─", 6 + 20 + 2));
}
