package poker

import (
	"bufio"
	"io"
	"strings"
	"time"
)

type CLI struct {
	playerStore PlayerStore
	in          *bufio.Scanner
	alerter     BlindAlerter
}

func NewCLI(store PlayerStore, in io.Reader, alerter BlindAlerter) *CLI {
	return &CLI{
		playerStore: store,
		in:          bufio.NewScanner(in),
		alerter:     alerter,
	}
}

func (cli *CLI) PlayPoker() {
	userInput := cli.ReadLine()
	cli.playerStore.RecordWin(extractWinner(userInput))
	cli.alerter.ScheduleAlertAt(5*time.Second, 100)
}

func extractWinner(userInput string) string {
	return strings.Replace(userInput, " wins", "", 1)
}

func (cli *CLI) ReadLine() string {
	cli.in.Scan()
	return cli.in.Text()
}

type BlindAlerter interface {
	ScheduleAlertAt(duration time.Duration, amount int)
}
