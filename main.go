package main

import (
	"fmt"
	"log"
	"os"

	"github.com/foxboron/gandictl/api"
	"github.com/spf13/cobra"
)

var (
	GandiAPIKey = os.Getenv("GANDICTL")
)

var rootCmd = &cobra.Command{
	Use:   "gandictl",
	Short: "Edit gandi stuff",
}

func listCmd(c *api.Client) *cobra.Command {
	return &cobra.Command{
		Use:   "list",
		Short: "list",
		Run: func(cmd *cobra.Command, args []string) {
			ret := api.GetDomains(c)
			for _, item := range ret {
				fmt.Println(item.Fqdn)
			}
		},
	}
}

func recordsCmd(c *api.Client) *cobra.Command {
	return &cobra.Command{
		Use:   "records",
		Short: "records",
		Run: func(cmd *cobra.Command, args []string) {
			ret := api.GetZonefile(c, args[0])
			fmt.Println(string(ret))
		},
	}
}

func editCmd(c *api.Client) *cobra.Command {
	return &cobra.Command{
		Use:   "edit",
		Short: "edit",
		Run: func(cmd *cobra.Command, args []string) {
			ret := api.GetZonefile(c, args[0])
			buf, err := TempFile(ret)
			if err != nil {
				log.Fatal(err)
			}
			resp := api.WriteZonefile(c, args[0], buf)
			fmt.Println(api.NewSuccsessResponse(resp).Message)
		},
	}
}

func main() {
	c := api.NewClient(GandiAPIKey)
	rootCmd.AddCommand(listCmd(c))
	rootCmd.AddCommand(editCmd(c))
	rootCmd.AddCommand(recordsCmd(c))
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
