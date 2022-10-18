package main

import (
	"fmt"
	"os"

	"github.com/invoicing3/core/pkg/consumers"
	"github.com/spf13/cobra"
)


func main() {
	var consumerGroupId string = "invoice-created";
	var rootCmd = &cobra.Command{
		Use:   "core",
		Short: "invoicing3 core is responsible for handle working with web3",
		Long: ``,
		Run: func(cmd *cobra.Command, args []string) {
			var kafkaConsumer = consumers.GetKafkaConsumer(consumerGroupId);
			kafkaConsumer.Events()
		},
	}
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}