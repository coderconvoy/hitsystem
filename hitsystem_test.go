package hitsystem

import (
	"testing"

	"engo.io/ecs"
)

func assert(t *testing.T, b bool, s string, args ...interface{}) {
	if !b {
		t.Errorf(s, args...)
	}
}

func Test_MinDistance(t *testing.T) {
	a := HitBox{20, 20, 20, 20}
	b := HitBox{8, 9, 20, 20}
	c := HitBox{9, 8, 20, 20}

	//right
	dx, dy := a.MinimumStepOffD(b)

	assert(t, dx == 8, "right: Dx wrong = %f", dx)
	assert(t, dy == 0, "right: Dy wrony = %f", dy)

	//left
	dx, dy = b.MinimumStepOffD(a)

	assert(t, dx == -8, "left wrong = %f", dx)
	assert(t, dy == 0, "left wrony = %f", dy)

	//top
	dx, dy = c.MinimumStepOffD(a)
	assert(t, dx == 0, "top wrong = %f", dx)
	assert(t, dy == -8, "top wrony = %f", dy)
	//bottom
	dx, dy = a.MinimumStepOffD(c)
	assert(t, dx == 0, "bottom wrong = %f", dx)
	assert(t, dy == 8, "bottom wrong = %f", dy)
}

type hitme struct {
	ecs.BasicEntity
	box         HitBox
	main, group HitGroup
}

func (hm *hitme) Push(dx, dy float32) {
	hm.box.x += dx
	hm.box.y += dy
}

func (hm *hitme) GetHitBox() HitBox {
	return hm.box
}

func (hm *hitme) HitGroups() (HitGroup, HitGroup) {
	return hm.main, hm.group
}

func Test_Update(t *testing.T) {
	hent := func(x, y, w, h float32, gm, gg HitGroup) *hitme {
		nb := ecs.NewBasic()
		return &hitme{
			BasicEntity: nb,
			box: HitBox{
				x: x, y: y, w: w, h: h},
			main:  gm,
			group: gg,
		}
	}

	ts := []*hitme{
		hent(20, 20, 20, 20, 1, 0),
		hent(7, 20, 20, 20, 0, 1),
	}

	sys := HitSystem{Solid: 1}

	for _, v := range ts {
		sys.Add(v)
	}
	sys.Update(0.001)

	hb := ts[0].GetHitBox()
	if hb.x != 27 {
		t.Errorf("No solid collision happened")
	}

}
