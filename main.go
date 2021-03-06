////////////////////////////////////////////////////////////////////////////
// Program: xmlfmt
// Purpose: XML Formatter
// Authors: Antonio Sun (c) 2016-2019, All rights reserved
////////////////////////////////////////////////////////////////////////////

package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/defektive/go-xmlfmt"
	"github.com/mkideal/cli"
	//clix "github.com/mkideal/cli/ext"
)

////////////////////////////////////////////////////////////////////////////
// Global variables definitions

var progname = "xmlfmt"
var date = "2019-12-08"

var rootArgv *rootT

////////////////////////////////////////////////////////////////////////////
// xmlfmt

func main() {
	if err := cli.Root(root).Run(os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
	fmt.Println("")
}

func xmlfmtC(ctx *cli.Context) error {
	// ctx.JSON(ctx.RootArgv())
	// ctx.JSON(ctx.Argv())
	// fmt.Println()

	argv := ctx.Argv().(*rootT)
	data, err := ioutil.ReadAll(argv.Filei)
	argv.Filei.Close()
	abortOn("Read input", err)
	//fmt.Println(string(data))
	fmt.Println(xmlfmt.FormatXML(string(data), argv.Prefix, argv.Indent))
	return nil
}

// abortOn will quit on anticipated errors gracefully without stack trace
func abortOn(errCase string, e error) {
	if e != nil {
		fmt.Printf("[%s] %s error: %v\n", progname, errCase, e)
		os.Exit(1)
	}
}
