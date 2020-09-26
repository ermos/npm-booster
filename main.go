package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"io/ioutil"
	"log"
	"math/rand"
	"os"
	"os/exec"
	"os/user"
	"path/filepath"
	"strconv"
	"time"
)

var iteration int
var count int
var dirBuild string
var usrName string

func main () {
	var err error

	_, err = exec.LookPath("npm")
	if err != nil {
		_error("you need to install npm for npm booster")
	}

	rand.Seed(time.Now().UnixNano())

	start := time.Now()

	if len(os.Args) != 3 {
		logger(WarningColor + "Usage : " + InfoColor + "npm-booster [count] [package]" + ResetColor)
		os.Exit(0)
	}

	usr, err := user.Current()
	if err != nil {
		log.Fatal( err )
	}

	dirBuild = filepath.Join(usr.HomeDir, ".npm-booster")
	usrName = usr.Name

	defer os.RemoveAll(dirBuild)

	packageName := os.Args[2]

	count, err = strconv.Atoi(os.Args[1])
	if err != nil {
		_error(err.Error())
	}

	var nbGoRoutine int
	if len(os.Args[1]) > 1 {
		nbGoRoutine, err = strconv.Atoi(os.Args[1][:len(os.Args[1])-1])
		if err != nil {
			_error(err.Error())
		}
	}

	var last int
	if os.Args[1][len(os.Args[1])-1] != 0 {
		last, err = strconv.Atoi(string(os.Args[1][len(os.Args[1])-1]))
		if err != nil {
			_error(err.Error())
		}
	}

	ctx := context.Background()

	if nbGoRoutine != 0 {
		nbPerLoop := 10
		for i := 0; i < nbGoRoutine; i++ {
			c := make(chan bool, nbPerLoop)
			for nb := 0; nb < nbPerLoop; nb++ {
				go npmInstall(ctx, c, packageName)
			}
			for w := 0; w < nbPerLoop; w++ {
				select {
				case <- c:
					iteration++
					loggerCount(iteration, count)
				}
			}
		}
	}

	if last != 0 {
		c := make(chan bool, last)
		for nb := 0; nb < last; nb++ {
			go npmInstall(ctx, c, packageName)
		}
		for w := 0; w < last; w++ {
			select {
			case <- c:
				iteration++
				loggerCount(iteration, count)
			}
		}
	}

	logger("Ending of process..")
	logger("For build %d times \"%s\", we take %.2fs", count, packageName, time.Since(start).Seconds())
}

func npmInstall (ctx context.Context, c chan bool, packageName string) {
	//https://registry.npmjs.org/css-vars-manager/-/css-vars-manager-1.1.0.tgz
	dirName := fmt.Sprintf("folder-%s", uuid.New().String())
	dirPath := filepath.Join(dirBuild, dirName)

	err := os.MkdirAll(dirPath, os.ModePerm)
	if err != nil {
		_error(err.Error())
	}

	packageJSON, _ := json.MarshalIndent(newPackageJSON(), "", "  ")
	_ = ioutil.WriteFile(filepath.Join(dirPath, "package.json"), packageJSON, 0644)

	execute(&dirPath, "npm", "install", packageName)

	c <- true
}

func execute(dir *string, name string, args... string) {
	ch := make(chan error)

	cmd := exec.Command(name, args...)
	if dir != nil {
		cmd.Dir = *dir
	}
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	go func(){
		ch <- cmd.Run()
	}()

	select{
	case err := <- ch:
		if err != nil {
			_error("command failed with %s\n", err)
		}
	}
}