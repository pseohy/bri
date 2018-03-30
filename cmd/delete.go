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
	"log"

	"github.com/pseohy/bri/conf"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// deleteCmd represents the delete command
var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete a device",
	Long:  `Delete a device from the configuration`,
	Run: func(cmd *cobra.Command, args []string) {
		h, err := conf.EncryptDevice(dtype, did)
		if err != nil {
			log.Fatal(err)
		}

		conf.DeviceData.Delete(h)
		conf.DeviceData.Dump()
	},
}

func init() {
	configCmd.AddCommand(deleteCmd)

	deleteCmd.Flags().StringVarP(&dtype, "type", "t", "", "Device type")
	deleteCmd.Flags().StringVarP(&did, "id", "i", "", "Device serial no.")

	viper.BindPFlag("type", deleteCmd.Flags().Lookup("type"))
	viper.BindPFlag("id", deleteCmd.Flags().Lookup("id"))
}
