package main

import (
	"github.com/manifoldco/promptui"
	"github.com/pterm/pterm"
	"github.com/wizedkyle/cvecli/config"
	"github.com/wizedkyle/cvecli/internal/authentication"
	configureCmd "github.com/wizedkyle/cvecli/internal/cmd/configure"
	"github.com/wizedkyle/cvecli/internal/cmd/root"
	"os"
)

func main() {
	if authentication.CheckCredentialsPath() == false {
		pterm.Warning.Println("There are no credentials set or accessible.")
		promptAccept := promptui.Prompt{
			Label:     "Would you like to set credentials now",
			IsConfirm: true,
		}
		promptResponse, err := promptAccept.Run()
		if err != nil {
			os.Exit(0)
		}
		if promptResponse == "y" {
			configureCmd.SetCredentials()
			client := authentication.GetCVEServicesSDKConfig()
			config.SetClient(client)
		}
	} else {
		client := authentication.GetCVEServicesSDKConfig()
		config.SetClient(client)
	}
	rootCmd := root.NewCmdRoot()
	rootCmd.Execute()
}
