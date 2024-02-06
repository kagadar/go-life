package main

import (
	"bufio"
	"context"
	"flag"
	"fmt"
	"io"
	"os"
	"os/signal"
	"strings"
	"time"

	"atomicgo.dev/cursor"
	"github.com/kagadar/go-loop"
	"github.com/kagadar/go-pipeline/slices"
)

var (
	tickrate = flag.Duration("main_loop_tick_delay", 16*time.Millisecond, "How long to wait between each tick of the main loop.")
)

func readInitialState(r io.Reader) [][]bool {
	var board [][]bool
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		board = append(board, slices.Transform([]rune(scanner.Text()), func(r rune) bool {
			return r == '■'
		}))
	}
	return board
}

func main() {
	board := readInitialState(os.Stdin)

	fmt.Print("\033[2J")
	defer cursor.Show()
	cursor.Hide()
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()
	loop.MainLoop(ctx, *tickrate, func(_ context.Context) {
		fmt.Printf("\033[0;0H%s\n", strings.Join(slices.Transform(board, func(row []bool) string {
			return string(slices.Transform(row, func(cell bool) rune {
				if cell {
					return '■'
				} else {
					return '□'
				}
			}))
		}), "\n"))
	})
}
