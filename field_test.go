package belajar_golang_logging

import (
	"github.com/sirupsen/logrus"
	"testing"
)

func TestField(t *testing.T) {
	logger := logrus.New()
	logger.SetFormatter(&logrus.JSONFormatter{})

	logger.WithField("username", "khannedy").Info("Hello World")

	logger.WithField("username", "eko").
		WithField("name", "Eko Kurniawan").
		Info("Hello World")
}

func TestFields(t *testing.T) {
	logger := logrus.New()
	logger.SetFormatter(&logrus.JSONFormatter{})

	logger.WithFields(logrus.Fields{
		"username": "eko",
		"name":     "Eko Kurniawan",
	}).Info("Hello World")
}
