package searcher

import (
	"bufio"
	"bytes"
	"compress/bzip2"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/TheManticoreProject/ExtractAS400LickeysFromDisk/licenceinfo"
	"github.com/TheManticoreProject/ExtractAS400LickeysFromDisk/results"
	"github.com/TheManticoreProject/Manticore/logger"
)

// SearchChunk searches a chunk of bytes for licenceinfo structures
func SearchChunk(baseOffset int64, chunk []byte, filename string, results *results.ResultsTable) {
	EBCDIC_V := byte('\xE5')
	EBCDIC_R := byte('\xD9')
	EBCDIC_M := byte('\xD4')
	EBCDIC_0 := byte('\xF0')
	EBCDIC_9 := byte('\xF9')

	currentIndex := 0
	for {
		if currentIndex >= len(chunk) {
			break
		}

		// Search for the next 'V' in the chunk
		// We look for the product id + option + version + release + modification + feature
		// For example: 5770SS1V7R5M05050 (product id: 5770, Option: SS1, Version: 7, Release: 5, Modification: 0, Feature: 5050)
		nextIndex := bytes.IndexByte(chunk[currentIndex:], EBCDIC_V)
		if nextIndex == -1 {
			break
		}

		// Check that the release and modification are valid (VxRxMx)
		if chunk[currentIndex+nextIndex] != EBCDIC_V || chunk[currentIndex+nextIndex+2] != EBCDIC_R || chunk[currentIndex+nextIndex+4] != EBCDIC_M {
			currentIndex += nextIndex + 1
			continue
		}
		// Check that the release and modification are valid (VxRxMx) with digits 0-9
		if chunk[currentIndex+nextIndex+1] < EBCDIC_0 || chunk[currentIndex+nextIndex+1] > EBCDIC_9 ||
			chunk[currentIndex+nextIndex+3] < EBCDIC_0 || chunk[currentIndex+nextIndex+3] > EBCDIC_9 ||
			chunk[currentIndex+nextIndex+5] < EBCDIC_0 || chunk[currentIndex+nextIndex+5] > EBCDIC_9 {
			currentIndex += nextIndex + 1
			continue
		}

		// Try to parse as LICENCEINFO2 first (1-byte header), then fall back to LICENCEINFO1 (20-byte header)

		// Candidate for LICENCEINFO2
		startLic2 := currentIndex + nextIndex - 7 - 1
		if startLic2 >= 0 && (startLic2+licenceinfo.LICENCEINFO2_STRUCT_SIZE) <= len(chunk) {
			endLic2 := startLic2 + licenceinfo.LICENCEINFO2_STRUCT_SIZE
			lic2 := licenceinfo.LICENCEINFO2{}
			if _, err := lic2.Unmarshal(chunk[startLic2:endLic2]); err == nil {
				results.ReportResult(baseOffset+int64(currentIndex+nextIndex), filename, &lic2)
				currentIndex += nextIndex + licenceinfo.LICENCEINFO2_STRUCT_SIZE
				continue
			}
		}

		// Candidate for LICENCEINFO1
		startLic1 := currentIndex + nextIndex - 7 - 20
		if startLic1 >= 0 && (startLic1+licenceinfo.LICENCEINFO1_STRUCT_SIZE) <= len(chunk) {
			endLic1 := startLic1 + licenceinfo.LICENCEINFO1_STRUCT_SIZE
			lic1 := licenceinfo.LICENCEINFO1{}
			if _, err := lic1.Unmarshal(chunk[startLic1:endLic1]); err == nil {
				results.ReportResult(baseOffset+int64(currentIndex+nextIndex), filename, &lic1)
				currentIndex += nextIndex + licenceinfo.LICENCEINFO1_STRUCT_SIZE
				continue
			}
		}

		// No parse success, move to next position
		currentIndex += nextIndex + 1
		continue
	}
}

// SearchFile searches a file for licenceinfo structures
func SearchFile(pathToFile string, results *results.ResultsTable, chunkSize int) error {
	// Open the file
	f, err := os.Open(pathToFile)
	if err != nil {
		return fmt.Errorf("failed to open file: %v", err)
	}
	defer f.Close()

	// Get file size for progress reporting
	fileInfo, err := f.Stat()
	if err != nil {
		return fmt.Errorf("failed to get file info: %v", err)
	}
	fileSize := fileInfo.Size()

	totalChunks := fileSize / int64(chunkSize)
	if fileSize%int64(chunkSize) != 0 {
		totalChunks++
	}

	// Create a reader for the file and decompress it if necessary
	var reader *bufio.Reader
	if strings.HasSuffix(strings.ToLower(pathToFile), ".bz2") {
		// To get the total decompressed size for a .bz2 file, we must decompress it.
		// This involves reading the entire stream once to determine its length,
		// and then resetting the file pointer and creating a new reader for actual processing.
		// This can be inefficient for very large files, as it requires two passes over the data.
		var decompressedSize int64
		// Create a temporary bz2Reader to calculate the decompressed size by reading to io.Discard
		tempBz2Reader := bzip2.NewReader(f)

		results.Lock()
		logger.Info(fmt.Sprintf("Calculating decompressed size for %s", pathToFile))
		results.Unlock()
		n, err := io.Copy(io.Discard, tempBz2Reader) // Read and discard to get decompressed length
		if err != nil {
			return fmt.Errorf("failed to calculate decompressed size for %s: %v", pathToFile, err)
		}
		decompressedSize = n
		results.Lock()
		logger.Info(fmt.Sprintf("Decompressed size for %s: %d", pathToFile, decompressedSize))
		results.Unlock()

		// Reset the file pointer to the beginning for the actual processing
		_, err = f.Seek(0, io.SeekStart)
		if err != nil {
			return fmt.Errorf("failed to seek file %s to start: %v", pathToFile, err)
		}
		bz2Reader := bzip2.NewReader(f)
		reader = bufio.NewReader(bz2Reader)

		results.AddTotalBytes(decompressedSize)
	} else {
		reader = bufio.NewReader(f)
		results.AddTotalBytes(fileSize)
	}

	// Create a buffer for the chunk
	buf := make([]byte, int64(chunkSize)+licenceinfo.LICENCEINFO1_STRUCT_SIZE) // overlap for partial matches

	chunkId := 1
	offset := int64(0)
	for {
		// Read the file into the buffer
		n, err := reader.Read(buf)
		if n == 0 {
			if err == io.EOF {
				break
			}
			return fmt.Errorf("error reading file %s at offset %d: %v", pathToFile, offset, err)
		}

		// Search the buffer for licenceinfo structures
		SearchChunk(offset, buf[:n], pathToFile, results)

		offset += int64(n)
		results.AddProcessedBytes(int64(n))
		chunkId++
	}

	return nil
}
