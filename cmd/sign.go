/*
Copyright © 2020 Ladislav Macoun <ladislavmacoun@gmail.com>

This program is free software; you can redistribute it and/or
modify it under the terms of the GNU General Public License
as published by the Free Software Foundation; either version 2
of the License, or (at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.

You should have received a copy of the GNU Lesser General Public License
along with this program. If not, see <http://www.gnu.org/licenses/>.
*/
package cmd

import (
	"errors"
	"fmt"
	"net/url"

	"github.com/spf13/cobra"
)

// Config holds CLI arguments settings.
type Config struct {
	rawURL string
	secret string
}

func (c *Config) compile() error {
	u, err := url.Parse(c.rawURL)
	if err != nil {
		return err
	}

	if u.Hostname() == "" {
		return errors.New("missing hostname")
	}

	if c.secret == "" {
		return errors.New("missing secret")
	}

	return nil
}

var cfg Config

// signCmd represents the sign command
var signCmd = &cobra.Command{
	Use:   "sign",
	Short: "A command for siging URLs with secure token",
	Long:  `TBA`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("sign called")
	},
}

func init() {
	rootCmd.AddCommand(signCmd)
	signCmd.Flags().StringVarP(&cfg.rawURL, "url", "u", "", "url to sign")
	signCmd.Flags().StringVarP(&cfg.secret, "secret", "s", "", "secret token to sign url with")
	signCmd.MarkFlagRequired("url")
	signCmd.MarkFlagRequired("secret")

	if err := cfg.compile(); err != nil {
		panic(err)
	}
}
