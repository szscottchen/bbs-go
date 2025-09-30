package repositories

import (
	"github.com/mlogclub/simple/sqls"
	"github.com/mlogclub/simple/web/params"
	"gorm.io/gorm"

	"bbs-go/internal/models"
)

var BDivRepository = newBDivRepository()

func newBDivRepository() *bdivRepository {
	return &bdivRepository{}
}

type bdivRepository struct {
}

func (r *bdivRepository) Get(db *gorm.DB, id int64) *models.BDiv {
	ret := &models.BDiv{}
	if err := db.First(ret, "id = ?", id).Error; err != nil {
		return nil
	}
	return ret
}

func (r *bdivRepository) Take(db *gorm.DB, where ...interface{}) *models.BDiv {
	ret := &models.BDiv{}
	if err := db.Take(ret, where...).Error; err != nil {
		return nil
	}
	return ret
}

func (r *bdivRepository) Find(db *gorm.DB, cnd *sqls.Cnd) (list []models.BDiv) {
	cnd.Find(db, &list)
	return
}

func (r *bdivRepository) FindOne(db *gorm.DB, cnd *sqls.Cnd) *models.BDiv {
	ret := &models.BDiv{}
	if err := cnd.FindOne(db, &ret); err != nil {
		return nil
	}
	return ret
}

func (r *bdivRepository) FindPageByParams(db *gorm.DB, params *params.QueryParams) (list []models.BDiv, paging *sqls.Paging) {
	return r.FindPageByCnd(db, &params.Cnd)
}

func (r *bdivRepository) FindPageByCnd(db *gorm.DB, cnd *sqls.Cnd) (list []models.BDiv, paging *sqls.Paging) {
	cnd.Find(db, &list)
	count := cnd.Count(db, &models.BDiv{})

	paging = &sqls.Paging{
		Page:  cnd.Paging.Page,
		Limit: cnd.Paging.Limit,
		Total: count,
	}
	return
}

func (r *bdivRepository) Create(db *gorm.DB, t *models.BDiv) (err error) {
	err = db.Create(t).Error
	return
}

func (r *bdivRepository) Update(db *gorm.DB, t *models.BDiv) (err error) {
	err = db.Save(t).Error
	return
}

func (r *bdivRepository) Updates(db *gorm.DB, id int64, columns map[string]interface{}) (err error) {
	err = db.Model(&models.BDiv{}).Where("id = ?", id).Updates(columns).Error
	return
}

func (r *bdivRepository) UpdateColumn(db *gorm.DB, id int64, name string, value interface{}) (err error) {
	err = db.Model(&models.BDiv{}).Where("id = ?", id).UpdateColumn(name, value).Error
	return
}

func (r *bdivRepository) Delete(db *gorm.DB, id int64) {
	db.Delete(&models.BDiv{}, "id = ?", id)
}

func (r *bdivRepository) GetById(db *gorm.DB, id string) *models.BDiv {
	return r.Take(db, "id = ?", id)
}

func (r *bdivRepository) GetAll(db *gorm.DB) []models.BDiv {
	var list []models.BDiv
	db.Find(&list)
	return list
}

// 在FindPageByCnd方法后添加Count方法
func (r *bdivRepository) Count(db *gorm.DB, cnd *sqls.Cnd) int64 {
	return cnd.Count(db, &models.BDiv{})
}