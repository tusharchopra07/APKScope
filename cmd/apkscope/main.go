package main

import (
	"log"

	"github.com/charmbracelet/bubbletea"

	"github.com/<your-username>/apkscope/internal/app"
)

func main() {
	p := bubbletea.NewProgram(
		app.New(),
		bubbletea.WithAltScreen(),
	)

	if _, err := p.Run(); err != nil {
		log.Fatal(err)
	}
}
