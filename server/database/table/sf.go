package table

import (
	"context"
	"database/sql"
	"github.com/kallydev/privacy/database"
	"github.com/kallydev/privacy/ent"
	"github.com/kallydev/privacy/ent/sfmodel"
)

var (
	_ database.Database = &SFDatabase{}
	_ database.Model    = &SFModel{}
)

type SFDatabase struct {
	Client *ent.Client
}

func (db *SFDatabase) QueryByQQNumber(ctx context.Context, qqNumber int64) ([]database.Model, error) {
	return []database.Model{}, nil
}

func (db *SFDatabase) QueryByEmail(ctx context.Context, email string) ([]database.Model, error) {
	return []database.Model{}, nil
}

func (db *SFDatabase) QueryByIDNumber(ctx context.Context, idNumber string) ([]database.Model, error) {
	return []database.Model{}, nil
}

func (db *SFDatabase) QueryByPhoneNumber(ctx context.Context, phoneNumber int64) ([]database.Model, error) {
	models, err := db.Client.SFModel.
		Query().
		Where(sfmodel.PhoneNumberEQ(phoneNumber)).
		All(ctx)
	if err != nil {
		return nil, err
	}
	return entModelsToSFModels(models), nil
}

type SFModel struct {
	Name        sql.NullString
	PhoneNumber sql.NullInt64
	Address     sql.NullString
}

func (model *SFModel) GetName() (name string, valid bool) {
	return model.Name.String, model.Name.Valid
}

func (model *SFModel) GetNickname() (nickname string, valid bool) {
	return "", false
}

func (model *SFModel) GetPassword() (password string, valid bool) {
	return "", false
}

func (model *SFModel) GetEmail() (email string, valid bool) {
	return "", false
}

func (model *SFModel) GetQQNumber() (qqNumber int64, valid bool) {
	return 0, false
}

func (model *SFModel) GetIDNumber() (idNumber string, valid bool) {
	return "", false
}

func (model *SFModel) GetPhoneNumber() (phoneNumber int64, valid bool) {
	return model.PhoneNumber.Int64, model.PhoneNumber.Valid
}

func (model *SFModel) GetAddress() (address string, valid bool) {
	return model.Address.String, model.Address.Valid
}

func entModelsToSFModels(endModels []*ent.SFModel) []database.Model {
	models := make([]database.Model, len(endModels))
	for i, model := range endModels {
		models[i] = &SFModel{
			Name: sql.NullString{
				String: model.Name,
				Valid:  model.Name != "",
			},
			PhoneNumber: sql.NullInt64{
				Int64: model.PhoneNumber,
				Valid: model.PhoneNumber != 0,
			},
			Address: sql.NullString{
				String: model.Address,
				Valid:  model.Name != "",
			},
		}
	}
	return models
}
