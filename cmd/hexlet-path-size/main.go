package main

import (
	"code"
	"context"
	"fmt"
	"os"

	"github.com/urfave/cli/v3"
)

func main() {
	cmd := &cli.Command{
		Name:      "hexlet-path-size",
		Usage:     "Print size of a file or directory",
		UsageText: "hexlet-path-size [global options] <path>",
		Action: func(ctx context.Context, c *cli.Command) error {
			if c.Args().Len() == 0 {
				return fmt.Errorf("path is required")
			}

			path := c.Args().First()
			size, err := code.GetPathSize(path)
			if err != nil {
				return fmt.Errorf("failed to get size: %w", err)
			}

			fmt.Printf("%d\t%s\n", size, path)
			return nil
		},
	}

	if err := cmd.Run(context.Background(), os.Args); err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}
}
