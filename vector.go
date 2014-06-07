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

func (v1 Vector) magnitude() float64 {
    return math.Sqrt(v1.x*v1.x + v1.y*v1.y + v1.z*v1.z)
}

func (v1 Vector) dot(v2 Vector) float64 {
    return v1.x*v2.x + v1.y*v2.y + v1.z*v2.z
}

func (v1 Vector) angle(v2 Vector) float64 {
    theta := v1.dot(v2) / (v1.magnitude() * v2.magnitude())
    return math.Acos(theta) * (180/math.Pi)
}

func (v1 Vector) orthogonal(v2 Vector) bool {
    return v1.dot(v2) == 0
}
