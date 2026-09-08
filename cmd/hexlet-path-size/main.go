package main

import (
	"code"
	"context"
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli/v3"
)

func main() {
	cmd := &cli.Command{
		Name:      "hexlet-path-size",
		Usage:     "print size of a file or directory",
		ArgsUsage: "<path>",
		Flags: []cli.Flag{
			&cli.BoolFlag{
				Name:    "human",
				Aliases: []string{"H"},
				Usage:   "human-readable sizes (auto-select unit)",
			},
			&cli.BoolFlag{
				Name:    "all",
				Aliases: []string{"a"},
				Usage:   "include hidden files and directories",
			},
		},
		Action: func(ctx context.Context, cmd *cli.Command) error {
			path := cmd.Args().First()
			human := cmd.Bool("human")
			all := cmd.Bool("all")
			size, err := code.GetPathSize(path, all)
			if err != nil {
				return err
			}
			s := code.FormatSize(size, human)
			fmt.Printf("%s\t%s\n", s, path)
			return nil
		},
	}

	if err := cmd.Run(context.Background(), os.Args); err != nil {
		log.Fatal(err)
	}
}
