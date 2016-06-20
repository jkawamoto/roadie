package util

import (
	"log"
	"net/http"
	"net/url"
	"os"

	"golang.org/x/net/context"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/storage/v1"
)

const gcsScope = storage.DevstorageFullControlScope

// Storage object.
type Storage struct {
	BucketName string
	client     *http.Client
	service    *storage.Service
}

// NewStorage creates a new Storage object named a given bucket name.
// If the given bucket does not exsits, it will be created.
func NewStorage(project, bucket string) (*Storage, error) {

	// Create a client.
	client, err := google.DefaultClient(context.Background(), gcsScope)
	if err != nil {
		return nil, err
	}

	// Create a servicer.
	service, err := storage.New(client)
	if err != nil {
		return nil, err
	}

	// Check the given bucket exists.
	if _, err := service.Buckets.Get(bucket).Do(); err != nil {

		if res, err := service.Buckets.Insert(project, &storage.Bucket{Name: bucket}).Do(); err == nil {
			log.Printf("Bucket %s was created at %s", res.Name, res.SelfLink)
		} else {
			return nil, err
		}

	}

	return &Storage{
		BucketName: bucket,
		client:     client,
		service:    service,
	}, nil

}

// Upload a file to a location.
func (s *Storage) Upload(filename string, location *url.URL) error {

	object := &storage.Object{Name: location.Path[1:]}
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	if res, err := s.service.Objects.Insert(s.BucketName, object).Media(file).Do(); err == nil {
		log.Printf("Object %s was uploaded at %s", res.Name, res.SelfLink)
	} else {
		return err
	}

	return nil

}
