package main

import (
	"bufio"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"

	clog "github.com/gregoryalbouy/go-custom-log"
)

var exts = []string{
	"csv",
	"json",
}

// Output struct
type Output struct {
	filename string
}

// NewOutput func
func NewOutput(filename string) Output {
	return Output{filename}
}

func (out Output) Write(fr FinalResult) error {
	rgx := regexp.MustCompile(fmt.Sprintf(`\.(%s)$`, strings.Join(exts, "|")))
	ext := string(rgx.Find([]byte(out.filename)))

	switch ext {
	case ".csv":
		return out.csv(fr)
	case ".json":
		return out.json(fr)
	}

	return errors.New("wrong output format, only .csv and .json accepted")
}

func (out Output) json(fr FinalResult) error {
	b, err := json.MarshalIndent(fr, "", "  ")
	if err != nil {
		return err
	}

	file, err := os.Create(out.filename)
	if err != nil {
		return err
	}

	w := bufio.NewWriter(file)
	_, err = w.Write(b)
	if err != nil {
		return err
	}

	if err := w.Flush(); err != nil {
		return err
	}

	out.confirm("JSON")

	return nil
}

func (out Output) csv(fr FinalResult) error {
	var (
		head  []string
		lines [][]string
	)

	// Set head line
	head = append(head, "algorithms")
	for _, size := range fr.Sizes {
		head = append(head, strconv.Itoa(size))
	}

	// Set body lines
	for _, tr := range fr.TestResults {
		var line []string

		line = append(line, tr.Name)
		for _, sizeRes := range tr.Results {
			line = append(line, strconv.Itoa(int(sizeRes.Duration.Nanoseconds())))
		}

		lines = append(lines, line)
	}

	// Create CSV file & writer
	file, err := os.Create(out.filename)
	if err != nil {
		return err
	}
	defer file.Close()

	w := bufio.NewWriter(file)

	// Write head line
	_, err = w.WriteString(strings.Join(head, ",") + "\n")
	if err != nil {
		return err
	}

	// Write body lines
	for _, line := range lines {
		_, err := w.WriteString(strings.Join(line, ",") + "\n")
		if err != nil {
			return err
		}
	}

	if err := w.Flush(); err != nil {
		return err
	}

	out.confirm("CSV")

	return nil
}

func (out Output) confirm(ext string) {
	fmt.Printf("%s file %s generated\n", ext, clog.Blue(out.filename))
}
