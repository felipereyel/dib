package tui

import (
	"dib/internal/utils"

	tea "github.com/charmbracelet/bubbletea"
)

func Tui(csv_path string) error {
	data, err := utils.ReadCsvFile(csv_path)
	if err != nil {
		return err
	}

	t, err := csvTable(data)
	if err != nil {
		return err
	}

	m := model{t}
	_, err = tea.NewProgram(m).Run()
	return err
}
