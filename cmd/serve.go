package cmd

import (
	"ecommerce/config"
	"log"
	"net/http"
)

func GetMe(w http.ResponseWriter, r *http.Request) {
	log.Println("AMI Data dibo")
}

func Server() {
	cnf := config.GetConfig()

}
