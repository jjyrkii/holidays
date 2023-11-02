package model

import "os"

type ExportInterface interface {
	Export() (*os.File, error)
}
