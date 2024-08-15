package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/spf13/cobra"

	"github.com/freesrz93/ask-gpt/consts"
)

var (
	showConfig   bool
	editConfig   bool
	listSessions bool
	listRoles    bool
	interactive  bool
	showHistory  bool
	showVersion  bool
	sessionName  string
	roleName     string
	backendName  string
	newRole      bool
)

var rootCmd = &cobra.Command{
	Use:   "ag [FLAGS] <PROMPT>",
	Short: "Ask GPT through CLI",
	Long:  "",
	Run:   root,
}

func root(cmd *cobra.Command, args []string) {
	var err error
	switch {
	case showVersion:
		P(consts.VerInfo)
	case showConfig:
		P(Config.String())
	case editConfig:
		err = EditCfg()
	case listSessions:
		for _, s := range ListSessions() {
			P(s.ID)
			Pln()
		}
	case listRoles:
		for name, r := range Config.Roles {
			P(fmt.Sprintf("name: %s\ndescription: %s\nprompt: %s\n", name, r.Description, r.Prompt))
			Pln()
		}
	case newRole:
		createRole()
	default:
		err = handleSession(cmd, args)
	}

	if err != nil {
		PFatal(err)
	}
}

func handleSession(_ *cobra.Command, args []string) error {
	c, s, err := initClient()
	if err != nil {
		return err
	}

	if showHistory {
		P(s.String())
	}

	input := strings.Join(args, " ")
	stat, _ := os.Stdin.Stat()
	if (stat.Mode() & os.ModeCharDevice) == 0 {
		pipe, err := io.ReadAll(os.Stdin)
		if err == nil {
			input = string(pipe) + "\n" + input
		}
	}
	if input != "" {
		if interactive {
			P(AIPrefix)
		}
		err = c.Stream(s, input)
		if err != nil {
			return err
		}
	}

	if interactive {
		interactiveMode(c, s, input)
		return nil
	}
	return nil
}

func initClient() (*Client, *Session, error) {
	if backendName == "" {
		backendName = Config.DefaultBackend
	}
	opt, ok := Config.Backends[backendName]
	if !ok {
		return nil, nil, errors.New("backend not exist")
	}
	if roleName == "" {
		roleName = opt.DefaultRole
	}
	r, err := GetRole(roleName)
	if err != nil {
		return nil, nil, err
	}
	client := NewClient(opt)
	s, err := GetSession(sessionName)
	if err != nil {
		return nil, nil, err
	}
	s.UseRole(r)
	return client, s, nil
}

func createRole() {
	scanner := bufio.NewScanner(os.Stdin)
	P("Please input role name:\n")
	scanner.Scan()
	name := scanner.Text()
	P("Role description:\n")
	scanner.Scan()
	desc := scanner.Text()
	P("Role system prompt:\n")
	scanner.Scan()
	prompt := scanner.Text()
	_, err := GetRole(name)
	if err == nil {
		P("role with this name already exists, overwrite? [Y/n] ")
		scanner.Scan()
		answer := scanner.Text()
		if answer == "n" || answer == "N" {
			return
		}
	}
	err = CreateRole(name, desc, prompt)
	if err != nil {
		PFatal(err)
	}
	P("Role " + name + " created!")
}

func interactiveMode(client *Client, s *Session, input string) {
	scanner := bufio.NewScanner(os.Stdin)
	P(UserPrefix)
	for scanner.Scan() {
		input = scanner.Text()
		if input == "exit" || input == "q" {
			break
		}
		P(AIPrefix)
		err := client.Stream(s, input)
		if err != nil {
			PFatal(err)
		}
		P(UserPrefix)
	}
	return
}

func init() {
	rootCmd.PersistentFlags().BoolVarP(&showConfig, "config", "c", false, "Show the configuration.")
	rootCmd.PersistentFlags().BoolVarP(&editConfig, "edit-config", "e", false, "Open config file in an editor.")
	rootCmd.PersistentFlags().BoolVarP(&listSessions, "list-sessions", "l", false, "List all sessions.")
	rootCmd.PersistentFlags().BoolVar(&listRoles, "list-roles", false, "List all roles.")
	rootCmd.PersistentFlags().BoolVarP(&interactive, "interactive", "i", false, "Use interactive mode. (default: false)")
	rootCmd.PersistentFlags().BoolVarP(&showHistory, "history", "h", false, "Show session history. (default: false)")
	rootCmd.PersistentFlags().StringVarP(&sessionName, "session", "s", tempSession, "Create or retrieve a session. If not set, create a temp session that won't be saved.")
	rootCmd.PersistentFlags().StringVarP(&roleName, "role", "r", "", "Specify the role to be used. Only valid for a new or temp session.")
	rootCmd.PersistentFlags().StringVarP(&backendName, "backend", "b", "", "Specify the backend to be used.")
	rootCmd.PersistentFlags().BoolVarP(&newRole, "new-role", "n", false, "Create a new role.")
	rootCmd.PersistentFlags().BoolP("help", "", false, "Show command usage.")
	rootCmd.PersistentFlags().BoolVarP(&showVersion, "version", "v", false, "Show app version.")

	if err := os.MkdirAll(CfgDir, os.ModePerm); err != nil {
		PFatal(err)
	}
	if err := os.MkdirAll(SessionDir, os.ModePerm); err != nil {
		PFatal(err)
	}
	LoadCfg()
}

func main() {
	err := rootCmd.Execute()
	if err != nil {
		PFatal(err)
	}
}
