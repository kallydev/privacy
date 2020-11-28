package service

import (
	"github.com/labstack/echo/v4"
)

type QueryResponse struct {
	Names        []string `json:"names"`
	Nicknames    []string `json:"nicknames"`
	PhoneNumbers []string `json:"phone_numbers"`
	IDNumbers    []string `json:"id_numbers"`
	QQNumbers    []string `json:"qq_numbers"`
	Passwords    []string `json:"passwords"`
	Emails       []string `json:"emails"`
	Addresses    []string `json:"addresses"`
}

func NewQueryResponse() *QueryResponse {
	return &QueryResponse{
		Names:        make([]string, 0),
		Nicknames:    make([]string, 0),
		PhoneNumbers: make([]string, 0),
		IDNumbers:    make([]string, 0),
		QQNumbers:    make([]string, 0),
		Passwords:    make([]string, 0),
		Emails:       make([]string, 0),
		Addresses:    make([]string, 0),
	}
}

func (svc *Service) queryHandlerFunc(ctx echo.Context) error {
	value := ctx.QueryParam("value")
	if value == "" {
		return InvalidParameterError
	}
	result := NewQueryResult()
	types := whatType(value)
	for _, t := range types {
		switch value := t.(type) {
		case QQNumber:
			result.addQQNumber(int64(value))
		case PhoneNumber:
			result.addPhoneNumber(int64(value))
		case Email:
			result.addEmail(string(value))
		case IDNumber:
			result.addIDNumber(string(value))
		}
	}
	ok := false
	for !ok {
		for qqNumber, checked := range result.QQNumbers {
			if !checked {
				if err := result.queryQQNumber(svc.databases, int64(qqNumber)); err != nil {
					return err
				}
			}
			continue
		}
		for phoneNumber, checked := range result.PhoneNumbers {
			if !checked {
				if err := result.queryPhoneNumber(svc.databases, int64(phoneNumber)); err != nil {
					return err
				}
			}
			continue
		}
		for idNumber, checked := range result.IDNumbers {
			if !checked {
				if err := result.queryIDNumber(svc.databases, string(idNumber)); err != nil {
					return err
				}
			}
			continue
		}
		for email, checked := range result.Emails {
			if !checked {
				if err := result.queryEmail(svc.databases, string(email)); err != nil {
					return err
				}
			}
			continue
		}
		ok = true
	}
	return NewResponse(ctx, nil, result.Build(svc.config.Mask))
}
