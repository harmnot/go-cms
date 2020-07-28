package resolver

import (
	"cms/generated/gqlgen"
	prisma "cms/generated/prisma-client"
	"cms/helper"
	"context"
	"fmt"
	"github.com/vektah/gqlparser/gqlerror"
)

// THIS CODE IS A STARTING POINT ONLY. IT WILL NOT BE UPDATED WITH SCHEMA CHANGES.

type Resolver struct{
	DB *prisma.Client
}

func (r *Resolver) Content() gqlgen.ContentResolver {
	return &contentResolver{r}
}
func (r *Resolver) Mutation() gqlgen.MutationResolver {
	return &mutationResolver{r}
}
func (r *Resolver) Person() gqlgen.PersonResolver {
	return &personResolver{r}
}
func (r *Resolver) Query() gqlgen.QueryResolver {
	return &queryResolver{r}
}

type contentResolver struct{ *Resolver }

func (r *contentResolver) Company(ctx context.Context, obj *prisma.Content) (*prisma.Company, error) {
	panic("not implemented")
}

type mutationResolver struct{ *Resolver }

func (r *mutationResolver) CreatePerson(ctx context.Context, name string, email string, password string, company string) (*prisma.Person, error) {
	hash, _ := helper.HashPassword(password)
	person, err := r.DB.CreatePerson(prisma.PersonCreateInput{
		Name:     name,
		Email:    email,
		Password: hash,
		Company:  &prisma.CompanyCreateOneInput{
			Create:  &prisma.CompanyCreateInput{
				Name: company,
			},
		},
	}).Exec(ctx)
	if err != nil {
		return nil, gqlerror.Errorf("System Error")
	} else {
		return person, nil
	}
}
func (r *mutationResolver) CreateContent(ctx context.Context, title *string, text string, category string, typeArg string, limit *int, note *string) (*prisma.Content, error) {
	//jwt := ctx.Value("Token")
	//initialUser, err := helper.ExtractToken(fmt.Sprintf("%v", jwt))
}
func (r *mutationResolver) EditContent(ctx context.Context, title *string, text *string, category *string, typeArg *string, limit *int, note *string) (*prisma.Content, error) {
	panic("not implemented")
}

type personResolver struct{ *Resolver }

func (r *personResolver) Company(ctx context.Context, obj *prisma.Person) (*prisma.Company, error) {
	return r.DB.Person(prisma.PersonWhereUniqueInput{
		ID:    &obj.ID,
	}).Company().Exec(ctx)
}

type queryResolver struct{ *Resolver }

func (r *queryResolver) Me(ctx context.Context) (*prisma.Person, error) {
	panic("not implemented")
}
func (r *queryResolver) FindOneContent(ctx context.Context, contentID string) (*prisma.Content, error) {
	panic("not implemented")
}
func (r *queryResolver) FindAllContent(ctx context.Context) ([]prisma.Content, error) {
	panic("not implemented")
}
func (r *queryResolver) FindContentBy(ctx context.Context, category *string, typeArg *string) ([]prisma.Content, error) {
	panic("not implemented")
}
