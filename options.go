package main

import "flag"

type Options struct {
	ctags  string
	cscope string
	raw    bool
	o      string
}

func parseOptions() Options {
	opts := Options{}
	flag.StringVar(&opts.ctags, "ctags", "./tags",
		"path to tags file")
	flag.StringVar(&opts.cscope, "cscope", "./cscope.out",
		"path to cscope.out file")
	flag.StringVar(&opts.o, "o", "-",
		"output file (- for stdout)")
	flag.BoolVar(&opts.raw, "raw", false,
		"output raw json graph")
	flag.Parse()
	return opts
}
