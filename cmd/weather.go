/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"

	"github.com/spf13/cobra"
)

// Flags
var apikey = "ef87d951c241d199580955e7e7478ed7"
var city string
var zipcode string // api.openweathermap.org/data/2.5/weather?zip={zip code},{country code}&appid={API key}


// weatherCmd represents the weather command
var weatherCmd = &cobra.Command{
	Use:   "weather",
	Short: "Weather Checker",
	Long: `Weather Checker CLI Application.`,
	Run: func(cmd *cobra.Command, args []string) {
		var URL  = "http://api.openweathermap.org/data/2.5/weather?q=" + city + "&appid=" + apikey
		weather(URL)
	},

}

func weather(url string) {
	response, err := http.Get(url)
	if err != nil{
		panic(err)
	}
	fmt.Println("Response Status: ", response.Status)
	data, err := io.ReadAll(response.Body)
	if err != nil{
		panic(err)
	}
	_, jsondata := strconv.Atoi(string(data))
	jsonResults, err := json.MarshalIndent(jsondata, "", " ")
	if err != nil {
		panic(err)
	}
	//os.Stdout.Write(jsonResults)
	fmt.Printf(string(jsonResults))
}

func init() {
	rootCmd.AddCommand(weatherCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	//weatherCmd.PersistentFlags().StringVar(&apikey, "apikey", "a", "Enter your API-Key")

	weatherCmd.PersistentFlags().StringVar(&city, "city", "c", "Enter your City eg. Dallas")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// weatherCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
