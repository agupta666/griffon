package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	shellwords "github.com/mattn/go-shellwords"
	readline "gopkg.in/readline.v1"
)

func sanitize(name string) string {
	if strings.HasSuffix(name, ".") {
		return name
	}
	return name + "."
}

type CmdHandler func(args []string)

func addCmd(args []string) {

	if len(args) != 3 {
		fmt.Println("ERROR:", "Illegal arguments")
		return
	}

	name := args[0]
	ip := args[1]
	port, err := strconv.Atoi(args[2])

	if err != nil {
		fmt.Println("ERROR:", err)
		return
	}

	name = sanitize(name)

	entry := &Entry{Name: name, IP: ip, Port: port}
	err = saveEntry(entry)

	if err != nil {
		fmt.Println("ERROR:", err)
	}
}

func deleteCmd(args []string) {
	if len(args) != 1 {
		fmt.Println("ERROR:", "Illegal arguments")
		return
	}

	name := sanitize(args[0])

	err := deleteEntry(name)

	if err != nil {
		fmt.Println("ERROR:", err)
	}
}

func listCmd(args []string) {
	for _, e := range allEntries() {
		fmt.Println(e)
	}
}

func exitCmd(args []string) {
	os.Exit(0)
}

var commands = map[string]CmdHandler{
	"add":    addCmd,
	"delete": deleteCmd,
	"list":   listCmd,
	"exit":   exitCmd,
}

func process(line string) {
	xs, err := shellwords.Parse(line)

	if err != nil {
		fmt.Println("ERROR:", "illegal command")
	}

	if len(xs) == 0 {
		return
	}

	cmd := xs[0]
	hx, ok := commands[cmd]

	if !ok {
		fmt.Println("ERROR:", "illegal command")
		return
	}

	hx(xs[1:])
}

func entries() func(string) []string {
	return func(line string) []string {
		exs := allEntries()
		var result []string
		for _, e := range exs {
			result = append(result, e.Name)
		}

		return result
	}
}

var completer = readline.NewPrefixCompleter(
	readline.PcItem("add"),
	readline.PcItem("list"),
	readline.PcItem("delete", readline.PcItemDynamic(entries())),
	readline.PcItem("exit"),
)

func startShell() {
	rl, err := readline.NewEx(&readline.Config{
		Prompt:       "griffon> ",
		HistoryFile:  ".griffon.hist",
		AutoComplete: completer,
	})

	if err != nil {
		fmt.Println("ERROR:", err)
	}
	defer rl.Close()

	for {
		line, err := rl.Readline()
		if err != nil {
			break
		}
		process(line)
	}
}
