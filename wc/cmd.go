package main

import (
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/spf13/cobra"
)

var (
	rootCmd = &cobra.Command{
		Use:   "wc",
		Short: "print newline, word, and byte counts for each file",
		RunE:  Run,
	}

	config struct {
		Bytes bool
		File  string
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

	return nil
}

func init() {
	rootCmd.Flags().BoolVarP(&config.Bytes, "bytes", "c", false, "print the byte counts")
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
}
