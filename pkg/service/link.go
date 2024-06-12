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
	"github.com/sergkim13/short-link-backend-go/pkg/repository"
)

type LinkService struct {
	repo repository.Link
}

func (s *LinkService) CreateLink(originalURL string) (string, error) {

	shortLink := s.generateHash(originalURL)

	res, err := s.repo.AddLink(originalURL, shortLink)
	if err != nil {
		if pqErr, ok := err.(*pq.Error); ok {
			if pqErr.Code == "23505" {
				if strings.Contains(pqErr.Message, "links_original_key") {
					return "", fmt.Errorf("original link %s already exists", originalURL)
				} else if strings.Contains(pqErr.Message, "links_short_key")  {
					return "", fmt.Errorf("short link %s already exists", shortLink)
				}
			}
		}
		return "", err
	}

	if res != "Ok" {
		return "", err
	}

	return shortLink, nil
}

func (s *LinkService) generateHash(value string) string {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	randomNumber := strconv.FormatInt(r.Int63(), 10)
	hash := sha256.Sum256([]byte(value + randomNumber))
	shortLink := base64.URLEncoding.EncodeToString(hash[:6])

	return shortLink
}

func (s *LinkService) checkIfOriginalExists(originalURL string) bool {
	return true
}

func NewLinkService(repo repository.Link) *LinkService {
	return &LinkService{repo: repo}
}
