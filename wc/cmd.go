package main

import (
	"errors"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

var (
	rootCmd = &cobra.Command{
		Use:   "wc [flags] file",
		Short: "print newline, word, and byte counts for each file",
		RunE:  Run,
	}

	config struct {
		Bytes bool
		Lines bool
		Words bool
	}
)

func Run(cmd *cobra.Command, args []string) error {

	if len(args) < 1 {
		return errors.New("not enough arguments")
	}

	file, err := os.ReadFile(args[0])
	if err != nil {
		return err
	}

	if config.Bytes {
		fmt.Printf("%d %s", len(file), args[0])
		return nil
	}

	if config.Words {
		fields := strings.Fields(string(file))
		fmt.Printf("%d %s", len(fields), args[0])
		return nil
	}

	var count int = 0
	for _, b := range file {
		if b == '\n' {
			count += 1
		}
	}
	fmt.Printf("%d %s", count, args[0])

	return nil
}

func init() {
	rootCmd.Flags().BoolVarP(&config.Bytes, "bytes", "c", false, "print the byte counts")
	rootCmd.Flags().BoolVarP(&config.Lines, "lines", "l", false, "print the newline counts")
	rootCmd.Flags().BoolVarP(&config.Words, "words", "w", false, "print the word counts")
	rootCmd.MarkFlagsMutuallyExclusive("bytes", "lines", "words")
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
}
