package repo

import (
	"context"
	"neuro-most/tags-service/internal/entities"
)

type tagsGORM struct {
	Id   int64  `gorm:"primaryKey"`
	Name string `gorm:"not null"`
}

type TagsRepo struct {
	db GSQL
}

func NewTagsRepo(db GSQL) TagsRepo {
	return TagsRepo{db: db}
}

func (r TagsRepo) Create(ctx context.Context, tag entities.Tags) error {
	var tagGORM tagsGORM
	tagGORM.Name = tag.Name()
	if err := r.db.Create(ctx, &tagGORM); err != nil {
		return err
	}
	return nil
}

func (r TagsRepo) Update(ctx context.Context, tag entities.Tags) error {
	updates := map[string]interface{}{
		"name": tag.Name(),
	}
	if err := r.db.UpdateOne(ctx, &updates, &tagsGORM{Id: tag.ID()}, &tagsGORM{}); err != nil {
		return err
	}
	return nil
}

func (r TagsRepo) Delete(ctx context.Context, tag entities.Tags) error {
	if err := r.db.Delete(ctx, &tagsGORM{}, &tagsGORM{Id: tag.ID()}); err != nil {
		return err
	}
	return nil
}

func (r TagsRepo) Fetch(ctx context.Context, page, pageSize int64) ([]entities.Tags, int64, error) {
	var tags []tagsGORM
	query := r.db.BeginFind(ctx, &tags)
	var total int64
	query.Count(&total)
	query = query.Page(int(page), int(pageSize)).OrderBy("id desc")
	err := query.Find(&tags)
	if err != nil {
		return nil, 0, entities.ErrorTagsFetch
	}
	var result []entities.Tags
	for _, tag := range tags {
		result = append(result, r.convertToTag(tag))
	}
	return result, total, nil
}

func (r TagsRepo) GetByID(ctx context.Context, id int64) (entities.Tags, error) {
	var media tagsGORM
	if err := r.db.BeginFind(ctx, &media).Where(&tagsGORM{Id: id}).First(&media); err != nil {
		return entities.Tags{}, entities.ErrTagsNotFound
	}
	return r.convertToTag(media), nil
}

func (r TagsRepo) convertToTag(tagsGORM tagsGORM) entities.Tags {
	return entities.NewTag(
		tagsGORM.Id,
		tagsGORM.Name,
	)
}
