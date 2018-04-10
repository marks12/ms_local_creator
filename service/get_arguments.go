package service

import (
	"ms_local_creator/types"
	"os"
)

func GetArguments() []types.Arguments {

	result := []types.Arguments{}


	if len(os.Args) > 1 {

		for ind, arg := range os.Args {

			if ind == 0 {
				continue
			}

			argument := types.Arguments{arg, ""}
			result = append(result, argument)
		}
	}


	return result

}