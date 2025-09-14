package cmd

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"strings"

	"github.com/kkrypt0nn/overscry/core"
	"github.com/kkrypt0nn/overscry/models"
	"github.com/kkrypt0nn/overscry/settings"
	"github.com/spf13/cobra"
)

const OVERPASS_API_URL string = "https://overpass-api.de/api/interpreter"

type ExecFlags struct {
	settingsFile string
	outFile      string
}

var execFlags ExecFlags

var execCmd = &cobra.Command{
	Use:   "exec",
	Short: "Executes an Overpass query with the given settings",
	Run: func(cmd *cobra.Command, args []string) {
		s, err := settings.LoadSettings(execFlags.settingsFile)
		if err != nil {
			core.Logger.Error("Failed to read settings: " + err.Error())
			return
		}

		res := fmt.Sprintf(models.JSON_OUTPUT, s.Node.ToOQL())

		formBody := url.Values{
			"data": []string{res},
		}.Encode()

		core.Logger.Success("Sending request to Overpass, waiting for the response...")
		resp, err := http.Post(OVERPASS_API_URL, "application/x-www-form-urlencoded", strings.NewReader(formBody))
		if err != nil {
			core.Logger.Error("Failed making POST request: " + err.Error())
			return
		}

		defer func() {
			err := resp.Body.Close()
			if err != nil {
				core.Logger.Error("Failed closing response body: " + err.Error())
			}
		}()
		if resp.StatusCode != http.StatusOK {
			core.Logger.Error("Received non-OK HTTP status: " + resp.Status)
			return
		}

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			core.Logger.Error("Failed reading response body: " + err.Error())
			return
		}

		var temp struct {
			Elements []map[string]any `json:"elements"`
		}

		err = json.Unmarshal(body, &temp)
		if err != nil {
			core.Logger.Error("Failed unmarshalling response body: " + err.Error())
			return
		}

		prettyJSON, err := json.MarshalIndent(temp.Elements, "", "  ")
		if err != nil {
			core.Logger.Error("Failed marshalling the elements with indentation: " + err.Error())
			return
		}

		fmt.Println(string(prettyJSON))

		if execFlags.outFile != "" {
			err := os.WriteFile(execFlags.outFile, prettyJSON, 0644)
			if err != nil {
				core.Logger.Error(err.Error())
				return
			}
			core.Logger.Success(fmt.Sprintf("Saved the query results to the %s file", execFlags.outFile))
		}
	},
}

func init() {
	execCmd.Flags().StringVar(&execFlags.settingsFile, "settings", "", "File path to the YML file for the settings")
	_ = execCmd.MarkFlagRequired("settings")
	execCmd.Flags().StringVar(&execFlags.outFile, "out", "", "Output file to save the query results in")
	rootCmd.AddCommand(execCmd)
}
