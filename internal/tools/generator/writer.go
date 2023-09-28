package generator

import (
	"bytes"
	"fmt"
	"go/format"
	"os"
	"strings"
)

// Writer and formatter of golang source file
type outputFile struct {
	*bytes.Buffer
	name string
}

func NewOutputFile(name string, pkgName string, imports []string) *outputFile {
	o := outputFile{
		Buffer: new(bytes.Buffer),
		name:   name,
	}

	fmt.Fprintln(o, "// Code generated by generate.sh, DO NOT EDIT.")
	fmt.Fprintln(o, "")
	fmt.Fprintf(o, "package %s\n", pkgName)
	fmt.Fprintln(o, "")
	fmt.Fprintf(o, "import (\n\n%s\n)\n", strings.Join(imports, "\n"))
	fmt.Fprintln(o, "")

	return &o
}

func (o *outputFile) Close() (err error) {
	// Output to file
	out, err := format.Source(o.Bytes())
	if err != nil {
		return fmt.Errorf("format error: %w\n%s\n", err, o.String())
	}
	fWrite, err := os.Create(o.name)
	if err != nil {
		return err
	}
	defer func() {
		errClose := fWrite.Close()
		if errClose != nil && err == nil {
			err = errClose
		}
	}()
	_, err = fWrite.Write(out)
	if err != nil {
		return err
	}
	return nil
}
