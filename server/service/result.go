package service

import (
	"context"
	"github.com/kallydev/privacy/database"
)

type Result interface {
	queryPhoneNumber(databases []database.Database, phoneNumber int64) error
	queryIDNumber(databases []database.Database, idNumber string) error
	queryQQNumber(databases []database.Database, qqNumber int64) error
	queryEmail(databases []database.Database, email string) error

	addName(name string)
	addNickname(nickname string)
	addPhoneNumber(phoneNumber int64)
	addIDNumber(idNumber string)
	addQQNumber(qqNumber int64)
	addPassword(password string)
	addEmail(email string)
	addAddress(address string)
}

var (
	_ Result = &QueryResult{}
)

type QueryResult struct {
	Names        map[Name]bool        `json:"names"`
	Nicknames    map[Nickname]bool    `json:"nicknames"`
	PhoneNumbers map[PhoneNumber]bool `json:"phone_numbers"`
	IDNumbers    map[IDNumber]bool    `json:"id_numbers"`
	QQNumbers    map[QQNumber]bool    `json:"qq_numbers"`
	Passwords    map[Password]bool    `json:"passwords"`
	Emails       map[Email]bool       `json:"emails"`
	Addresses    map[Address]bool     `json:"addresses"`
}

func NewQueryResult() *QueryResult {
	return &QueryResult{
		Names:        make(map[Name]bool),
		Nicknames:    make(map[Nickname]bool),
		PhoneNumbers: make(map[PhoneNumber]bool),
		IDNumbers:    make(map[IDNumber]bool),
		QQNumbers:    make(map[QQNumber]bool),
		Passwords:    make(map[Password]bool),
		Emails:       make(map[Email]bool),
		Addresses:    make(map[Address]bool),
	}
}

func (result *QueryResult) addModel(model database.Model) {
	if name, valid := model.GetName(); valid {
		result.addName(name)
	}
	if nickname, valid := model.GetNickname(); valid {
		result.addNickname(nickname)
	}
	if phoneNumber, valid := model.GetPhoneNumber(); valid {
		result.addPhoneNumber(phoneNumber)
	}
	if idNumber, valid := model.GetIDNumber(); valid {
		result.addIDNumber(idNumber)
	}
	if qqNumber, valid := model.GetQQNumber(); valid {
		result.addQQNumber(qqNumber)
	}
	if password, valid := model.GetPassword(); valid {
		result.addPassword(password)
	}
	if email, valid := model.GetEmail(); valid {
		result.addEmail(email)
	}
	if address, valid := model.GetAddress(); valid {
		result.addAddress(address)
	}
}

func (result *QueryResult) queryPhoneNumber(databases []database.Database, phoneNumber int64) error {
	for _, db := range databases {
		models, err := db.QueryByPhoneNumber(context.Background(), phoneNumber)
		if err != nil {
			return err
		}
		for _, model := range models {
			result.addModel(model)
		}
	}
	result.PhoneNumbers[PhoneNumber(phoneNumber)] = true
	return nil
}

func (result *QueryResult) queryIDNumber(databases []database.Database, idNumber string) error {
	for _, db := range databases {
		models, err := db.QueryByIDNumber(context.Background(), idNumber)
		if err != nil {
			return err
		}
		for _, model := range models {
			result.addModel(model)
		}
	}
	result.IDNumbers[IDNumber(idNumber)] = true
	return nil
}

func (result *QueryResult) queryQQNumber(databases []database.Database, qqNumber int64) error {
	for _, db := range databases {
		models, err := db.QueryByQQNumber(context.Background(), qqNumber)
		if err != nil {
			return err
		}
		for _, model := range models {
			result.addModel(model)
		}
	}
	result.QQNumbers[QQNumber(qqNumber)] = true
	return nil
}

func (result *QueryResult) queryEmail(databases []database.Database, email string) error {
	for _, db := range databases {
		models, err := db.QueryByEmail(context.Background(), email)
		if err != nil {
			return err
		}
		for _, model := range models {
			result.addModel(model)
		}
	}
	result.Emails[Email(email)] = true
	return nil
}

func (result *QueryResult) addName(value string) {
	if value == "" {
		return
	}
	name := Name(value)
	if _, ok := result.Names[name]; !ok {
		result.Names[name] = false
	}
}

func (result *QueryResult) addNickname(value string) {
	if value == "" {
		return
	}
	nickname := Nickname(value)
	if _, ok := result.Nicknames[nickname]; !ok {
		result.Nicknames[nickname] = false
	}
}

func (result *QueryResult) addPhoneNumber(value int64) {
	if value == 0 {
		return
	}
	phoneNumber := PhoneNumber(value)
	if _, ok := result.PhoneNumbers[phoneNumber]; !ok {
		result.PhoneNumbers[phoneNumber] = false
	}
}

func (result *QueryResult) addIDNumber(value string) {
	if value == "" {
		return
	}
	idNumber := IDNumber(value)
	if _, ok := result.IDNumbers[idNumber]; !ok {
		result.IDNumbers[idNumber] = false
	}
}

func (result *QueryResult) addQQNumber(value int64) {
	if value == 0 {
		return
	}
	qqNumber := QQNumber(value)
	if _, ok := result.QQNumbers[qqNumber]; !ok {
		result.QQNumbers[qqNumber] = false
	}
}

func (result *QueryResult) addPassword(value string) {
	if value == "" {
		return
	}
	password := Password(value)
	if _, ok := result.Passwords[password]; !ok {
		result.Passwords[password] = false
	}
}

func (result *QueryResult) addEmail(value string) {
	if value == "" {
		return
	}
	email := Email(value)
	if _, ok := result.Emails[email]; !ok {
		result.Emails[email] = false
	}
}

func (result *QueryResult) addAddress(value string) {
	if value == "" {
		return
	}
	address := Address(value)
	if _, ok := result.Addresses[address]; !ok {
		result.Addresses[address] = false
	}
}

func (result *QueryResult) Build(mask bool) *QueryResponse {
	response := NewQueryResponse()
	for name := range result.Names {
		result := name.String()
		if mask {
			result = name.Masking()
		}
		response.Names = append(response.Names, result)
	}
	for nickname := range result.Nicknames {
		result := nickname.String()
		if mask {
			result = nickname.Masking()
		}
		response.Nicknames = append(response.Nicknames, result)
	}
	for phoneNumber := range result.PhoneNumbers {
		result := phoneNumber.String()
		if mask {
			result = phoneNumber.Masking()
		}
		response.PhoneNumbers = append(response.PhoneNumbers, result)
	}
	for idNumber := range result.IDNumbers {
		result := idNumber.String()
		if mask {
			result = idNumber.Masking()
		}
		response.IDNumbers = append(response.IDNumbers, result)
	}
	for qqNumber := range result.QQNumbers {
		result := qqNumber.String()
		if mask {
			result = qqNumber.Masking()
		}
		response.QQNumbers = append(response.QQNumbers, result)
	}
	for password := range result.Passwords {
		result := password.String()
		if mask {
			result = password.Masking()
		}
		response.Passwords = append(response.Passwords, result)
	}
	for email := range result.Emails {
		result := email.String()
		if mask {
			result = email.Masking()
		}
		response.Emails = append(response.Emails, result)
	}
	for address := range result.Addresses {
		result := address.String()
		if mask {
			result = address.Masking()
		}
		response.Addresses = append(response.Addresses, result)
	}
	return response
}
