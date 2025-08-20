// SPDX-FileCopyrightText: 2025 Jacques Supcik <jacques.supcik@hefr.ch>
//
// SPDX-License-Identifier: MIT OR Apache-2.0

package cmd

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/spf13/cobra"
)

func wait(url string, timeout time.Duration, interval time.Duration) error {
	client := &http.Client{}
	start := time.Now()
	last := time.Now().Add(-interval)
	for time.Since(start) < timeout {
		now := time.Now()
		if now.Sub(last) < interval {
			time.Sleep(interval - now.Sub(last))
		}
		last = time.Now()

		resp, err := client.Get(url)
		if err != nil {
			continue
		}
		err = resp.Body.Close()
		if err != nil {
			continue
		}

		if resp.StatusCode == http.StatusOK {
			return nil
		}
	}
	return fmt.Errorf("timed out waiting for %s to become available", url)
}

var rootCmd = &cobra.Command{
	Use:   "webwait <url>",
	Short: "Wait for web server response with timeout",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		timeout, err := time.ParseDuration(cmd.Flag("timeout").Value.String())
		if err != nil {
			fmt.Fprintf(os.Stderr, "Invalid timeout: %v\n", err)
			os.Exit(1)
		}
		interval, err := time.ParseDuration(cmd.Flag("interval").Value.String())
		if err != nil {
			fmt.Fprintf(os.Stderr, "Invalid interval: %v\n", err)
			os.Exit(1)
		}
		err = wait(args[0], timeout, interval)
		if err != nil {
			os.Exit(1)
		}
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().StringP("timeout", "t", "1m", "Timeout duration for waiting")
	rootCmd.Flags().StringP("interval", "i", "10s", "Minimum interval between checks")
}
