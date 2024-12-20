package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.42

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"github.com/google/uuid"
	"github.com/keviinliuu/leetlist/graph"
	"github.com/keviinliuu/leetlist/graph/model"
	"github.com/keviinliuu/leetlist/util"
)

// CreateQuestion is the resolver for the createQuestion field.
func (r *mutationResolver) CreateQuestion(ctx context.Context, input model.NewQuestion) (*model.Question, error) {
	question := model.Question{
		ID:         uuid.New().String(),
		Title:      input.Title,
		URL:        input.URL,
		Difficulty: input.Difficulty,
		Complete:   false,
	}

	err := r.DB.Create(&question).Error
	if err != nil {
		return nil, err
	}

	return &question, nil
}

// UpdateQuestion is the resolver for the updateQuestion field.
func (r *mutationResolver) UpdateQuestion(ctx context.Context, id string, input model.UpdateQuestion) (*model.Question, error) {
	var question model.Question

	err := r.DB.First(&question, "id = ?", id).Error
	if err != nil {
		return nil, err
	}

	if input.Title != nil {
		question.Title = *input.Title
	}
	if input.URL != nil {
		question.URL = *input.URL
	}
	if input.Difficulty != nil {
		question.Difficulty = *input.Difficulty
	}

	err = r.DB.Save(&question).Error
	if err != nil {
		return nil, err
	}

	return &question, nil
}

// CreateList is the resolver for the createList field.
func (r *mutationResolver) CreateList(ctx context.Context, input model.NewList) (*model.List, error) {
	var user model.User
	if err := r.DB.Where("ID = ?", input.UserID).First(&user).Error; err != nil {
		return nil, fmt.Errorf("user not found")
	}

	list := model.List{
		ID:          uuid.New().String(),
		Title:       input.Title,
		Description: input.Description,
		UserID:      user.ID,
	}

	entries := []*model.Question{}

	for _, qInput := range input.Entries {
		if qInput == nil {
			continue
		}

		question, err := r.CreateQuestion(ctx, *qInput)
		if err != nil {
			return nil, err
		}
		entries = append(entries, question)
	}

	list.Entries = entries

	err := r.DB.Create(&list).Error
	if err != nil {
		return nil, err
	}

	user.Lists = append(user.Lists, &list)
	if err := r.DB.Save(&user).Error; err != nil {
		return nil, err
	}

	return &list, err
}

// UpdateList is the resolver for the updateList field.
func (r *mutationResolver) UpdateList(ctx context.Context, id string, input model.UpdateList) (*model.List, error) {
	var list model.List

	err := r.DB.Preload("Entries").First(&list, "id = ?", id).Error
	if err != nil {
		return nil, err
	}

	if input.Title != nil {
		list.Title = *input.Title
	}
	if input.Description != nil {
		list.Description = input.Description
	}

	for _, qInput := range input.AddQuestions {
		if qInput == nil {
			continue
		}

		question, err := r.CreateQuestion(ctx, *qInput)
		if err != nil {
			return nil, err
		}
		list.Entries = append(list.Entries, question)
	}

	for _, qID := range input.RemoveQuestionIds {
		err := r.DB.Model(&list).Association("Entries").Delete(&model.Question{ID: qID})
		if err != nil {
			return nil, err
		}

		var question model.Question

		err = r.DB.Where("id = ?", qID).Delete(&question).Error
		if err != nil {
			return nil, err
		}
	}

	err = r.DB.Save(&list).Error
	if err != nil {
		return nil, err
	}

	return &list, nil
}

// DeleteQuestion is the resolver for the deleteQuestion field.
func (r *mutationResolver) DeleteQuestion(ctx context.Context, id string) (*model.Question, error) {
	var question model.Question

	err := r.DB.First(&question, "id = ?", id).Error
	if err != nil {
		return nil, fmt.Errorf("Question not found: %v", err)
	}

	if err := r.DB.Exec("DELETE FROM list_questions WHERE question_id = ?", id).Error; err != nil {
		return nil, fmt.Errorf("failed to disassociate question from lists: %v", err)
	}

	err = r.DB.Delete(&question).Error
	if err != nil {
		return nil, fmt.Errorf("failed to delete question: %v", err)
	}

	return nil, nil
}

// DeleteList is the resolver for the deleteList field.
func (r *mutationResolver) DeleteList(ctx context.Context, id string) (*model.List, error) {
	var list model.List

	err := r.DB.Preload("Entries").First(&list, "id = ?", id).Error
	if err != nil {
		return nil, fmt.Errorf("List not found: %v", err)
	}

	for _, question := range list.Entries {
		_, err := r.DeleteQuestion(ctx, question.ID)
		if err != nil {
			return nil, fmt.Errorf("failed to delete associated question with ID %s: %v", question.ID, err)
		}
	}

	err = r.DB.Delete(&list).Error
	if err != nil {
		return nil, fmt.Errorf("failed to delete list: %v", err)
	}

	return &list, nil
}

// Register is the resolver for the register field.
func (r *mutationResolver) Register(ctx context.Context, input model.NewUser) (*model.AuthPayload, error) {
	var existingUser model.User

	err := r.DB.Where("email = ?", input.Email).First(&existingUser).Error
	if err == nil {
		return nil, fmt.Errorf("a user with this email already exists")
	}

	hashedPassword, err := util.GenerateHashPassword(input.Password)
	if err != nil {
		return nil, err
	}

	user := model.User{
		ID:       uuid.New().String(),
		Email:    input.Email,
		Password: hashedPassword,
	}

	err = r.DB.Create(&user).Error
	if err != nil {
		return nil, err
	}

	token, err := util.GenerateToken(user)
	if err != nil {
		return nil, fmt.Errorf("failed to generate token: %v", err)
	}

	return &model.AuthPayload{
		Token: &token,
		User:  &user,
	}, nil
}

// Login is the resolver for the login field.
func (r *mutationResolver) Login(ctx context.Context, email string, password string) (*model.AuthPayload, error) {
	var user model.User

	err := r.DB.Where("email = ?", email).First(&user).Error
	if err != nil {
		return nil, fmt.Errorf("user not found")
	}

	err = util.CompareHashPassword(user.Password, password)
	if err != nil {
		return nil, fmt.Errorf("invalid password")
	}

	token, err := util.GenerateToken(user)
	if err != nil {
		return nil, fmt.Errorf("failed to generate token: %v", err)
	}

	return &model.AuthPayload{
		Token: &token,
		User:  &user,
	}, nil
}

// Question is the resolver for the question field.
func (r *queryResolver) Question(ctx context.Context, id string) (*model.Question, error) {
	var question model.Question

	err := r.DB.First(&question, "id = ?", id).Error
	if err != nil {
		return nil, err
	}

	return &question, nil
}

// Questions is the resolver for the questions field.
func (r *queryResolver) Questions(ctx context.Context) ([]*model.Question, error) {
	// email, ok := ctx.Value(auth.UserCtxKey).(string)
	// if !ok || email == "" {
	// 	return nil, errors.New("Unauthorized: must be logged in")
	// }

	var questions []*model.Question

	err := r.DB.Find(&questions).Error
	if err != nil {
		return nil, err
	}

	return questions, nil
}

// List is the resolver for the list field.
func (r *queryResolver) List(ctx context.Context, id string) (*model.List, error) {
	var list model.List

	err := r.DB.Preload("Entries").First(&list, "id = ?", id).Error
	if err != nil {
		return nil, err
	}

	return &list, nil
}

// Lists is the resolver for the lists field.
func (r *queryResolver) Lists(ctx context.Context) ([]*model.List, error) {
	var lists []*model.List

	err := r.DB.Preload("Entries").Find(&lists).Error
	if err != nil {
		return nil, err
	}

	return lists, nil
}

func (r *queryResolver) User(ctx context.Context, id string) (*model.User, error) {
	var user model.User

	if err := r.DB.Preload("Lists.Entries").First(&user, "id = ?", id).Error; err != nil {
		return nil, fmt.Errorf("user not found: %v", err)
	}

	return &user, nil
}

// ScrapeQuestion is the resolver for the scrapeQuestion field.
func (r *queryResolver) ScrapeQuestion(ctx context.Context, url string) (*model.QuestionInfo, error) {
	if !strings.HasPrefix(url, "https://leetcode.com/problems/") {
		return nil, errors.New("not a valid Leetcode problem URL")
	}

	title, difficulty := util.GetQuestionInfo(url, r.Browser)

	return &model.QuestionInfo{
		Title:      title,
		Difficulty: difficulty,
	}, nil
}

// Mutation returns graph.MutationResolver implementation.
func (r *Resolver) Mutation() graph.MutationResolver { return &mutationResolver{r} }

// Query returns graph.QueryResolver implementation.
func (r *Resolver) Query() graph.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
