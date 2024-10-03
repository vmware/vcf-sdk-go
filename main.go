// © Broadcom. All Rights Reserved.
// The term “Broadcom” refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: BSD-2-Clause

//go:generate go run github.com/oapi-codegen/oapi-codegen/v2/cmd/oapi-codegen -config cfg.yaml swagger.json
//go:generate go run main.go -add-header

package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
    addHeader := flag.Bool("add-header", false, "Add header to generated SDK.")
    flag.Parse()

    if *addHeader {
        const headerFile = "copyright.txt"
        const generatedFile = "./vcf/vcf.gen.go"

        header, err := os.ReadFile(headerFile)
        if err != nil {
            fmt.Println(fmt.Errorf("error reading header file: %w", err))
            os.Exit(1)
        }

        content, err := os.ReadFile(generatedFile)
        if err != nil {
            fmt.Println(fmt.Errorf("error reading generated file: %w", err))
            os.Exit(1)
        }

        newContent := string(header) + string(content)

        err = os.WriteFile(generatedFile, []byte(newContent), 0644)
        if err != nil {
            fmt.Println(fmt.Errorf("error writing to generated file: %w", err))
            os.Exit(1)
        }

        fmt.Println("Header added to", generatedFile)
    }
}
