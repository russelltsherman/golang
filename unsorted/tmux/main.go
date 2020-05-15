package main

import (
	"fmt"
	"os"
	"os/exec"
	"os/signal"
	"syscall"
	"time"
)

func init() {
	fmt.Println("Initializing application data")
}

func main() {
	cmd := exec.Command("/bin/sh", "-c", "tmux new -s foo;")
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	go func(c *exec.Cmd) {
		err := c.Run()
		if err != nil {
			fmt.Println(err)
		}
	}(cmd)

	tmuxLayout := `
    tmux send-keys -t foo.0 "tmux split-window -v -p 5" ENTER;
    tmux send-keys -t foo.0 "tmux selectp -t 0" ENTER;
    tmux send-keys -t foo.0 "tmux split-window -v -p 95" ENTER;
    tmux send-keys -t foo.0 "tmux selectp -t 1" ENTER;
    tmux send-keys -t foo.0 "tmux split-window -h -p 85" ENTER;
    tmux send-keys -t foo.0 "tmux selectp -t 2" ENTER;
    tmux send-keys -t foo.0 "tmux split-window -v -p 70" ENTER;
    tmux send-keys -t foo.0 "tmux selectp -t 2" ENTER;
    tmux send-keys -t foo.0 "tmux split-window -h -p 33" ENTER;
    tmux send-keys -t foo.0 "tmux selectp -t 2" ENTER;
    tmux send-keys -t foo.0 "tmux split-window -h -p 50" ENTER;

    tmux send-keys -t foo.0 "tmux selectp -t 0" ENTER;
    tmux send-keys -t foo.0 "clear" ENTER;
    tmux send-keys -t foo.0 "tmux selectp -t 1" ENTER;
    tmux send-keys -t foo.0 "clear" ENTER;
    tmux send-keys -t foo.0 "tmux selectp -t 2" ENTER;
    tmux send-keys -t foo.0 "clear" ENTER;
    tmux send-keys -t foo.0 "tmux selectp -t 3" ENTER;
    tmux send-keys -t foo.0 "clear" ENTER;
    tmux send-keys -t foo.0 "tmux selectp -t 4" ENTER;
    tmux send-keys -t foo.0 "clear" ENTER;
    tmux send-keys -t foo.0 "tmux selectp -t 5" ENTER;
    tmux send-keys -t foo.0 "clear" ENTER;
    tmux send-keys -t foo.0 "tmux selectp -t 6" ENTER;
    tmux send-keys -t foo.0 "clear" ENTER;

  `

	<-time.After(1 * time.Second)
	cmd2 := exec.Command("/bin/sh", "-c", tmuxLayout)
	err := cmd2.Run()
	if err != nil {
		fmt.Println(err)
	}

	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM, syscall.SIGTSTP, syscall.SIGQUIT)

	for {
		<-c
		fmt.Println("killing app")
		break
	}
}
