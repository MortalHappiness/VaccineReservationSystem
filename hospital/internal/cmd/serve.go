package cmd

import (
	"log"

	"github.com/MortalHappiness/VaccineReservationSystem/hospital/internal/worker"
	"github.com/spf13/cobra"
)

// serveCmd represents the serve command.
var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Parent command for starting public and internal HTTP/2 APIs",
	Long: `
	Worker exposes one port for handling HTTP requests. The Internal APIs is splitted by root path "/internal".
	`,
	Run: func(cmd *cobra.Command, args []string) {
		log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
		a := worker.New()
		a.Run()
	},
}

func init() {
	RootCmd.AddCommand(serveCmd)
}
