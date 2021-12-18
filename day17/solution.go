package main

type Target struct {
	minx, maxx int
	miny, maxy int
}

func Simulate(vx, vy int, target Target) (int, bool) {
	x, y := 0, 0
	highestY := 0

	for {
		x += vx
		y += vy
		if vx > 0 {
			vx--
		}
		if vx < 0 {
			vx++
		}
		vy -= 1
		if y > highestY {
			highestY = y
		}
		if x >= target.minx && x <= target.maxx && y >= target.miny && y <= target.maxy {
			return highestY, true
		}
		if x > target.maxx || y < target.miny {
			break
		}
	}
	return highestY, false
}

func BruteForceHighestY(target Target) int {
	h := 0
	for vx := 0; vx < (target.maxx / 2); vx++ {
		for vy := target.miny; vy < (target.maxx / 2); vy++ {
			highest, hit := Simulate(vx, vy, target)
			if hit && highest > h {
				h = highest
			}
		}
	}
	return h
}

func BruteForceDistinctVectors(target Target) (count int) {
	for vx := -target.maxx; vx <= target.maxx; vx++ {
		for vy := -target.maxx; vy <= target.maxx; vy++ {
			_, hit := Simulate(vx, vy, target)
			if hit {
				count++
			}
		}
	}
	return
}
