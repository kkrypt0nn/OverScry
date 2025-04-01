package cmd

import (
	"fmt"
	"os"

	"github.com/kkrypt0nn/overscry/core"
	"github.com/kkrypt0nn/overscry/models"
	"github.com/kkrypt0nn/overscry/settings"
	"github.com/spf13/cobra"
)

var settingsFile string
var outFile string

var genCmd = &cobra.Command{
	Use:   "gen",
	Short: "Generate an Overpass query with the given settings",
	Run: func(cmd *cobra.Command, args []string) {
		s, err := settings.LoadSettings(settingsFile)
		if err != nil {
			core.Logger.Error("Failed to read settings: " + err.Error())
			return
		}

		res := fmt.Sprintf(models.JSON_OUTPUT, s.Node.ToOQL())
		fmt.Println(res)

		if outFile != "" {
			err := os.WriteFile(outFile, []byte(res), 0644)
			if err != nil {
				core.Logger.Error(err.Error())
			} else {
				core.Logger.Success(fmt.Sprintf("Saved the resulting query to the %s file", outFile))
			}
		}
	},
}

func init() {
	genCmd.Flags().StringVar(&settingsFile, "settings", "", "File path to the YML file for the settings")
	_ = genCmd.MarkFlagRequired("settings")
	genCmd.Flags().StringVar(&outFile, "out", "", "Output file to save the query in")
	rootCmd.AddCommand(genCmd)
}
