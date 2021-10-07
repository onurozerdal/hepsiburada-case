package service

import (
	"errors"
	"hepsiburada-case/api/model"
	"hepsiburada-case/api/repository"
)

const (
	personalized    = "personalized"
	nonPersonalized = "non-personalized"
	maxArraySize = 5
)

type ApiService struct {
	r repository.ApiRepository
}

func NewApiService(repository repository.ApiRepository) *ApiService {
	return &ApiService{r: repository}
}

func (service *ApiService) BrowsingHistories(userId string) (model.Response, error) {
	if len(userId) < 1 {
		return model.Response{}, errors.New("userId should not nil or blank")
	}
	r, _ := service.r.BrowsingHistories(userId)

	checkListSize(&r)

	res := model.Response{UserId: userId, Products: r, Type: personalized}
	return res, nil
}

func (service *ApiService) BestsellerProducts(userId string) (model.Response, error) {
	if len(userId) < 1 {
		return model.Response{}, errors.New("userId should not nil or blank")
	}
	c := service.r.CheckBrowsingHistories(userId)
	if c {
		r, _ := service.r.Bestseller10Products()
		checkListSize(&r)
		res := model.Response{UserId: userId, Products: r, Type: nonPersonalized}
		return res, nil
	}
	r, _ := service.r.Bestseller10ProductsByUserInterest(userId)
	checkListSize(&r)

	res := model.Response{UserId: userId, Products: r, Type: nonPersonalized}
	return res, nil
}

func (service *ApiService) DeleteHistory(userId, productId string) (string, error) {
	if len(userId) < 1 {
		return "", errors.New("userId should not nil or blank")
	}
	if len(productId) < 1 {
		return "", errors.New("productId should not nil or blank")
	}
	service.r.DeleteHistory(userId, productId)
	return "Successfuly deleted.", nil
}

func checkListSize(r *[]string) {
	if len(*r) < maxArraySize {
		*r = []string{}
	}
}