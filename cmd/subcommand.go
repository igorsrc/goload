package cmd

import (
	"fmt"

	"github.com/goload/domain"
	"github.com/spf13/cobra"
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
				req := &domain.Request{
					Address:    address,
					Count:      count,
					Concurrent: concurrent,
					Backoff:    backoff,
				}
				return Process(req, Get)
			}
			fmt.Println(run())
		},
	}

	cmdPost = &cobra.Command{
		Use:   "post",
		Short: "Process http POST",
		Long:  "Process http POST requests",
		Run: func(cmd *cobra.Command, args []string) {
			run := func() string {
				req := &domain.Request{
					Address:    address,
					Count:      count,
					Concurrent: concurrent,
					Backoff:    backoff,
				}
				return Process(req, Post)
			}
			fmt.Println(run())
		},
	}
)

// config flags
func init() {
	// get
	cmdGet.Flags().StringVarP(
		&address,
		"address", "a",
		"",
		"target api address")
	cmdGet.Flags().IntVarP(
		&count,
		"number", "n",
		defaultNumber,
		"number of requests")
	cmdGet.Flags().IntVarP(
		&concurrent,
		"concurrent", "c",
		defaultConcurrency,
		"concurrency rate")
	cmdGet.Flags().IntVarP(
		&backoff,
		"backoff", "b",
		defaultBackOff,
		"backoff")
	cmdGet.Flags().BoolVarP(
		&debug,
		"debug", "d",
		defaultDebug,
		"debug mode")
	cmdGet.Flags().StringVarP(
		&token,
		"token",
		"t",
		"",
		"token")

	_ = cmdGet.MarkFlagRequired("address")

	// post
	cmdPost.Flags().StringVarP(
		&address,
		"address", "a",
		"",
		"target api address")
	cmdPost.Flags().IntVarP(
		&count,
		"number", "n",
		defaultNumber,
		"number of requests")
	cmdPost.Flags().IntVarP(
		&concurrent,
		"concurrent", "c",
		defaultConcurrency,
		"concurrency rate")
	cmdPost.Flags().IntVarP(
		&backoff,
		"backoff", "b",
		defaultBackOff,
		"backoff")
	cmdPost.Flags().BoolVarP(
		&debug,
		"debug", "d",
		defaultDebug,
		"debug mode")
	cmdPost.Flags().StringVarP(
		&payload,
		"payload",
		"p",
		"",
		"request body")
	cmdPost.Flags().StringVarP(
		&token,
		"token",
		"t",
		"",
		"token")
	_ = cmdPost.MarkFlagRequired("address")
}
