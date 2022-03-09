package compress

import (
	"archive/zip"
	"bytes"
	"io"
)

type File struct {
	Name string
	Data *[]byte
}

func GenerateZipFromByte(files *[]File) (*bytes.Buffer, error) {
	var zipFile *bytes.Buffer

	zipWriter := zip.NewWriter(zipFile)

	for _, file := range *files {
		ioWriter, err := zipWriter.Create(file.Name)
		if err != nil {
			return nil, err
		}
		_, err = io.Copy(ioWriter, bytes.NewReader(*file.Data))
		if err != nil {
			return nil, err
		}
	}

	err := zipWriter.Close()
	if err != nil {
		return nil, err
	}

	return zipFile, nil
}
