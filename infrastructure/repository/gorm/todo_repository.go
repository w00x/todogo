package gorm

import (
	"gorm.io/gorm"
	"todogohexa/domain"
	"todogohexa/infrastructure/errors"
	"todogohexa/infrastructure/repository/gorm/mappers"
	"todogohexa/infrastructure/repository/gorm/models"
)

type TodoRepository struct {
	postgresBase *PostgresBase
}

func NewTodoRepository() *TodoRepository {
	postgresBase := NewPostgresBase()
	return &TodoRepository{postgresBase}
}

func (r TodoRepository) All() (*[]domain.Todo, errors.IBaseError) {
	var todos []models.Todo
	result := r.postgresBase.DB.Model(&models.Todo{}).Scan(&todos)
	if err := result.Error; err != nil {
		return nil, errors.NewInternalServerError(err.Error())
	}
	return mappers.NewTodoListDomainFromModel(&todos), nil
}

func (r TodoRepository) FindById(id string) (*domain.Todo, errors.IBaseError) {
	var todo models.Todo

	result := r.postgresBase.DB.First(&todo, "id = ?", id)

	if err := result.Error; err == gorm.ErrRecordNotFound {
		return nil, errors.NewNotFoundError("Repository not found")
	} else if err != nil {
		return nil, errors.NewInternalServerError(err.Error())
	}

	return mappers.FromTodoModelToDomain(&todo), nil
}

func (r TodoRepository) Create(todo *domain.Todo) (*domain.Todo, errors.IBaseError) {
	model := mappers.FromTodoDomainToModel(todo)
	result:= r.postgresBase.DB.Create(model)
	if err := result.Error; err != nil {
		return nil, errors.NewInternalServerError(err.Error())
	}
	count := result.RowsAffected

	if count == 0 {
		return nil, errors.NewNotFoundError("Repository not created")
	}

	return r.FindById(model.ID)
}

func (r TodoRepository) Update(todo *domain.Todo) (*domain.Todo, errors.IBaseError) {
	d, err := r.FindById(todo.Id())
	if err != nil {
		return nil, err
	}

	result := r.postgresBase.DB.Model(mappers.FromTodoDomainToModel(d)).Updates(mappers.FromTodoDomainToModel(todo))

	if err := result.Error; err != nil {
		return nil, errors.NewInternalServerError(err.Error())
	}
	count := result.RowsAffected

	if count == 0 {
		return nil, errors.NewNotFoundError("Repository not updated")
	}

	return r.FindById(todo.Id())
}

func (r TodoRepository) Destroy(todo *domain.Todo) errors.IBaseError {
	model := mappers.FromTodoDomainToModel(todo)
	result := r.postgresBase.DB.Delete(model)

	if err := result.Error; err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	count := result.RowsAffected

	if count == 0 {
		return errors.NewNotFoundError("Repository not deleted")
	}

	return nil
}