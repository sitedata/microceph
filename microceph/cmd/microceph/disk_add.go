package main

import (
	"context"

	"github.com/canonical/microcluster/microcluster"
	"github.com/spf13/cobra"

	"github.com/canonical/microceph/microceph/api/types"
	"github.com/canonical/microceph/microceph/client"
)

type cmdDiskAdd struct {
	common *CmdControl
	disk   *cmdDisk

	flagWipe    bool
	flagEncrypt bool
	walDevice   string
	walEncrypt  bool
	walWipe     bool
	dbDevice    string
	dbEncrypt   bool
	dbWipe      bool
}

func (c *cmdDiskAdd) Command() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "add <SPEC>",
		Short: "Add a new Ceph disk (OSD)",
		Long: `Add a new Ceph disk (OSD) for <SPEC>.
    SPEC is either a path to a block device such as /dev/sdb or a specification for a loop file.
    The specification for a loop file is of the form "loop,<size>,<nr>"
    The size is a number with M, G, or T suffixes for megabytes, gigabytes, or terabytes.
    The nr is the number of loop OSDs to create.
    A spec of "loop,8G,3" will create 3 loop OSDs of 8GB each.`,
		RunE: c.Run,
	}

	cmd.PersistentFlags().BoolVar(&c.flagWipe, "wipe", false, "Wipe the disk prior to use")
	cmd.PersistentFlags().BoolVar(&c.flagEncrypt, "encrypt", false, "Encrypt the disk prior to use")
	cmd.PersistentFlags().StringVar(&c.walDevice, "wal-device", "", "The device used for WAL")
	cmd.PersistentFlags().BoolVar(&c.walWipe, "wal-wipe", false, "Wipe the WAL device prior to use")
	cmd.PersistentFlags().BoolVar(&c.walEncrypt, "wal-encrypt", false, "Encrypt the WAL device prior to use")
	cmd.PersistentFlags().StringVar(&c.dbDevice, "db-device", "", "The device used for the DB")
	cmd.PersistentFlags().BoolVar(&c.dbWipe, "db-wipe", false, "Wipe the DB device prior to use")
	cmd.PersistentFlags().BoolVar(&c.dbEncrypt, "db-encrypt", false, "Encrypt the DB device prior to use")

	return cmd
}

func (c *cmdDiskAdd) Run(cmd *cobra.Command, args []string) error {
	if len(args) != 1 {
		return cmd.Help()
	}

	m, err := microcluster.App(context.Background(), microcluster.Args{StateDir: c.common.FlagStateDir, Verbose: c.common.FlagLogVerbose, Debug: c.common.FlagLogDebug})
	if err != nil {
		return err
	}

	cli, err := m.LocalClient()
	if err != nil {
		return err
	}

	req := &types.DisksPost{
		Path:    args[0],
		Wipe:    c.flagWipe,
		Encrypt: c.flagEncrypt,
	}

	if c.walDevice != "" {
		req.WALDev = &c.walDevice
		req.WALWipe = c.walWipe
		req.WALEncrypt = c.walEncrypt
	}

	if c.dbDevice != "" {
		req.DBDev = &c.dbDevice
		req.DBWipe = c.dbWipe
		req.DBEncrypt = c.dbEncrypt
	}

	err = client.AddDisk(context.Background(), cli, req)
	if err != nil {
		return err
	}

	return nil
}
