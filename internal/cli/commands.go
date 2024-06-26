package cli

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"text/tabwriter"

	"github.com/fabian-gubler/pokedexcli/pkg/config"
)

type cliCommand struct {
	name        string
	description string
	callback    func(*config.Config) error
}


func (c *cliCommand) execute(cfg *config.Config) error {
	return c.callback(cfg)
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

func commandHelp(_ *config.Config) error {
	// Create a tab writer to format the output as a table
	writer := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', 0)

	// Print header
	fmt.Fprintln(writer, "Command\tDescription")

	// Print all commands in table format
	for _, cmd := range initializeCommands() {
		fmt.Fprintf(writer, "%s\t%s\n", cmd.name, cmd.description)
	}

	// Flush the writer to ensure all data is written
	writer.Flush()
	return nil
}

func commandExit(_ *config.Config) error {
	fmt.Println("Exiting Pokedex")
	return nil
}

func commandMap(cfg *config.Config) error {

	resp, err := cfg.PokeapiClient.ListLocationAreas(cfg.NextLocationURL)

	if err != nil {
		log.Fatal(err)
	}

	// Print results
	fmt.Println("Locations:")
	for _, area := range resp.Results {
		fmt.Printf(" - %s\n", area.Name)
	}

	cfg.NextLocationURL = resp.Next
	cfg.PreviousLocationURL = resp.Previous

	return nil
}

func commandMapb(cfg *config.Config) error {

	if cfg.PreviousLocationURL == nil {
		return errors.New("you're on the first page")
	}

	resp, err := cfg.PokeapiClient.ListLocationAreas(cfg.PreviousLocationURL)

	if err != nil {
		log.Fatal(err)
	}

	// Print results
	fmt.Println("Locations:")
	for _, area := range resp.Results {
		fmt.Printf(" - %s\n", area.Name)
	}

	cfg.NextLocationURL = resp.Next
	cfg.PreviousLocationURL = resp.Previous

	return nil
}

func readInput(prompt string) string {
	fmt.Print(prompt)
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	return scanner.Text()
}

func RunCLI(cfg *config.Config) {
	commands := initializeCommands()

	for {
		command := readInput("Pokedex > ")

		cmd, found := commands[command]
		if !found {
			fmt.Println("Command not found")
			continue
		}

		err := cmd.execute(cfg)
		if err != nil {
			fmt.Println("Error executing command:", err)
		}

		if command == "exit" {
			break
		}
	}
}
