package main

import (
	"fmt"
	"strings"
	"strconv"
	"os"
)

func (reg *Register) run() {
	f := open_db();
	defer f.Close();
	m := parse_db(f);

	_, ok := m[reg.Name];
	if ok {
		fmt.Print("Project with name '", reg.Name, "' already exists");
		return;
	}

	p := Project{name: reg.Name, desc: reg.Desc, path: reg.Path, tags: strings.Split(reg.Tags, ",")};
	m[reg.Name] = p;

	s := write_db(m);
	_, err := f.WriteString(s);
	if err != nil {
		fmt.Println(err);
	}
	f.Sync();
	f.Close();

	fmt.Println("Regisitered: ", p);
	p.display();
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
		projects := []Project{};
		for _, elem := range(m) {
			ham := hammingDistance(elem.name, list.Name);
			if ham < 4 {
				projects = append(projects, elem);
			}
		}
		for _, p := range(projects) {
			p.display();
		}
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

func getuserint(maxChoice int) int {
	var choice string
	for {
		fmt.Print("Enter choice (default=0): ");
		_, err := fmt.Scanln(&choice);
		if choice == "" {
			return 0;
		}
		i, err := strconv.Atoi(choice);
		if err != nil {
			fmt.Println("Enter valid number");
		} else if i < 0 || i >= maxChoice {
			fmt.Println("Out of range");
		} else {
			return i;
		}
	}
}

func modifyProject(project Project) Project {
	fmt.Println("Modifying project");
	project.display();

	var tags string
	fmt.Print("Name (", project.name, "): ");
	fmt.Scanln(&project.name);
	fmt.Print("Path (", project.path, "): ");
	fmt.Scanln(&project.path);
	fmt.Print("Desc (", project.desc, "): ");
	fmt.Scanln(&project.desc);
	fmt.Print("Tags (", project.tags, "): ");
	fmt.Scanln(&tags);
	if tags != "" {
		project.tags = strings.Split(tags, ",");
	}

	return project;
}

func (modify *Modify) run() {
	f := open_db();
	m := parse_db(f);
	if len(m) == 0 {
		fmt.Println("No projects registered to modiy");
		os.Exit(1);
	}
	defer f.Close();

	if modify.Name != "" {
		projects := []Project{};
		// Get projects to suggest
		for _, elem := range(m) {
			ham := hammingDistance(elem.name, modify.Name);
			if ham < 4 {
				projects = append(projects, elem);
			}
		}

		// Get modified project
		fmt.Println("Which project would you like to modify?");
		for i, p := range(projects) {
			fmt.Println(" ", i, " - ", p.name);
		}
		choice := getuserint(len(projects));
		project := modifyProject(projects[choice]);

		// Serialize changes
		delete(m, projects[choice].name);
		m[modify.Name] = project;
		s := write_db(m);
		if err := f.Truncate(0); err != nil {
			fmt.Println("Couldnt empty file");
			os.Exit(2);
		}
		_, err := f.WriteString(s);
		if err != nil {
			fmt.Println(err);
		}
		f.Sync();
		f.Close();

		fmt.Println("Successfully modified project");
	}
}
