package main

import (
	"bufio"
	"bytes"
	"context"
	_ "embed"
	"flag"
	"fmt"
	"io"
	"os"
	"os/signal"
	"strings"
	"time"

	"atomicgo.dev/cursor"
	"github.com/kagadar/go-life/game"
	"github.com/kagadar/go-loop"
	"github.com/kagadar/go-pipeline/slices"
)

const (
	version = "Alpha 0.0.1"
)

var (
	//go:embed initial_state.txt
	initialState string
	tickrate     = flag.Duration("main_loop_tick_delay", 100*time.Millisecond, "How long to wait between each tick of the main loop.")
)

func readInitialState(r io.Reader) [][]bool {
	var board [][]bool
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		board = append(board, slices.Transform(slices.Filter([]rune(scanner.Text()), func(r rune) bool { return r != ' ' }), func(cell rune) bool {
			return cell == '■'
		}))
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}
	return board
}

func main() {
	board := game.NewMapBoard(readInitialState(bytes.NewBufferString(initialState)))

	fmt.Print("\033[2J")
	fmt.Printf("\033[0;0HGo Life %s", version)
	defer cursor.Show()
	cursor.Hide()
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()
	loop.MainLoop(ctx, *tickrate, func(_ context.Context) {
		fmt.Printf("\033[2;0H%s\n", strings.Join(slices.Transform(board.State(), func(row []bool) string {
			return strings.Join(slices.Transform(row, func(cell bool) string {
				if cell {
					return "■"
				} else {
					return "□"
				}
			}), " ")
		}), "\n"))
		board.Tick()
	})
}
