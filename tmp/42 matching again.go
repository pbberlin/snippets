package transposablematrix

import "github.com/pbberlin/tools/subsort"

type Bumper struct {
	l1, l2 int
	pness  int
}

func (ar *Reservoir) Match(perspective int, l []Point, clns [][]int,
	sct []int) (Point, Point, Point, []subsort.SortedByIntVal) {

	// pf("clns %v\n", clns)
	pf("sect %v\n", sct)

	pstt := Point{} // point stitch to
	pstt.x = l[sct[0]].x
	pstt.y = l[sct[0]].y

	krfl := Point{} // kerf left
	krfr := Point{} // kerf right

	//
	// *Preliminary* distinction of cases:
	var sl, sr []int // section left, section right, being fused into supersection
	{
		// probing the immediate flanks of base.
		// *fusing* two neighboring sections according to case.
		_, idxNeighborRight := findYourNeighbor(clns, sct)
		dyw, dye := sct[3], sct[6]
		if dyw <= 0 && dye > 0 { // Westside: descending or empty, east: ascending
			sl = sct
			sr = clns[idxNeighborRight]
			if idxNeighborRight == -1 { // impossible 'cause dye > 0
				sr = []int{0, 0, 0, 0, 0, 0, 0, 0, 0}
			}
		} else if dyw > 0 && dye <= 0 {
			// invert above case - mirror vertically
		} else if dyw <= 0 && dye <= 0 { // top of the world - all convex

		} else if dyw > 0 && dye > 0 { // concave case

		}
	}

	//
	// Continue on the basis of sl,sr
	yw, xw := sl[3], sl[4]
	x1 := sl[2]
	y1 := sl[6]
	x2 := sr[2]
	y2 := sr[6]
	xe, ye := sr[7], sr[8] // not sr[6], sr[7]
	_ = y2
	_, _, _, _ = yw, xw, ye, xe

	slack := y1 * x2
	pf("srch w%v(%v+%v)s%v(%v*%v): ", x1+x2,
		x1, x2, slack, y1, x2)

	bumperSW := Bumper{yw, xw, 0} // y first
	bumperNE := Bumper{xe, ye, 0} // x first
	bumperSW.permissiveness(true, perspective, ar)
	bumperNE.permissiveness(false, perspective+1, ar)
	pf("\nbumperSW %v bumperNE %v\n", bumperSW, bumperNE)

	//
	//
	var idxSortedAmorphs []subsort.SortedByIntVal
	if bumperSW.pness > 0 && bumperNE.pness > 0 {
		if bumperSW.pness >= bumperNE.pness {
			// more room for south-west expansion => go for smallest height
			pf(" restr/min height  ")
			idxSortedAmorphs = ar.SortedSection(3, slack, slack)
			if len(idxSortedAmorphs) > 0 {
				horizOff := ar.Amorphs[idxSortedAmorphs[0].IdxOrig].Cols -
					(x1 + x2)
				pstt.x -= horizOff
				krfr.x, krfr.y = x2, y1
			}
		} else {
			// north-east is most open => restrain width - minimize width
			pf(" restr/min width  ")
			idxSortedAmorphs = ar.SortedSection(2, slack, slack)
		}
	}
	ar.ApplySortedSectionDump(idxSortedAmorphs)

	{
		i := divideFloated(ar.AvgSlack, x2, ar.SmallestDesirableHeight, y1)
		pf("iNE %3.2v  ", i)
	}
	{
		i := divideFloated(ar.AvgSlack, y1, ar.SmallestDesirableWidth, x2)
		pf("iSW %3.2v  ", i)
	}

	return pstt, krfl, krfr, idxSortedAmorphs
	// pf("\n")
}

func divideFloated(div, divisor, su1, su2 int) int {
	r := float64(div) / float64(divisor)
	return int(10 * r)
}

//
// Permissiveness measures available room for.
// Permissiveness states "how free is the bumper"
// High pness means less restriction.
// values < 0 mean
func (b *Bumper) permissiveness(isWest bool, perspective int, ar *Reservoir) {

	absL1 := b.l1
	if absL1 < 0 { // l1 might be negative - as opposed to l2 which is horizontal and therefore always positive
		absL1 = -absL1
	}

	if absL1 == 0 || b.l2 <= 0 {
		b.pness = 100
	} else {
		if absL1 > ar.SmallestDesirableGap(perspective+1) {
			b.pness = 50
		} else if absL1 == ar.SmallestDesirableGap(perspective+1) {
			b.pness = 8
		} else if absL1 < ar.SmallestDesirableGap(perspective+1) &&
			b.l2 < 2*ar.SmallestDesirableGap(perspective) {
			b.pness = 2 // expand - but only if absolutely necessary
		} else {
			b.pness = -2 // do not expand
		}

	}

	if isWest && b.l1 > 0 {
		b.pness = -100 // westward concave boundary => no room at all
	}
	if !isWest && b.l2 < 0 { // => north-east
		b.pness = 100 // an north-eastward bumper - with convex side => no restriction at all
	}

}
