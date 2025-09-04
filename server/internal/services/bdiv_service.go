package services

import (
	"bbs-go/internal/models"
	"bbs-go/internal/repositories"

	"github.com/mlogclub/simple/sqls"
	"github.com/mlogclub/simple/web/params"
)

var BDivService = newBDivService()

func newBDivService() *bdivService {
	return &bdivService{}
}

type bdivService struct {
}

func (s *bdivService) Get(id int64) *models.BDiv {
	return repositories.BDivRepository.Get(sqls.DB(), id)
}

func (s *bdivService) Take(where ...interface{}) *models.BDiv {
	return repositories.BDivRepository.Take(sqls.DB(), where...)
}

func (s *bdivService) Find(cnd *sqls.Cnd) []models.BDiv {
	return repositories.BDivRepository.Find(sqls.DB(), cnd)
}

func (s *bdivService) FindOne(cnd *sqls.Cnd) *models.BDiv {
	return repositories.BDivRepository.FindOne(sqls.DB(), cnd)
}

func (s *bdivService) FindPageByParams(params *params.QueryParams) (list []models.BDiv, paging *sqls.Paging) {
	return repositories.BDivRepository.FindPageByParams(sqls.DB(), params)
}

func (s *bdivService) FindPageByCnd(cnd *sqls.Cnd) (list []models.BDiv, paging *sqls.Paging) {
	return repositories.BDivRepository.FindPageByCnd(sqls.DB(), cnd)
}

func (s *bdivService) Count(cnd *sqls.Cnd) int64 {
	return repositories.BDivRepository.Count(sqls.DB(), cnd)
}

func (s *bdivService) Create(t *models.BDiv) error {
	return repositories.BDivRepository.Create(sqls.DB(), t)
}

func (s *bdivService) Update(t *models.BDiv) error {
	return repositories.BDivRepository.Update(sqls.DB(), t)
}

func (s *bdivService) Updates(id int64, columns map[string]interface{}) error {
	return repositories.BDivRepository.Updates(sqls.DB(), id, columns)
}

func (s *bdivService) UpdateColumn(id int64, name string, value interface{}) error {
	return repositories.BDivRepository.UpdateColumn(sqls.DB(), id, name, value)
}

func (s *bdivService) Delete(id int64) {
	repositories.BDivRepository.Delete(sqls.DB(), id)
}

func (s *bdivService) GetById(id string) *models.BDiv {
	return repositories.BDivRepository.GetById(sqls.DB(), id)
}

func (s *bdivService) GetAll() []models.BDiv {
	return repositories.BDivRepository.GetAll(sqls.DB())
}