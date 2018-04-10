package main

import (
	"ms_local_creator/commands"
	//"ms_local_creator/service"
	//"fmt"
	"ms_local_creator/service"
	"fmt"
)

func main() {

	cmd := commands.GetInstance()

	cmd.AddCommand(commands.CreateProject,		"create-project", "cp")
	cmd.AddCommand(commands.CreateDatabase,		"create-db", "create-database", "cdb")

	arguments := service.GetArguments()

	err, res := cmd.CallCommand(arguments)

	if(err != nil) {
		fmt.Println("Error: ", res)
	}

	if(res != nil) {
		fmt.Println("Command executed. ", res)
	}


	return
}
