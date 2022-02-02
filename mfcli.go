package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"
)

const check_dir string = "/tmp/mfchecks/"

func main() {

	newpath := filepath.Join(check_dir)
	err := os.MkdirAll(newpath, os.ModePerm)

	// https://stackoverflow.com/questions/14668850/list-directory-in-go
	files, err := ioutil.ReadDir(check_dir)
	if err != nil {
		log.Fatal(err)
	}

	for _, filename := range files {
		//buf := bytes.NewBuffer(nil)
		// https://stackoverflow.com/questions/13514184/how-can-i-read-a-whole-file-into-a-string-variable
		//b, err := ioutil.ReadFile(check_dir + filename.Name()) // just pass the file name
		//if err != nil {
		//	fmt.Print(err)
		//}
		secsCheck, err := readFile(check_dir + filename.Name())
		if err != nil {
			panic(err)
		}

		//var secsNow, ago uint64
		var secsNow, agoSecs int
		now := time.Now() // current local time
		secsNow = int(now.Unix())
		//secsCheck := binary.BigEndian.Uint64(b)
		//secsCheck = int(b)
		//secsCheck, _ = strconv.Atoi(string(b))
		//secsCheck = int(b)

		//secsCheck, err := ReadInts(strings.NewReader(b))

		agoSecs = secsNow - secsCheck[0]
		//fmt.Println(filename.Name() + " - " + string(secsNow) + " - " + string(secsCheck[0]) + " - " + string(ago))
		fmt.Printf("%s -  %d - %d secs (%d minutes)\n", filename.Name(), secsNow, agoSecs, agoSecs/60)
	}

}

func readFile(fname string) (nums []int, err error) {
	b, err := ioutil.ReadFile(fname)
	if err != nil {
		return nil, err
	}

	lines := strings.Split(string(b), "\n")
	// Assign cap to avoid resize on every append.
	nums = make([]int, 0, len(lines))

	for _, l := range lines {
		// Empty line occurs at the end of the file when we use Split.
		if len(l) == 0 {
			continue
		}
		// Atoi better suits the job when we know exactly what we're dealing
		// with. Scanf is the more general option.
		n, err := strconv.Atoi(l)
		if err != nil {
			return nil, err
		}
		nums = append(nums, n)
	}

	return nums, nil
}
