package main

import (
	"bufio"
	"strings"
	"testing"
)

func TestStringSquare(t *testing.T) {
	w := NewWorld(3, 3)
	expected := `
...
...
...`
	assertEquals(t, w.String(), expected)
}

func TestStringCols(t *testing.T) {
	w := NewWorld(1, 3)
	expected := `
...`
	assertEquals(t, w.String(), expected)
}

func TestStringRows(t *testing.T) {
	w := NewWorld(3, 1)
	expected := `
.
.
.`
	assertEquals(t, w.String(), expected)
}

func TestStringRectangle(t *testing.T) {
	w := NewWorld(3, 5)
	expected := `
.....
.....
.....`
	assertEquals(t, w.String(), expected)
}

func TestAwakenCorner(t *testing.T) {
	w := NewWorld(3, 5)
	w.Awaken(0, 0)
	expected := `
*....
.....
.....`
	assertEquals(t, w.String(), expected)
}

func TestAwakenMiddle(t *testing.T) {
	w := NewWorld(3, 5)
	w.Awaken(1, 3)
	expected := `
.....
...*.
.....`
	assertEquals(t, w.String(), expected)
}

func TestAwakenAboveBoundsY(t *testing.T) {
	w := NewWorld(3, 5)
	w.Awaken(3, 3)
	expected := `
.....
.....
.....`
	assertEquals(t, w.String(), expected)
}

func TestAwakenAboveBoundsX(t *testing.T) {
	w := NewWorld(3, 5)
	w.Awaken(1, 5)
	expected := `
.....
.....
.....`
	assertEquals(t, w.String(), expected)
}

func TestAwakenBelowBoundsY(t *testing.T) {
	w := NewWorld(3, 5)
	w.Awaken(-1, 3)
	expected := `
.....
.....
.....`
	assertEquals(t, w.String(), expected)
}

func TestAwakenBelowBoundsX(t *testing.T) {
	w := NewWorld(3, 5)
	w.Awaken(1, -1)
	expected := `
.....
.....
.....`
	assertEquals(t, w.String(), expected)
}

func TestParseEmptyWorld(t *testing.T) {
	worldStr := `3 5
.....
.....
.....`
	w := parseWorld(worldStr)
	expected := `
.....
.....
.....`
	assertEquals(t, w.String(), expected)
}

func TestParseWorldWithLiveCells(t *testing.T) {
	worldStr := `3 5
.*...
..*..
...*.`
	w := parseWorld(worldStr)
	expected := `
.*...
..*..
...*.`
	assertEquals(t, w.String(), expected)
}

func TestLoneCellDies(t *testing.T) {
	w := NewWorld(3, 5)
	w.Awaken(1, 3)
	w.Advance()
	expected := `
.....
.....
.....`
	assertEquals(t, w.String(), expected)
}

func TestCellsWithFewerThanTwoLiveNeighboursDie(t *testing.T) {
	worldStr := `3 5
.*...
..*..
...*.`
	w := parseWorld(worldStr)

	w.Advance()
	expectedNextGen := `
.....
..*..
.....`
	assertEquals(t, w.String(), expectedNextGen)
}

func TestCellsWithMoreThanThreeLiveNeighboursDie(t *testing.T) {
	worldStr := `3 5
.*.*.
..*..
.*.*.`
	w := parseWorld(worldStr)

	w.Advance()
	expectedNextGen := `
..*..
.*.*.
..*..`
	assertEquals(t, w.String(), expectedNextGen)
}

func TestDeadCellWithExactlyThreeLiveNeighboursAwakes(t *testing.T) {
	worldStr := `3 5
..*..
.**..
.....`
	w := parseWorld(worldStr)

	w.Advance()
	expectedNextGen := `
.**..
.**..
.....`
	assertEquals(t, w.String(), expectedNextGen)
}

func parseWorld(worldStr string) World {
	return ParseWorld(bufio.NewReader(strings.NewReader(worldStr)))
}

func assertEquals(t *testing.T, actual, expected string) {
	if actual != expected {
		t.Errorf("Wanted '%s', got '%s'", expected, actual)
	}
}
