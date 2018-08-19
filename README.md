## Solution
An API `GetType` that will take as input side lengths and will compute whether the triangle is equilateral, isosceles or scalene. Wrapped up in a `cli` tool so that binary built from this code can be used to test the api with various inputs.
Alternatively can also be integrated with any service etc.

## Input validation
We use gopkg.in/urfave/cli.v1 to get user inputs for the sides and there we check if side lengths are greater than 0

## How to run it (macOS)
```
$ brew install go
$ brew install dep
$ go get -u -v github.com/lake-of-dreams/triangles
$ go install -v github.com/lake-of-dreams/triangles
$ triangles 1 2 3

```

## Code structure
**main.go** - Contains code for setting up a command line app<br />
**typeCmd.go** - Code for setting up command line parameters, validation etc. and invoking the actual logic<br />
**triangle/triangle.go** - Contains logic for the API<br />
**triangle/triangle_test.go** - Contains unit test for the API<br />
**Gopkg.lock,Gopkg.toml** - generated files for golang dependency tool dep

## Sample input/output


```
Ravis-MacBook-Pro-5:triangles rsinghal$ triangles 
NAME:
   Triangles - App to determine if a triangle is equilateral, isosceles or scalene based on input side lengths

USAGE:
   triangles [global options] command [command options] [arguments...]

VERSION:
   0.0.0

COMMANDS:
     type     get type of triangle based on input side lengths
     help, h  Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --help, -h     show help
   --version, -v  print the version
```

```
Ravis-MacBook-Pro-5:triangles rsinghal$ triangles type --h
NAME:
   triangles type - get type of triangle based on input side lengths

USAGE:
   Usage: type <side1> <side2> <side3>
 <side1>, <side2> and <side3> are side lengths of triangle
```

```
Ravis-MacBook-Pro-5:triangles rsinghal$ triangles type 3 2 1
ERRO[0000] The given input is invalid as it does not conform to the triangle inequality (sum of any 2 sides of a triangle must be greater than the measure of the third side) 
```

```
Ravis-MacBook-Pro-5:triangles rsinghal$ triangles type 3 3 3
INFO[0000] Triangle type: Equilateral     
```

```
Ravis-MacBook-Pro-5:triangles rsinghal$ triangles type 3 3 2
INFO[0000] Triangle type: Isosceles    
```

```
Ravis-MacBook-Pro-5:triangles rsinghal$ triangles type 3 4 5
INFO[0000] Triangle type: Scalene  
```

























































