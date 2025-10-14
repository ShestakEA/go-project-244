package main

import (
	"code"
	"context"
	"fmt"
	"os"

	"github.com/urfave/cli/v3"
)

func main() {
	app := &cli.Command{
		Name:  "gendiff",
		Usage: "Compares two configuration files and shows a difference.",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "format",
				Value:   "stylish",
				Aliases: []string{"f"},
				Usage:   "output format",
			},
		},
		Action: func(_ context.Context, cmd *cli.Command) error {
			if cmd.Args().Len() < 2 {
				fmt.Println("Файлы не переданны")
				return nil
			}
			file1 := cmd.Args().Get(0)
			file2 := cmd.Args().Get(1)
			format := cmd.String("format")

			result, err := code.GenDiff(file1, file2, format)
			if err != nil {
				return fmt.Errorf("невозможно сравнить файлы: %v", err)
			}
			fmt.Println(result)
			return nil
		},
	}

	err := app.Run(context.Background(), os.Args)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}
