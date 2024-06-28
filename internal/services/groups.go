package services

import (
	DB "github.com/Yakov-Varnaev/ft/internal"
	"github.com/Yakov-Varnaev/ft/internal/models"
	"github.com/doug-martin/goqu/v9"
	"github.com/google/uuid"
)

type Groups struct{}

func (s *Groups) GetById(id uuid.UUID) (*models.Group, error) {
	return DB.GetById[models.Group](DB.GROUPS_TABLE, id, nil)
}

func (s *Groups) Create(data *models.WriteGroup) (*models.Group, error) {
	group := models.Group{
		ID:   uuid.New(),
		Name: data.Name,
	}
	db := DB.GetDB()
	_, err := db.Insert(DB.GROUPS_TABLE).Rows(group).Executor().Exec()
	if err != nil {
		return nil, err
	}

	resGroup, err := s.GetById(group.ID)
	if err != nil {
		return nil, err
	}
	if resGroup == nil {
		panic("created group not found")
	}

	return resGroup, nil
}

// TODO: add pagination
func (s *Groups) List() (*[]models.Group, error) {
	groups, err := DB.List[models.Group](DB.GROUPS_TABLE, nil)
	if err != nil {
		return nil, err
	}
	return groups, nil
}

func (s *Groups) Update(id uuid.UUID, group models.WriteGroup) (*models.Group, error) {
	return DB.Update[models.WriteGroup, models.Group](DB.GROUPS_TABLE, id, &group)
}

func (s *Groups) Delete(id uuid.UUID) (uuid.UUID, error) {
	db := DB.GetDB()
	_, err := db.Delete(DB.GROUPS_TABLE).Where(goqu.C("id").Eq(id)).Executor().Exec()
	if err != nil {
		return id, err
	}
	return id, nil
}
