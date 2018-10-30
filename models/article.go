package models

import (
	"time"
	"github.com/jinzhu/gorm"
	"fmt"
	"errors"
)

//go:generate goqueryset -in article.go

// Article struct represent article model. Next line (gen:qs) is needed to autogenerate ArticleQuerySet.
// gen:qs
type Article struct {
	ID        int64     `json:"id"`
	Title     string    `json:"title" validate:"required"`
	Content   string    `json:"content" validate:"required"`
	Thumbnail string    `json:"thumbnail" validate:"required"`
	UpdatedAt time.Time `json:"updated_at"`
	CreatedAt time.Time `json:"created_at"`
}

// ===== BEGIN of all query sets

// ===== BEGIN of query set ArticleQuerySet

// ArticleQuerySet is an queryset type for Article
type ArticleQuerySet struct {
	db *gorm.DB
}

// NewArticleQuerySet constructs new ArticleQuerySet
func NewArticleQuerySet(db *gorm.DB) ArticleQuerySet {
	return ArticleQuerySet{
		db: db.Model(&Article{}),
	}
}

func (qs ArticleQuerySet) w(db *gorm.DB) ArticleQuerySet {
	return NewArticleQuerySet(db)
}

// All is an autogenerated method
// nolint: dupl
func (qs ArticleQuerySet) All(ret *[]Article) error {
	return qs.db.Find(ret).Error
}

// ContentEq is an autogenerated method
// nolint: dupl
func (qs ArticleQuerySet) ContentEq(content string) ArticleQuerySet {
	return qs.w(qs.db.Where("content = ?", content))
}

// ContentIn is an autogenerated method
// nolint: dupl
func (qs ArticleQuerySet) ContentIn(content ...string) ArticleQuerySet {
	if len(content) == 0 {
		qs.db.AddError(errors.New("must at least pass one content in ContentIn"))
		return qs.w(qs.db)
	}
	return qs.w(qs.db.Where("content IN (?)", content))
}

// ContentNe is an autogenerated method
// nolint: dupl
func (qs ArticleQuerySet) ContentNe(content string) ArticleQuerySet {
	return qs.w(qs.db.Where("content != ?", content))
}

// ContentNotIn is an autogenerated method
// nolint: dupl
func (qs ArticleQuerySet) ContentNotIn(content ...string) ArticleQuerySet {
	if len(content) == 0 {
		qs.db.AddError(errors.New("must at least pass one content in ContentNotIn"))
		return qs.w(qs.db)
	}
	return qs.w(qs.db.Where("content NOT IN (?)", content))
}

// Count is an autogenerated method
// nolint: dupl
func (qs ArticleQuerySet) Count() (int, error) {
	var count int
	err := qs.db.Count(&count).Error
	return count, err
}

// Create is an autogenerated method
// nolint: dupl
func (o *Article) Create(db *gorm.DB) error {
	return db.Create(o).Error
}

// CreatedAtEq is an autogenerated method
// nolint: dupl
func (qs ArticleQuerySet) CreatedAtEq(createdAt time.Time) ArticleQuerySet {
	return qs.w(qs.db.Where("created_at = ?", createdAt))
}

// CreatedAtGt is an autogenerated method
// nolint: dupl
func (qs ArticleQuerySet) CreatedAtGt(createdAt time.Time) ArticleQuerySet {
	return qs.w(qs.db.Where("created_at > ?", createdAt))
}

// CreatedAtGte is an autogenerated method
// nolint: dupl
func (qs ArticleQuerySet) CreatedAtGte(createdAt time.Time) ArticleQuerySet {
	return qs.w(qs.db.Where("created_at >= ?", createdAt))
}

// CreatedAtLt is an autogenerated method
// nolint: dupl
func (qs ArticleQuerySet) CreatedAtLt(createdAt time.Time) ArticleQuerySet {
	return qs.w(qs.db.Where("created_at < ?", createdAt))
}

// CreatedAtLte is an autogenerated method
// nolint: dupl
func (qs ArticleQuerySet) CreatedAtLte(createdAt time.Time) ArticleQuerySet {
	return qs.w(qs.db.Where("created_at <= ?", createdAt))
}

// CreatedAtNe is an autogenerated method
// nolint: dupl
func (qs ArticleQuerySet) CreatedAtNe(createdAt time.Time) ArticleQuerySet {
	return qs.w(qs.db.Where("created_at != ?", createdAt))
}

// Delete is an autogenerated method
// nolint: dupl
func (o *Article) Delete(db *gorm.DB) error {
	return db.Delete(o).Error
}

// Delete is an autogenerated method
// nolint: dupl
func (qs ArticleQuerySet) Delete() error {
	return qs.db.Delete(Article{}).Error
}

// DeleteNum is an autogenerated method
// nolint: dupl
func (qs ArticleQuerySet) DeleteNum() (int64, error) {
	db := qs.db.Delete(Article{})
	return db.RowsAffected, db.Error
}

// DeleteNumUnscoped is an autogenerated method
// nolint: dupl
func (qs ArticleQuerySet) DeleteNumUnscoped() (int64, error) {
	db := qs.db.Unscoped().Delete(Article{})
	return db.RowsAffected, db.Error
}

// GetUpdater is an autogenerated method
// nolint: dupl
func (qs ArticleQuerySet) GetUpdater() ArticleUpdater {
	return NewArticleUpdater(qs.db)
}

// IDEq is an autogenerated method
// nolint: dupl
func (qs ArticleQuerySet) IDEq(ID int64) ArticleQuerySet {
	return qs.w(qs.db.Where("id = ?", ID))
}

// IDGt is an autogenerated method
// nolint: dupl
func (qs ArticleQuerySet) IDGt(ID int64) ArticleQuerySet {
	return qs.w(qs.db.Where("id > ?", ID))
}

// IDGte is an autogenerated method
// nolint: dupl
func (qs ArticleQuerySet) IDGte(ID int64) ArticleQuerySet {
	return qs.w(qs.db.Where("id >= ?", ID))
}

// IDIn is an autogenerated method
// nolint: dupl
func (qs ArticleQuerySet) IDIn(ID ...int64) ArticleQuerySet {
	if len(ID) == 0 {
		qs.db.AddError(errors.New("must at least pass one ID in IDIn"))
		return qs.w(qs.db)
	}
	return qs.w(qs.db.Where("id IN (?)", ID))
}

// IDLt is an autogenerated method
// nolint: dupl
func (qs ArticleQuerySet) IDLt(ID int64) ArticleQuerySet {
	return qs.w(qs.db.Where("id < ?", ID))
}

// IDLte is an autogenerated method
// nolint: dupl
func (qs ArticleQuerySet) IDLte(ID int64) ArticleQuerySet {
	return qs.w(qs.db.Where("id <= ?", ID))
}

// IDNe is an autogenerated method
// nolint: dupl
func (qs ArticleQuerySet) IDNe(ID int64) ArticleQuerySet {
	return qs.w(qs.db.Where("id != ?", ID))
}

// IDNotIn is an autogenerated method
// nolint: dupl
func (qs ArticleQuerySet) IDNotIn(ID ...int64) ArticleQuerySet {
	if len(ID) == 0 {
		qs.db.AddError(errors.New("must at least pass one ID in IDNotIn"))
		return qs.w(qs.db)
	}
	return qs.w(qs.db.Where("id NOT IN (?)", ID))
}

// Limit is an autogenerated method
// nolint: dupl
func (qs ArticleQuerySet) Limit(limit int) ArticleQuerySet {
	return qs.w(qs.db.Limit(limit))
}

// Offset is an autogenerated method
// nolint: dupl
func (qs ArticleQuerySet) Offset(offset int) ArticleQuerySet {
	return qs.w(qs.db.Offset(offset))
}

// One is used to retrieve one result. It returns gorm.ErrRecordNotFound
// if nothing was fetched
func (qs ArticleQuerySet) One(ret *Article) error {
	return qs.db.First(ret).Error
}

// OrderAscByCreatedAt is an autogenerated method
// nolint: dupl
func (qs ArticleQuerySet) OrderAscByCreatedAt() ArticleQuerySet {
	return qs.w(qs.db.Order("created_at ASC"))
}

// OrderAscByID is an autogenerated method
// nolint: dupl
func (qs ArticleQuerySet) OrderAscByID() ArticleQuerySet {
	return qs.w(qs.db.Order("id ASC"))
}

// OrderAscByUpdatedAt is an autogenerated method
// nolint: dupl
func (qs ArticleQuerySet) OrderAscByUpdatedAt() ArticleQuerySet {
	return qs.w(qs.db.Order("updated_at ASC"))
}

// OrderDescByCreatedAt is an autogenerated method
// nolint: dupl
func (qs ArticleQuerySet) OrderDescByCreatedAt() ArticleQuerySet {
	return qs.w(qs.db.Order("created_at DESC"))
}

// OrderDescByID is an autogenerated method
// nolint: dupl
func (qs ArticleQuerySet) OrderDescByID() ArticleQuerySet {
	return qs.w(qs.db.Order("id DESC"))
}

// OrderDescByUpdatedAt is an autogenerated method
// nolint: dupl
func (qs ArticleQuerySet) OrderDescByUpdatedAt() ArticleQuerySet {
	return qs.w(qs.db.Order("updated_at DESC"))
}

// SetContent is an autogenerated method
// nolint: dupl
func (u ArticleUpdater) SetContent(content string) ArticleUpdater {
	u.fields[string(ArticleDBSchema.Content)] = content
	return u
}

// SetCreatedAt is an autogenerated method
// nolint: dupl
func (u ArticleUpdater) SetCreatedAt(createdAt time.Time) ArticleUpdater {
	u.fields[string(ArticleDBSchema.CreatedAt)] = createdAt
	return u
}

// SetID is an autogenerated method
// nolint: dupl
func (u ArticleUpdater) SetID(ID int64) ArticleUpdater {
	u.fields[string(ArticleDBSchema.ID)] = ID
	return u
}

// SetThumbnail is an autogenerated method
// nolint: dupl
func (u ArticleUpdater) SetThumbnail(thumbnail string) ArticleUpdater {
	u.fields[string(ArticleDBSchema.Thumbnail)] = thumbnail
	return u
}

// SetTitle is an autogenerated method
// nolint: dupl
func (u ArticleUpdater) SetTitle(title string) ArticleUpdater {
	u.fields[string(ArticleDBSchema.Title)] = title
	return u
}

// SetUpdatedAt is an autogenerated method
// nolint: dupl
func (u ArticleUpdater) SetUpdatedAt(updatedAt time.Time) ArticleUpdater {
	u.fields[string(ArticleDBSchema.UpdatedAt)] = updatedAt
	return u
}

// ThumbnailEq is an autogenerated method
// nolint: dupl
func (qs ArticleQuerySet) ThumbnailEq(thumbnail string) ArticleQuerySet {
	return qs.w(qs.db.Where("thumbnail = ?", thumbnail))
}

// ThumbnailIn is an autogenerated method
// nolint: dupl
func (qs ArticleQuerySet) ThumbnailIn(thumbnail ...string) ArticleQuerySet {
	if len(thumbnail) == 0 {
		qs.db.AddError(errors.New("must at least pass one thumbnail in ThumbnailIn"))
		return qs.w(qs.db)
	}
	return qs.w(qs.db.Where("thumbnail IN (?)", thumbnail))
}

// ThumbnailNe is an autogenerated method
// nolint: dupl
func (qs ArticleQuerySet) ThumbnailNe(thumbnail string) ArticleQuerySet {
	return qs.w(qs.db.Where("thumbnail != ?", thumbnail))
}

// ThumbnailNotIn is an autogenerated method
// nolint: dupl
func (qs ArticleQuerySet) ThumbnailNotIn(thumbnail ...string) ArticleQuerySet {
	if len(thumbnail) == 0 {
		qs.db.AddError(errors.New("must at least pass one thumbnail in ThumbnailNotIn"))
		return qs.w(qs.db)
	}
	return qs.w(qs.db.Where("thumbnail NOT IN (?)", thumbnail))
}

// TitleEq is an autogenerated method
// nolint: dupl
func (qs ArticleQuerySet) TitleEq(title string) ArticleQuerySet {
	return qs.w(qs.db.Where("title = ?", title))
}

// TitleIn is an autogenerated method
// nolint: dupl
func (qs ArticleQuerySet) TitleIn(title ...string) ArticleQuerySet {
	if len(title) == 0 {
		qs.db.AddError(errors.New("must at least pass one title in TitleIn"))
		return qs.w(qs.db)
	}
	return qs.w(qs.db.Where("title IN (?)", title))
}

// TitleNe is an autogenerated method
// nolint: dupl
func (qs ArticleQuerySet) TitleNe(title string) ArticleQuerySet {
	return qs.w(qs.db.Where("title != ?", title))
}

// TitleNotIn is an autogenerated method
// nolint: dupl
func (qs ArticleQuerySet) TitleNotIn(title ...string) ArticleQuerySet {
	if len(title) == 0 {
		qs.db.AddError(errors.New("must at least pass one title in TitleNotIn"))
		return qs.w(qs.db)
	}
	return qs.w(qs.db.Where("title NOT IN (?)", title))
}

// Update is an autogenerated method
// nolint: dupl
func (u ArticleUpdater) Update() error {
	return u.db.Updates(u.fields).Error
}

// UpdateNum is an autogenerated method
// nolint: dupl
func (u ArticleUpdater) UpdateNum() (int64, error) {
	db := u.db.Updates(u.fields)
	return db.RowsAffected, db.Error
}

// UpdatedAtEq is an autogenerated method
// nolint: dupl
func (qs ArticleQuerySet) UpdatedAtEq(updatedAt time.Time) ArticleQuerySet {
	return qs.w(qs.db.Where("updated_at = ?", updatedAt))
}

// UpdatedAtGt is an autogenerated method
// nolint: dupl
func (qs ArticleQuerySet) UpdatedAtGt(updatedAt time.Time) ArticleQuerySet {
	return qs.w(qs.db.Where("updated_at > ?", updatedAt))
}

// UpdatedAtGte is an autogenerated method
// nolint: dupl
func (qs ArticleQuerySet) UpdatedAtGte(updatedAt time.Time) ArticleQuerySet {
	return qs.w(qs.db.Where("updated_at >= ?", updatedAt))
}

// UpdatedAtLt is an autogenerated method
// nolint: dupl
func (qs ArticleQuerySet) UpdatedAtLt(updatedAt time.Time) ArticleQuerySet {
	return qs.w(qs.db.Where("updated_at < ?", updatedAt))
}

// UpdatedAtLte is an autogenerated method
// nolint: dupl
func (qs ArticleQuerySet) UpdatedAtLte(updatedAt time.Time) ArticleQuerySet {
	return qs.w(qs.db.Where("updated_at <= ?", updatedAt))
}

// UpdatedAtNe is an autogenerated method
// nolint: dupl
func (qs ArticleQuerySet) UpdatedAtNe(updatedAt time.Time) ArticleQuerySet {
	return qs.w(qs.db.Where("updated_at != ?", updatedAt))
}

// ===== END of query set ArticleQuerySet

// ===== BEGIN of Article modifiers

// ArticleDBSchemaField describes database schema field. It requires for method 'Update'
type ArticleDBSchemaField string

// String method returns string representation of field.
// nolint: dupl
func (f ArticleDBSchemaField) String() string {
	return string(f)
}

// ArticleDBSchema stores db field names of Article
var ArticleDBSchema = struct {
	ID        ArticleDBSchemaField
	Title     ArticleDBSchemaField
	Content   ArticleDBSchemaField
	Thumbnail ArticleDBSchemaField
	UpdatedAt ArticleDBSchemaField
	CreatedAt ArticleDBSchemaField
}{

	ID:        ArticleDBSchemaField("id"),
	Title:     ArticleDBSchemaField("title"),
	Content:   ArticleDBSchemaField("content"),
	Thumbnail: ArticleDBSchemaField("thumbnail"),
	UpdatedAt: ArticleDBSchemaField("updated_at"),
	CreatedAt: ArticleDBSchemaField("created_at"),
}

// Update updates Article fields by primary key
// nolint: dupl
func (o *Article) Update(db *gorm.DB, fields ...ArticleDBSchemaField) error {
	dbNameToFieldName := map[string]interface{}{
		"id":         o.ID,
		"title":      o.Title,
		"content":    o.Content,
		"thumbnail":  o.Thumbnail,
		"updated_at": o.UpdatedAt,
		"created_at": o.CreatedAt,
	}
	u := map[string]interface{}{}
	for _, f := range fields {
		fs := f.String()
		u[fs] = dbNameToFieldName[fs]
	}
	if err := db.Model(o).Updates(u).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return err
		}

		return fmt.Errorf("can't update Article %v fields %v: %s",
			o, fields, err)
	}

	return nil
}

// ArticleUpdater is an Article updates manager
type ArticleUpdater struct {
	fields map[string]interface{}
	db     *gorm.DB
}

// NewArticleUpdater creates new Article updater
// nolint: dupl
func NewArticleUpdater(db *gorm.DB) ArticleUpdater {
	return ArticleUpdater{
		fields: map[string]interface{}{},
		db:     db.Model(&Article{}),
	}
}

// ===== END of Article modifiers

// ===== END of all query sets