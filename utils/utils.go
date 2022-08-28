package utils

import (
	"encoding/json"
	"net/http"
	"regexp"
	"time"
)

var myClient = &http.Client{Timeout: 10 * time.Second}

const Url = "https://run.mocky.io/v3/d02168c6-d88d-4ff2-aac6-9e9eb3425e31"

func Authorization(url string, target interface{}) bool {
	r, err := myClient.Get(url)
	if err != nil {
		return false
	}
	defer r.Body.Close()

	err = json.NewDecoder(r.Body).Decode(&target)
	if err != nil {
		return false
	}
	authorizeResponse := make(map[string]any)
	authorizeResponse = target.(map[string]any)
	if authorizeResponse["authorization"] == false {
		return false
	}
	return true
}

func ValidateCPF(cpf string) bool {
	const pattern = `^\d{3}\.?\d{3}\.?\d{3}-?\d{2}$`
	matched, _ := regexp.MatchString(pattern, cpf)
	return matched
}

func ValidateCNPJ(cnpj string) bool {
	const pattern = `^^\d{2}\.?\d{3}\.?\d{3}\/?(:?\d{3}[1-9]|\d{2}[1-9]\d|\d[1-9]\d{2}|[1-9]\d{3})-?\d{2}$$`
	matched, _ := regexp.MatchString(pattern, cnpj)
	return matched
}

func ValidateEmail(email string) bool {
	const pattern = `^[a-zA-Z0-9.!#$%&'*+/=?^_` + "`" + `{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$`
	matched, _ := regexp.MatchString(pattern, email)
	return matched
}
