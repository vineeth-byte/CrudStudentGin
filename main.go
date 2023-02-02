package main

import (
	db "task/db"
	rpc "task/rpc"
)

func main() {
	db.Connect()
	rpc.Connect()
}
