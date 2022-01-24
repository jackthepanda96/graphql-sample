package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"Project/research/sample-gql/entities"
	entities1 "Project/research/sample-gql/entities/model"
	"Project/research/sample-gql/util/graph/generated"
	"context"
	"errors"
	"fmt"
	"strconv"
)

func (r *mutationResolver) AddPerson(ctx context.Context, input entities1.NewPerson) (*entities1.Person, error) {
	res, err := r.personRepo.Create(entities.Person{Nama: input.Nama, HP: *input.Hp, Umur: input.Umur, Password: input.Password})

	if err != nil {
		return nil, errors.New("not found")
	}

	// r.tmpList = append(r.tmpList, &model.Person{ID: strconv.Itoa(int(res.ID)), Nama: res.Nama, Hp: &res.HP, Umur: res.Umur})

	// r.mu.Lock()

	// for _, osv := range r.Observer {
	// 	osv <- &entities1.Person{ID: strconv.Itoa(int(res.ID)), Nama: res.Nama, Hp: &res.HP, Umur: res.Umur}
	// }

	// r.mu.Unlock()

	return &entities1.Person{ID: strconv.Itoa(int(res.ID)), Nama: res.Nama, Hp: &res.HP, Umur: res.Umur}, nil
}

func (r *mutationResolver) AddBook(ctx context.Context, input entities1.NewBook) (*entities1.Book, error) {
	res, err := r.bookRepo.Create(entities.Book{Title: input.Title, Author: uint(input.Author)})

	if err != nil {
		return nil, errors.New("not found")
	}
	convID := int(res.ID)

	return &entities1.Book{ID: &convID, Title: res.Title}, nil
	// return &res, nil

	// panic(fmt.Errorf("not implemented"))

	// res, err := repo.Create(entities.Book{Title: input.Title, Author: })
}

func (r *queryResolver) Books(ctx context.Context) ([]*entities1.Book, error) {
	res, err := r.bookRepo.Get()

	preload := GetPreloads(ctx)
	fmt.Println(preload)

	pointRes := []*entities1.Book{}
	// pointRes := []*entities.Book{}

	// for _, v := range res {
	// 	pointRes = append(pointRes, &model.Book{Title: v.Title})
	// }

	for i := 0; i < len(res); i++ {
		convID := int(res[i].ID)
		pointRes = append(pointRes, &entities1.Book{ID: &convID, Title: res[i].Title})
	}

	// for i := 0; i < len(res); i++ {
	// 	// convID := int(res[i].ID)
	// 	pointRes = append(pointRes, &res[i])
	// }

	if err != nil {
		return nil, errors.New("not found")
	}

	return pointRes, nil
}

func (r *queryResolver) Persons(ctx context.Context) ([]*entities1.Person, error) {
	res, err := r.personRepo.Get()

	if err != nil {
		return nil, errors.New("not found")
	}
	resArr := []*entities1.Person{}
	for i := 0; i < len(res); i++ {
		resArr = append(resArr, &entities1.Person{ID: strconv.Itoa(int(res[i].ID)), Nama: res[i].Nama, Umur: res[i].Umur, Hp: &res[i].HP})
	}

	return resArr, err
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
