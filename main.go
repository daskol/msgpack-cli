// Tool msgpack-cli allows to encode MsgPack objects to JSON and decode MsgPack
// object from JSON.
//
//
// Usage examples
//
//  $ msgpack-cli encode <<< "Hello, World!" > hello.msgpack
//  $ msgpack-cli decode < hello.msgpack
//  "Hello, World!"
package main

import (
	"encoding/json"
	"flag"
	"io"
	"log"
	"os"

	"github.com/vmihailenco/msgpack"
)

const usage = `Convert MsgPack objects to JSON objects and vice versa.
Usage:  msgpack-cli [-h | --help] [-v | --vesion] <command> [<args>]
        msgpack-cli {enc,encode} <input> <output>
        msgpack-cli {dec,decode} <input> <output>
        msgpack-cli help <input> <output>

Common options
    -h, --help           show usage
    -v, --version        show version

Commands
    enc,encode           encode MsgPack objects from JSON
    dec,decode           decode MsgPack objects to JSON
    help                 show this message

Visit https://github.com/daskol/msgpack-cli for feedback or source code.
`

var flagVersion bool

func msgpack2json(r io.Reader, w io.Writer) error {
	var dec = msgpack.NewDecoder(r)
	var enc = json.NewEncoder(w)
	var msg interface{}

	if err := dec.Decode(&msg); err != nil {
		return err
	}

	if err := enc.Encode(&msg); err != nil {
		return err
	}

	return nil
}

func json2msgpack(r io.Reader, w io.Writer) error {
	var dec = json.NewDecoder(r)
	var enc = msgpack.NewEncoder(w)
	var msg interface{}

	if err := dec.Decode(&msg); err != nil {
		return err
	}

	if err := enc.Encode(&msg); err != nil {
		return err
	}

	return nil
}

func init() {
	flag.BoolVar(&flagVersion, "v", false, "")
	flag.BoolVar(&flagVersion, "version", false, "")
	flag.Usage = func() { print(usage) }
	flag.Parse()
}

func main() {
	if flagVersion {
		println("msgpack-cli version 0.0.0")
		return
	}

	switch flag.Arg(0) {
	case "enc", "encode":
		if err := json2msgpack(os.Stdin, os.Stdout); err != nil {
			log.Fatalf("failed to convert msgpack to json: %s", err)
		}
	case "dec", "decode":
		if err := msgpack2json(os.Stdin, os.Stdout); err != nil {
			log.Fatalf("failed to convert msgpack to json: %s", err)
		}
	case "help":
		flag.Usage()
	default:
		flag.Usage()
	}
}
