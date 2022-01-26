package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"Project/research/sample-gql/delivery/middleware"
	"Project/research/sample-gql/entities"
	"Project/research/sample-gql/entities/model"
	"Project/research/sample-gql/util/graph/generated"
	"context"
	"errors"
	"fmt"
)

func (r *mutationResolver) AddPerson(ctx context.Context, input model.NewPerson) (*model.Person, error) {
	res, err := r.personRepo.Create(entities.Person{Nama: input.Nama, HP: *input.Hp, Umur: input.Umur, Password: input.Password})

	if err != nil {
		return nil, errors.New("not found")
	}

	return &model.Person{ID: int(res.ID), Nama: res.Nama, Hp: &res.HP, Umur: res.Umur}, nil
}

func (r *mutationResolver) AddBook(ctx context.Context, input model.NewBook) (*model.Book, error) {
	var res1 entities.Person
	var err error
	if input.Author != nil {
		res1, err = r.personRepo.Create(entities.Person{Nama: input.Author.Nama, Umur: input.Author.Umur, HP: *input.Author.Hp, Password: input.Author.Password})
		if err != nil {
			return nil, errors.New("not found")
		}
	}

	res, err := r.bookRepo.Create(entities.Book{Title: input.Title, Author: uint(res1.ID)})
	if err != nil {
		return nil, errors.New("not found")
	}

	convID := int(res.ID)

	return &model.Book{ID: &convID, Title: res.Title, Author: &model.Person{ID: int(res1.ID), Nama: res1.Nama, Hp: &res1.HP}}, nil
}

func (r *mutationResolver) BuatPerson(ctx context.Context, nama string, hp *string, umur *int, password string) (*model.Person, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Books(ctx context.Context) ([]*model.Book, error) {
	res2, err := r.bookRepo.GraphGet()

	if err != nil {
		return nil, errors.New("not found")
	}

	return res2, nil
}

func (r *queryResolver) BooksByID(ctx context.Context, id int) (*model.Book, error) {
	res, err := r.bookRepo.GraphGetByID(id)

	if err != nil {
		return nil, errors.New("not found")
	}

	return res, nil
}

func (r *queryResolver) Persons(ctx context.Context) ([]*model.Person, error) {
	raw := ctx.Value("EchoContextKey").(*middleware.User)
	convRaw := *raw
	fmt.Println(convRaw.Id)

	res, err := r.personRepo.Get()

	if err != nil {
		return nil, errors.New("not found")
	}
	resArr := []*model.Person{}
	for i := 0; i < len(res); i++ {
		resArr = append(resArr, &model.Person{ID: int(res[i].ID), Nama: res[i].Nama, Umur: res[i].Umur, Hp: &res[i].HP})
	}

	return resArr, err
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
