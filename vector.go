package vector

import ("math")

type Vector struct {
    x, y, z float64
}

func (v1 Vector) add(v2 Vector) Vector {
    return Vector{
        v1.x + v2.x,
        v1.y + v2.y,
        v1.z + v2.z,
    }
}

func (v1 Vector) sub(v2 Vector) Vector {
    return Vector{
        v1.x - v2.x,
        v1.y - v2.y,
        v1.z - v2.z,
    }
}

func (v1 Vector) mul(v2 Vector) Vector {
    return Vector{
        v1.x * v2.x,
        v1.y * v2.y,
        v1.z * v2.z,
    }
}

func (v1 Vector) mag() float64 {
    return math.Sqrt(v1.x*v1.x + v1.y*v1.y + v1.z*v1.z)
}

func (v1 Vector) dot(v2 Vector) float64 {
    return v1.x*v2.x + v1.y*v2.y + v1.z*v2.z
}

func (v1 Vector) angle(v2 Vector) float64 {
    theta := v1.dot(v2) / (v1.mag() * v2.mag())
    return math.Acos(theta) * (180/math.Pi)
}

func (v1 Vector) orthogonal(v2 Vector) bool {
    return v1.dot(v2) == 0
}

func (v1 Vector) distance(v2 Vector) float64 {
    return v1.sub(v2).mag()
}

func (v1 Vector) norm() Vector {
    mag := v1.mag()

    return Vector{
        v1.x / mag,
        v1.y / mag,
        v1.z / mag,
    }
}

func (v1 Vector) negative() Vector {
    return Vector{
        -1*v1.x,
        -1*v1.y,
        -1*v1.z,
    }
}
