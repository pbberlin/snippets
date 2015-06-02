package transposablematrix

// FindClosestMatch takes slice with sections of format
// [idx section][idxP1 idxP2 DX westwDY westwDX eastwDY eastwDX].
func (ar *Reservoir) FindClosestMatch(perspective int, l []Point, sections [][]int) (*Amorph, Point, Point) {

	var r *Amorph
	var PStitchTo, POffset, PKerch Point

	for i := 0; i < len(sections); i++ {
		dxb, dyw, dxw, dyw2, dye, dxe, dye2 := sections[i][2], sections[i][3], sections[i][4], sections[i][5], sections[i][6], sections[i][7], sections[i][8]
		_, _ = dxw, dyw2

		westwTraj, eastwTraj := convexConcaveSectionSides(sections[i])
		// pf("%v\t\t %v %v\n", sections[i], westwTraj, eastwTraj)

		if westwTraj == Convex && eastwTraj == Concave {
			pf("%v|%v srch w%v(%v+%v)s%v: ", westwTraj, eastwTraj, dxb+dxe,
				dxb, dxe, dye*dxe)
			r, PKerch, POffset = ar.searchAvenue12(perspective, dxb, dye, dxe, dye2, dyw)
			PStitchTo.x, PStitchTo.y = l[i].x+POffset.x, l[i].y+POffset.y

			if false || true {
				// Avenue 3
				dxb = dye
				dye = dxb
				dxe = dyw
				dye2 = dxw
				dyw = dxe
				pf("%v|%v srch w%v(%v+%v)s%v: ", westwTraj, eastwTraj, dxb+dxe,
					dxb, dxe, dye*dxe)
				r, PKerch, POffset = ar.searchAvenue12(perspective+1, dxb, dye, dxe, dye2, dyw)
				PStitchTo.x, PStitchTo.y = l[i].x+POffset.x, l[i].y+POffset.y
			}

			// if r == nil {
			// 	pf("WestwardHoFor %v (%v+%v %v) %v %v ... ", dxb+dxw, dxb, dxw, dyw, westwTraj, eastwTraj)
			// 	r, PKerch, POffset = ar.searchAvenue12(perspective, dxb, dyw, dxw, dyw2, dyw)
			// }
			pf("\n")
		}

		break // only first

	}
	// r = f(initSize+2, 0, 0) // this must cause panic
	return r, PStitchTo, PKerch
}

func (ar *Reservoir) searchAvenue12(perspective, xb, dy, dx, dy2, dyw int) (*Amorph, Point, Point) {

	// I am not sure, why width always works?
	// It's possibly, because of "diagonal symmetry"?
	// Because amorph w4h7_s2 is equivalent w7h4_s2 ?
	// And because Amorph.transpose() is normalizing inner coordinates anyway?
	f := ar.WidthOver

	// The previous code was:
	// f := ar.HeightOver
	// if perspective%2 == 0 {
	// 	f = ar.WidthOver
	// }

	// Avenue 1
	// ========================
	pf("Av.1. pfct ")
	if false {
		match := f(xb+dx, dy*dx, 0, true)
		if match != nil {
			pf("found w%vs%v(%v)\n", match.Cols, match.Slack, match.Rows)
			return match, Point{dx, dy}, Point{}
		}

		pf("Av.1 incr...")
		for i := dx - ar.SmallestDesirableGap(perspective); i >= 0; i-- { // columnwise - smaller columns
			pf("%v-", xb+i)
			match := f(xb+i, i*dy, 0, true)
			if match != nil {
				pf("found w%vs%v(%v)\n", match.Cols, match.Slack, match.Rows)
				return match, Point{i, dy}, Point{}
			}
		}
	}

	// Avenue 2
	// ========================
	hjoMin, hjoMax := 0, 0 // horizontalJuttingOut
	if dyw < ar.SmallestDesirableGap(perspective+1) {
		hjoMin, hjoMax = 1, ar.SmallestDesirableGap(perspective)+1 // prevent tiny stripes
	} else if dyw >= ar.SmallestDesirableGap(perspective+1) || dyw == 0 {
		hjoMin, hjoMax = 1, 2*ar.SmallestDesirableGap(perspective)
	}
	pf("Av.2 pfct...")
	for j := hjoMin; j <= hjoMax; j++ {
		pf("%v-", j)
		match := f(xb+dx+j, dy*dx, 0, true)
		if match != nil {
			pf("found w%vs%v(%v)\n", match.Cols, match.Slack, match.Rows)
			return match, Point{dx, dy}, Point{-j, 0}
		}
	}
	pf("Av.2 incr...")
	for j := hjoMin; j <= hjoMax; j++ {
		pf(" j%v|", j)
		for i := dx - ar.SmallestDesirableGap(perspective); i >= 0; i-- { // columnwise - smaller columns
			pf("%v-", xb+i+j)
			match := f(xb+i+j, i*dy, 0, true)
			if match != nil {
				pf("found w%vs%v(%v)\n", match.Cols, match.Slack, match.Rows)
				return match, Point{i, dy}, Point{-j, 0}
			}
		}
	}

	return nil, Point{}, Point{}
}
