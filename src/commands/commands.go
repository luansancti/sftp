package commands

import (
	"fmt"
	"helper"
	"io/ioutil"
	"log"
	models "models"
	"os"
	"path"
	"strconv"
	"strings"
	"syscall"
	"time"
	"user"
	//"strings"
	//"os"
	//"log"
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
		if person.Expiration == -1 {
			if !helper.Execute(fmt.Sprint("useradd --groups sftpgroup --shell=/bin/false -d", " ", person.PathUser, " ", person.User)) {
				return models.ResponseDefault(fmt.Sprint("Error to create user: ", person.User), false)
			}
		} else {
			if !helper.Execute(fmt.Sprint("useradd --groups sftpgroup --shell=/bin/false -d", " ", person.PathUser, " ", person.User, " ", "-e ", time.Now().Local().AddDate(0, 0, 7).Format("2006-01-02"))) {
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

func DeleteUser(person user.User) models.DefaultResponse {

	person = user.NewUser(person.User, "", 0)
	if !helper.Execute(fmt.Sprint("pkill -u", " ", person.User)) {
		return models.ResponseDefault(fmt.Sprint("Error to delete user: ", person.User), false)
	}
	if !helper.Execute(fmt.Sprint("userdel -f", " ", person.User)) {
		return models.ResponseDefault(fmt.Sprint("Error to delete user: ", person.User), false)
	}
	if err := syscall.Unmount(person.PathUserPublic, 0); err != nil {
		return models.ResponseDefault(fmt.Sprint("Error to delete user: ", person.User), false)
	}
	os.RemoveAll(person.PathUser)
	RemoveCacheUser(person)
	return models.ResponseDefault(fmt.Sprint("User removed: ", person.User), true)
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
			userUp, err := os.Stat(person.PathUserUp)
			if err != nil {
				models.ResponseListUsers(fmt.Sprint("List update"), false, []models.UserDetails{})
			}

			userDetails.UserName = username
			userDetails.Expiration = time.Unix((expired * 86400), 0)

			userSsh, err := os.Stat(path.Join(person.PathUser, ".ssh"))
			if err != nil {

				userDetails.Key = false
				userDetails.Size = userUp.Size()
				fmt.Println(userDetails.Expiration)

			} else {
				userDetails.Key = true
				userDetails.Size = userUp.Size() + userSsh.Size()
			}
			listUserDeta = append(listUserDeta, userDetails)
		}

	}
	return models.ResponseListUsers(fmt.Sprint("List update"), true, listUserDeta)
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

		if person.Expiration == -1 {
			if !helper.Execute(fmt.Sprint("useradd --groups sftpgroup --shell=/bin/false -d", " ", person.PathUser, " ", person.User)) {
				return models.ResponseDefault(fmt.Sprint("Error to create user: ", person.User), false)
			}
		} else {
			if !helper.Execute(fmt.Sprint("useradd --groups sftpgroup --shell=/bin/false -d", " ", person.PathUser, " ", person.User, " ", "-e ", time.Now().Local().AddDate(0, 0, 7).Format("2006-01-02"))) {
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

		if !helper.Execute(fmt.Sprint("ssh-keygen -b 4096 -N ", " ", `""`, " ", "-t rsa -f", " ", pathSshId)) {
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
		log.Fatalf("File reading error", err)
	}

	indexString := strings.Index(string(data), person.User)
	dataRune := []rune(string(data))
	shadowArray := strings.Split(string(dataRune[indexString:len(string(data))]), "\n")[0]
	cacheLocale := helper.GetConfigPaths().CacheUsers

	f, err := os.OpenFile(cacheLocale,
		os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)

	if err != nil {
		log.Fatalln(err)
	}

	defer f.Close()

	if _, err := f.WriteString(shadowArray + "\n"); err != nil {
		log.Fatalln(err)
	}
}

func RemoveCacheUser(person user.User) {

	cacheLocale := helper.GetConfigPaths().CacheUsers

	input, err := ioutil.ReadFile(cacheLocale)
	if err != nil {
		log.Fatalln(err)
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
		log.Fatalln(err)
	}
}
