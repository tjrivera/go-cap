package main

import (
    "github.com/spf13/cobra"
    "github.com/spf13/viper"
)

var mainCmd = &cobra.Command{
    Use: "gocap",

    Short: "GoCap CLI",

    Long: "GoCap CLI: A command-line client to REDCap.",

    Run: func(cmd *cobra.Command, args []string){
        cmd.Help()
    },
}

func main(){
    mainCmd.AddCommand(formsToCsv)

    viper.SetEnvPrefix("gocap")
    viper.AutomaticEnv()

    flags := mainCmd.PersistentFlags()

    flags.String("token", "", "REDCap API token")
    flags.String("host", "", "REDCap host URL")

    viper.BindPFlag("token", flags.Lookup("token"))
    viper.BindPFlag("host", flags.Lookup("host"))

    mainCmd.Execute()



}
