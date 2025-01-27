package cmd

import (
	"io"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var createDataFile string

var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create resources or events from a file or stdin",
	Long:  "Create resources or events from a file or stdin",
}

func init() {
	rootCmd.AddCommand(createCmd)

	createCmd.PersistentFlags().StringVarP(&createDataFile, "file", "f", "-", "File to read data from. If '.' then reads from './data.yaml'. Defaults to reading from stdin.")
	viper.BindPFlags(createCmd.Flags())
}

func hasStdin() bool {
	stat, err := os.Stdin.Stat()
	if err != nil {
		return false
	}
	return stat.Size() > 0
}

func readCreateFile() ([]byte, error) {
	if hasStdin() {
		data, err := io.ReadAll(os.Stdin)
		if err != nil {
			return nil, err
		}
		return data, nil
	} else {
		file, err := os.Open(createDataFile)
		if err != nil {
			return nil, err
		}
		defer file.Close()

		data, err := io.ReadAll(file)
		if err != nil {
			return nil, err
		}
		return data, nil
	}
}

func readCreateConfigFile() {
	if createDataFile != "" {
		if createDataFile == "-" {
			viper.SetConfigType("yaml")
			if hasStdin() {
				viper.ReadConfig(os.Stdin)
			}
			return
		} else if createDataFile == "." {
			viper.SetConfigFile("./data.yaml")
		} else {
			viper.SetConfigFile(createDataFile)
		}
	} else {
		home, err := os.UserHomeDir()
		cobra.CheckErr(err)

		viper.SetConfigName("data")
		viper.SetConfigType("yaml")
		viper.AddConfigPath(".")
		viper.AddConfigPath(home)
	}
	viper.ReadInConfig()
}
