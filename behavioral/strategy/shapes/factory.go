package shapes

import (
	"fmt"
	"os"

	"github.com/MrGru/GoDesignPattern/behavioral/strategy"
)

const (
	TEXT_STRATEGY  = "text"
	IMAGE_STRATEGY = "image"
)

func Factory(s string) (strategy.Output, error) {
	switch s {
	case TEXT_STRATEGY:
		return &TextSquare{
			DrawOutput: strategy.DrawOutput{
				LogWriter: os.Stdout,
			},
		}, nil
	case IMAGE_STRATEGY:
		return &ImageSquare{
			DrawOutput: strategy.DrawOutput{
				LogWriter: os.Stdout,
			},
		}, nil
	default:
		return nil, fmt.Errorf("Strategy '%s' not found\n", s)
	}

}
