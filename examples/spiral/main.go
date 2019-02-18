//-----------------------------------------------------------------------------
/*

Spirals

*/
//-----------------------------------------------------------------------------

package main

import (
	. "github.com/deadsy/sdfx/sdf"
)

//-----------------------------------------------------------------------------

func main() {
	s2 := ArcSpiral2D(1.0, 5.0, 0, 10*Tau, 1.0)
	RenderDXFSlow(s2, 400, "spiral.dxf")
}

//-----------------------------------------------------------------------------
