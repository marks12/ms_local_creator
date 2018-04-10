package commands

import (
	"ms_local_creator/types"
	"errors"
)

// создаем хранилище команд
type commander struct {

	commands 		map[string]func(arguments []types.Arguments) (error, interface{})
}

// инстанс командера как единственный экземпляр командера
var instance *commander

// синглатон командера
func GetInstance() *commander {

	if instance == nil {
		instance = &commander{}   // <--- NOT THREAD SAFE
		instance.commands = make(map[string]func(arguments []types.Arguments) (error, interface{}))
	}

	return instance
}

// добавление множества алиасов
func (cmd *commander) AddCommand(f func(arguments []types.Arguments) (error, interface{}),  names ...string) {

	for _, name := range names {

		cmd.defineCommandCall(f, name)
	}
}

// привязка алиса к выполняемой команде.
// Команда должна на входе принимать массив аргументов,
// а на выходе возвращять ошибку и все что хочет или ничего
func (cmd *commander) defineCommandCall(f func(arguments []types.Arguments) (error, interface{}),  name string) {

	cmd.commands[name] = f

}

// вызываем команду
func (cmd *commander) CallCommand(args []types.Arguments) (error, interface{}) {

	for _, arg := range args {

		function, ok := cmd.commands[arg.Param]
		if ok {
			return function(args)
		} else {
			return errors.New("command not exists"), nil
		}

		return cmd.commands[arg.Param](args)
		break
	}

	return errors.New("command is not set"), nil
}