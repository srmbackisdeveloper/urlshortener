package repository

import (
	"log"
	"time"
	"urlshortener/internal/model"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type URLRepository struct {
	DB *gorm.DB
}

func NewPostgresDB(dsn string) (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	sqlDB, err := db.DB()
	if err != nil {
        return nil, err
    }
	
	if err := sqlDB.Ping(); err != nil {
        return nil, err
    }

	log.Println("Connected to PostgreSQL")
    return db, nil
}

func NewURLRepository(db *gorm.DB) *URLRepository {
    return &URLRepository{DB: db}
}

// ---

func (repo *URLRepository) SaveURL(originalURL, shortCode string) (*model.URL, error) {
    url := &model.URL{
        OriginalURL: originalURL,
        ShortCode:   shortCode,
        CreatedAt:   time.Now(),
    }

    if err := repo.DB.Create(url).Error; err != nil {
        return nil, err
    }
    return url, nil
}

func (repo *URLRepository) GetURLByShortCode(shortCode string) (*model.URL, error) {
    var url model.URL
    if err := repo.DB.Where("short_code = ?", shortCode).First(&url).Error; err != nil {
        return nil, err
    }
    return &url, nil
}

func (repo *URLRepository) IncrementClickCount(shortCode string) error {
    return repo.DB.Model(&model.URL{}).
        Where("short_code = ?", shortCode).
        UpdateColumn("clicks", gorm.Expr("clicks + ?", 1)).
        Error
}
