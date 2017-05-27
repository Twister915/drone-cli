package build

import (
	"errors"
	"fmt"

	"github.com/drone/drone-cli/drone/internal"
	"github.com/urfave/cli"
)

var buildStartCmd = cli.Command{
	Name:      "start",
	Usage:     "start a build",
	ArgsUsage: "<repo/name> [build]",
	Action:    buildStart,
	Flags: []cli.Flag{
		cli.StringSliceFlag{
			Name:  "param, p",
			Usage: "custom parameters to be injected into the job environment. Format: KEY=value",
		},
	},
}

func buildStart(c *cli.Context) (err error) {
	repo := c.Args().First()
	owner, name, err := internal.ParseRepo(repo)
	if err != nil {
		return err
	}

	client, err := internal.NewClient(c)
	if err != nil {
		return err
	}

<<<<<<< HEAD
	buildArg := c.Args().Get(1)
	var number int
	if buildArg == "last" {
		// Fetch the build number from the last build
		build, err := client.BuildLast(owner, name, "")
		if err != nil {
			return err
		}
		number = build.Number
	} else {
		if len(buildArg) == 0 {
			return errors.New("missing job number")
		}
		number, err = strconv.Atoi(buildArg)
		if err != nil {
			return err
		}
	}

=======
	_, number, err := getBuildWithArg(c.Args().Get(1), owner, name, client)
>>>>>>> add clearer error messages when drone fails to parse the build number
	params := internal.ParseKeyPair(c.StringSlice("param"))

	build, err := client.BuildStart(owner, name, number, params)
	if err != nil {
		return err
	}

	fmt.Printf("Starting build %s/%s#%d\n", owner, name, build.Number)
	return nil
}
