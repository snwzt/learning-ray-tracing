package main

import (
	"fmt"
)

type Color = Vec3

func WriteColor(c Color) string {
	r := c.X()
	g := c.Y()
	b := c.Z()

	rByte := int(255.999 * r)
	gByte := int(255.999 * g)
	bByte := int(255.999 * b)

	return fmt.Sprintf("%d %d %d\n", rByte, gByte, bByte)
}
