package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/go-easygen/cli"
	"github.com/spf13/viper"

	"github.com/ver13/ava/tools/avaEnum/generator"

	"github.com/ver13/ava/pkg/common/commands"
)

func init() {
	commands.RootCmd.AddCommand(commands.VersionCmd)

	commands.RootCmd.PersistentFlags().Bool("noprefix", false, "Prevents the constants generated from having the Enum as a prefix.")
	commands.RootCmd.PersistentFlags().Bool("lower", false, "Adds lowercase variants of the enum strings for lookup.")
	commands.RootCmd.PersistentFlags().Bool("marshal", false, "Adds text (and inherently json) marshalling functions.")
	commands.RootCmd.PersistentFlags().Bool("sql", false, "Adds SQL database scan and value functions.")
	commands.RootCmd.PersistentFlags().Bool("flag", false, "Adds golang flag functions.")
	commands.RootCmd.PersistentFlags().Bool("names", false, "Generates a 'Names() []string' function, and adds the possible enum values in the error response during parsing.")
	commands.RootCmd.PersistentFlags().Bool("nocamel", false, "Removes the snake_case to CamelCase name changing.")

	viper.BindPFlag("file", commands.RootCmd.PersistentFlags().Lookup("file"))
}

type rootT struct {
	cli.Helper
	FileNames      []string `cli:"*f,file" usage:"The file(s) to generate enums.  Use more than one flag for more files."`
	NoPrefix       bool     `cli:"noprefix" usage:"Prevents the constants generated from having the Enum as a prefix."`
	Lowercase      bool     `cli:"lower" usage:"Adds lowercase variants of the enum strings for lookup."`
	Marshal        bool     `cli:"marshal" usage:"Adds text (and inherently json) marshalling functions."`
	SQL            bool     `cli:"sql" usage:"Adds SQL database scan and value functions."`
	Flag           bool     `cli:"flag" usage:"Adds golang flag functions."`
	Prefix         string   `cli:"prefix" usage:"Replaces the prefix with a user one."`
	Names          bool     `cli:"names" usage:"Generates a 'Names() []string' function, and adds the possible enum values in the error response during parsing"`
	LeaveSnakeCase bool     `cli:"nocamel" usage:"Removes the snake_case to CamelCase name changing"`
}

func main() {
	cli.Run(new(rootT), func(ctx *cli.Context) error {
		argv := ctx.Argv().(*rootT)

		for _, fileName := range argv.FileNames {

			g := generator.NewGenerator()

			if argv.NoPrefix {
				g.WithNoPrefix()
			}
			if argv.Lowercase {
				g.WithLowercaseVariant()
			}
			if argv.Marshal {
				g.WithMarshal()
			}
			if argv.SQL {
				g.WithSQLDriver()
			}
			if argv.Flag {
				g.WithFlag()
			}
			if argv.Names {
				g.WithNames()
			}
			if argv.LeaveSnakeCase {
				g.WithoutSnakeToCamel()
			}
			if argv.Prefix != "" {
				g.WithPrefix(argv.Prefix)
			}

			originalName := fileName

			ctx.String("ava-enum started. file: %s\n", ctx.Color().Cyan(originalName))
			fileName, _ = filepath.Abs(fileName)
			outFilePath := fmt.Sprintf("%s_enum.go", strings.TrimSuffix(fileName, filepath.Ext(fileName)))

			// Parse the file given in arguments
			raw, err := g.GenerateFromFile(fileName)
			if err != nil {
				return fmt.Errorf("failed generating enums\nInputFile=%s\nError=%s", ctx.Color().Cyan(fileName), ctx.Color().RedBg(err))
			}

			mode := int(0644)
			errWrite := ioutil.WriteFile(outFilePath, raw, os.FileMode(mode))
			if errWrite != nil {
				return fmt.Errorf("failed writing to file %s: %s", ctx.Color().Cyan(outFilePath), ctx.Color().Red(errWrite))
			}
			ctx.String("ava-enum finished. file: %s\n", ctx.Color().Cyan(originalName))
		}

		return nil
	})
}
