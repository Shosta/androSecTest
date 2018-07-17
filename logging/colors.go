package logging

import "github.com/shosta/androSecTest/variables"

func Orange(str string) string {
	return variables.Orange + str + variables.Endc
}

func Green(str string) string {
	return variables.Green + str + variables.Endc
}

func Red(str string) string {
	return variables.Red + str + variables.Endc
}

func Blue(str string) string {
	return variables.Blue + str + variables.Endc
}

func Bold(str string) string {
	return variables.Bold + str + variables.Endc
}
