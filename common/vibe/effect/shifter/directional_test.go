package shifter

// import (
// 	"testing"

// 	"github.com/jmbarzee/show/common"
// 	"github.com/jmbarzee/show/common/space"
// 	"github.com/jmbarzee/show/common/testutil"
// 	"github.com/jmbarzee/show/common/vibe/effect/bender"
// 	helper "github.com/jmbarzee/show/common/vibe/testhelper"
// )

// func TestDirectionalShift(t *testing.T) {
// 	aFloat := 1.1
// 	cases := []ShiftTest{
// 		{
// 			Name: "One shift per second",
// 			Shifter: &Directional{
// 				PhiBender: &bender.Static{
// 					TheBend: &aFloat,
// 				},
// 				ThetaBender: &bender.Static{
// 					TheBend: &aFloat,
// 				},
// 			},
// 			Instants: []Instant{
// 				{
// 					Light: &testutil.Light{
// 						Orientation: space.Spherical{
// 							R: 1,
// 							P: 1,
// 							T: 2,
// 						},
// 					},
// 					ExpectedShift: aFloat * 2,
// 				},
// 				{
// 					Light: &testutil.Light{
// 						Orientation: space.Spherical{
// 							R: 1,
// 							P: -1,
// 							T: -2,
// 						},
// 					},
// 					ExpectedShift: aFloat * 2,
// 				},
// 				{
// 					Light: &testutil.Light{
// 						Orientation: space.Spherical{
// 							R: 1,
// 							P: 0,
// 							T: 0,
// 						},
// 					},
// 					ExpectedShift: aFloat * 2,
// 				},
// 			},
// 		},
// 	}
// 	RunShifterTests(t, cases)
// }
// func TestDirectionalGetStabilizeFuncs(t *testing.T) {
// 	aFloat := 1.1
// 	c := helper.StabilizeableTest{
// 		Stabilizable: &Directional{},
// 		ExpectedVersions: []common.Stabilizable{
// 			&Directional{
// 				PhiBender: &bender.Static{},
// 			},
// 			&Directional{
// 				PhiBender: &bender.Static{
// 					TheBend: &aFloat, // This is a little dirty. The Benders are both/all pointing to the same struct, so TheBend is set with the first bender
// 				},
// 			},
// 			&Directional{
// 				PhiBender: &bender.Static{
// 					TheBend: &aFloat,
// 				},
// 				ThetaBender: &bender.Static{
// 					TheBend: &aFloat,
// 				},
// 			},
// 		},
// 		Palette: helper.TestPalette{
// 			Bender: &bender.Static{},
// 			Shift:  aFloat,
// 		},
// 	}
// 	helper.RunStabilizeableTest(t, c)
// }
