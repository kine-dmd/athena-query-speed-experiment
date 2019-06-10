package main

import (
	"github.com/kine-dmd/athena-query-speed-test/appleWatch3Row"
	"github.com/kine-dmd/athena-query-speed-test/parquetHandler"
	"github.com/kine-dmd/athena-query-speed-test/s3Connection"
	"log"
	"math/rand"
	"os"
	"runtime"
	"strconv"
	"sync"
	"time"
)

const (
	totalRows       = 100000000
	bucketName      = "athena-query-size-test"
	localFilePrefix = "/data/rp1615/"
)

func main() {
	// Constants
	numRowsPerFile, _ := strconv.Atoi(os.Args[1])
	numFiles := totalRows / numRowsPerFile

	// Make a file uploader
	s3Conn := s3Connection.MakeS3Connection()

	// Sync variables to limit number of threads
	wg := sync.WaitGroup{}
	routinePool := make(chan struct{}, runtime.NumCPU()*3)

	// Make required number of files and upload them on multiple threads
	for i := 0; i < numFiles; i++ {
		// Send an empty struct to the routine pool (blocks if routine pool is full)
		routinePool <- struct{}{}
		wg.Add(1)

		go func() {
			// Make the parquet file and open it for upload
			filePath := makeFileWithRandomRows(numRowsPerFile)
			file, _ := os.Open(filePath)
			s3FilePath := strconv.Itoa(numRowsPerFile) + "/" + filePath[len(localFilePrefix):]

			// Upload the file and delete it to save space
			err := s3Conn.UploadFile(bucketName, s3FilePath, file)
			if err != nil {
				log.Print("Error uploading file: ", err)
			}
			_ = file.Close()
			_ = os.Remove(filePath)

			// Read out from routine pool before uploading so another file can begin data generation
			<-routinePool
			wg.Done()
		}()
	}

	// Wait for all threads to finish
	wg.Wait()
}

func generateRandomRow() appleWatch3Row.AppleWatch3Row {
	return appleWatch3Row.AppleWatch3Row{
		Ts: rand.Uint64(),
		Rx: rand.Float64(),
		Ry: rand.Float64(),
		Rz: rand.Float64(),
		Rl: rand.Float64(),
		Pt: rand.Float64(),
		Yw: rand.Float64(),
		Ax: rand.Float64(),
		Ay: rand.Float64(),
		Az: rand.Float64(),
		Hr: rand.Float64(),
	}
}

func makeFileWithRandomRows(numRows int) string {
	// Make a new file
	filePath := localFilePrefix + strconv.Itoa(int(time.Now().UnixNano())) + ".parquet"
	pqFile, _ := parquetHandler.MakeParquetFile(filePath)

	// Write rows to the file
	for i := 0; i < numRows; i++ {
		err := pqFile.WriteRow(generateRandomRow())
		if err != nil {
			log.Println("Error writing row to parquet file. ", err)
		}
	}

	// Close the file
	err := pqFile.CloseFile()
	if err != nil {
		log.Println("Unable to close file: ", filePath, err)
	}
	return filePath
}
