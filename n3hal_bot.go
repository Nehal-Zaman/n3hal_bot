package main

import (
	"bytes"
	"flag"
	"fmt"
	"os"
	"os/exec"
	"strings"
	"sync"

	"github.com/Nehal-Zaman/n3hal_bot/colors"
	"github.com/NicoNex/echotron/v3"
	"github.com/joho/godotenv"
)

var wg sync.WaitGroup

const start_msg = `Hi, this is ` + "`n3hal_bot`" + `.
This bot is created by Nehal Zaman (` + "`n3hal_`" + `) for recon automation.

1. ` + "`run <passcode> <cmd>`" + `: to run a command.
2. ` + "`subenum <passcode> <target>`" + `: start a subdomain enumneration of target.
`

func RunCommand(passcode string, server_passcode string, cmd string) string {
	if server_passcode != passcode {
		return "Invalid passcode"
	} else {
		cmd_to_run := exec.Command("bash", "-c", cmd)

		var stdoutBuff bytes.Buffer
		cmd_to_run.Stdout = &stdoutBuff

		err := cmd_to_run.Run()
		if err != nil {
			return err.Error()
		}

		return stdoutBuff.String()
	}
}

func printToStdout(user string, msg string) {
	fmt.Println(colors.WhiteBold("Message from ") + colors.GreenBold(user) + colors.WhiteBold(": ") + colors.YellowBold(msg))
}

func printBanner() {
	fmt.Println(colors.CyanBold("n3hal_bot") + ":" + colors.WhiteBold(" telegram bot for automated recon"))
}

func getCliArgs() int {
	threadsPtr := flag.Int("threads", 10, "specify number of threads to utilize")

	flag.Parse()

	return *threadsPtr
}

func main() {
	printBanner()

	err := godotenv.Load()
	if err != nil {
		panic(err)
	}

	token := os.Getenv("BOT_KEY")
	run_key := os.Getenv("RUN_KEY")
	scripts_path := os.Getenv("SCRIPTS_PATH")
	threads := getCliArgs()

	api := echotron.NewAPI(token)
	msgChan := make(chan *echotron.Update)

	for i := 0; i < threads; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for u := range msgChan {
				analyzeMessage(u, run_key, scripts_path, api)
			}
		}()
	}

	for u := range echotron.PollingUpdates(token) {
		msgChan <- u
	}

	close(msgChan)
	wg.Wait()

}

func analyzeMessage(u *echotron.Update, run_key string, scripts_path string, api echotron.API) {
	printToStdout(u.Message.From.Username, u.Message.Text)

	if u.Message.Text == "/start" {
		api.SendMessage(start_msg, u.ChatID(), &echotron.MessageOptions{ParseMode: "Markdown"})
	}

	// to run a command on running server
	if strings.HasPrefix(u.Message.Text, "run ") {
		args := strings.Split(u.Message.Text, " ")
		if len(args) < 3 {
			api.SendMessage("Invalid number of arguments to 'run'", u.ChatID(), nil)
		} else {
			cmd := strings.Join(args[2:], " ")
			api.SendMessage(RunCommand(args[1], run_key, cmd), u.ChatID(), nil)
		}
	}

	// to start a subdomain enumeration of a target
	if strings.HasPrefix(u.Message.Text, "subenum ") {
		args := strings.Split(u.Message.Text, " ")
		if len(args) < 3 {
			api.SendMessage("Invalid number of arguments to 'subenum'", u.ChatID(), nil)
		} else {
			cmd := scripts_path + "/subenum/subenum.sh " + args[2]
			api.SendMessage("Target "+args[2]+" is added for recon by "+u.Message.From.Username, u.ChatID(), nil)
			output := fmt.Sprintf("**Subdomains discovered for __%v__:**\n\n```\n%v\n```", args[2], RunCommand(args[1], run_key, cmd))
			api.SendMessage(output, u.ChatID(), &echotron.MessageOptions{ParseMode: "Markdown"})
		}
	}
}
