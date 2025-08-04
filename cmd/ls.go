/*
Copyright © 2025 Ken'ichiro Oyama <k1lowxb@gmail.com>

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/
package cmd

import (
	"fmt"

	"github.com/k1LoW/deck"
	"github.com/spf13/cobra"
)

var lsCmd = &cobra.Command{
	Use:   "ls",
	Short: "list Google Slides presentations",
	Long:  `list Google Slides presentations.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		opts := []deck.Option{
			deck.WithProfile(profile),
			deck.WithSupportAllDrives(supportAllDrives),
		}
		slides, err := deck.List(cmd.Context(), opts...)
		if err != nil {
			return err
		}
		for _, slide := range slides {
			fmt.Printf("%s\t%s\n", slide.ID, slide.Title)
		}
		return nil
	},
}

func init() {
	rootCmd.AddCommand(lsCmd)
	lsCmd.Flags().BoolVar(&supportAllDrives, "support-all-drives", true, "Whether to include shared drives in the list")
}
