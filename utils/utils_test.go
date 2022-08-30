package utils

import (
	"fmt"
	"testing"
)

// Test ValidateCPF
func TestValidateCPF_return_true(t *testing.T) {
	cpf := "02651132291"
	if !ValidateCPF(cpf) {
		t.Error("CPF is valid")
	}
}

func TestInValidateCPF_return_false(t *testing.T) {
	cpf := "026511322913"
	if ValidateCPF(cpf) {
		t.Error("CPF is invalid")
	}
}

func TestValidateCNPJ(t *testing.T) {
	cnpj := "97.560.303/0001-09"
	if !ValidateCNPJ(cnpj) {
		t.Error("CNPJ is invalid")
	}
}

func TestInvalidCNPJ(t *testing.T) {
	cnpj := "97.560.303/0001-091"
	if ValidateCNPJ(cnpj) {
		t.Error("CNPJ is valid")
	}
}

func TestValidateEmail(t *testing.T) {
	email := "junior.moura10@hotmail.com"
	if !ValidateEmail(email) {
		t.Error("Email is invalid")
	}
}

func TestInvalidEmail(t *testing.T) {
	email := "junior.moura10hotmail.com"
	if ValidateEmail(email) {
		t.Error("Email is valid")
	}
}

// Test UnMaskCPFCNPJ
func TestUnMaskCPFCNPJ_return_cpf_masked(t *testing.T) {
	cpf := "026.511.322-91"
	if UnMaskCPFCNPJ(cpf) == "02651132291" {
		fmt.Println("CPF is masked")
		return
	} else {
		t.Error("CPF failed to be unmasked")
		return
	}
}

// Test Authorization request return true
func TestAuthorization_return_true(t *testing.T) {
	auth, _ := Authorization(Url, &ResponseAuth{})
	if auth {
		fmt.Println("Authorization is valid")
		return
	} else {
		t.Error("Authorization is invalid")
	}
}
func TestAuthorization_return_false(t *testing.T) {
	auth, _ := Authorization(Url+"a", &ResponseAuth{})
	if !auth {
		fmt.Println("Authorization is invalid")
		return
	} else {
		t.Error("Authorization is valid")
	}
}
