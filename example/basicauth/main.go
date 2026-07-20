package main

import (
	"bufio"
	"context"
	"errors"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/google/go-github/v89/github"
	"golang.org/x/term"
)

func main() {
	r := bufio.NewReader(os.Stdin)
	fmt.Print("GitHub Username: ")
	username, _ := r.ReadString('\n')

	fmt.Print("GitHub Password: ")
	password, _ := term.ReadPassword(int(os.Stdin.Fd()))

	tp := github.BasicAuthTransport{
		Username: strings.TrimSpace(username),
		Password: strings.TrimSpace(string(password)),
	}

	client, err := github.NewClient(github.WithHTTPClient(tp.Client()))
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	ctx := context.Background()
	user, _, err := client.Users.Get(ctx, "")

	if errors.As(err, new(*github.TwoFactorAuthError)) {
		fmt.Print("\nGitHub OTP: ")
		otp, _ := r.ReadString('\n')
		tp.OTP = strings.TrimSpace(otp)
		user, _, err = client.Users.Get(ctx, "")
	}

	if err != nil {
		log.Fatalf("error: %v", err)
	}

	fmt.Printf("\n%v\n", github.Stringify(user))
}
