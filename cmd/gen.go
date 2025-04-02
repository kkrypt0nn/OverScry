package cmd

import (
	"fmt"
	"os"

	"github.com/kkrypt0nn/overscry/core"
	"github.com/kkrypt0nn/overscry/models"
	"github.com/kkrypt0nn/overscry/settings"
	"github.com/spf13/cobra"
)

type GenFlags struct {
	settingsFile string
	outFile      string
}

var genFlags GenFlags

var genCmd = &cobra.Command{
	Use:   "gen",
	Short: "Generate an Overpass query with the given settings",
	Run: func(cmd *cobra.Command, args []string) {
		s, err := settings.LoadSettings(genFlags.settingsFile)
		if err != nil {
			core.Logger.Error("Failed to read settings: " + err.Error())
			return
		}

		res := fmt.Sprintf(models.JSON_OUTPUT, s.Node.ToOQL())
		fmt.Println(res)

		if genFlags.outFile != "" {
			err := os.WriteFile(genFlags.outFile, []byte(res), 0644)
			if err != nil {
				core.Logger.Error(err.Error())
				return
			}
			core.Logger.Success(fmt.Sprintf("Saved the resulting query to the %s file", genFlags.outFile))
		}
	},
}

func init() {
	genCmd.Flags().StringVar(&genFlags.settingsFile, "settings", "", "File path to the YML file for the settings")
	_ = genCmd.MarkFlagRequired("settings")
	genCmd.Flags().StringVar(&genFlags.outFile, "out", "", "Output file to save the query in")
	rootCmd.AddCommand(genCmd)
}
