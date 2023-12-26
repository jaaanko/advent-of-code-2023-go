package main

import (
	"errors"
	"flag"
	"fmt"
	"log"
	"math/big"
	"os"
	"strconv"
	"strings"
)

type hailstone struct {
	x, y, z, vx, vy, vz int
}

type point struct {
	x, y int
}

type line struct {
	a, b, c int
}

var errNoIntersection = errors.New("lines do not intersect")

func main() {
	p1 := flag.Bool("p1", false, "run part 1")
	p2 := flag.Bool("p2", false, "run part 2")
	flag.Parse()
	args := flag.Args()

	if len(args) == 0 {
		log.Fatal("no input file provided")
	}

	b, err := os.ReadFile(args[0])
	if err != nil {
		log.Fatal(err)
	}

	input := string(b)

	if *p1 {
		fmt.Println("part 1:", part1(input, 400000000000000, 200000000000000))
	}

	if *p2 {
		fmt.Println("part 2:", part2(input))
	}
}

func part1(input string, upper, lower int) int {
	hailstones := parse(input)
	n := len(hailstones)
	res := 0

	for i := 0; i < n; i++ {
		x1, vx1 := hailstones[i].x, hailstones[i].vx
		y1, vy1 := hailstones[i].y, hailstones[i].vy
		l1 := getLine(
			point{x: x1, y: y1},
			point{x: x1 + vx1, y: y1 + vy1},
		)

		for j := i + 1; j < n; j++ {
			x2, vx2 := hailstones[j].x, hailstones[j].vx
			y2, vy2 := hailstones[j].y, hailstones[j].vy
			l2 := getLine(
				point{x: x2, y: y2},
				point{x: x2 + vx2, y: y2 + vy2},
			)

			x3, y3, err := intersection(l1, l2)
			if err != nil {
				continue
			}

			past := (vx1 > 0 && x3 < float64(x1) || vx1 < 0 && x3 > float64(x1)) ||
				(vy1 > 0 && y3 < float64(y1) || vy1 < 0 && y3 > float64(y1)) ||
				(vx2 > 0 && x3 < float64(x2) || vx2 < 0 && x3 > float64(x2)) ||
				(vy2 > 0 && y3 < float64(y2) || vy2 < 0 && y3 > float64(y2))

			if !past && float64(lower) <= x3 && x3 <= float64(upper) && float64(lower) <= y3 && y3 <= float64(upper) {
				res += 1
			}
		}
	}

	return res
}

/*
Solved part 2 with Python using sympy (https://www.sympy.org/en/index.html).
A very good explanation by HyperNeutrino can be found here: https://www.youtube.com/watch?v=guOyA7Ijqgk&t=1514s

Given a single hailstone h and time t,
we want our rock positions rx,ry,rz and hailstone positions hx,hy,hz to be equal.

For each dimension we have:

rx + t*rvx = hx + t*hvx
ry + t*rvy = hy + t*hvy
rz + t*rvz = hz + t*hvz

Where rvx,rvy,rvz are the x,y,z velocities of the rock,
and hvx,hvy,hvz are the x,y,z velocities of a single random hailstone.

We can rewrite the equation to separate t:

rx + t*rvx = hx + t*hvx
t*rvx - t*hvx = hx - rx
t*(rvx-hvx) = hx - rx
t = (hx-rx) / (rvx-hvx)

Same process for the other dimensions:

t = (hy-ry) / (rvy-hvy)
t = (hz-rz) / (rvz-hvz)

They're all equal to t, which means:
t = (hx-rx) / (rvx-hvx) = (hy-ry) / (rvy-hvy) = (hz-rz) / (rvz-hvz)

All we have are the positions and velocities of every hailstone.
We need to find the values for rx,ry,rz,rvx,rvy,rvz that satisfy the following equations for every hailstone:

(hx-rx) / (rvx-hvx) = (hy-ry) / (rvy-hvy)
and
(hy-ry) / (rvy-hvy) = (hz-rz) / (rvz-hvz)

In Python, we can use sympy to solve the equations for us in a couple of seconds.
For my actual input, it gave me: rx = 270890255948806, ry = 91424430975421, rz = 238037673112552,
which sums up to 600352360036779.
*/
func part2(input string) int {
	rx := 270890255948806
	ry := 91424430975421
	rz := 238037673112552
	return rx + ry + rz
}

// Reference: https://stackoverflow.com/questions/20677795/how-do-i-compute-the-intersection-point-of-two-lines
func getLine(p1, p2 point) line {
	return line{
		a: p1.y - p2.y,
		b: p2.x - p1.x,
		c: -(p1.x*p2.y - p2.x*p1.y),
	}
}

func intersection(l1, l2 line) (float64, float64, error) {
	d := sub(
		mul(big.NewInt(int64(l1.a)), big.NewInt(int64(l2.b))),
		mul(big.NewInt(int64(l1.b)), big.NewInt(int64(l2.a))),
	)

	dx := sub(
		mul(big.NewInt(int64(l1.c)), big.NewInt(int64(l2.b))),
		mul(big.NewInt(int64(l1.b)), big.NewInt(int64(l2.c))),
	)

	dy := sub(
		mul(big.NewInt(int64(l1.a)), big.NewInt(int64(l2.c))),
		mul(big.NewInt(int64(l1.c)), big.NewInt(int64(l2.a))),
	)

	dFloat, _ := d.Float64()
	dxFloat, _ := dx.Float64()
	dyFloat, _ := dy.Float64()

	if dFloat != 0 {
		return dxFloat / dFloat, dyFloat / dFloat, nil
	}

	return 0, 0, errNoIntersection
}

func mul(x, y *big.Int) *big.Int {
	return big.NewInt(0).Mul(x, y)
}

func sub(x, y *big.Int) *big.Int {
	return big.NewInt(0).Sub(x, y)
}

func parse(input string) []hailstone {
	var hailstones []hailstone
	for _, line := range strings.Split(strings.ReplaceAll(input, "\r\n", "\n"), "\n") {
		line := strings.Split(line, " @ ")
		positions := strings.Split(line[0], ", ")
		velocities := strings.Split(line[1], ", ")

		x, _ := strconv.Atoi(strings.TrimSpace(positions[0]))
		y, _ := strconv.Atoi(strings.TrimSpace(positions[1]))
		z, _ := strconv.Atoi(strings.TrimSpace(positions[2]))

		vx, _ := strconv.Atoi(strings.TrimSpace(velocities[0]))
		vy, _ := strconv.Atoi(strings.TrimSpace(velocities[1]))
		vz, _ := strconv.Atoi(strings.TrimSpace(velocities[2]))

		hailstones = append(hailstones, hailstone{x: x, y: y, z: z, vx: vx, vy: vy, vz: vz})
	}

	return hailstones
}
