// wrengo v0.0.1-dev
//
// (c) Harry Lawrence 2015
//
// @package wrengo
// @version 0.0.1-dev
//
// @author Harry Lawrence <http://github.com/hazbo>
//
// License: MIT
//
// For the full copyright and license information, please view the LICENSE
// file that was distributed with this source code.

package main

// #cgo CFLAGS: -std=c99 -Wall -Werror -I./wren/include
// #cgo LDFLAGS: -L. wren/libwren.a
// #include <wren.h>
import "C"

import (
    "os"
    "fmt"
    "./classes"
    "github.com/hazbo/cli"
)

// WrengoVM is An instance of the VM
var WrengoVM VM

func main() {
    app := cli.NewApp()

    // constants located in meta.go
    app.Name    = appName
    app.Version = version
    app.Usage   = usage
    app.Author  = author
    app.Email   = email

    app.Commands = []cli.Command {
        {
            Name: "run",
            Usage: "your wren file!",
            Action: func (c *cli.Context) {
                s, err := NewScript(os.Args[2]);
                if err != nil {
                    fmt.Println(err); os.Exit(1)
                }
                WrengoVM = NewVM()
                WrengoVM.Script = s
                WrengoVM.Script.readApi([]string{
                    "src/api/file.wren",
                    "src/api/markdown.wren",
                })
                WrengoVM.Script.initApiMain()
                WrengoVM.Interpret()
            },
        },
    }

    app.Run(os.Args)
}

//export class_markdown_parse
func class_markdown_parse(vm *C.WrenVM) {
    m := class.NewMarkdown()
    a := C.wrenGetArgumentString(vm, 1)
    r := m.Parse(C.GoString(a))
    C.wrenReturnString(vm, C.CString(r), -1)
}

//export class_file_read
func class_file_read(vm *C.WrenVM) {
    f := class.NewFile()
    a := C.wrenGetArgumentString(vm, 1)
    r := f.Read(C.GoString(a))
    C.wrenReturnString(vm, C.CString(r), -1)
}

//export class_file_write
func class_file_write(vm *C.WrenVM) {
    f := class.NewFile()
    filename := C.wrenGetArgumentString(vm, 1)
    data     := C.wrenGetArgumentString(vm, 2)
    perm     := C.wrenGetArgumentDouble(vm, 3)
    f.Write(C.GoString(filename), C.GoString(data), int(perm))
}
