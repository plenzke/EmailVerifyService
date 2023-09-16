package decorator

import (
	"github.com/sirupsen/logrus"
)

func ApplyCommandDecorators[H any](handler CommandHandler[H], logger *logrus.Entry)

type CommandHandler[C any] interface {
	Han
}
