package service

import (
	"strconv"
	"strings"
)

type Data interface {
	String() string
	Masking() string
}

var (
	_ Data = Address("")
	_ Data = Nickname("")
	_ Data = PhoneNumber(0)
	_ Data = IDNumber("")
	_ Data = QQNumber(0)
	_ Data = Password("")
	_ Data = Email("")
	_ Data = Address("")
)

type (
	Name        string
	Nickname    string
	PhoneNumber int64
	IDNumber    string
	QQNumber    int64
	Password    string
	Email       string
	Address     string
)

func (address Address) Masking() string {
	return maskLeft(string(address), 1)
}

func (address Address) String() string {
	return string(address)
}

func (email Email) Masking() string {
	return maskLeft(email.String(), len(strings.Split(string(email), "@")[0])+1)
}

func (email Email) String() string {
	return string(email)
}

func (password Password) Masking() string {
	return maskLeft(password.String(), 4)
}

func (password Password) String() string {
	return string(password)
}

func (qqNumber QQNumber) Masking() string {
	return maskLeft(qqNumber.String(), 3)
}

func (qqNumber QQNumber) String() string {
	return strconv.Itoa(int(qqNumber))
}

func (idNumber IDNumber) Masking() string {
	return mask(maskLeft(idNumber.String(), 3), 16, 16)
}

func (idNumber IDNumber) String() string {
	return string(idNumber)
}

func (phoneNumber PhoneNumber) Masking() string {
	return mask(phoneNumber.String(), 3, 6)
}

func (phoneNumber PhoneNumber) String() string {
	return strconv.Itoa(int(phoneNumber))
}

func (nickname Nickname) Masking() string {
	return maskLeft(nickname.String(), 1)
}

func (nickname Nickname) String() string {
	return string(nickname)
}

func (name Name) Masking() string {
	return maskLeft(name.String(), 1)
}

func (name Name) String() string {
	return string(name)
}

func isPhoneNumber(number int) bool {
	numberStr := strconv.Itoa(number)
	if len(numberStr) != 11 {
		return false
	}
	// https://www.miit.gov.cn/n1146295/n1652858/n1652930/n3757020/c5623267/part/5623278.doc
	s := []int{
		130, 131, 132, 133, 134, 135, 136, 137, 138, 139,
		145, 146, 147, 148, 148,
		150, 151, 152, 153, 154, 155, 156, 157, 158, 159,
		161, 162, 164, 165, 167,
		170, 172, 173, 175, 176, 177, 178, 179,
		181, 182, 183, 184, 185, 186, 187, 188, 189,
		190, 191, 192, 193, 194, 195, 196, 197, 198, 199,
	}
	header, _ := strconv.Atoi(numberStr[:3])
	for _, num := range s {
		if num == header {
			return true
		}
	}
	return false
}

func whatType(value string) (result []interface{}) {
	if number, err := strconv.ParseInt(value, 10, 64); err == nil {
		if number <= 19999999999 && isPhoneNumber(int(number)) {
			result = append(result, PhoneNumber(number))
		}
		if number <= 99999999999 && len(strconv.Itoa(int(number))) >= 5 {
			result = append(result, QQNumber(number))
		}
	}
	if len(value) == 18 {
		result = append(result, IDNumber(value))
	}
	if strings.Contains(value, "@") {
		result = append(result, Email(value))
	}
	return result
}
