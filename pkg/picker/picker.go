package picker

import (
	"errors"
	"fmt"
	"strings"

	fzf "github.com/junegunn/fzf/src"
	"github.com/nobbmaestro/tmux-tether/pkg/config"
)

var ErrInterrupted = errors.New("picker: interrupted")

type Picker[T any] struct {
	cfg   config.PickerConfig
	label func(T) string
}

func New[T any](
	cfg config.PickerConfig,
	label func(T) string,
) Picker[T] {
	return Picker[T]{cfg: cfg, label: label}
}

func (p Picker[T]) Pick(items []T) (T, error) {
	var zero T

	if len(items) == 0 {
		return zero, fmt.Errorf("no items to pick")
	}

	labels := make([]string, len(items))
	for i, item := range items {
		labels[i] = p.label(item)
	}

	inputChan := make(chan string, len(labels))
	for _, l := range labels {
		inputChan <- l
	}
	close(inputChan)

	outputChan := make(chan string, 1)

	options, err := fzf.ParseOptions(
		true,
		[]string{
			"--no-sort",
			"--cycle",
			"--track",
			"--exact",
			"--pointer=" + p.cfg.Pointer,
			"--prompt=" + p.cfg.Prompt,
		},
	)
	if err != nil {
		return zero, fmt.Errorf("pick: %w", err)
	}

	options.Input = inputChan
	options.Output = outputChan

	code, err := fzf.Run(options)
	if err != nil {
		return zero, fmt.Errorf("pick: %w", err)
	}

	if code == fzf.ExitInterrupt {
		return zero, ErrInterrupted
	}

	if code != fzf.ExitOk {
		return zero, fmt.Errorf("pick: exit code %d", code)
	}

	selected := strings.TrimSpace(<-outputChan)
	for i, l := range labels {
		if l == selected {
			return items[i], nil
		}
	}

	return zero, fmt.Errorf("pick: selected item not found")
}
