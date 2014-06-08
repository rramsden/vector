package vector

import "testing"
import "math"
import "fmt"

func assert(cond bool, t *testing.T) {
    if cond != true {
        t.Errorf("expected %t to be true", cond, cond) 
    }
}

func TestVectorAdd(t *testing.T) {
    v1 := Vector{1,1,1}
    v2 := Vector{1,1,1}
    v3 := v1.add(v2)

    assert(v3.x == 2, t)
    assert(v3.y == 2, t)
    assert(v3.z == 2, t)
}

func TestVectorSub(t *testing.T) {
    v1 := Vector{1,1,1}
    v2 := Vector{1,1,1}
    v3 := v1.sub(v2)

    assert(v3.x == 0, t)
    assert(v3.y == 0, t)
    assert(v3.z == 0, t)
}

func TestVectorMul(t *testing.T) {
    v1 := Vector{2,2,2}
    v2 := Vector{2,2,2}
    v3 := v1.mul(v2)
   
    assert(v3.x == 4, t)
    assert(v3.y == 4, t)
    assert(v3.z == 4, t)
}

func TestVectorMag(t *testing.T) {
    v1 := Vector{2,2,2}
    magnitude := v1.mag()

    expected := math.Sqrt(v1.x+v1.x + v1.x*v1.x + v1.y*v1.y)
    assert(magnitude == expected, t)
}

func TestVectorDot(t *testing.T) {
    v1 := Vector{1,2,3}
    v2 := Vector{1,2,3}
    val := v1.dot(v2)

    expected := v1.x*v2.x + v1.y*v2.y + v1.z*v2.z
    assert(val == expected, t)
}

func TestVectorAngle(t *testing.T) {
    v2 := Vector{4,4,0}
    fmt.Println(v2.mag())
}

func TestVectorDistance(t *testing.T) {
    v1 := Vector{1,1,0}
    v2 := Vector{4,4,0}
    v3 := v1.sub(v2)

    expected := math.Sqrt(v3.x*v3.x + v3.y*v3.y + v3.z*v3.z)
    assert(v3.mag() == expected, t)
}