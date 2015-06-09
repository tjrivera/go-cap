package main

import (
    "fmt"
    "github.research.chop.edu/cbmi/go-cap/redcap"
    "github.com/spf13/cobra"
    "github.com/spf13/viper"
)

var formsToCsv = &cobra.Command{
    Use: "formstocsv",

    Short: "Output all of the forms from a REDCap project to csv files.",

    Run: func(cmd *cobra.Command, args []string){

        token := viper.GetString("token")
        host := viper.GetString("host")
        project := redcap.NewRedcapProject(
    		host,
    		token)

        if len(args) == 0 {
            fmt.Println("Please specify an output path")
            return
        }
    	project.FormsToCsv(args[0])
    },
}


func init(){

}
