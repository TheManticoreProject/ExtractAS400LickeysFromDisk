package results

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"time"

	"github.com/TheManticoreProject/ExtractAS400LickeysFromDisk/licenceinfo"
	"github.com/TheManticoreProject/ExtractAS400LickeysFromDisk/products"
	"github.com/TheManticoreProject/ExtractAS400LickeysFromDisk/progress"
)

type ResultsTable struct {
	LicenceInfos []licenceinfo.LicenceInfoInterface

	PathToOutputCLFile   string
	PathToOutputJsonFile string

	// Internal
	debug       bool
	startTime   time.Time
	elapsedTime time.Duration
	lock        sync.Mutex
	done        bool
	progressBar *progress.Progress
}

func NewResultsTable(debug bool) *ResultsTable {
	return &ResultsTable{
		debug:       debug,
		startTime:   time.Now(),
		lock:        sync.Mutex{},
		progressBar: progress.NewProgress(),
	}
}

// SetPathToOutputFile sets the path to the output file
func (r *ResultsTable) SetPathToOutputCLFile(path string) {
	r.lock.Lock()
	defer r.lock.Unlock()

	outputDir := filepath.Dir(path)
	if err := os.MkdirAll(outputDir, 0755); err != nil {
		fmt.Printf("Error creating output directory %s: %v\n", outputDir, err)
	}

	r.PathToOutputCLFile = path
}

// SetPathToOutputJsonFile sets the path to the output JSON file
func (r *ResultsTable) SetPathToOutputJsonFile(path string) {
	r.lock.Lock()
	defer r.lock.Unlock()

	outputDir := filepath.Dir(path)
	if err := os.MkdirAll(outputDir, 0755); err != nil {
		fmt.Printf("Error creating output directory %s: %v\n", outputDir, err)
	}

	r.PathToOutputJsonFile = path
}

// ReportResult reports a found LICKEY and adds it to the results table
func (r *ResultsTable) ReportResult(offset int64, filename string, lic licenceinfo.LicenceInfoInterface) {
	r.lock.Lock()

	r.LicenceInfos = append(r.LicenceInfos, lic)

	color_green := "\x1b[38;2;0;204;102m"
	color_reset := "\x1b[0m"

	licenceInfoType := "LICENCEINFOx"
	if _, ok := lic.(*licenceinfo.LICENCEINFO2); ok {
		licenceInfoType = "LICENCEINFO2"
	} else if _, ok := lic.(*licenceinfo.LICENCEINFO1); ok {
		licenceInfoType = "LICENCEINFO1"
	} else {
		licenceInfoType = "LICENCEINFOx"
	}

	if r.debug {
		fmt.Printf("\x1b[2K\r[%s] Found LICKEY for '%s' at offset: 0x%012x (%d) in %s\n",
			time.Now().Format("2006-01-02 15:04:05"),
			products.GetNameFromPrdidAndFeature(lic.GetPrdidString(), lic.GetFeatureString()),
			offset,
			offset,
			filename,
		)

		fmt.Printf("[%s] ├── PRDID=[%s] LICTRM=[%s] FEATURE=[%s] LICKEY=[%s] SERIAL=[%s] PRCGRP=[%s] (struct type: %s)\n",
			time.Now().Format("2006-01-02 15:04:05"),
			color_green+lic.GetPrdidString()+color_reset,
			color_green+lic.GetLictrmString()+color_reset,
			color_green+lic.GetFeatureString()+color_reset,
			color_green+lic.GetLickeyString()+color_reset,
			color_green+lic.GetSerialString()+color_reset,
			color_green+lic.GetPrcgrpString()+color_reset,
			licenceInfoType,
		)
		fmt.Printf("[%s] └── raw structure: %s\n", time.Now().Format("2006-01-02 15:04:05"), strings.ToUpper(hex.EncodeToString(lic.GetRawStructure()[:])))
	} else {
		fmt.Printf("\x1b[2K\r[%s] Found LICKEY for '%s':\n", time.Now().Format("2006-01-02 15:04:05"),
			products.GetNameFromPrdidAndFeature(lic.GetPrdidString(), lic.GetFeatureString()),
		)

		fmt.Printf("[%s] └── PRDID=[%s] LICTRM=[%s] FEATURE=[%s] LICKEY=[%s] SERIAL=[%s] PRCGRP=[%s]\n",
			time.Now().Format("2006-01-02 15:04:05"),
			color_green+lic.GetPrdidString()+color_reset,
			color_green+lic.GetLictrmString()+color_reset,
			color_green+lic.GetFeatureString()+color_reset,
			color_green+lic.GetLickeyString()+color_reset,
			color_green+lic.GetSerialString()+color_reset,
			color_green+lic.GetPrcgrpString()+color_reset,
		)
	}
	r.lock.Unlock()

	r.WriteResultsToFile()
}

// WriteResultsToFile writes the results to the output file
func (r *ResultsTable) WriteResultsToFile() {
	r.lock.Lock()
	defer r.lock.Unlock()

	if r.PathToOutputCLFile != "" {
		outputFile, err := os.OpenFile(r.PathToOutputCLFile, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0644)
		if err != nil {
			fmt.Printf("Error opening output file %s: %v\n", r.PathToOutputCLFile, err)
			return
		}
		defer outputFile.Close()

		for _, lic := range r.LicenceInfos {
			outputFile.WriteString(lic.ToCLAddLickey() + "\n")
		}

		outputFile.Sync()
	}

	if r.PathToOutputJsonFile != "" {
		outputFile, err := os.OpenFile(r.PathToOutputJsonFile, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0644)
		if err != nil {
			fmt.Printf("Error opening output file %s: %v\n", r.PathToOutputJsonFile, err)
			return
		}
		defer outputFile.Close()

		data := make(map[string]map[string]map[string]map[string][]map[string]string)
		for _, lic := range r.LicenceInfos {
			serialValue := lic.GetSerialString()
			lictrmValue := lic.GetLictrmString()
			prdidValue := lic.GetPrdidString()
			featureValue := lic.GetFeatureString()
			lickeyValue := lic.GetLickeyString()

			// Ensure the serial map exists
			if _, ok := data[serialValue]; !ok {
				data[serialValue] = make(map[string]map[string]map[string][]map[string]string)
			}
			// Ensure the lictrm map exists within the serial map
			if _, ok := data[serialValue][prdidValue]; !ok {
				data[serialValue][prdidValue] = make(map[string]map[string][]map[string]string)
			}
			// Ensure the prdid map exists within the lictrm map
			if _, ok := data[serialValue][prdidValue][lictrmValue]; !ok {
				data[serialValue][prdidValue][lictrmValue] = make(map[string][]map[string]string)
			}
			// Assign the lickey to the feature key
			if _, ok := data[serialValue][prdidValue][lictrmValue][featureValue]; !ok {
				data[serialValue][prdidValue][lictrmValue][featureValue] = []map[string]string{}
			}

			lickeyData := map[string]string{
				"name":           products.GetNameFromPrdidAndFeature(prdidValue, featureValue),
				"prdid":          prdidValue,
				"lictrm":         lictrmValue,
				"feature":        featureValue,
				"lickey":         lickeyValue,
				"serial":         serialValue,
				"prcgrp":         lic.GetPrcgrpString(),
				"structure_data": strings.ToUpper(hex.EncodeToString(lic.GetRawStructure()[:])),
			}
			data[serialValue][prdidValue][lictrmValue][featureValue] = append(data[serialValue][prdidValue][lictrmValue][featureValue], lickeyData)
		}

		jsonData, err := json.MarshalIndent(data, "", "    ")
		if err != nil {
			fmt.Printf("Error marshalling JSON data: %v\n", err)
			return
		}

		outputFile.Write(jsonData)

		outputFile.Sync()
	}
}

func (r *ResultsTable) Lock() {
	r.lock.Lock()
}

func (r *ResultsTable) Unlock() {
	r.lock.Unlock()
}

// Done writes the results to the output file and sets the done flag
func (r *ResultsTable) Done() {
	r.lock.Lock()
	defer r.lock.Unlock()

	r.done = true
	r.elapsedTime = time.Since(r.startTime)

	progressPercentage := float64(r.progressBar.ProcessedBytes) / float64(r.progressBar.TotalBytes) * 100

	var eta string
	if r.progressBar.TotalBytes == 0 || r.progressBar.ProcessedBytes == 0 {
		eta = "N/A"
	} else {
		// Calculate estimated total duration based on current progress
		// Formula: estimatedTotalDuration = elapsed / (processedBytes / totalBytes)
		// Which simplifies to: estimatedTotalDuration = elapsed * (totalBytes / processedBytes)

		// Use float64 for calculations to maintain precision, then convert back to time.Duration.
		// elapsed is a time.Duration, which is an int64 representing nanoseconds.
		estimatedTotalDurationFloat := float64(r.elapsedTime) * (float64(r.progressBar.TotalBytes) / float64(r.progressBar.ProcessedBytes))
		estimatedTotalDuration := time.Duration(estimatedTotalDurationFloat)

		// Calculate remaining duration (ETA)
		etaDuration := estimatedTotalDuration - r.elapsedTime

		// Ensure ETA is not negative (can happen due to floating point inaccuracies or if progress jumps)
		if etaDuration < 0 {
			etaDuration = 0
		}
		eta = etaDuration.Round(time.Second).String()
	}

	barString := r.progressBar.GetProgressBarString(50)
	fmt.Printf(
		"\x1b[2K\r[%s] Progress: %s %5.2f%% | elapsed: %s | eta: %s\n",
		time.Now().Format("2006-01-02 15:04:05"),
		barString,
		progressPercentage,
		time.Since(r.startTime).Round(time.Second),
		eta,
	)
}

func (r *ResultsTable) StartProgressPrinter() {
	go func() {
		ticker := time.NewTicker(500 * time.Millisecond) // Update progress every 5 seconds

		for range ticker.C {
			r.lock.Lock()

			elapsed := time.Since(r.startTime)

			progressPercentage := float64(r.progressBar.ProcessedBytes) / float64(r.progressBar.TotalBytes) * 100

			var eta string
			if r.progressBar.TotalBytes == 0 || r.progressBar.ProcessedBytes == 0 {
				eta = "N/A"
			} else {
				// Calculate estimated total duration based on current progress
				// Formula: estimatedTotalDuration = elapsed / (processedBytes / totalBytes)
				// Which simplifies to: estimatedTotalDuration = elapsed * (totalBytes / processedBytes)

				// Use float64 for calculations to maintain precision, then convert back to time.Duration.
				// elapsed is a time.Duration, which is an int64 representing nanoseconds.
				estimatedTotalDurationFloat := float64(elapsed) * (float64(r.progressBar.TotalBytes) / float64(r.progressBar.ProcessedBytes))
				estimatedTotalDuration := time.Duration(estimatedTotalDurationFloat)

				// Calculate remaining duration (ETA)
				etaDuration := estimatedTotalDuration - elapsed

				// Ensure ETA is not negative (can happen due to floating point inaccuracies or if progress jumps)
				if etaDuration < 0 {
					etaDuration = 0
				}
				eta = etaDuration.Round(time.Second).String()
			}

			barString := r.progressBar.GetProgressBarString(50)
			fmt.Printf(
				"\x1b[2K\r[%s] Progress: %s %5.2f%% | elapsed: %s | eta: %s",
				time.Now().Format("2006-01-02 15:04:05"),
				barString,
				progressPercentage,
				elapsed.Round(time.Second),
				eta,
			)

			r.lock.Unlock()
		}
	}()
}

// AddTotalBytes adds the total bytes to the progress bar
func (r *ResultsTable) AddTotalBytes(sizeInBytes int64) {
	r.lock.Lock()
	defer r.lock.Unlock()

	r.progressBar.AddTotalBytes(sizeInBytes)
}

// AddProcessedBytes adds the processed bytes to the progress bar
func (r *ResultsTable) AddProcessedBytes(sizeInBytes int64) {
	r.lock.Lock()
	defer r.lock.Unlock()

	r.progressBar.AddProcessedBytes(sizeInBytes)
}
