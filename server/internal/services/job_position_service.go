package services

import (
	"bbs-go/internal/models"
	"bbs-go/internal/repositories"

	"github.com/mlogclub/simple/sqls"
	"github.com/mlogclub/simple/web/params"
)

var JobPositionService = newJobPositionService()

func newJobPositionService() *jobPositionService {
	return &jobPositionService{}
}

type jobPositionService struct {
}

func (s *jobPositionService) Get(id int64) *models.JobPosition {
	return repositories.JobPositionRepository.Get(sqls.DB(), id)
}

func (s *jobPositionService) Take(where ...interface{}) *models.JobPosition {
	return repositories.JobPositionRepository.Take(sqls.DB(), where...)
}

func (s *jobPositionService) Find(cnd *sqls.Cnd) []models.JobPosition {
	return repositories.JobPositionRepository.Find(sqls.DB(), cnd)
}

func (s *jobPositionService) FindOne(cnd *sqls.Cnd) *models.JobPosition {
	return repositories.JobPositionRepository.FindOne(sqls.DB(), cnd)
}

func (s *jobPositionService) FindPageByParams(params *params.QueryParams) (list []models.JobPosition, paging *sqls.Paging) {
	return repositories.JobPositionRepository.FindPageByParams(sqls.DB(), params)
}

func (s *jobPositionService) FindPageByCnd(cnd *sqls.Cnd) (list []models.JobPosition, paging *sqls.Paging) {
	return repositories.JobPositionRepository.FindPageByCnd(sqls.DB(), cnd)
}

func (s *jobPositionService) Count(cnd *sqls.Cnd) int64 {
	return repositories.JobPositionRepository.Count(sqls.DB(), cnd)
}

func (s *jobPositionService) Create(t *models.JobPosition) error {
	return repositories.JobPositionRepository.Create(sqls.DB(), t)
}

func (s *jobPositionService) Update(t *models.JobPosition) error {
	return repositories.JobPositionRepository.Update(sqls.DB(), t)
}

func (s *jobPositionService) Updates(id int64, columns map[string]interface{}) error {
	return repositories.JobPositionRepository.Updates(sqls.DB(), id, columns)
}

func (s *jobPositionService) UpdateColumn(id int64, name string, value interface{}) error {
	return repositories.JobPositionRepository.UpdateColumn(sqls.DB(), id, name, value)
}

func (s *jobPositionService) Delete(id int64) {
	repositories.JobPositionRepository.Delete(sqls.DB(), id)
}

func (s *jobPositionService) GetById(id string) *models.JobPosition {
	return repositories.JobPositionRepository.GetById(sqls.DB(), id)
}

func (s *jobPositionService) GetAll() []models.JobPosition {
	return repositories.JobPositionRepository.GetAll(sqls.DB())
}