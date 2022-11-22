package cmd

import (
	"fmt"
	"os"

	"github.com/hashicorp/go-version"
	"github.com/spf13/cobra"
)

var (
	Semverv1 string
	Semverv2 string
	Semversw int
)

var semverCmd = &cobra.Command{
	GroupID: groupMain,
	Use:     "semver",
	Short:   "Validation of semver",
	Long:    ``,
	Run: func(cmd *cobra.Command, args []string) {

		vv, err := cmd.Flags().GetInt("e")
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		fmt.Println(vv)

		v1, err := version.NewVersion(Semverv1)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
			return
		}

		v2, err := version.NewVersion(Semverv2)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
			return
		}

		switch Semversw {
		case 1:
			if v1.Equal(v2) {
				fmt.Println("true")
			} else {
				fmt.Println("false")
			}
		case 2:
			if v1.GreaterThan(v2) || v1.Equal(v2) {
				fmt.Println("true")
			} else {
				fmt.Println("false")
			}
		case 3:
			if v1.LessThan(v2) || v1.Equal(v2) {
				fmt.Println("true")
			} else {
				fmt.Println("false")
			}
		case 4:
			if v1.GreaterThan(v2) {
				fmt.Println("true")
			} else {
				fmt.Println("false")
			}
		case 5:
			if v1.LessThan(v2) {
				fmt.Println("true")
			} else {
				fmt.Println("false")
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(semverCmd)
	semverCmd.Flags().StringVarP(&Semverv1, "v1", "", "", "Version 1")
	semverCmd.Flags().StringVarP(&Semverv2, "v2", "", "", "Version 2")
	semverCmd.Flags().IntVarP(&Semversw, "e", "", 1, "Equal")
	semverCmd.Flags().IntVarP(&Semversw, "eg", "", 2, "Equal or greater")
	semverCmd.Flags().IntVarP(&Semversw, "el", "", 3, "Equal or less")
	semverCmd.Flags().IntVarP(&Semversw, "g", "", 4, "Greater")
	semverCmd.Flags().IntVarP(&Semversw, "l", "", 5, "Less")
	semverCmd.MarkFlagRequired("v1")
	semverCmd.MarkFlagRequired("v2")
	semverCmd.Flags().Lookup("e").NoOptDefVal = "1"
	semverCmd.Flags().Lookup("eg").NoOptDefVal = "2"
	semverCmd.Flags().Lookup("el").NoOptDefVal = "3"
	semverCmd.Flags().Lookup("g").NoOptDefVal = "4"
	semverCmd.Flags().Lookup("l").NoOptDefVal = "5"

}
