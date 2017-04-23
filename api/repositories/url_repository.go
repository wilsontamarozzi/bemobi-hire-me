package repositories

import (
	"github.com/wilsontamarozzi/bemobi-hire-me/api/database"
	"github.com/wilsontamarozzi/bemobi-hire-me/api/helpers"
	"github.com/wilsontamarozzi/bemobi-hire-me/api/models"
	"log"
	"net/url"
	"strconv"
)

type URLRepositoryInterface interface {
	GetAllRanking(q url.Values) models.URLS
	GetByAlias(id string) models.URL
	Create(url *models.URL) error
	UpdateView(url *models.URL) error
}

type urlRepository struct{}

func NewURLRepository() *urlRepository {
	return new(urlRepository)
}

func (repository urlRepository) GetAllRanking(q url.Values) models.URLS {
	db := database.GetInstance()

	currentPage, _ := strconv.Atoi(q.Get("page"))
	itemPerPage, _ := strconv.Atoi(q.Get("per_page"))
	pagination := helpers.MakePagination(repository.CountRows(), currentPage, itemPerPage)

	var urls models.URLS
	urls.Meta.Pagination = pagination

	db.Limit(pagination.ItemPerPage).
		Offset(pagination.StartIndex).
		Order("view desc").
		Find(&urls.Urls)

	return urls
}

func (repository urlRepository) GetByAlias(alias string) models.URL {
	db := database.GetInstance()

	var url models.URL
	db.Where("alias = ?", alias).
		First(&url)

	return url
}

func (repository urlRepository) Get(uuid string) models.URL {
	db := database.GetInstance()

	var url models.URL
	db.Where("uuid = ?", uuid).
		First(&url)

	return url
}

func (repository urlRepository) Create(url *models.URL) error {
	db := database.GetInstance()

	err := db.Save(&url).Error
	if err != nil {
		log.Print(err.Error())
		return err
	}

	*(url) = repository.Get(url.UUID)

	return nil
}

func (repository urlRepository) UpdateView(url *models.URL) error {
	db := database.GetInstance()

	url.View++
	err := db.Model(&models.URL{}).
		Where("alias = ?", url.Alias).
		Updates(&url).Error

	if err != nil {
		log.Print(err.Error())
	}

	return err
}

func (repository urlRepository) CountRows() int {
	db := database.GetInstance()
	var count int
	db.Model(&models.URL{}).Count(&count)

	return count
}
