package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"goload/service"
	"time"
)

const (
	defaultNumber      = 1
	defaultConcurrency = 1
	defaultBackOff     = 100
	defaultDebug       = false
)

var (
	address    string
	token      string
	payload    string
	count      int
	concurrent int
	backoff    int
	debug      bool

	cmdGet = &cobra.Command{
		Use:   "get",
		Short: "Process http GET",
		Long:  "Process http GET requests",
		Run: func(cmd *cobra.Command, args []string) {
			run := func() string {
				start := time.Now()
				res := service.Get(address, concurrent, backoff, count)
				ms := time.Since(start).Milliseconds()
				msg := "%d/%d in %vms."
				return fmt.Sprintf(msg, len(res), count, ms)
			}
			fmt.Println(run())
		},
	}

	cmdPost = &cobra.Command{
		Use:   "post",
		Short: "Process http POST",
		Long:  "Process http POST requests",
	}

	cmdPut = &cobra.Command{
		Use:   "put",
		Short: "Process http PUT",
		Long:  "Process http PUT requests",
	}
)

func init() {
	// get flags section
	cmdGet.Flags().StringVarP(
		&address,
		"address",
		"a",
		"",
		"target api address")
	cmdGet.Flags().IntVarP(
		&count,
		"number",
		"n",
		defaultNumber,
		"number of requests")
	cmdGet.Flags().IntVarP(
		&concurrent,
		"concurrent",
		"c",
		defaultConcurrency,
		"concurrency rate")
	cmdGet.Flags().IntVarP(
		&backoff,
		"backoff",
		"b",
		defaultBackOff,
		"backoff")
	cmdGet.Flags().BoolVarP(&debug,
		"debug",
		"d",
		defaultDebug,
		"debug mode")

	_ = cmdGet.MarkFlagRequired("address")

	// post flag section
}
