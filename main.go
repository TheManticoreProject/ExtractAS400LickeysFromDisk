package main

import (
	"fmt"
	"os"
	"runtime"
	"sync"
	"time"

	"github.com/TheManticoreProject/ExtractAS400LickeysFromDisk/results"
	"github.com/TheManticoreProject/ExtractAS400LickeysFromDisk/searcher"
	"github.com/TheManticoreProject/goopts/parser"
)

var (
	// Configuration
	debug bool

	// Input files
	inputFiles []string
	maxWorkers int

	// Output settings
	outputCLFile   string
	outputJsonFile string

	chunkSize int
)

func parseArgs() {
	ap := parser.ArgumentsParser{
		Banner: "ExtractAS400LickeysFromDisk - by Remi GASCOU (Podalirius) @ TheManticoreProject - v1.0.0",
	}
	ap.SetOptShowBannerOnRun(true)
	ap.SetOptShowBannerOnHelp(true)

	// Configuration flags
	ap.NewBoolArgument(&debug, "", "--debug", false, "Debug mode.")

	// Input settings
	group_input, err := ap.NewArgumentGroup("Input")
	if err == nil {
		group_input.NewListOfStringsArgument(&inputFiles, "-i", "--input", []string{}, true, "Path(s) to the disk image file(s) to scan. Can be specified multiple times.")
	} else {
		fmt.Printf("[error] Error creating input argument group: %v\n", err)
		os.Exit(1)
	}

	// Configuration settings
	group_configuration, err := ap.NewArgumentGroup("Configuration")
	if err == nil {
		group_configuration.NewIntArgument(&chunkSize, "-c", "--chunk-size", 100*1024*1024, false, "Chunk size in bytes for reading the file (default: 100MB).")
		group_configuration.NewIntArgument(&maxWorkers, "-w", "--workers", runtime.NumCPU(), false, "Maximum number of worker threads (default: number of files).")
	} else {
		fmt.Printf("[error] Error creating configuration argument group: %v\n", err)
		os.Exit(1)
	}

	// Output settings
	group_output, err := ap.NewArgumentGroup("Output")
	if err == nil {
		group_output.NewStringArgument(&outputCLFile, "-c", "--output-cl", "./addlickey.cl", false, "Output file for control language ADDLICKEY commands (default: ./addlickey.cl).")
		group_output.NewStringArgument(&outputJsonFile, "-j", "--output-json", "./lickeys.json", false, "Output file for JSON output (default: ./lickeys.json).")
	} else {
		fmt.Printf("[error] Error creating output argument group: %v\n", err)
		os.Exit(1)
	}

	ap.Parse()
}

func main() {
	parseArgs()

	// Check if all input files exist
	keptFiles := []string{}
	for _, file := range inputFiles {
		if _, err := os.Stat(file); os.IsNotExist(err) {
			fmt.Printf("[warn] Skipping input file: %s (does not exist)\n", file)
			continue
		}
		keptFiles = append(keptFiles, file)
	}

	resultsTable := results.NewResultsTable(debug)
	resultsTable.SetPathToOutputCLFile(outputCLFile)
	resultsTable.SetPathToOutputJsonFile(outputJsonFile)
	resultsTable.StartProgressPrinter()

	var wg sync.WaitGroup
	for _, file := range keptFiles {
		wg.Add(1)
		go func(f string) {
			defer wg.Done()
			err := searcher.SearchFile(f, resultsTable, chunkSize)
			if err != nil {
				fmt.Printf("[error] Error processing file %s: %v\n", f, err)
			}
		}(file)
	}
	wg.Wait()

	resultsTable.Done()

	fmt.Printf("[%s] Done! Found %d licence keys\n", time.Now().Format("2006-01-02 15:04:05"), len(resultsTable.LicenceInfos))
}
