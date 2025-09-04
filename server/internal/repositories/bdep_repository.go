package repositories

import (
	"github.com/mlogclub/simple/sqls"
	"github.com/mlogclub/simple/web/params"
	"gorm.io/gorm"

	"bbs-go/internal/models"
)

var BDepRepository = newBDepRepository()

func newBDepRepository() *bdepRepository {
	return &bdepRepository{}
}

type bdepRepository struct {
}

func (r *bdepRepository) Get(db *gorm.DB, id int64) *models.BDep {
	ret := &models.BDep{}
	if err := db.First(ret, "id = ?", id).Error; err != nil {
		return nil
	}
	return ret
}

func (r *bdepRepository) Take(db *gorm.DB, where ...interface{}) *models.BDep {
	ret := &models.BDep{}
	if err := db.Take(ret, where...).Error; err != nil {
		return nil
	}
	return ret
}

func (r *bdepRepository) Find(db *gorm.DB, cnd *sqls.Cnd) (list []models.BDep) {
	cnd.Find(db, &list)
	return
}

func (r *bdepRepository) FindOne(db *gorm.DB, cnd *sqls.Cnd) *models.BDep {
	ret := &models.BDep{}
	if err := cnd.FindOne(db, &ret); err != nil {
		return nil
	}
	return ret
}

func (r *bdepRepository) FindPageByParams(db *gorm.DB, params *params.QueryParams) (list []models.BDep, paging *sqls.Paging) {
	return r.FindPageByCnd(db, &params.Cnd)
}

// 正确的方法定义 - 将Count方法移到FindPageByCnd方法外部
func (r *bdepRepository) FindPageByCnd(db *gorm.DB, cnd *sqls.Cnd) (list []models.BDep, paging *sqls.Paging) {
	cnd.Find(db, &list)
	count := cnd.Count(db, &models.BDep{})

	paging = &sqls.Paging{
		Page:  cnd.Paging.Page,
		Limit: cnd.Paging.Limit,
		Total: count,
	}
	return
}

// Count方法作为独立方法定义
func (r *bdepRepository) Count(db *gorm.DB, cnd *sqls.Cnd) int64 {
	return cnd.Count(db, &models.BDep{})
}

func (r *bdepRepository) Create(db *gorm.DB, t *models.BDep) (err error) {
	err = db.Create(t).Error
	return
}

func (r *bdepRepository) Update(db *gorm.DB, t *models.BDep) (err error) {
	err = db.Save(t).Error
	return
}

func (r *bdepRepository) Updates(db *gorm.DB, id int64, columns map[string]interface{}) (err error) {
	err = db.Model(&models.BDep{}).Where("id = ?", id).Updates(columns).Error
	return
}

func (r *bdepRepository) UpdateColumn(db *gorm.DB, id int64, name string, value interface{}) (err error) {
	err = db.Model(&models.BDep{}).Where("id = ?", id).UpdateColumn(name, value).Error
	return
}

func (r *bdepRepository) Delete(db *gorm.DB, id int64) {
	db.Delete(&models.BDep{}, "id = ?", id)
}

func (r *bdepRepository) GetById(db *gorm.DB, id string) *models.BDep {
	return r.Take(db, "id = ?", id)
}

func (r *bdepRepository) GetAll(db *gorm.DB) []models.BDep {
	var list []models.BDep
	db.Find(&list)
	return list
}