package utils

import "log"

func CheckErr(err error, msg string) {
	if err != nil {
		log.Fatalln(msg, err)
	}
}

func Debug(data interface{}) {
	log.Println("+++++++++++", data)
}

func Fatal(data interface{}) {
	log.Fatal(data)
}