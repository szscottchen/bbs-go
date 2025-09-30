package services

import (
	"bbs-go/internal/models"
	"bbs-go/internal/repositories"

	"github.com/mlogclub/simple/sqls"
	"github.com/mlogclub/simple/web/params"
)

var CRoleService = newCRoleService()

func newCRoleService() *croleService {
	return &croleService{}
}

type croleService struct {
}

func (s *croleService) Get(id int64) *models.CRole {
	return repositories.CRoleRepository.Get(sqls.DB(), id)
}

func (s *croleService) Take(where ...interface{}) *models.CRole {
	return repositories.CRoleRepository.Take(sqls.DB(), where...)
}

func (s *croleService) Find(cnd *sqls.Cnd) []models.CRole {
	return repositories.CRoleRepository.Find(sqls.DB(), cnd)
}

func (s *croleService) FindOne(cnd *sqls.Cnd) *models.CRole {
	return repositories.CRoleRepository.FindOne(sqls.DB(), cnd)
}

func (s *croleService) FindPageByParams(params *params.QueryParams) (list []models.CRole, paging *sqls.Paging) {
	return repositories.CRoleRepository.FindPageByParams(sqls.DB(), params)
}

func (s *croleService) FindPageByCnd(cnd *sqls.Cnd) (list []models.CRole, paging *sqls.Paging) {
	return repositories.CRoleRepository.FindPageByCnd(sqls.DB(), cnd)
}

func (s *croleService) Count(cnd *sqls.Cnd) int64 {
	return repositories.CRoleRepository.Count(sqls.DB(), cnd)
}

func (s *croleService) Create(t *models.CRole) error {
	return repositories.CRoleRepository.Create(sqls.DB(), t)
}

func (s *croleService) Update(t *models.CRole) error {
	return repositories.CRoleRepository.Update(sqls.DB(), t)
}

func (s *croleService) Updates(id int64, columns map[string]interface{}) error {
	return repositories.CRoleRepository.Updates(sqls.DB(), id, columns)
}

func (s *croleService) UpdateColumn(id int64, name string, value interface{}) error {
	return repositories.CRoleRepository.UpdateColumn(sqls.DB(), id, name, value)
}

func (s *croleService) Delete(id int64) {
	repositories.CRoleRepository.Delete(sqls.DB(), id)
}

func (s *croleService) GetById(id string) *models.CRole {
	return repositories.CRoleRepository.GetById(sqls.DB(), id)
}

func (s *croleService) GetAll() []models.CRole {
	return repositories.CRoleRepository.GetAll(sqls.DB())
}