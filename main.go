package main

import (
    "bearLinks/controller"
    "bearLinks/datastore"
    "bearLinks/service"
)

func main() {

    datastore.InitRedisClient()

    service.InitAnalytics()

    controller.Init()
}