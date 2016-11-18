package easel

import (
	"errors"
	"fmt"

	log "github.com/Sirupsen/logrus"

	"github.com/go-gl/gl/v4.1-core/gl"
)

func checkGLError(msg string) error {
	errID := gl.GetError()
	if errID != gl.NO_ERROR {
		errStr := fmt.Sprintf("%s: %08x", msg, errID)
		log.Error(errStr)
		return errors.New(errStr)
	}
	return nil
}
