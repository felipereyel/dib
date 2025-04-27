package tui

import (
	"errors"

	"github.com/charmbracelet/bubbles/table"
	"github.com/charmbracelet/lipgloss"
)

// "Track URI","Track Name","Artist URI(s)","Artist Name(s)","Album URI","Album Name","Album Artist URI(s)","Album Artist Name(s)","Album Release Date","Album Image URL","Disc Number","Track Number","Track Duration (ms)","Track Preview URL","Explicit","Popularity","ISRC","Added By","Added At"

func csvTable(data [][]string) (table.Model, error) {
	columns := []table.Column{
		{Title: "Track Name", Width: 70},
		{Title: "Album Name", Width: 50},
		{Title: "Artist Name(s)", Width: 40},
		{Title: "Album Release Date", Width: 20},
	}

	firstRow := data[0]
	headersIdx := []int{}
	for _, column := range columns {
		found := false
		for idx, cell := range firstRow {
			if cell == column.Title {
				headersIdx = append(headersIdx, idx)
				found = true
				break
			}
		}

		if !found {
			return table.Model{}, errors.New("could not match required headers")
		}
	}

	rows := []table.Row{}
	for _, row := range data[1:] {
		tableRow := table.Row{}
		for _, v := range headersIdx {
			tableRow = append(tableRow, row[v])
		}

		rows = append(rows, tableRow)
	}

	t := table.New(
		table.WithColumns(columns),
		table.WithRows(rows),
		table.WithFocused(true),
		table.WithHeight(10),
	)

	s := table.DefaultStyles()

	s.Header = s.Header.
		BorderStyle(lipgloss.NormalBorder()).
		BorderForeground(lipgloss.Color("240")).
		BorderBottom(true).
		Bold(false)

	s.Selected = s.Selected.
		Foreground(lipgloss.Color("229")).
		Background(lipgloss.Color("57")).
		Bold(false)

	t.SetStyles(s)

	return t, nil
}
