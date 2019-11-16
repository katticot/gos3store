package main

import "context"
type fileType int
const (
 PDF  fileType =iota +1
 DOCX
 PPT
 JPEG
)
type File struct {
	fileBinary byte
	fileType fileType
	fileName	string
}

type FileStorageService interface {
	Create(ctx context.Context, file File) (File, error)
	Get(ctx context.Context, id string) (File, error)
	Delete(ctx context.Context, id string) (string, error)
}

type ImplementedFileStoreService struct {


}
// Repository describes the persistence on order model
func main() {

	
}
