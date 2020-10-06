package helper

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path"
	"strings"
	"syscall"
	"time"
)

type Config struct {
	UsersPath  string `json:"usersPath"`
	PublicPath string `json:"publicPath"`
	CacheUsers string `json:"cacheUsers"`
}

type DiskStatus struct {
	All  uint64 `json:"all"`
	Used uint64 `json:"used"`
	Free uint64 `json:"free"`
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

func SizedDisk(currentPath string, info os.FileInfo) int64 {
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
		size += SizedDisk(currentPath+"/"+file.Name(), file)
	}

	fmt.Printf("Size in bytes : [%d] : [%s]\n", size, currentPath)

	return size
}

func DiskUsage(path string) float64 {
	disk := DiskStatus{}
	fs := syscall.Statfs_t{}
	err := syscall.Statfs(path, &fs)
	if err != nil {
		return -1.0
	}
	disk.All = fs.Blocks * uint64(fs.Bsize)
	disk.Free = fs.Bfree * uint64(fs.Bsize)
	disk.Used = disk.All - disk.Free
	return (float64(disk.Used) / float64(disk.All)) * float64(100)
}

func DiffDate(t1 time.Time) int {
	t2 := time.Now()
	return int(t2.Sub(t1).Hours() / 24)
}
