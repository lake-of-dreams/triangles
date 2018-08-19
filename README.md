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





















































