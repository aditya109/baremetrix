package helper

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"bitbucket.org/exotel/ent-leadassist-test/baremetrix/internal/models"
	logger "github.com/sirupsen/logrus"
)

// GetAbsolutePath provides absolute path for a relative path -
func GetAbsolutePath(relPath string) (string, error) {
	if relPath[0] == '/' {
		relPath = relPath[1:]
	}
	cwd, err := os.Getwd()
	var path string
	if err != nil {
		logger.Error(err)
		return "", err
	}
	projectLocation := strings.Split(cwd, "baremetrix")
	path = filepath.Join(projectLocation[0], "baremetrix", relPath)
	return path, nil
}

// GetFormattedFileName gets a formatted filename
func GetFormattedFileName(directives models.SummarySpecificDirectives) (string, error) {
	path := directives.FileSpecs.ContainerDirectory
	absolutePath, err := GetAbsolutePath(path)
	absolutePath = strings.ReplaceAll(absolutePath, "<DIFFERENTIATOR_STAMP>", directives.TimeStamp)
	if directives.ShouldUseDirectives {
		absolutePath = fmt.Sprintf("%s/%s/%s/iteration-%s/", absolutePath, directives.Tenant, directives.PlayName, directives.Iteration)
	}
	if err != nil {
		logger.Error(err)
		return "", err
	}
	if _, err := os.Stat(absolutePath); errors.Is(err, os.ErrNotExist) {
		err := os.MkdirAll(absolutePath, os.ModePerm)
		if err != nil {
			logger.Error(err)
			return "", err
		}
	}
	var fileName = fmt.Sprintf("%s/%s", absolutePath, directives.FileSpecs.Name[0])
	if directives.ShouldUseGraphIndicator {
		fileName += fmt.Sprintf("_%s", directives.GraphType)
	}
	fileName += directives.FileSpecs.Extension
	return fileName, nil
}
