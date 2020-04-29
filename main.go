package main

import (
	collections_uc "learning/collections"
	if_uc "learning/conditionals/ifstatement"
	switch_uc "learning/conditionals/switchstatement"
	dpc_uc "learning/defer-panic-recover"
	functions_uc "learning/functions"
	iota_uc "learning/iotauc"
	structs_uc "learning/structs"
)

func main() {
	iota_uc.Run()
	collections_uc.Run()
	structs_uc.Run()
	if_uc.Run()
	switch_uc.Run()
	dpc_uc.Run()
	functions_uc.Run()
}
