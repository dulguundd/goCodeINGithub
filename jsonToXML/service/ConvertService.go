package service

import (
	"goCodeINGithub/jsonToXML/domain"
	"goCodeINGithub/jsonToXML/dto"
)

type ConvertService interface {
	XMLToJson(dto.XMLData) (dto.JsonData, error)
	JsonToXML(dto.JsonData) (dto.XMLData, error)
}

type DefaultConvertService struct {
	repo domain.ConvertRepository
}

func (s DefaultConvertService) XMLToJson(data dto.XMLData) (dto.JsonData, error) {
	return dto.JsonData{
		INACTENDD:   data.Channel.Post.INACTENDD,
		Error:       data.Channel.Post.Error,
		RETAILER:    data.Channel.Post.RETAILER,
		CLASS:       data.Channel.Post.CLASS,
		ACTENDD:     data.Channel.Post.ACTENDD,
		ADMINST:     data.Channel.Post.ADMINST,
		CREDITVIOCE: data.Channel.Post.CREDITVIOCE,
		CODE:        data.Channel.Post.CODE,
		PHONE:       data.Channel.Post.PHONE,
		RBAL:        data.Channel.Post.RBAL,
	}, nil
}

func (s DefaultConvertService) JsonToXML(data dto.JsonData) (dto.XMLData, error) {
	var response dto.XMLData
	response.Channel.Post = data
	return response, nil
}

func NewConvertService(repository domain.ConvertRepository) DefaultConvertService {
	return DefaultConvertService{(repository)}
}
