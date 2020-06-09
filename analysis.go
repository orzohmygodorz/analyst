package main

import (
    "./api/v1"
)

func main() {
    //queries.Instant_Query("container_network_receive_bytes_total", "", "2m")
    query.Range_Query("container_network_receive_bytes_total", "", "2m")
}


