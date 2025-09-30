package services

import (
	"bbs-go/internal/models"
	"bbs-go/internal/repositories"

	"github.com/mlogclub/simple/sqls"
	"github.com/mlogclub/simple/web/params"
)

var BDepService = newBDepService()

func newBDepService() *bdepService {
	return &bdepService{}
}

type bdepService struct {
}

func (s *bdepService) Get(id int64) *models.BDep {
	return repositories.BDepRepository.Get(sqls.DB(), id)
}

func (s *bdepService) Take(where ...interface{}) *models.BDep {
	return repositories.BDepRepository.Take(sqls.DB(), where...)
}

func (s *bdepService) Find(cnd *sqls.Cnd) []models.BDep {
	return repositories.BDepRepository.Find(sqls.DB(), cnd)
}

func (s *bdepService) FindOne(cnd *sqls.Cnd) *models.BDep {
	return repositories.BDepRepository.FindOne(sqls.DB(), cnd)
}

func (s *bdepService) FindPageByParams(params *params.QueryParams) (list []models.BDep, paging *sqls.Paging) {
	return repositories.BDepRepository.FindPageByParams(sqls.DB(), params)
}

func (s *bdepService) FindPageByCnd(cnd *sqls.Cnd) (list []models.BDep, paging *sqls.Paging) {
	return repositories.BDepRepository.FindPageByCnd(sqls.DB(), cnd)
}

func (s *bdepService) Count(cnd *sqls.Cnd) int64 {
	return repositories.BDepRepository.Count(sqls.DB(), cnd)
}

func (s *bdepService) Create(t *models.BDep) error {
	return repositories.BDepRepository.Create(sqls.DB(), t)
}

func (s *bdepService) Update(t *models.BDep) error {
	return repositories.BDepRepository.Update(sqls.DB(), t)
}

func (s *bdepService) Updates(id int64, columns map[string]interface{}) error {
	return repositories.BDepRepository.Updates(sqls.DB(), id, columns)
}

func (s *bdepService) UpdateColumn(id int64, name string, value interface{}) error {
	return repositories.BDepRepository.UpdateColumn(sqls.DB(), id, name, value)
}

func (s *bdepService) Delete(id int64) {
	repositories.BDepRepository.Delete(sqls.DB(), id)
}

func (s *bdepService) GetById(id string) *models.BDep {
	return repositories.BDepRepository.GetById(sqls.DB(), id)
}

func (s *bdepService) GetAll() []models.BDep {
	return repositories.BDepRepository.GetAll(sqls.DB())
}