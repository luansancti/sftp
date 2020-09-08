package commands

import (
	"fmt"
	"helper"
	"io/ioutil"
	"log"
	"models"
	"os"
	"path"
	"strings"
	"syscall"
	"user"
	//"strings"
	//"os"
	//"log"
)

func CreateUser(person user.User) models.DefaultResponse {

	if helper.CheckUserExists(person.User) {
		return models.ResponseDefault(fmt.Sprint("User exists: ", person.User), false)
	} else {
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

		if !helper.Execute(fmt.Sprint("useradd --groups sftpgroup --shell=/bin/false -d", " ", person.PathUser, " ", person.User)) {
			return models.ResponseDefault(fmt.Sprint("Error to create user: ", person.User), false)
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

func ListUsers() bool {
	data, err := ioutil.ReadFile("/etc/shadow")
	if err != nil {
		fmt.Println("File reading error", err)
		return true
	}

	indexString := strings.Index(string(data), "admin")

	dataRune := []rune(string(data))

	shadowArray := strings.Split(string(dataRune[indexString:len(string(data))]), "\n")

	for _, line := range shadowArray[1:] {
		if line != "" {
			username := strings.Split(line, ":")[0]
			//sexpired := strings.Split(line, ":")[7]
			person := user.NewUser(username, "", 0)
			fmt.Println(username)
			userUp, err := os.Stat(person.PathUserUp)
			if err != nil {
				return false
			}
			userSsh, err := os.Stat(path.Join(person.PathUser, ".ssh"))
			if err != nil {
				fmt.Println("cai no erro")
				fmt.Println(userUp.Size())
				return false

			}
			fmt.Println("up size: ", userUp.Name())
			fmt.Println("ssh size: ", userSsh.Name())

		}

	}
	return false

}

func CreateUserKey(person user.User) models.DefaultResponse {

	person = user.NewUser(person.User, person.Password, person.Expiration)
	err := os.MkdirAll(person.PathUserUp, 0755)
	if err != nil {
		return models.ResponseDefault(fmt.Sprint("Error to create user: ", person.User), false)
	}
	err = os.MkdirAll(person.PathUserPublic, 0755)
	if err != nil {
		return models.ResponseDefault(fmt.Sprint("Error to create user: ", person.User), false)
	}

	if !helper.Execute(fmt.Sprint("useradd --groups sftpgroup --shell=/bin/false -d", " ", person.PathUser, " ", person.User)) {
		return models.ResponseDefault(fmt.Sprint("Error to create user: ", person.User), false)
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
