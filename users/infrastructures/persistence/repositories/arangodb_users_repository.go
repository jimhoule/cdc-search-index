package repositories

import (
	"context"
	"main/database"
	"main/users/domain/models"
)

type ArangodbUsersRepository struct {
	Db *database.Db
}

func (aur *ArangodbUsersRepository) GetAll() ([]*models.User, error) {
	query := "FOR doc IN users RETURN doc"
	bindVars := map[string]any{}
	
	cursor, err:= aur.Db.Client.Query(context.Background(), query, bindVars)
	if err != nil {
		return nil, err
	}
	defer cursor.Close()

	users := []*models.User{}
	for cursor.HasMore() {
		user := &models.User{}
		
		_, err := cursor.ReadDocument(context.Background(), &user)
		if err != nil {
			return nil, err
		}

		users = append(users, user)
	}

	return users, nil
}

func (aur *ArangodbUsersRepository) GetById(id string) (*models.User, error) {
	query := "FOR doc IN users FILTER doc._key == @key RETURN doc"
	bindVars:= map[string]interface{}{
		"key":id,
	}

	cursor, err := aur.Db.Client.Query(context.Background(), query, bindVars)
	if err != nil {
		return nil, err
	}
	defer cursor.Close()

	user := &models.User{}
	for cursor.HasMore() {
		_, err := cursor.ReadDocument(context.Background(), &user)
		if err != nil {
			return nil, err
		}
	}
	
	return user, nil
}

func (aur *ArangodbUsersRepository) Create(user *models.User) (*models.User, error) {
	query := "INSERT { _key: @key, id: @id, firstname: @firstname, lastname: @lastname } INTO users"
	bindVars := map[string]any{
		"key": user.Id,
		"id": user.Id,
		"firstname": user.Firstname,
		"lastname": user.LastName,
	}

	_, err := aur.Db.Client.Query(context.Background(), query, bindVars)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (aur *ArangodbUsersRepository) Update(updatedUser *models.User) (*models.User, error) {
	query := "UPDATE @key WITH { firstname: @firstname, lastname: @lastname }"
	bindVars := map[string]any{
		"key": updatedUser.Id,
		"firstname": updatedUser.Firstname,
		"lastname": updatedUser.LastName,
	}

	_, err := aur.Db.Client.Query(context.Background(), query, bindVars)
	if err != nil {
		return nil, err
	}

	return updatedUser, nil
}

func (aur *ArangodbUsersRepository) Delete(id string) (string, error) {
	query := "REMOVE @key IN users"
	bindVars := map[string]any{
		"key": id,
	}

	_, err := aur.Db.Client.Query(context.Background(), query, bindVars)
	if err != nil {
		return "", err
	}

	return id, nil
}