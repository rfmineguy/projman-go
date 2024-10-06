# Projman (project management)
This is a program I wrote for myself due to having a ton of local programming projects sitting on my harddrive.
I struggle sometimes to find them so I decided to write a program to manage their locations and attributes.
This program provides a simple CSV database that stores the name, description, and tags of the projects you register to it.

One feature I particularly like is the ability to put the database anywhere. Through the `PROJDB_LOC` environment variable you can
put the database anywhere. [See Database section](#database) This also means you can use multiple databases in different locations of you choose to.

It is written in go, so you should have that installed.

# Database
projman supplies a feature that lets you change where the active database is located. So lets say you want a database that manages only your c projects. You could put a database in that folder by setting `PROJDB_LOC` to your c projects folder.
Then say you want one that manages go projects. You could initialize a new database there!

Now whenever you want to switch what database projman uses you just set `PROJDB_LOC` and that's it!

# Installing
```bash
$ git clone --depth=1 https://github.com/rfmineguy/projman-go
$ go build
$ go env -w GOBIN=/usr/local/bin
$ go install   # May require sudo priviledges
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
      -n, --name= The name of the project
      -d, --desc= The description of the project
      -p, --path= The path of the project on your computer
      -t, --tags= The tags you wish to associate with your project
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
      -a, --all      Listing of all registered projects (can be long)
      -t, --by-tags= Listing of projects registered with the supplied tags
      -n, --by-name= Listing of projects registered with the supplied name
```
