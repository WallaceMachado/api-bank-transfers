package utils

// validação feita com base no site da receita federal http://www.receita.fazenda.gov.br/aplicacoes/atcta/cpf/funcoes.js
import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

func Format(s string) string {
	s = strings.TrimSpace(s)
	s = strings.Replace(s, ".", "", -1)
	s = strings.Replace(s, "-", "", -1)
	s = strings.Replace(s, "/", "", -1)
	return s
}

func ValidateCPF(cpfIn string) (string, error) {
	cpf := Format(cpfIn)
	if len(cpf) != 11 || cpf == "00000000000" ||
		cpf == "11111111111" ||
		cpf == "22222222222" ||
		cpf == "33333333333" ||
		cpf == "44444444444" ||
		cpf == "55555555555" ||
		cpf == "66666666666" ||
		cpf == "77777777777" ||
		cpf == "88888888888" ||
		cpf == "99999999999" {
		return "", errors.New("invalid CPF")
	}

	cpfFirstDigitTable := []int{10, 9, 8, 7, 6, 5, 4, 3, 2}
	cpfSecondDigitTable := []int{11, 10, 9, 8, 7, 6, 5, 4, 3, 2}

	firstPart := cpf[0:9]
	sum, err := sumDigit(firstPart, cpfFirstDigitTable)
	if err != nil {
		return "", errors.New("invalid CPF")
	}
	r1 := sum % 11
	d1 := 0

	if r1 >= 2 {
		d1 = 11 - r1
	}

	secondPart := firstPart + strconv.Itoa(d1)

	dsum, err := sumDigit(secondPart, cpfSecondDigitTable)
	if err != nil {
		return "", errors.New("invalid CPF")
	}

	r2 := dsum % 11
	d2 := 0

	if r2 >= 2 {
		d2 = 11 - r2
	}

	finalPart := fmt.Sprintf("%s%d%d", firstPart, d1, d2)

	if finalPart != cpf {
		return "", errors.New("invalid CPF")
	}

	return cpf, nil
}

func sumDigit(s string, table []int) (int, error) {

	if len(s) != len(table) {
		return 0, errors.New("invalid CPF")
	}

	sum := 0

	for i, v := range table {
		c := string(s[i])
		d, err := strconv.Atoi(c)
		if err != nil {
			return 0, err
		}
		sum += v * d
	}

	return sum, nil
}
