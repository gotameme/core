/*
Copyright (c) 2024 Sebastian Kroczek <me@xbug.de>

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/

package simulation

import "github.com/gotameme/core/ant"

type Properties struct {
	Speed    AntSpeed
	Rotation AntRotation
	Load     AntLoad
	Vision   AntVision
	Range    int
	Energy   AntEnergy
	Attack   AntAttack
}

func NewDefaultProperties(defaultRange float64) Properties {
	return Properties{
		Speed:    DefaultSpeed,
		Rotation: DefaultRotation,
		Load:     DefaultLoad,
		Vision:   DefaultVision,
		Range:    int(defaultRange * DefaultRange),
		Energy:   DefaultEnergy,
		Attack:   DefaultAttack,
	}
}

func ApplyAntRole(defaultRange float64, role ant.Adjustments) Properties {
	a := Properties{}
	a.Speed = ToAntSpeed(role.Speed)
	a.Rotation = ToAntRotation(role.Rotation)
	a.Load = ToAntLoad(role.Load)
	a.Vision = ToAntVision(role.Vision)
	a.Range = int(defaultRange * ToAntRange(role.Range))
	a.Energy = ToAntEnergy(role.Energy)
	a.Attack = ToAntAttack(role.Attack)
	return a
}

// region AntSpeed

// AntSpeed Adjustments for speed
type AntSpeed float64

const (
	DecreasedSpeed AntSpeed = 3
	DefaultSpeed            = 4
	IncreasedSpeed          = 5
	BoostedSpeed            = 6
)

func ToAntSpeed(a ant.Level) AntSpeed {
	switch a {
	case ant.Decreased:
		return DecreasedSpeed
	case ant.Default:
		return DefaultSpeed
	case ant.Increased:
		return IncreasedSpeed
	case ant.Boosted:
		return BoostedSpeed
	default:
		return DefaultSpeed
	}
}

// endregion

// region AntRotation

// AntRotation Adjustments for rotation
type AntRotation int

const (
	DecreasedRotation AntRotation = 6
	DefaultRotation               = 8
	IncreasedRotation             = 12
	BoostedRotation               = 16
)

func ToAntRotation(a ant.Level) AntRotation {
	switch a {
	case ant.Decreased:
		return DecreasedRotation
	case ant.Default:
		return DefaultRotation
	case ant.Increased:
		return IncreasedRotation
	case ant.Boosted:
		return BoostedRotation
	default:
		return DefaultRotation
	}
}

// endregion

// region AntLoad

// AntLoad Adjustments for load
type AntLoad int

const (
	DecreasedLoad AntLoad = 4
	DefaultLoad           = 5
	IncreasedLoad         = 7
	BoostedLoad           = 10
)

func ToAntLoad(a ant.Level) AntLoad {
	switch a {
	case ant.Decreased:
		return DecreasedLoad
	case ant.Default:
		return DefaultLoad
	case ant.Increased:
		return IncreasedLoad
	case ant.Boosted:
		return BoostedLoad
	default:
		return DefaultLoad
	}
}

// endregion

// region AntVision

// AntVision Adjustments for vision
type AntVision int

const (
	DecreasedVision AntVision = 45
	DefaultVision             = 60
	IncreasedVision           = 75
	BoostedVision             = 90
)

func ToAntVision(a ant.Level) AntVision {
	switch a {
	case ant.Decreased:
		return DecreasedVision
	case ant.Default:
		return DefaultVision
	case ant.Increased:
		return IncreasedVision
	case ant.Boosted:
		return BoostedVision
	default:
		return DefaultVision
	}
}

// endregion

// region AntRange

// AntRange Adjustments for range
type AntRange float64

const (
	DecreasedRange AntRange = 0.75
	DefaultRange            = 1.0
	IncreasedRange          = 1.5
	BoostedRange            = 2
)

func ToAntRange(a ant.Level) float64 {
	switch a {
	case ant.Decreased:
		return float64(DecreasedRange)
	case ant.Default:
		return DefaultRange
	case ant.Increased:
		return IncreasedRange
	case ant.Boosted:
		return BoostedRange
	default:
		return DefaultRange
	}
}

// endregion

// region AntEnergy

// AntEnergy Adjustments for energy
type AntEnergy int

const (
	DecreasedEnergy AntEnergy = 50
	DefaultEnergy             = 100
	IncreasedEnergy           = 175
	BoostedEnergy             = 250
)

func ToAntEnergy(a ant.Level) AntEnergy {
	switch a {
	case ant.Decreased:
		return DecreasedEnergy
	case ant.Default:
		return DefaultEnergy
	case ant.Increased:
		return IncreasedEnergy
	case ant.Boosted:
		return BoostedEnergy
	default:
		return DefaultEnergy
	}
}

// endregion

// region AntAttack

// AntAttack Adjustments for attack
type AntAttack int

const (
	DecreasedAttack AntAttack = 0
	DefaultAttack             = 10
	IncreasedAttack           = 20
	BoostedAttack             = 30
)

func ToAntAttack(a ant.Level) AntAttack {
	switch a {
	case ant.Decreased:
		return DecreasedAttack
	case ant.Default:
		return DefaultAttack
	case ant.Increased:
		return IncreasedAttack
	case ant.Boosted:
		return BoostedAttack
	default:
		return DefaultAttack
	}
}

// endregion
