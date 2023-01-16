/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"encoding/binary"
	"fmt"
	"net/netip"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "ipcalc [IP Address]",
	DisableFlagsInUseLine: true,
	Short: "convert IP address to decimal and hexadecimal",
	PreRunE: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return fmt.Errorf("too few arguments")
		}
		return nil
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		ip, err := netip.ParseAddr(args[0])
		if err != nil {
			return err
		}
		if ip.Is4() {
			fmt.Printf("%d\n0x%x\n",
				binary.BigEndian.Uint32(ip.AsSlice()),
				binary.BigEndian.Uint32(ip.AsSlice()))
			return nil
		}
		if ip.Is6() {
			fmt.Printf("%d%020d\n0x%x%016x\n",
				binary.BigEndian.Uint64(ip.AsSlice()[:8]),
				binary.BigEndian.Uint64(ip.AsSlice()[8:]),
				binary.BigEndian.Uint64(ip.AsSlice()[:8]),
				binary.BigEndian.Uint64(ip.AsSlice()[8:]))
			return nil
		}

		return nil
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
}
