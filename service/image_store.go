package service

import (
	"bytes"
	"fmt"
	"os"
	"sync"

	"github.com/google/uuid"
)

type ImageStore interface {
	Save(LaptopID string, imagType string, imageData bytes.Buffer) (string, error)
}

type DiskImageStore struct {
	mutex       sync.RWMutex
	imagefolder string
	images      map[string]*ImageInfo
}

type ImageInfo struct {
	LaptopID string
	Type     string
	Path     string
}

func NewDiskImageStore(imageFolder string) *DiskImageStore {
	return &DiskImageStore{
		imagefolder: imageFolder,
		images:      make(map[string]*ImageInfo),
	}
}

func (store *DiskImageStore) Save(LaptopID string, imagType string, imageData bytes.Buffer) (string, error) {
	imageID, err := uuid.NewRandom()

	if err != nil {
		return "", fmt.Errorf("cannot create uuid")
	}

	imagepath := fmt.Sprintf("%s/%s%s", store.imagefolder, imageID, imagType)
	file, err := os.Create(imagepath)
	if err != nil {
		return "", fmt.Errorf("Error creatig ImageFile %w", err)
	}

	_, err = imageData.WriteTo(file)
	if err != nil {
		return "", fmt.Errorf("Error writing ImageFile %w", err)
	}
	store.mutex.Lock()
	defer store.mutex.Unlock()
	store.images[imageID.String()] = &ImageInfo{
		LaptopID: LaptopID,
		Type:     imagType,
		Path:     imagepath,
	}

	return imageID.String(), nil
}
