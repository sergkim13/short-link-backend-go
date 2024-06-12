package service

import (
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"time"

	"github.com/lib/pq"
	"github.com/sergkim13/short-link-backend-go/configs"
	"github.com/sergkim13/short-link-backend-go/pkg/repository"
	"github.com/sirupsen/logrus"
)

type LinkService struct {
	repo repository.Link
}

func (s *LinkService) MakeShort(originalURL string) (string, error) {
	LinksHost := configs.EnvConfig.LinksHost

	shortURL, err := s.CreateLink(originalURL)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%s/%s", LinksHost, shortURL), nil
}

func (s *LinkService) CreateLink(originalURL string) (string, error) {

	shortURL := s.generateHash(originalURL)

	_, err := s.repo.AddLink(originalURL, shortURL)

	if err != nil {
		if pqErr, ok := err.(*pq.Error); ok {
			if pqErr.Code == "23505" {
				if strings.Contains(pqErr.Message, "links_original_key") {
					logrus.Infof("original url %s already exists, returning it's short url", originalURL)

					shortURL, err = s.repo.GetShortByOriginalURL(originalURL)

					if err != nil {
						logrus.Errorf("error while getting link by existnig original url %s: %s", originalURL, err.Error())
						return "", err
					}

					return shortURL, nil

				} else if strings.Contains(pqErr.Message, "links_short_key")  {
					logrus.Infof("short url %s already exists, generating new one", originalURL)
					shortURL, err := s.MakeShort(originalURL)

					if err != nil {
						logrus.Errorf("error while generating new short link for existing %s: %s", shortURL, err.Error())
						return shortURL, err
					}
				}
			}
		}
		return "", err
	}

	return shortURL, nil
}

func (s *LinkService) generateHash(value string) string {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	randomNumber := strconv.FormatInt(r.Int63(), 10)
	hash := sha256.Sum256([]byte(value + randomNumber))
	shortURL := base64.URLEncoding.EncodeToString(hash[:6])

	return shortURL
}

func NewLinkService(repo repository.Link) *LinkService {
	return &LinkService{repo: repo}
}
