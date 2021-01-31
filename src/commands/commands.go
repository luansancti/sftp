package commands

import (
	"fmt"
	"helper"
	"io/ioutil"
	"log"
	models "models"
	"os"
	"path"
	"path/filepath"
	"strconv"
	"strings"
	"syscall"
	"time"
	"user"
)

func CreateUser(person user.User) models.DefaultResponse {

	if helper.CheckUserExists(person.User) {
		return models.ResponseDefault(fmt.Sprint("User exists: ", person.User), false)
	} else {
		fmt.Println("not usuario existe")
		person = user.NewUser(person.User, person.Password, person.Expiration)
		err := os.MkdirAll(person.PathUserUp, 0755)
		if err != nil {
			log.Fatalln(err)
			return models.ResponseDefault(string(err.Error()), false)
		}
		err = os.MkdirAll(person.PathUserPublic, 0755)
		if err != nil {
			log.Fatalln(err)
			return models.ResponseDefault(string(err.Error()), false)
		}
		if person.Expiration <= 0 {
			if !helper.Execute(fmt.Sprint("useradd --groups sftpgroup --shell=/bin/false -d", " ", person.PathUser, " ", person.User)) {
				return models.ResponseDefault(fmt.Sprint("Error to create user: ", person.User), false)
			}
		} else {
			if !helper.Execute(fmt.Sprint("useradd --groups sftpgroup --shell=/bin/false -d", " ", person.PathUser, " ", person.User, " ", "-e ", time.Now().Local().AddDate(0, 0, person.Expiration).Format("2006-01-02"))) {
				return models.ResponseDefault(fmt.Sprint("Error to create user: ", person.User), false)
			}
		}

		if !helper.Execute(fmt.Sprint("echo -e", " ", `"`, person.Password, `\n`, person.Password, `"`, " ", "|", " ", "passwd", " ", person.User)) {
			return models.ResponseDefault(fmt.Sprint("Error to create user: ", person.User), false)
		}
		FixPermission(person)
		if err = syscall.Mount(person.PathPublic, person.PathUserPublic, "none", syscall.MS_BIND, ""); err != nil {
			return models.ResponseDefault("erro no mount sangue bom", false)
		}
		AddCacheuser(person)
		return models.ResponseDefault("Usuário criado com sucesso", true)
	}

}

func FixPermission(person user.User) models.DefaultResponse {

	person = user.NewUser(person.User, "", 0)
	if !helper.Execute(fmt.Sprint("chown root", " ", person.PathUser)) {
		return models.ResponseDefault(fmt.Sprint("Error to fix permission user: ", person.User), false)
	}
	if !helper.Execute(fmt.Sprint("chmod go-w", " ", person.PathUser)) {
		return models.ResponseDefault(fmt.Sprint("Error to fix permission user: ", person.User), false)
	}
	if !helper.Execute(fmt.Sprint("chown -R", " ", person.User, ":", "sftpgroup", " ", person.PathUserUp)) {
		return models.ResponseDefault(fmt.Sprint("Error to fix permission user: ", person.User), false)
	}
	if !helper.Execute(fmt.Sprint("chmod ug+rwX", " ", person.PathUserUp)) {
		return models.ResponseDefault(fmt.Sprint("Error to fix permission user: ", person.User), false)
	}
	return models.ResponseDefault(fmt.Sprint("Fix permission user: ", person.User), true)

}

func GetUsersLogged() models.ResponseData {

	outputCommand := helper.ExecuteReturn("ps -ef | grep '[s]shd' | grep -v ^root | awk '{print$9}' | grep 'sftp' | sed 's/@internal-sftp//g'")

	//outputCommand := helper.ExecuteReturn("netstat -tnpa | grep 'ESTABLISHED.*sshd' | awk '{print $8}'")
	fmt.Println(string(outputCommand))
	return models.DataResponse(true, "", helper.Delete_Empty(strings.Split(outputCommand, "\n")))
}

func Unlink_User(person user.User) models.DefaultResponse {
	person = user.NewUser(person.User, "", 0)
	if !helper.Execute(fmt.Sprint("pkill -u", " ", person.User, " ", "||", " ", "true")) {
		return models.ResponseDefault(fmt.Sprint("Error to unlink user: ", person.User), false)
	}
	return models.ResponseDefault(fmt.Sprint("User unlinked: ", person.User), true)
}

func ChangeExpiration(person user.User) models.DefaultResponse {

	if !helper.Execute(fmt.Sprint("chage", " ", "-E", time.Now().Local().AddDate(0, 0, person.Expiration).Format("2006-01-02"), " ", person.User)) {
		return models.ResponseDefault(fmt.Sprint("Error to change expiration: ", person.User), false)
	}

	return models.ResponseDefault(fmt.Sprint("Expiration changed: ", person.User), true)
}

func ChangePassword(person user.User) models.DefaultResponse {

	if !helper.Execute(fmt.Sprint("echo -e", " ", `"`, person.Password, `\n`, person.Password, `"`, " ", "|", " ", "passwd", " ", person.User)) {
		return models.ResponseDefault(fmt.Sprint("Error change password user: ", person.User), false)
	}
	return models.ResponseDefault(fmt.Sprint("Changed Password: ", person.User), true)
}

func DeleteUser(person user.User) models.DefaultResponse {

	person = user.NewUser(person.User, "", 0)
	if !helper.Execute(fmt.Sprint("pkill -u", " ", person.User, " ", "||", " ", "true")) {
		return models.ResponseDefault(fmt.Sprint("Error to delete user: ", person.User), false)
	}
	if !helper.Execute(fmt.Sprint("userdel -f", " ", person.User)) {
		return models.ResponseDefault(fmt.Sprint("Error to delete user: ", person.User), false)
	}
	if err := syscall.Unmount(person.PathUserPublic, 0); err != nil {
		fmt.Println(err)
		return models.ResponseDefault(fmt.Sprint("Error to delete user: ", person.User), false)
	}
	os.RemoveAll(person.PathUser)
	RemoveCacheUser(person)
	return models.ResponseDefault(fmt.Sprint("User removed: ", person.User), true)
}

func ReturnPathKey(person user.User) models.KeyResponse {
	person = user.NewUser(person.User, "", 0)
	pathSsh := path.Join(person.PathUser, ".ssh", "id_rsa")
	data, _ := ioutil.ReadFile(pathSsh)
	return models.ResponseKey(true, "Get Key", string(data))

}

func ListUsers() models.ListUser {
	data, err := ioutil.ReadFile("/etc/shadow")
	if err != nil {
		fmt.Println("File reading error", err)
		models.ResponseListUsers(fmt.Sprint("List update"), false, []models.UserDetails{})
	}

	indexString := strings.Index(string(data), "admin")
	dataRune := []rune(string(data))
	shadowArray := strings.Split(string(dataRune[indexString:len(string(data))]), "\n")
	listUserDeta := []models.UserDetails{}

	for _, line := range shadowArray[1:] {
		if line != "" {
			userDetails := models.UserDetails{}
			username := strings.Split(line, ":")[0]
			expired, _ := strconv.ParseInt(strings.Split(line, ":")[7], 10, 64)
			person := user.NewUser(username, "", 0)
			infoUp, err := os.Lstat(person.PathUserUp)

			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}

			userDetails.UserName = username

			if userDetails.Expiration = "Never"; expired == 0 {

			} else {
				userDetails.Expiration = fmt.Sprint(expired * 86400)
			}

			infoSsh, err := os.Lstat(path.Join(person.PathUser, ".ssh"))
			if err != nil {
				userDetails.Key = false
				userDetails.Size = helper.SizedDisk(person.PathUserUp, infoUp)

			} else {

				userDetails.Key = true
				userDetails.Size = helper.SizedDisk(person.PathUserUp, infoUp) + helper.SizedDisk(path.Join(person.PathUser, ".ssh"), infoSsh)
			}
			listUserDeta = append(listUserDeta, userDetails)
		}
	}
	fmt.Println(listUserDeta)
	return models.ResponseListUsers(fmt.Sprint("List update"), true, listUserDeta)
}

func ListDirectory(pathName string) []models.DirectoryInfo {

	arrayFolder := []models.DirectoryInfo{}
	fmt.Println(helper.GetConfigPaths().UsersPath)
	folder := models.DirectoryInfo{}

	files, err := filepath.Glob(pathName)
	if err != nil {
		log.Fatal(err)
	}

	for _, s := range files {
		fi, err := os.Stat(s)
		if err != nil {
			fmt.Println(err)
			return arrayFolder
		}
		switch mode := fi.Mode(); {
		case mode.IsDir():
			folder.IsDirectory = true
		case mode.IsRegular():
			folder.IsDirectory = false
		}
		//stat, _ := os.Lstat(s)
		//folder.Size = helper.SizedDisk(pathName, stat)
		folder.Size = fi.Size()
		folder.Name = pathName
		folder.ModTime = fi.ModTime()
		arrayFolder = append(arrayFolder, folder)
	}

	return arrayFolder
}

func DiskPercent() models.DirectoryPerc {
	configPath := helper.GetConfigPaths()
	listPercentageUsage := []models.PercentageUsed{}

	percentageUsers := models.PercentageUsed{}
	percentagePublic := models.PercentageUsed{}

	percentageUsers.DirectoryName = configPath.UsersPath
	percentageUsers.Percentage = helper.DiskUsage(configPath.UsersPath)

	percentagePublic.DirectoryName = configPath.PublicPath
	percentagePublic.Percentage = helper.DiskUsage(configPath.PublicPath)

	listPercentageUsage = append(listPercentageUsage, percentageUsers, percentagePublic)

	return models.ResponseDirectoryPerc("", true, listPercentageUsage)

}

func DirSize(path string) (int64, error) {
	var size int64
	err := filepath.Walk(path, func(_ string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			size += info.Size()
		}
		return err
	})
	return size, err
}

func CreateUserKey(person user.User) models.DefaultResponse {

	if helper.CheckUserExists(person.User) {
		return models.ResponseDefault(fmt.Sprint("User exists: ", person.User), false)
	} else {

		person = user.NewUser(person.User, person.Password, person.Expiration)
		err := os.MkdirAll(person.PathUserUp, 0755)
		if err != nil {
			return models.ResponseDefault(fmt.Sprint("Error to create user: ", person.User), false)
		}
		err = os.MkdirAll(person.PathUserPublic, 0755)
		if err != nil {
			return models.ResponseDefault(fmt.Sprint("Error to create user: ", person.User), false)
		}

		if person.Expiration <= 0 {
			if !helper.Execute(fmt.Sprint("useradd --groups sftpgroup --shell=/bin/false -d", " ", person.PathUser, " ", person.User)) {
				return models.ResponseDefault(fmt.Sprint("Error to create user: ", person.User), false)
			}
		} else {
			if !helper.Execute(fmt.Sprint("useradd --groups sftpgroup --shell=/bin/false -d", " ", person.PathUser, " ", person.User, " ", "-e ", time.Now().Local().AddDate(0, 0, person.Expiration).Format("2006-01-02"))) {
				return models.ResponseDefault(fmt.Sprint("Error to create user: ", person.User), false)
			}
		}

		if !helper.Execute(fmt.Sprint("passwd -l", " ", person.User)) {
			return models.ResponseDefault(fmt.Sprint("Error to create user: ", person.User), false)
		}
		FixPermission(person)
		if err = syscall.Mount(person.PathPublic, person.PathUserPublic, "none", syscall.MS_BIND, ""); err != nil {
			return models.ResponseDefault(fmt.Sprint("Error to create user: ", person.User), false)
		}

		person = user.NewUser(person.User, "", 0)

		pathSsh := path.Join(person.PathUser, ".ssh")
		pathSshId := path.Join(pathSsh, "id_rsa")
		fileAuthorized := path.Join(pathSsh, "authorized_keys")
		keyPub := path.Join(pathSsh, "id_rsa.pub")

		err = os.MkdirAll(pathSsh, 0700)
		if err != nil {
			return models.ResponseDefault(fmt.Sprint("Error to create user: ", person.User), false)
		}

		_, err = os.Create(fileAuthorized)

		if !helper.Execute(fmt.Sprint("ssh-keygen -b 2048 -N", " ", `""`, " ", "-t rsa -f", " ", pathSshId, " ", "<<<y")) {
			return models.ResponseDefault(fmt.Sprint("Error to create user: ", person.User), false)
		}

		if !helper.Execute(fmt.Sprint("chown -R", " ", person.User, ":", person.User, " ", pathSsh)) {
			return models.ResponseDefault(fmt.Sprint("Error to create user: ", person.User), false)
		}

		if !helper.Execute(fmt.Sprint("chmod 700", " ", pathSsh)) {
			return models.ResponseDefault(fmt.Sprint("Error to create user: ", person.User), false)
		}
		if !helper.Execute(fmt.Sprint("chmod 600", " ", fileAuthorized)) {
			return models.ResponseDefault(fmt.Sprint("Error to create user: ", person.User), false)
		}

		input, err := ioutil.ReadFile(keyPub)
		if err != nil {
			return models.ResponseDefault(fmt.Sprint("Error to create user: ", person.User), false)
		}

		err = ioutil.WriteFile(fileAuthorized, []byte(input), 0700)
		if err != nil {
			return models.ResponseDefault(fmt.Sprint("Error to create user: ", person.User), false)
		}
		AddCacheuser(person)
		return models.ResponseDefault(fmt.Sprint("Create user: ", person.User), true)
	}

}

func AddCacheuser(person user.User) {

	data, err := ioutil.ReadFile("/etc/shadow")

	if err != nil {
		fmt.Println("File reading error", err)
	}

	indexString := strings.Index(string(data), person.User)
	dataRune := []rune(string(data))
	shadowArray := strings.Split(string(dataRune[indexString:len(string(data))]), "\n")[0]
	cacheLocale := helper.GetConfigPaths().CacheUsers

	f, err := os.OpenFile(cacheLocale,
		os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)

	if err != nil {
		fmt.Println(err)
	}

	defer f.Close()

	if _, err := f.WriteString(shadowArray + "\n"); err != nil {
		fmt.Println(err)
	}
}

func RemoveCacheUser(person user.User) {

	cacheLocale := helper.GetConfigPaths().CacheUsers

	input, err := ioutil.ReadFile(cacheLocale)
	if err != nil {
		fmt.Println(err)
	}

	lines := strings.Split(string(input), "\n")

	var s []string

	for _, line := range lines {
		if strings.Split(line, ":")[0] != person.User {
			s = append(s, line)
		}
	}

	output := strings.Join(s, "\n")
	err = ioutil.WriteFile(cacheLocale, []byte(output), 0644)
	if err != nil {
		fmt.Println(err)
	}
}
