package table

import (
	"context"
	"database/sql"
	"github.com/kallydev/privacy/database"
	"github.com/kallydev/privacy/ent"
	"github.com/kallydev/privacy/ent/qqmodel"
)

var (
	_ database.Database = &QQDatabase{}
	_ database.Model    = &QQModel{}
)

type QQDatabase struct {
	Client *ent.Client
}

func (db *QQDatabase) QueryByQQNumber(ctx context.Context, qqNumber int64) ([]database.Model, error) {
	models, err := db.Client.QQModel.
		Query().
		Where(qqmodel.QqNumberEQ(qqNumber)).
		All(ctx)
	if err != nil {
		return nil, err
	}
	return entModelsToQQModels(models), nil
}

func (db *QQDatabase) QueryByEmail(ctx context.Context, email string) ([]database.Model, error) {
	return []database.Model{}, nil
}

func (db *QQDatabase) QueryByIDNumber(ctx context.Context, idNumber string) ([]database.Model, error) {
	return []database.Model{}, nil
}

func (db *QQDatabase) QueryByPhoneNumber(ctx context.Context, phoneNumber int64) ([]database.Model, error) {
	models, err := db.Client.QQModel.
		Query().
		Where(qqmodel.PhoneNumberEQ(phoneNumber)).
		All(ctx)
	if err != nil {
		return nil, err
	}
	return entModelsToQQModels(models), nil
}

type QQModel struct {
	QQNumber    sql.NullInt64
	PhoneNumber sql.NullInt64
}

func (model *QQModel) GetName() (name string, valid bool) {
	return "", false
}

func (model *QQModel) GetNickname() (nickname string, valid bool) {
	return "", false
}

func (model *QQModel) GetPassword() (password string, valid bool) {
	return "", false
}

func (model *QQModel) GetEmail() (email string, valid bool) {
	return "", false
}

func (model *QQModel) GetQQNumber() (qqNumber int64, valid bool) {
	return model.QQNumber.Int64, model.QQNumber.Valid
}

func (model *QQModel) GetIDNumber() (idNumber string, valid bool) {
	return "", false
}

func (model *QQModel) GetPhoneNumber() (phoneNumber int64, valid bool) {
	return model.PhoneNumber.Int64, model.PhoneNumber.Valid
}

func (model *QQModel) GetAddress() (address string, valid bool) {
	return "", false
}

func entModelsToQQModels(endModels []*ent.QQModel) []database.Model {
	models := make([]database.Model, len(endModels))
	for i, model := range endModels {
		models[i] = &QQModel{
			QQNumber: sql.NullInt64{
				Int64: model.QqNumber,
				Valid: model.QqNumber != 0,
			},
			PhoneNumber: sql.NullInt64{
				Int64: model.PhoneNumber,
				Valid: model.PhoneNumber != 0,
			},
		}
	}
	return models
}
