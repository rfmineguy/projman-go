package main

import (
    "github.com/jessevdk/go-flags"
)

func main() {
	flags.Parse(&opts);

	switch {
	case opts.List.Enabled: opts.List.run();
	case opts.Register.Enabled: opts.Register.run();
	case opts.Modify.Enabled: opts.Modify.run();
	}
}
