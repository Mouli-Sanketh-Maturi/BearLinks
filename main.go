package main

import (
    "bearLinks/controller"
    "bearLinks/datastore"
)

func main() {

    datastore.InitRedisClient()

    controller.Init()
}