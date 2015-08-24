package utils

import "log"

func CheckErr(err error, msg interface{}) {
	if err != nil {
		log.Println("+++++++++++", msg, err)
	}
}

func Debug(data interface{}) {
	log.Println("+++++++++++", data)
}

func Fatal(data interface{}) {
	log.Fatal(data)
}