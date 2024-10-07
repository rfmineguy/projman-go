package main

import (
	"os"
	"bufio"
	"strings"
	"fmt"
)

func open_db() *os.File {
	// Setup the projdb file
	projdir := os.Getenv("PROJDB_LOC")
	if projdir == "" {
		fmt.Println("Must assign PROJDB_LOC");
		fmt.Println("   i.e.   export PROJDB_LOC=path");
		os.Exit(1);
	}
	projloc := projdir + "/projdb";
	if _, err := os.Stat(projdir); os.IsNotExist(err) {
		os.MkdirAll(projdir, 0770);
	}
	if _, err := os.Stat(projloc); os.IsNotExist(err) {
		os.Create(projloc);
		fmt.Println("Create ", projloc);
	}

	f, err := os.OpenFile(projloc, os.O_APPEND | os.O_RDWR, 0666);
	if err != nil {
		panic(err);
	}
	return f;
}

func parse_db(f *os.File) map[string]Project {
	scanner := bufio.NewScanner(f);
	m := map[string]Project {}
	for scanner.Scan() {
		if scanner.Text() == "Empty" { break; }
		// fmt.Println(" +", scanner.Text());
		slice := strings.Split(scanner.Text(), ", ");
		p := Project{};
		m[slice[0]] = *p.read(scanner.Text())
	}
	return m;
}

func write_db(m map[string]Project) string {
	var out strings.Builder;
	for _, v := range(m) {
		out.WriteString(v.write());
		out.WriteString("\n");
	}
	return out.String();
}
