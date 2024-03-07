/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"github.com/olekukonko/tablewriter"
	"github.com/prometheus-community/pro-bing"
	"github.com/spf13/cobra"
	"os"
)

var count int
var region string

var lagCmd = &cobra.Command{
	Use:   "lag",
	Short: "Used to check your connection",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Checking connection...")

		pinger, err := setRegion(region)
		if err != nil {
			panic(err)
		}

		pinger.SetPrivileged(true)

		pinger.Count = count
		err = pinger.Run()
		if err != nil {
			panic(err)
		}
		stats := pinger.Statistics()
		showPing(stats)
	},
}

func init() {
	rootCmd.AddCommand(lagCmd)

	rootCmd.PersistentFlags().IntVarP(&count, "count", "c", 3, "Desired number of packages")
	rootCmd.PersistentFlags().StringVarP(&region, "region", "r", "sa", "Desired region name")
}

func showPing(stats *probing.Statistics) {
	fmt.Println("Ping Results:")

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Metric", "Value"})

	table.Append([]string{"Packets Sent", fmt.Sprintf("%d", stats.PacketsSent)})
	table.Append([]string{"Packets Received", fmt.Sprintf("%d", stats.PacketsRecv)})
	table.Append([]string{"Packet Loss", fmt.Sprintf("%.2f%%", stats.PacketLoss)})
	table.Append([]string{"Min Round-trip Time", fmt.Sprintf("%.2f ms", stats.MinRtt.Seconds()*1000)})
	table.Append([]string{"Avg Round-trip Time", fmt.Sprintf("%.2f ms", stats.AvgRtt.Seconds()*1000)})
	table.Append([]string{"Max Round-trip Time", fmt.Sprintf("%.2f ms", stats.MaxRtt.Seconds()*1000)})

	table.Render()
}

func setRegion(region string) (*probing.Pinger, error) {
	var pinger *probing.Pinger
	var err error

	if region == "sa" {
		pinger, err = probing.NewPinger("s3-sa-east-1.amazonaws.com")
	} else if region == "us-west" {
		pinger, err = probing.NewPinger("s3-us-west-2.amazonaws.com")
	} else if region == "us-east" {
		pinger, err = probing.NewPinger("s3.amazonaws.com")
	} else if region == "eu-west" {
		pinger, err = probing.NewPinger("s3-eu-west-1.amazonaws.com")
	} else if region == "sea" {
		pinger, err = probing.NewPinger("s3-ap-southeast-1.amazonaws.com")
	} else if region == "oc" {
		pinger, err = probing.NewPinger("s3-ap-southeast-2.amazonaws.com")
	} else if region == "ru" {
		pinger, err = probing.NewPinger("storage.yandexcloud.net")
	}

	return pinger, err
}
