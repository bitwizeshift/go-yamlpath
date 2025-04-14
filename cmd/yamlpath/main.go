package main

import (
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"gopkg.in/yaml.v3"
	"rodusek.dev/pkg/yamlpath"
	"rodusek.dev/pkg/yamlpath/internal/yamlconv"
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

type constantFlags struct {
	opts *[]yamlpath.Option
}

func (c *constantFlags) String() string {
	return ""
}
func (c *constantFlags) Set(flag string) error {
	kv := strings.SplitN(flag, "=", 2)
	if len(kv) != 2 {
		return fmt.Errorf("invalid constant format, expected key=value")
	}
	key := kv[0]
	value := kv[1]

	var kind string
	kind, key = c.parseKey(key)
	switch kind {
	case "bool":
		b, err := strconv.ParseBool(value)
		if err != nil {
			return err
		}
		*c.opts = append(*c.opts, yamlpath.WithConstant(key, yamlconv.Bool(b)))
	case "int":
		i, err := strconv.ParseInt(value, 10, 64)
		if err != nil {
			return err
		}
		*c.opts = append(*c.opts, yamlpath.WithConstant(key, yamlconv.Int(i)))
	case "float":
		f, err := strconv.ParseFloat(value, 64)
		if err != nil {
			return err
		}
		*c.opts = append(*c.opts, yamlpath.WithConstant(key, yamlconv.Float(f)))
	case "string":
		*c.opts = append(*c.opts, yamlpath.WithConstant(key, yamlconv.String(value)))
	case "auto":
		var node yaml.Node
		if err := yaml.Unmarshal([]byte(value), &node); err != nil {
			return err
		}
		*c.opts = append(*c.opts, yamlpath.WithConstant(key, &node))
	}

	return nil
}

func (c *constantFlags) parseKey(input string) (kind string, key string) {
	parts := strings.SplitN(input, ":", 2)
	if len(parts) == 1 {
		return "auto", parts[0]
	}
	switch parts[0] {
	case "bool", "int", "float", "string":
		return parts[0], parts[1]
	}
	return "auto", key
}

func (c *constantFlags) Type() string {
	return "constant"
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
	writer  io.WriteCloser
	reader  io.ReadCloser
	options []yamlpath.Option
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
	fs.VarP(&constantFlags{opts: &c.options}, "constant", "c", "constant value to use in the YAMLPath expression")
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

	yp, err := yamlpath.Compile(args[0], c.options...)
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
