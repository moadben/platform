// Copyright (c) 2016 Mattermost, Inc. All Rights Reserved.
// See License.txt for license information.
package main

import (
	"errors"
	"io/ioutil"

	"github.com/mattermost/platform/app"
	"github.com/spf13/cobra"
)

var licenseCmd = &cobra.Command{
	Use:   "license",
	Short: "Licensing commands",
}

var uploadLicenseCmd = &cobra.Command{
	Use:     "upload [license]",
	Short:   "Upload a license.",
	Long:    "Upload a license. Replaces current license.",
	Example: "  license upload /path/to/license/mylicensefile.mattermost-license",
	RunE:    uploadLicenseCmdF,
}

func init() {
	licenseCmd.AddCommand(uploadLicenseCmd)
}

func uploadLicenseCmdF(cmd *cobra.Command, args []string) error {
	initDBCommandContextCobra(cmd)

	if len(args) != 1 {
		return errors.New("Enter one license file to upload")
	}

	var fileBytes []byte
	var err error
	if fileBytes, err = ioutil.ReadFile(args[0]); err != nil {
		return err
	}

	if _, err := app.SaveLicense(fileBytes); err != nil {
		return err
	}

	CommandPrettyPrintln("Uploaded license file")

	return nil
}
