package utils

import "fmt"

const (
	HelpCommand      = "todo --help"
	HelpKeyword      = "--help"
	QuitCommand      = "todo quit"
	AddCommand       = "todo add"
	ListCommand      = "todo list"
	DeleteCommand    = "todo delete <id>"
	DoneCommand      = "todo done <id>"
	EditCommand      = "todo edit <id> \"<new title>\""
	HelpCommandShort = "todo <cmd> --help"
)

func AvailableCommands() {
	fmt.Println()
	fmt.Println("  📋 TODO CLI — Available Commands")
	fmt.Println("  " + repeat("─", 40))
	fmt.Printf("  %-30s %s\n", HelpCommand, "Show this help message")
	fmt.Printf("  %-30s %s\n", QuitCommand, "Exit the application")
	fmt.Printf("  %-30s %s\n", AddCommand, "Add a new TODO")
	fmt.Printf("  %-30s %s\n", ListCommand, "List all TODOs")
	fmt.Printf("  %-30s %s\n", DeleteCommand, "Delete a TODO by ID")
	fmt.Printf("  %-30s %s\n", DoneCommand, "Mark a TODO as done")
	fmt.Printf("  %-30s %s\n", EditCommand, "Edit a TODO title")
	fmt.Println()
	fmt.Printf("  %-30s %s\n", HelpCommandShort, "Get help for a specific command")
	fmt.Println("  " + repeat("─", 40))
	fmt.Println()
}

func CommandHelp(command string) {
	fmt.Println()
	switch command {
	case "add":
		printCommandHelp(
			AddCommand,
			"Create a new TODO item with a title.",
			"todo add \"<title>\"",
			[]string{
				"todo add \"Buy groceries\"",
				"todo add \"Call Tony at 6pm\"",
			},
		)
	case "delete":
		printCommandHelp(
			DeleteCommand,
			"Permanently remove a TODO by its ID.",
			"todo delete <id>",
			[]string{
				"todo delete 1",
				"todo delete 3",
			},
		)
	case "done":
		printCommandHelp(
			DoneCommand,
			"Mark a TODO as completed.",
			"todo done <id>",
			[]string{
				"todo done 1",
				"todo done 3",
			},
		)
	case "list":
		printCommandHelp(
			ListCommand,
			"Display all TODOs. Optionally filter by status.",
			"todo list [--filter=<status>]",
			[]string{
				"todo list",
				"todo list --filter=pending",
				"todo list --filter=done",
			},
		)
	case "edit":
		printCommandHelp(
			EditCommand,
			"Update the title of an existing TODO by its ID.",
			"todo edit <id> \"<new title>\"",
			[]string{
				"todo edit 1 \"Buy groceries and cook dinner\"",
				"todo edit 3 \"Call Tony at 8pm instead\"",
			},
		)
	default:
		fmt.Println("  ⚠️  Unknown command:", command)
		fmt.Println("  Run 'todo help' to see all available commands.")
		fmt.Println()
	}
}

func printCommandHelp(name, description, usage string, examples []string) {
	fmt.Println("  " + repeat("─", 40))
	fmt.Println("  Command    :", name)
	fmt.Println("  " + repeat("─", 40))
	fmt.Println("  Description:", description)
	fmt.Println()
	fmt.Println("  Usage      :", usage)
	fmt.Println()
	fmt.Println("  Examples")
	for _, example := range examples {
		fmt.Println("    →", example)
	}
	fmt.Println("  " + repeat("─", 40))
	fmt.Println()
}

func repeat(char string, count int) string {
	result := ""
	for i := 0; i < count; i++ {
		result += char
	}
	return result
}
