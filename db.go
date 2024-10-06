package main

import (
	"os"
	"bufio"
	"strings"
)

func open_db() *os.File {
	// Setup the projdb file
	projdir := os.Getenv("PROJDB_LOC")
	projloc := projdir + "/projdb";
	if _, err := os.Stat(projdir); os.IsNotExist(err) {
		os.MkdirAll(projdir, 0700);
	}
	if _, err := os.Stat(projloc); os.IsNotExist(err) {
		os.Create(projloc);
	}

	f, err := os.OpenFile(projloc, os.O_APPEND, 0600);
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
