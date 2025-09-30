package repositories

import (
	"github.com/mlogclub/simple/sqls"
	"github.com/mlogclub/simple/web/params"
	"gorm.io/gorm"

	"bbs-go/internal/models"
)

var CRoleRepository = newCRoleRepository()

func newCRoleRepository() *croleRepository {
	return &croleRepository{}
}

type croleRepository struct {
}

func (r *croleRepository) Get(db *gorm.DB, id int64) *models.CRole {
	ret := &models.CRole{}
	if err := db.First(ret, "id = ?", id).Error; err != nil {
		return nil
	}
	return ret
}

func (r *croleRepository) Take(db *gorm.DB, where ...interface{}) *models.CRole {
	ret := &models.CRole{}
	if err := db.Take(ret, where...).Error; err != nil {
		return nil
	}
	return ret
}

func (r *croleRepository) Find(db *gorm.DB, cnd *sqls.Cnd) (list []models.CRole) {
	cnd.Find(db, &list)
	return
}

func (r *croleRepository) FindOne(db *gorm.DB, cnd *sqls.Cnd) *models.CRole {
	ret := &models.CRole{}
	if err := cnd.FindOne(db, &ret); err != nil {
		return nil
	}
	return ret
}

func (r *croleRepository) FindPageByParams(db *gorm.DB, params *params.QueryParams) (list []models.CRole, paging *sqls.Paging) {
	return r.FindPageByCnd(db, &params.Cnd)
}

func (r *croleRepository) FindPageByCnd(db *gorm.DB, cnd *sqls.Cnd) (list []models.CRole, paging *sqls.Paging) {
	cnd.Find(db, &list)
	
	// 计算记录总数
	count := cnd.Count(db, &models.CRole{})
	
	paging = &sqls.Paging{
		Page:  cnd.Paging.Page,
		Limit: cnd.Paging.Limit,
		Total: count,
	}
	return
}

// Count方法保持不变
func (r *croleRepository) Count(db *gorm.DB, cnd *sqls.Cnd) int64 {
	return cnd.Count(db, &models.CRole{})
}

func (r *croleRepository) Create(db *gorm.DB, t *models.CRole) (err error) {
	err = db.Create(t).Error
	return
}

func (r *croleRepository) Update(db *gorm.DB, t *models.CRole) (err error) {
	err = db.Save(t).Error
	return
}

func (r *croleRepository) Updates(db *gorm.DB, id int64, columns map[string]interface{}) (err error) {
	err = db.Model(&models.CRole{}).Where("id = ?", id).Updates(columns).Error
	return
}

func (r *croleRepository) UpdateColumn(db *gorm.DB, id int64, name string, value interface{}) (err error) {
	err = db.Model(&models.CRole{}).Where("id = ?", id).UpdateColumn(name, value).Error
	return
}

func (r *croleRepository) Delete(db *gorm.DB, id int64) {
	db.Delete(&models.CRole{}, "id = ?", id)
}

func (r *croleRepository) GetById(db *gorm.DB, id string) *models.CRole {
	return r.Take(db, "id = ?", id)
}

func (r *croleRepository) GetAll(db *gorm.DB) []models.CRole {
	var list []models.CRole
	db.Find(&list)
	return list
}