package main

import (
	"fmt"
	"strings"
)

func (reg *Register) run() {
	fmt.Println("Register: ", *reg);
}

func (list *List) run() {
	f := open_db();
	m := parse_db(f);
	defer f.Close()

	if list.All {
		for _, elem := range(m) {
			elem.display();
		}
	}
	if list.Name != "" {
		val, ok := m[list.Name];
		if ok {
			m[val.name].display();
		} else {
			for _, elem := range(m) {
				ham := hammingDistance(elem.name, list.Name);
				if ham < 4 {
					fmt.Println("Did you mean '" + elem.name + "'?");
					return;
				}
			}
		}
		fmt.Println("No project with name close to '" + list.Name + "'?");
	}
	if list.Tags != "" {
		taglist := strings.Split(list.Tags, ",");
		for _, elem := range(m) {
			if isSubset(taglist, elem.tags) {
				elem.display();
			}
		}
	}
}
