# Projman (project management)
This is a program I wrote for myself due to having a ton of local programming projects sitting on my harddrive.
I struggle sometimes to find them so I decided to write a program to manage their locations and attributes.
This program provides a simple CSV database that stores the name, description, and tags of the projects you register to it.

One feature I particularly like is the ability to put the database anywhere. Through the `PROJDB_LOC` environment variable you can
put the database anywhere. This also means you can use multiple databases in different locations of you choose to.

It is written in go, so you should have that installed.

# Installing
```bash
$ git clone --depth=1 https://github.com/rfmineguy/projman-go
$ go build
$ go env -w GOBIN=/usr/local/bin
$ go install
$ projman <args>
```

# Features
```
Usage:
  projman [OPTIONS] <list | register>

Help Options:
  -h, --help  Show this help message

Available commands:
  list
  register
```

## Register
Register allows you put a new entry into the database<br>
You supply a `name`, a `path`, and a set of `tags` that describe your project

```sh
Usage:
  projman [OPTIONS] register [register-OPTIONS]

Help Options:
  -h, --help      Show this help message

[register command options]
      -n, --name=
      -p, --path=
      -t, --tags=
```

## List
List allows you to search your database of projects either by name, or by tag
- When you search by name it will suggest potential projects if you made a typo
- When you search by tag it will do the same, except if you know you want to search with more than one tag you can comma separate them

```
Usage:
  projman [OPTIONS] list [list-OPTIONS]

Help Options:
  -h, --help         Show this help message

[list command options]
      -a, --all
      -t, --by-tags=
      -n, --by-name=
```
