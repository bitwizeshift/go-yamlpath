package main

import (
	"fmt"
	"io"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"gopkg.in/yaml.v3"
	"rodusek.dev/pkg/yamlpath"
)

type inputFlag struct {
	stream *io.ReadCloser
}

func (s *inputFlag) String() string {
	return ""
}

func (s *inputFlag) Set(path string) error {
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	*s.stream = file
	return nil
}

func (s *inputFlag) Type() string {
	return "path"
}

type outputFlag struct {
	stream io.WriteCloser
}

func (s *outputFlag) String() string {
	return ""
}

func (s *outputFlag) Set(path string) error {
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	s.stream = file
	return nil
}

func (s *outputFlag) Type() string {
	return "path"
}

type nopCloser struct {
	io.Writer
}

func (n *nopCloser) Close() error {
	return nil
}

var _ pflag.Value = (*outputFlag)(nil)

type Command struct {
	writer io.WriteCloser
	reader io.ReadCloser
}

func (c *Command) Reader(cmd *cobra.Command) io.ReadCloser {
	if c.reader == nil {
		return io.NopCloser(cmd.InOrStdin())
	}
	return c.reader
}

func (c *Command) Writer(cmd *cobra.Command) io.WriteCloser {
	if c.writer == nil {
		return &nopCloser{cmd.OutOrStdout()}
	}
	return c.writer
}

func (c *Command) RegisterFlags(fs *pflag.FlagSet) {
	fs.VarP(&inputFlag{stream: &c.reader}, "input", "i", "input path to YAML file")
	fs.VarP(&outputFlag{stream: c.writer}, "output", "o", "output path to YAML file")
}

func (c *Command) RegisterCommand(cmd *cobra.Command) {
	cmd.Args = func(cmd *cobra.Command, args []string) error {
		if len(args) == 1 {
			return nil
		}
		return fmt.Errorf("requires exactly one argument")
	}
	cmd.Run = c.Run

}

func (c *Command) Run(cmd *cobra.Command, args []string) {
	in := c.Reader(cmd)
	out := c.Writer(cmd)
	defer in.Close()
	defer out.Close()

	yp, err := yamlpath.Compile(args[0])
	if err != nil {
		_, _ = fmt.Fprintf(cmd.ErrOrStderr(), "error: %v\n", err)
		os.Exit(1)
		return
	}

	decoder := yaml.NewDecoder(in)
	for i := 0; ; i++ {
		var node yaml.Node
		if err := decoder.Decode(&node); err != nil {
			if err == io.EOF {
				break
			}
			_, _ = fmt.Fprintf(cmd.ErrOrStderr(), "error: %v\n", err)
			os.Exit(1)
			return
		}
		if i > 0 {
			_, _ = fmt.Fprintln(out, "---")
		}
		result, err := yp.Match(&node)
		if err != nil {
			_, _ = fmt.Fprintf(cmd.ErrOrStderr(), "error: %v\n", err)
			os.Exit(1)
			return
		}

		encoder := yaml.NewEncoder(out)
		for _, node := range result {
			if err := encoder.Encode(node); err != nil {
				_, _ = fmt.Fprintf(cmd.ErrOrStderr(), "error: %v\n", err)
				os.Exit(1)
				return
			}
		}
	}

}

func main() {
	cmd := &Command{}
	root := &cobra.Command{
		Use:   "yamlpath",
		Short: "yamlpath is a command-line tool for querying YAML documents",
	}
	cmd.RegisterFlags(root.Flags())
	cmd.RegisterCommand(root)

	if err := root.Execute(); err != nil {
		os.Exit(1)
	}
}
