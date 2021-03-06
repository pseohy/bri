// Copyright © 2018 Seonghyun Park <pseohy@gmail.com>
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package cmd

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"

	"github.com/pseohy/bri/conf"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	device_dtype string
	device_did   string
	device_uInfo []string
	devide_msg   string

	deviceURL string
)

// deviceCmd represents the device command
var deviceCmd = &cobra.Command{
	Use:   "device",
	Short: "A device sends data to the server",
	Long: `A device notifies the server with information about its usage.
Only data from the authenticated devices are collected`,

	Run: func(cmd *cobra.Command, args []string) {
		var dmsg = conf.DeviceMsg{
			Dtype: device_dtype,
			Did:   device_did,
			UInfo: device_uInfo,
			Msg:   devide_msg,
		}

		jsonDmsg, err := json.Marshal(&dmsg)
		if err != nil {
			log.Fatal(err)
		}

		req, err := http.NewRequest("POST", deviceURL, bytes.NewBuffer(jsonDmsg))
		req.Header.Set("Cnotent-type", "application/json")

		client := http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			panic(err)
		}
		defer resp.Body.Close()
	},
}

func init() {
	rootCmd.AddCommand(deviceCmd)

	deviceCmd.Flags().StringVarP(&device_dtype, "type", "t", "", "Device type")
	deviceCmd.Flags().StringVarP(&device_did, "id", "i", "", "Device serial no.")
	deviceCmd.Flags().StringSliceVarP(&device_uInfo, "user", "u",
		[]string{"", ""}, "User name and phone number")
	deviceCmd.Flags().StringVarP(&devide_msg, "msg", "m", "on", "Message")
	deviceCmd.Flags().StringVar(&deviceURL, "url", "http://localhost:4000/device", "URL to the server")

	viper.BindPFlag("type", deviceCmd.Flags().Lookup("type"))
	viper.BindPFlag("id", deviceCmd.Flags().Lookup("id"))
	viper.BindPFlag("user", deviceCmd.Flags().Lookup("user"))
	viper.BindPFlag("msg", deviceCmd.Flags().Lookup("msg"))
	viper.BindPFlag("url", deviceCmd.Flags().Lookup("url"))
}
