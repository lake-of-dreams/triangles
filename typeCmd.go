package main

import (
	"fmt"
	"strconv"

	"github.com/pkg/errors"

	log "github.com/sirupsen/logrus"

	"github.com/lake-of-dreams/triangles/triangle"
	cli "gopkg.in/urfave/cli.v1"
)

var (
	argNameSide1      = "side1"
	argNameSide2      = "side2"
	argNameSide3      = "side3"
	defaultSideLength = float64(0)
	typeCmdUsage      = fmt.Sprintf("Usage: type <%s> <%s> <%s>\n <%s>, <%s> and <%s> are side lengths of triangle", argNameSide1, argNameSide2, argNameSide3, argNameSide1, argNameSide2, argNameSide3)
)

var typeCommand = cli.Command{
	Name:      "type",
	Usage:     "get type of triangle based on input side lengths",
	Action:    typeAction,
	UsageText: typeCmdUsage,
}

var typeAction = func(c *cli.Context) error {
	log.Debug("Parsing type command arguments")
	o, err := parseTypeArguments(c)
	if err != nil {
		log.WithError(err).Error("Unable to validate arguments")
		return cli.NewExitError(fmt.Sprintf("Invalid arguments.\n%s", typeCmdUsage), 1)
	}
	triangleType, err := triangle.GetType(o.Side1, o.Side2, o.Side3)
	if err != nil {
		log.Error(err.Error())
		return err
	}
	log.Infof("Triangle type: %s", *triangleType)

	return nil
}

type typeCommandArguments struct {
	Side1 float64
	Side2 float64
	Side3 float64
}

func parseTypeArguments(c *cli.Context) (*typeCommandArguments, error) {
	numberOfArgs := c.NArg()
	if numberOfArgs == 0 {
		return nil, fmt.Errorf("type requires 3 arguments")
	}
	if numberOfArgs != 3 {
		return nil, fmt.Errorf("Invalid no. of arguments, expected: %d, found: %d", 3, numberOfArgs)
	}
	args := c.Args()
	side1, err := validateAndGetSideLength(args.Get(0))
	if err != nil {
		return nil, err
	}

	side2, err := validateAndGetSideLength(args.Get(1))
	if err != nil {
		return nil, err
	}

	side3, err := validateAndGetSideLength(args.Get(2))
	if err != nil {
		return nil, err
	}

	return &typeCommandArguments{
		Side1: side1,
		Side2: side2,
		Side3: side3,
	}, nil
}

func validateAndGetSideLength(arg string) (float64, error) {
	sideLength, err := strconv.ParseFloat(arg, 64)
	if err != nil {
		return 0, errors.Wrap(err, "Invalid input")
	}
	if !(sideLength > defaultSideLength) {
		return 0, fmt.Errorf("Invalid side length for %f", sideLength)
	}
	return sideLength, nil
}
