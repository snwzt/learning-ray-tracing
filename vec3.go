package main

import (
	"math"
)

type Vec3 struct {
	e [3]float64
}

func NewVec3(e0, e1, e2 float64) Vec3 {
	return Vec3{e: [3]float64{e0, e1, e2}}
}

func (v *Vec3) X() float64 {
	return v.e[0]
}

func (v *Vec3) Y() float64 {
	return v.e[1]
}

func (v *Vec3) Z() float64 {
	return v.e[2]
}

func (v *Vec3) Neg() Vec3 {
	return Vec3{e: [3]float64{-v.e[0], -v.e[1], -v.e[2]}}
}

func (v *Vec3) At(i int) float64 {
	return v.e[i]
}

func (v *Vec3) Add(v2 Vec3) {
	v.e[0] += v2.e[0]
	v.e[1] += v2.e[1]
	v.e[2] += v2.e[2]
}

func (v *Vec3) Mul(t float64) {
	v.e[0] *= t
	v.e[1] *= t
	v.e[2] *= t
}

func (v *Vec3) Div(t float64) {
	v.Mul(1 / t)
}

func (v *Vec3) Length() float64 {
	return math.Sqrt(v.LengthSquared())
}

func (v *Vec3) LengthSquared() float64 {
	return v.e[0]*v.e[0] + v.e[1]*v.e[1] + v.e[2]*v.e[2]
}

func AddVec3(u, v Vec3) Vec3 {
	return NewVec3(u.e[0]+v.e[0], u.e[1]+v.e[1], u.e[2]+v.e[2])
}

func SubVec3(u, v Vec3) Vec3 {
	return NewVec3(u.e[0]-v.e[0], u.e[1]-v.e[1], u.e[2]-v.e[2])
}

func MulVec3(u, v Vec3) Vec3 {
	return NewVec3(u.e[0]*v.e[0], u.e[1]*v.e[1], u.e[2]*v.e[2])
}

func ScaleVec3(t float64, v Vec3) Vec3 {
	return NewVec3(t*v.e[0], t*v.e[1], t*v.e[2])
}

func Dot(u, v Vec3) float64 {
	return u.e[0]*v.e[0] + u.e[1]*v.e[1] + u.e[2]*v.e[2]
}

func Cross(u, v Vec3) Vec3 {
	return NewVec3(u.e[1]*v.e[2]-u.e[2]*v.e[1], u.e[2]*v.e[0]-u.e[0]*v.e[2], u.e[0]*v.e[1]-u.e[1]*v.e[0])
}

func UnitVector(v Vec3) Vec3 {
	return ScaleVec3(1/v.Length(), v)
}
