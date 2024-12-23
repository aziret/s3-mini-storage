package converter

import (
	"github.com/aziret/s3-mini-internal/pkg/api/filetransfer_v1"
	"github.com/aziret/s3-mini-storage/internal/model"
)

func ToFileFromApi(req *filetransfer_v1.FileChunkUpload) *model.File {
	return &model.File{
		Data: req.GetData(),
		UUID: req.GetUuid(),
	}
}
