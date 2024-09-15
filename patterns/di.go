package main

import "fmt"

type SafetyPlacer interface {
	placeSafeties()
}

type RockClimber struct {
	kind         int
	rocksClimbed int
	sp           SafetyPlacer
}

func newRockClimber(sp SafetyPlacer) *RockClimber {
	return &RockClimber{
		sp: sp,
	}
}

type IceSafetyPlacer struct {
	//db
}

func (sp IceSafetyPlacer) placeSafeties() {
	fmt.Println("Placing my ICE Safeties...")
}

type NOPSafetyPlacer struct{}

func (sp NOPSafetyPlacer) placeSafeties() {
	fmt.Println("placing NO safeties...")
}

func (rc *RockClimber) climbRock() {
	rc.rocksClimbed++
	if rc.rocksClimbed == 10 {
		rc.sp.placeSafeties()
	}
}
