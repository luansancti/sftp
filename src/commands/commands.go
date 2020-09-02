package commands

import (
	"fmt"
	"helper"
	"io/ioutil"
	"log"
	"os"
	"strings"
	"syscall"
	"user"
	//"strings"
	//"os"
	//"log"
)

func CreateUser(person user.User) bool {

	if !helper.CheckUserExists(person.User) {

		person = user.NewUser(person.User, person.Password, person.Expiration)
		err := os.MkdirAll(person.PathUserUp, os.ModePerm)
		if err != nil {
			fmt.Println(err.Error())
			return false
		}
		err = os.MkdirAll(person.PathUserPublic, os.ModePerm)
		if err != nil {
			fmt.Println(err.Error())
			return false
		}

		helper.Execute(fmt.Sprint("useradd --groups sftpgroup --shell=/bin/false -d", " ", person.PathUser, " ", person.User))
		helper.Execute(fmt.Sprint("echo -e", " ", `"`, person.Password, `\n`, person.Password, `"`, " ", "|", " ", "passwd", " ", person.User))
		FixPermission(person)
		syscall.Mount(person.PathPublic, person.PathUserPublic, "none", syscall.MS_BIND, "")
		AddCacheuser(person)

	} else {
		fmt.Println("usuário existe")
		return false
	}
	return true
}

func FixPermission(person user.User) {

	helper.Execute(fmt.Sprint("chown root", " ", person.PathUser))

	helper.Execute(fmt.Sprint("chmod go-w", " ", person.PathUser))

	helper.Execute(fmt.Sprint("chown -R", " ", person.User, ":", "sftpgroup", " ", person.PathUserUp))

	helper.Execute(fmt.Sprint("chmod ug+rwX", " ", person.PathUserUp))
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
		log.Println(err)
	}

	defer f.Close()

	if _, err := f.WriteString(shadowArray + "\n"); err != nil {
		log.Println(err)
	}
}
