package main

import (
	"fmt"
	"testing"
)

func TestDay22(t *testing.T) {
	const (
		caveDepth = 510
		targetX   = 10
		targetY   = 10
	)
	cavePlan := createPlan(caveDepth, targetX, targetY)
	sum := erosionSum(cavePlan)
	if sum != 114 {
		t.Error("Expected erosion sum 114, got", sum)
	}
	dist, newPlan := findRescuePathLength(caveDepth, targetX, targetY, cavePlan)
	fmt.Println("Distance ", dist)
	displayCave(targetX, targetY, newPlan)
	if dist != 45 {
		t.Error("Expected distance 45, got", dist)
	}
}
