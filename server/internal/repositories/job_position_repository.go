package repositories

import (
	"github.com/mlogclub/simple/sqls"
	"github.com/mlogclub/simple/web/params"
	"gorm.io/gorm"

	"bbs-go/internal/models"
)

var JobPositionRepository = newJobPositionRepository()

func newJobPositionRepository() *jobPositionRepository {
	return &jobPositionRepository{}
}

type jobPositionRepository struct {
}

func (r *jobPositionRepository) Get(db *gorm.DB, id int64) *models.JobPosition {
	ret := &models.JobPosition{}
	if err := db.First(ret, "id = ?", id).Error; err != nil {
		return nil
	}
	return ret
}

func (r *jobPositionRepository) Take(db *gorm.DB, where ...interface{}) *models.JobPosition {
	ret := &models.JobPosition{}
	if err := db.Take(ret, where...).Error; err != nil {
		return nil
	}
	return ret
}

func (r *jobPositionRepository) Find(db *gorm.DB, cnd *sqls.Cnd) (list []models.JobPosition) {
	cnd.Find(db, &list)
	return
}

func (r *jobPositionRepository) FindOne(db *gorm.DB, cnd *sqls.Cnd) *models.JobPosition {
	ret := &models.JobPosition{}
	if err := cnd.FindOne(db, &ret); err != nil {
		return nil
	}
	return ret
}

func (r *jobPositionRepository) FindPageByParams(db *gorm.DB, params *params.QueryParams) (list []models.JobPosition, paging *sqls.Paging) {
	return r.FindPageByCnd(db, &params.Cnd)
}

func (r *jobPositionRepository) FindPageByCnd(db *gorm.DB, cnd *sqls.Cnd) (list []models.JobPosition, paging *sqls.Paging) {
	cnd.Find(db, &list)
	
	// 计算记录总数
	count := cnd.Count(db, &models.JobPosition{})
	
	paging = &sqls.Paging{
		Page:  cnd.Paging.Page,
		Limit: cnd.Paging.Limit,
		Total: count,
	}
	return
}

// Count方法保持不变
func (r *jobPositionRepository) Count(db *gorm.DB, cnd *sqls.Cnd) int64 {
	return cnd.Count(db, &models.JobPosition{})
}

func (r *jobPositionRepository) Create(db *gorm.DB, t *models.JobPosition) (err error) {
	err = db.Create(t).Error
	return
}

func (r *jobPositionRepository) Update(db *gorm.DB, t *models.JobPosition) (err error) {
	err = db.Save(t).Error
	return
}

func (r *jobPositionRepository) Updates(db *gorm.DB, id int64, columns map[string]interface{}) (err error) {
	err = db.Model(&models.JobPosition{}).Where("id = ?", id).Updates(columns).Error
	return
}

func (r *jobPositionRepository) UpdateColumn(db *gorm.DB, id int64, name string, value interface{}) (err error) {
	err = db.Model(&models.JobPosition{}).Where("id = ?", id).UpdateColumn(name, value).Error
	return
}

func (r *jobPositionRepository) Delete(db *gorm.DB, id int64) {
	db.Delete(&models.JobPosition{}, "id = ?", id)
}

func (r *jobPositionRepository) GetById(db *gorm.DB, id string) *models.JobPosition {
	return r.Take(db, "id = ?", id)
}

func (r *jobPositionRepository) GetAll(db *gorm.DB) []models.JobPosition {
	var list []models.JobPosition
	db.Find(&list)
	return list
}