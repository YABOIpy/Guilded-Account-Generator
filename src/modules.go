package src

import (
	"bytes"
	crnd "crypto/rand"
	"encoding/hex"
	"encoding/json"
	"log"
	"math/rand"
	"os"
	"strconv"
	"time"

	"github.com/Pallinder/go-randomdata"
)

func (Gu *Guilded) Err(err error) (er error) {
	if err != nil {
		log.Fatal(err)
	}
	return er
}

func (Gu *Guilded) Load(data map[string]string) (payload *bytes.Buffer, err error) {
	d, err := json.Marshal(data)
	Err := Gu.Err(err)
	payload = bytes.NewBuffer(d)
	return payload, Err
}

func (Gu *Guilded) RandHex(n int) string {
	bytes := make([]byte, n)
	if _, err := crnd.Read(bytes); err != nil {
		log.Fatal(err)
	}
	return hex.EncodeToString(bytes)
}

func (Gu *Guilded) RandMail() (string, string) {
	rand.Seed(time.Now().Unix())
	domains := []string{
		"@gmail.com",
		"@hotmail.com",
		"@yahoo.ca",
		"@verizon.net",
		"@msn.com",
		"@aol.com",
		"@outlook.com",
	}
	date := randomdata.Number(1940, 2000)
	name := randomdata.FirstName(randomdata.Male) + strconv.Itoa(date)
	x := rand.Int() % len(domains)
	email := name + domains[x]

	return email, name
}

func (Gu *Guilded) WriteFile(files string, item string) {
	f, err := os.OpenFile(files, os.O_RDWR|os.O_APPEND, 0660)
	Gu.Err(err)
	defer f.Close()
	_, ers := f.WriteString(item + "\n")
	Gu.Err(ers)
}

func (Gu *Guilded) Config() Guilded {
	var config Guilded
	conf, err := os.Open("config.json")
	defer conf.Close()
	Gu.Err(err)
	xp := json.NewDecoder(conf)
	xp.Decode(&config)
	return config

}
