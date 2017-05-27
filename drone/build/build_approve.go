package build

import (
	"fmt"

	"github.com/drone/drone-cli/drone/internal"
	"github.com/urfave/cli"
	"strconv"
)

var buildApproveCmd = cli.Command{
	Name:      "approve",
	Usage:     "approve a build",
	ArgsUsage: "<repo/name> <build>",
	Action:    buildApprove,
}

func buildApprove(c *cli.Context) (err error) {
	repo := c.Args().First()
	owner, name, err := internal.ParseRepo(repo)
	if err != nil {
		return err
	}
	number, err := strconv.Atoi(c.Args().Get(1))
	if err != nil {
		return errInvalidBuildNumber
	}

	client, err := internal.NewClient(c)
	if err != nil {
		return err
	}

	_, err = client.BuildApprove(owner, name, number)
	if err != nil {
		return err
	}

	fmt.Printf("Approving build %s/%s#%d\n", owner, name, number)
	return nil
}
