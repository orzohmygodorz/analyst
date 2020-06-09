package query

import (
    "context"
    "fmt"
    "os"
    "time"

    "github.com/prometheus/client_golang/api"
    "github.com/prometheus/client_golang/api/prometheus/v1"

    "../../pkg/value"
)

func Range_Query(metrics_name string,
                 filter string,
                 range_selector string) {
    client, err := api.NewClient(api.Config{
        Address: value.Ip_prometheus_server,
    })
    if err != nil {
        fmt.Printf("Error creating client: %v\n", err)
        os.Exit(1)
    }

    v1api := v1.NewAPI(client)
    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()

    /*
     * Deal with string: metrics name
     */
    target_metrics := metrics_name + "{" + filter + "}"
    if range_selector != "" {target_metrics += "[" + range_selector + "]"}
    fmt.Println(target_metrics)

    r := v1.Range{
        Start: time.Now().Add(-time.Hour),
        End:   time.Now(),
        Step:  time.Minute,
    }
    result, warnings, err := v1api.QueryRange(ctx, "rate("+target_metrics+")", r)
    if err != nil {
        fmt.Printf("Error querying Prometheus: %v\n", err)
        os.Exit(1)
    }
    if len(warnings) > 0 {
        fmt.Printf("Warnings: %v\n", warnings)
    }
    fmt.Printf("Result:\n%v\n", result)
}

