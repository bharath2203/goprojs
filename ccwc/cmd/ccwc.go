package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
	"unicode"
)

var ccwcCommand = &cobra.Command{
	Use:   "ccwc",
	Short: "word count",
	Long:  "word count implementation to provide various functionalities with sequence of text",
	Args:  cobra.MinimumNArgs(1),
	RunE:  handle,
}

var options struct {
	runInAsync bool
}

func handle(cmd *cobra.Command, args []string) error {
	file, err := os.ReadFile(args[0])
	if err != nil {
		return err
	}
	if options.runInAsync {
		runAsync(file)
		return nil
	}
	sizeInBytes := computeSizeInBytes(file)
	wordCount := computeWordCount(file)
	lineCount := computeLineCount(file)
	fmt.Println(sizeInBytes, wordCount, lineCount)
	return nil
}

func runAndReturn(fun func(s []byte) int64, stream []byte) chan int64 {
	ch := make(chan int64)
	go func() {
		ch <- fun(stream)
	}()
	return ch
}

func runAsync(stream []byte) {
	sizeInBytesCh := runAndReturn(computeSizeInBytes, stream)
	wordCountCh := runAndReturn(computeWordCount, stream)
	lineCountCh := runAndReturn(computeLineCount, stream)
	fmt.Println(<-sizeInBytesCh, <-wordCountCh, <-lineCountCh)
}

func computeSizeInBytes(stream []byte) int64 {
	return int64(len(stream))
}

func computeWordCount(stream []byte) int64 {
	count := 0
	streamInStr := string(stream)
	for i := 0; i < len(streamInStr); i++ {
		if unicode.IsSpace(rune(streamInStr[i])) {
			count++
		}
	}
	return int64(count)
}

func computeLineCount(stream []byte) int64 {
	count := 1
	streamInStr := string(stream)
	for i := 0; i < len(streamInStr); i++ {
		if streamInStr[i] == '\n' {
			count++
		}
	}
	return int64(count)
}

func Execute() {
	ccwcCommand.Flags().BoolVar(&options.runInAsync, "parallel", false, "Run in parallel")
	if err := ccwcCommand.Execute(); err != nil {
		fmt.Println("error while executing ccwcCommand")
		os.Exit(-1)
	}
}
