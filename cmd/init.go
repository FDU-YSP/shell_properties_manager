/*
Copyright © 2022 Shaopeng <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/FDU-YSP/spmanager/util"
	"path/filepath"
)

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "spmanger config file initialization.",
	Long: `Usage: spm init / spm init -f.
spm init: means you want to initialize spm config file, but if config file exist, it will not override.
spm init -f: means no mater spmanager config file exist, this cli will generate new config files`,

	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("init called ... ")
		if f, err := cmd.Flags().GetBool("force"); err == nil {
			if f == false {
				fmt.Println("`force` flag is false, not override config files...")
				GenerateConfigFilesForSPM(false)
			} else {
				fmt.Println("`force` flag is true, will override config files...")
				GenerateConfigFilesForSPM(true)
			}
		} else {
			fmt.Println("spm cli execute error when get `force` flag.")
		}
	},
}

func GenerateConfigFilesForSPM(force bool) {
	var spmconfig string
	if confighome := util.HomeDir(); confighome != "" {
		// fmt.Println(confighome)
		spmconfig = filepath.Join(confighome, ".spmanager", "spm.conf")
		fmt.Println(spmconfig)
	} else {
		fmt.Println("config home is null")
	}
}

func init() {
	rootCmd.AddCommand(initCmd)
	// initCmd.Flags().StringVar("force", "f", "", "force to init config")
	initCmd.Flags().BoolP("force", "f", false, "force to init config")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// initCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// initCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
