package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"github.com/gregdel/pushover"
)

const check_dir string = "/tmp/mfchecks/"
const thresholdSecs_Warning int = 70
const thresholdSecs_Critical int = 120

type Settings struct {
	App_key       string `json:"app_key"`
	Recipient_key string `json:"recipient_key"`
}

var settings Settings

func main() {
	//jsonFile, err := os.Open("config.json")
	jsonFile, err := os.Open("/home/josh/.mf/mfcli.config")

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer jsonFile.Close()

	// read our opened jsonFile as a byte array.
	byteValue, _ := ioutil.ReadAll(jsonFile)

	// we initialize our Users array
	//var users Users

	// we unmarshal our byteArray which contains our
	// jsonFile's content into 'users' which we defined above
	json.Unmarshal(byteValue, &settings)

	app := pushover.New(settings.App_key)
	recipient := pushover.NewRecipient(settings.Recipient_key)

	newpath := filepath.Join(check_dir)
	err = os.MkdirAll(newpath, os.ModePerm)

	// https://stackoverflow.com/questions/14668850/list-directory-in-go
	files, err := ioutil.ReadDir(check_dir)
	if err != nil {
		log.Fatal(err)
	}

	for _, filename := range files {

		secsCheck, err := readFile(check_dir + filename.Name())
		if err != nil {
			panic(err)
		}

		//var secsNow, ago uint64
		var secsNow, agoSecs int
		now := time.Now() // current local time
		secsNow = int(now.Unix())
		agoSecs = secsNow - secsCheck[0]
		//fmt.Println(filename.Name() + " - " + string(secsNow) + " - " + string(secsCheck[0]) + " - " + string(ago))
		fmt.Printf("%s -  %d - %d secs (%d minutes)\n", filename.Name(), secsNow, agoSecs, agoSecs/60)

		// TODO: implement a critical time period too.
		if agoSecs > thresholdSecs_Warning {

			// TODO: check when the last notification was sent
			// TODO: have a duration of wait time between notifications

			s := strings.Split(filename.Name(), ".")
			fmt.Printf("******* WARNING: %s *******\n", s[0])

			msg := "WARNING: " + s[0] + " has not checked in"
			message := pushover.NewMessage(msg)

			// Send the message to the recipient
			_, err := app.SendMessage(message, recipient)
			if err != nil {
				log.Panic(err)
			}

			// TODO: keep track of when the last notification was sent

		}

		// TODO: implement a recovery alert when it goes from bad to good

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
