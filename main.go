package main

import (
	"os"

	cli "gopkg.in/urfave/cli.v1"
)

func main() {
	app := cli.NewApp()

	app.Name = "Triangles"
	app.Usage = "App to determine if a triangle is equilateral, isosceles or scalene based on input side lengths"

	app.Commands = []cli.Command{
		typeCommand,
	}

	app.Run(os.Args)
}
