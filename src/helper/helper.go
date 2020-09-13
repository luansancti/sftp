package helper

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path"
	"strings"
)

type Config struct {
	UsersPath  string `json:"usersPath"`
	PublicPath string `json:"publicPath"`
	CacheUsers string `json:"cacheUsers"`
}

func GetConfigPaths() Config {

	config := Config{}
	mydir, err := os.Getwd()
	configFolder := path.Join(mydir, "conf/appsettings.json")
	jsonFile, err := os.Open(configFolder)

	if err != nil {
		fmt.Println("error capture Json config", err.Error())
	}

	byteValue, _ := ioutil.ReadAll(jsonFile)

	json.Unmarshal(byteValue, &config)
	return config
}

func CheckUserExists(username string) bool {

	data, err := ioutil.ReadFile("/etc/shadow")
	if err != nil {
		fmt.Println("File reading error", err)
		return true
	}

	indexString := strings.Index(string(data), "admin")

	dataRune := []rune(string(data))

	shadowArray := strings.Split(string(dataRune[indexString:len(string(data))]), "\n")

	for _, line := range shadowArray {
		if strings.Split(line, ":")[0] == username {
			return true
		}
	}
	return false

}

func Execute(command string) bool {
	_, err := exec.Command("bash", "-c", command).Output()

	if err != nil {
		fmt.Println("error", command, err.Error())
		return false
	}

	return true
}

func DiskUsage(currentPath string, info os.FileInfo) int64 {
	size := info.Size()

	if !info.IsDir() {
		return size
	}

	dir, err := os.Open(currentPath)

	if err != nil {
		fmt.Println(err)
		return size
	}
	defer dir.Close()

	files, err := dir.Readdir(-1)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	for _, file := range files {
		if file.Name() == "." || file.Name() == ".." {
			continue
		}
		size += DiskUsage(currentPath+"/"+file.Name(), file)
	}

	fmt.Printf("Size in bytes : [%d] : [%s]\n", size, currentPath)

	return size
}
