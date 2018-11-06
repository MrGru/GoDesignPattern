package shapes

import (
	"github.com/MrGru/GoDesignPattern/behavioral/strategy"
)

type TextSquare struct {
	strategy.DrawOutput
}

func (t *TextSquare) Draw() error {
	t.Writer.Write([]byte("Circle"))
	return nil
}
