package main

type Vector struct {
	x int
	y int
}

func checkInBounds(X Vector, bounds Vector) bool {
	return X.x >= 0 && X.x < bounds.x && X.y >= 0 && X.y < bounds.y
}

func (X Vector) Equals(Y Vector) bool {
	return X.x == Y.x && X.y == Y.y
}

func (X Vector) Add(Y Vector) Vector {
	return Vector{X.x + Y.x, X.y + Y.y}
}

func (X Vector) times(x int) Vector {
	return Vector{X.x * x, X.y * x}
}

func (X Vector) Revert() Vector {
	return Vector{-X.x, -X.y}
}

func (X Vector) Dot(Y Vector) int {
	return X.x*Y.y + X.y*Y.y
}

func (X Vector) Mod2() int {
	return X.Dot(X)
}

func (X Vector) Mod(vec Vector) Vector {
	return Vector{(X.x%vec.x + vec.x) % vec.x, (X.y%vec.y + vec.y) % vec.y}
}

func (X Vector) Orientation(Y Vector) int {
	product := X.x*Y.y + X.y*Y.x
	if product > 0 {
		return 1
	} else if product == 0 {
		return 0
	} else {
		return -1
	}
}

func (X Vector) Rotate() Vector {
	return Vector{-X.y, X.x}
}

func (X Vector) RotateInverse() Vector {
	return Vector{X.y, -X.x}
}

func (X Vector) inList(vecList []Vector) bool {
	for _, pos := range vecList {
		if X.x == pos.x && X.y == pos.y {
			return true
		}
	}
	return false

}
