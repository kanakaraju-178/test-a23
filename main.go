package main

// gronit
// [go]cron[monitor]

import (
	_ "flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	args := os.Args[1:]
	if len(args) < 1 {
		help()
		os.Exit(1)
	}

	sys := defaultSys()
	opts := parseOptions(sys, args)
	tasks := getTasks(sys, opts)

	if opts.LoadYaml != EMPTYSTR || opts.LoadJson != EMPTYSTR || opts.LoadCron != EMPTYSTR {
		tasksToCron(tasks, sys)
	}

	os.Exit(1)

}

// update TODO
func update() {
	// PROTOTYPE:
	// okay now we can write to cron, need to pass in arguments like:
	// 	-- yaml,json file or text
	// need to check if cron exists already
	// Writing tmp file
	fakeCron := []byte("* * * * * touch $HOME/yooo\n")
	tmpfile, err := ioutil.TempFile("", "example")
	if err != nil {
		log.Fatal(err)
	}
	defer os.Remove(tmpfile.Name())
	fmt.Println(tmpfile.Name())
	if _, err := tmpfile.Write(fakeCron); err != nil {
		log.Fatal(err)
	}
	if err := tmpfile.Close(); err != nil {
		log.Fatal(err)
	}

	// Reading file
	content, err := ioutil.ReadFile("/var/spool/cron/crontabs/root")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%s", append(content[:], fakeCron[:]...))
	if err := ioutil.WriteFile("/var/spool/cron/crontabs/root",
		append(content[:], fakeCron[:]...), 0777); err != nil {
		log.Fatal(err)
	}
}
