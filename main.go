package main

import (
	"bufio"
	"context"
	"embed"
	"flag"
	"fmt"
	"os"
	"os/signal"
	"time"

	"atomicgo.dev/cursor"
	"github.com/kagadar/go-life/game"
	"github.com/kagadar/go-loop"
	"github.com/kagadar/go-pipeline/slices"
)

const (
	version = "Alpha 0.0.2"
)

var (
	//go:embed initial_states/*
	states   embed.FS
	tickrate = flag.Duration("delay", 250*time.Millisecond, "How long to wait between each tick of the main loop.")
	width    = flag.Int("width", 64, "Width of the play area when using a random seed")
	height   = flag.Int("height", 32, "Height of the play area when using a random seed")
	chance   = flag.Float64("chance", 0.05, "Initial chance for each cell to be alive")
	file     = flag.String("file", "", "If set, will load the named playfield from `initial_states` rather than generating one randomly")
)

func loadState(name string) (state game.State) {
	f, err := states.Open(fmt.Sprintf("initial_states/%s.txt", name))
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		state = append(state, slices.Transform(slices.Filter([]rune(scanner.Text()), func(r rune) bool { return r != ' ' }), func(cell rune) bool {
			return cell == '■'
		}))
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}
	return state
}

func init() {
	flag.Parse()
}

func main() {
	var state game.State
	if *file != "" {
		state = loadState(*file)
	} else {
		state = game.RandomState(*width, *height, *chance)
	}
	board := game.NewSlice1_Unroll(state)
	// Clear terminal
	fmt.Printf("\033[2J\033[0;0HGo Life %s", version)
	defer cursor.Show()
	cursor.Hide()
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()
	l := loop.New(*tickrate)
	l.EnableStats(true)
	l.Run(ctx, func(ctx context.Context) {
		stats := l.Stats()
		fmt.Printf(`%s%s%s

Loop running for %s
Avg tick: %s
Max tick: %s
Min tick: %s
Memory usage: %s`,
			// Move cursor after title
			"\033[2;0H",
			board.String(),
			// Clear old debug stats from terminal
			"\033[0J",
			stats.Duration(),
			stats.AvgTick(),
			stats.MaxTick(),
			stats.MinTick(),
			stats.HeapAllocFmt(),
		)
		board.Tick()
	})
}
