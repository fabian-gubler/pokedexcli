package cli

import (
	"bufio"
	"fmt"
	"os"

	"github.com/fabian-gubler/pokedexcli/internal/api"
)

type cliCommand struct {
	name        string
	description string
	callback    func() error
}

type config struct {
	nextURL     string
	previousURL string
	apiClient   *api.PokeAPIClient
}

func (c *cliCommand) execute() error {
	return c.callback()
}

func initializeCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"map": {
			name:        "map",
			description: "Displays the next 20 locations",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Displays the previous 20 locations",
			callback:    commandMapb,
		},
	}
}

func commandHelp() error {
	fmt.Println("Available commands:")
	fmt.Println("help - Displays a help message")
	fmt.Println("exit - Exit the Pokedex")
	return nil
}

func commandExit() error {
	fmt.Println("Exiting Pokedex")
	return nil
}

func commandMap() error {
	fmt.Println("Displaying the next 20 locations")
	return nil
}

func commandMapb() error {
	fmt.Println("Displaying the previous 20 locations")
	return nil
}

func readInput(prompt string) string {
	fmt.Print(prompt)
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	return scanner.Text()
}

func RunCLI() {
	commands := initializeCommands()

	for {
		command := readInput("Pokedex > ")

		cmd, found := commands[command]
		if !found {
			fmt.Println("Command not found")
			continue
		}

		err := cmd.execute()
		if err != nil {
			fmt.Println("Error executing command:", err)
		}

		if command == "exit" {
			break
		}
	}
}
