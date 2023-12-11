package apipgu

import (
	"archive/zip"
	"bytes"
	"fmt"
)

// PushAttachment - файл вложения к создаваемому заявлению (см [PushArchive])
type PushAttachment struct {
	Filename string // Имя файла с расширением. Пример: "req_346ee59c-a428-42f6-342e-c780dd2e278e.xml"
	Data     []byte // Содержимое файла
}

// PushArchive - архив вложений к создаваемому заявлению.
// Используется для методов [Client.OrderPush] и [Client.OrderPushChunked].
type PushArchive struct {
	Name  string           // Имя архива (без расширения). Пример: "35002123456-archive"
	Files []PushAttachment // Файлы вложений
}

// Zip - формирует zip-архив из файлов вложений.
//
// В случае успеха возвращает байты архива.
// В случае ошибки возвращает [ErrZipCreate], [ErrZipWrite] или [ErrZipClose].
func (a *PushArchive) Zip() ([]byte, error) {
	if len(a.Files) == 0 {
		return nil, ErrNoFiles
	}

	var b bytes.Buffer
	zipWriter := zip.NewWriter(&b)

	for _, file := range a.Files {
		fileWriter, err := zipWriter.Create(file.Filename)
		if err != nil {
			return nil, fmt.Errorf("%w: %w", ErrZipCreate, err)
		}
		if _, err = fileWriter.Write(file.Data); err != nil {
			return nil, fmt.Errorf("%w: %w", ErrZipWrite, err)
		}
	}

	if err := zipWriter.Close(); err != nil {
		return nil, fmt.Errorf("%w: %w", ErrZipClose, err)
	}

	return b.Bytes(), nil
}
