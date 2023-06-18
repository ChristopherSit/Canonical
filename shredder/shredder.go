package shredder

import (
	"io/ioutil"
	"os"
	"crypto/rand"
	"math/big"
	// "fmt"
)

func Shred(filePath string) error {
	// get file permissions
	fileInfo, err := os.Stat(filePath)
	if err != nil {
		return err
	}

	// get file size
	fileSize := fileInfo.Size()

	// loop for scrambling data 3 times
	for x := 0; x < 3; x++ {
		// get random size to determine data to write based on fileSize * 10
		dataSize, err := rand.Int(rand.Reader, big.NewInt(fileSize*10))
		if err != nil {
			return err
		}
		randomSize := dataSize.Int64()

		// create byte slice equal to random size
		data := make([]byte, randomSize)
		// replace all 0s with actual random data using rand.Read, we ignore the # bytes read
		_, err = rand.Read(data)
		// fmt.Println(data) if you want to check
		if err != nil {
			return err
		}

		// preserve permissions when writting new garbage data
		err = ioutil.WriteFile(filePath, data, fileInfo.Mode())
		if err != nil {
			return err
		}
	}

	// delete file
	err = os.Remove(filePath)
	if err != nil {
		return err
	}

	return nil
}