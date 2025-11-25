package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/rcrowley/mergician/html"
	"golang.org/x/net/html/atom"
)

func Main(args []string, stdin io.Reader, stdout io.Writer) {
	flags := flag.NewFlagSet(args[0], flag.ExitOnError)
	flags.Usage = func() {
		fmt.Fprint(os.Stderr, `Usage: critaique <input>
  <input> input HTML document

Synopsis: critaique prompts an LLM to make suggestions for improving your writing, ask follow-up questions, etc.
`)
	}
	flags.Parse(args[1:])

	in := must2(html.ParseFile(flags.Arg(0)))

	claude := &Claude{
		APIKey: os.Getenv("ANTHROPIC_API_KEY"),
		Model:  "claude-3-7-sonnet-latest",
	}

	fmt.Println(must2(claude.Prompt(fmt.Sprintf(
		"Read and consider the rest of this prompt (after this paragraph) as if you were an editor preparing an article for publication. Provide feedback, make suggestions, and ask follow-up questions with the goal of helping the author improve their writing.\n\n%s",
		html.Text(in).String(),
	))))
}

func init() {
	log.SetFlags(0)
}

func main() {
	Main(os.Args, os.Stdin, os.Stdout)
}

func must(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func must2[T any](v T, err error) T {
	must(err)
	return v
}

func wrap(in *html.Node) *html.Node {
	out := must2(html.ParseString(`<!DOCTYPE html>
<html lang="en">
<head>
<meta charset="utf-8">
</head>
<body>
</body>
</html>
`))
	html.Find(out, html.IsAtom(atom.Body)).AppendChild(html.CopyNode(in))
	return out
}
