package vibe

import (
	"testing"
	"time"

	"github.com/jmbarzee/show/common"
	"github.com/jmbarzee/show/common/color"
	"github.com/jmbarzee/show/common/repeat"
	"github.com/jmbarzee/show/common/vibe/effect"
	"github.com/jmbarzee/show/common/vibe/effect/bender"
	"github.com/jmbarzee/show/common/vibe/effect/painter"
	"github.com/jmbarzee/show/common/vibe/effect/shifter"
	"github.com/jmbarzee/show/common/vibe/palette"
	"github.com/jmbarzee/show/common/vibe/span"
	"github.com/jmbarzee/show/common/vibe/testutil"
)

func TestBasicStabilize(t *testing.T) {
	aTime1 := time.Date(2009, 11, 17, 20, 34, 50, 651387237, time.UTC)
	theTruth := true
	aDuration := time.Nanosecond * 2245197264
	aSpan := span.New(aTime1, aTime1.Add(time.Hour))
	aSeed := repeat.NewSeed(aTime1)

	cases := []StabilizeTest{
		{
			Name: "Basic Vibe",
			ActualVibe: &Basic{
				Span:    aSpan,
				Palette: palette.NewRandom(aSeed),
			},
			ExpectedVibes: []common.Vibe{
				&Basic{
					Span:    aSpan,
					Palette: palette.NewRandom(aSeed),
					effects: []common.Effect{
						&effect.Future{
							BasicEffect: effect.BasicEffect{Spanner: aSpan},
						},
					},
				},
				&Basic{
					Span:    aSpan,
					Palette: palette.NewRandom(aSeed),
					effects: []common.Effect{
						&effect.Future{
							BasicEffect:  effect.BasicEffect{Spanner: aSpan},
							TimePerLight: &aDuration,
						},
					},
				},
				&Basic{
					Span:    aSpan,
					Palette: palette.NewRandom(aSeed),
					effects: []common.Effect{
						&effect.Future{
							BasicEffect:  effect.BasicEffect{Spanner: aSpan},
							TimePerLight: &aDuration,
							Painter:      &painter.Bounce{},
						},
					},
				},
				&Basic{
					Span:    aSpan,
					Palette: palette.NewRandom(aSeed),
					effects: []common.Effect{
						&effect.Future{
							BasicEffect:  effect.BasicEffect{Spanner: aSpan},
							TimePerLight: &aDuration,
							Painter: &painter.Bounce{
								ColorEnd: color.RedMagenta,
							},
						},
					},
				},
				&Basic{
					Span:    aSpan,
					Palette: palette.NewRandom(aSeed),
					effects: []common.Effect{
						&effect.Future{
							BasicEffect:  effect.BasicEffect{Spanner: aSpan},
							TimePerLight: &aDuration,
							Painter: &painter.Bounce{
								ColorEnd: color.RedMagenta,
								Shifter:  &shifter.Combo{},
							},
						},
					},
				},
				&Basic{
					Span:    aSpan,
					Palette: palette.NewRandom(aSeed),
					effects: []common.Effect{
						&effect.Future{
							BasicEffect:  effect.BasicEffect{Spanner: aSpan},
							TimePerLight: &aDuration,
							Painter: &painter.Bounce{
								ColorEnd: color.RedMagenta,
								Shifter: &shifter.Combo{
									B: &shifter.Locational{},
								},
							},
						},
					},
				},
				&Basic{
					Span:    aSpan,
					Palette: palette.NewRandom(aSeed),
					effects: []common.Effect{
						&effect.Future{
							BasicEffect:  effect.BasicEffect{Spanner: aSpan},
							TimePerLight: &aDuration,
							Painter: &painter.Bounce{
								ColorStart: color.Red,
								ColorEnd:   color.RedMagenta,
								Shifter: &shifter.Combo{
									B: &shifter.Locational{},
								},
							},
						},
					},
				},
				&Basic{
					Span:    aSpan,
					Palette: palette.NewRandom(aSeed),
					effects: []common.Effect{
						&effect.Future{
							BasicEffect:  effect.BasicEffect{Spanner: aSpan},
							TimePerLight: &aDuration,
							Painter: &painter.Bounce{
								ColorStart: color.Red,
								ColorEnd:   color.RedMagenta,
								Up:         &theTruth,
								Shifter: &shifter.Combo{
									B: &shifter.Locational{},
								},
							},
						},
					},
				},
				&Basic{
					Span:    aSpan,
					Palette: palette.NewRandom(aSeed),
					effects: []common.Effect{
						&effect.Future{
							BasicEffect:  effect.BasicEffect{Spanner: aSpan},
							TimePerLight: &aDuration,
							Painter: &painter.Bounce{
								ColorStart: color.Red,
								ColorEnd:   color.RedMagenta,
								Up:         &theTruth,
								Shifter: &shifter.Combo{
									B: &shifter.Locational{
										ZBender: &bender.Linear{Coefficient: 1.1},
									},
								},
							},
						},
					},
				},
				&Basic{
					Span:    aSpan,
					Palette: palette.NewRandom(aSeed),
					effects: []common.Effect{
						&effect.Future{
							BasicEffect:  effect.BasicEffect{Spanner: aSpan},
							TimePerLight: &aDuration,
							Painter: &painter.Bounce{
								ColorStart: color.Red,
								ColorEnd:   color.RedMagenta,
								Up:         &theTruth,
								Shifter: &shifter.Combo{
									B: &shifter.Locational{
										YBender: &bender.Linear{Coefficient: 1.1},
										ZBender: &bender.Linear{Coefficient: 1.1},
									},
								},
							},
						},
					},
				},
				&Basic{
					Span:    aSpan,
					Palette: palette.NewRandom(aSeed),
					effects: []common.Effect{
						&effect.Future{
							BasicEffect:  effect.BasicEffect{Spanner: aSpan},
							TimePerLight: &aDuration,
							Painter: &painter.Bounce{
								ColorStart: color.Red,
								ColorEnd:   color.RedMagenta,
								Up:         &theTruth,
								Shifter: &shifter.Combo{
									A: &shifter.Positional{},
									B: &shifter.Locational{
										YBender: &bender.Linear{Coefficient: 1.1},
										ZBender: &bender.Linear{Coefficient: 1.1},
									},
								},
							},
						},
					},
				},
				&Basic{
					Span:    aSpan,
					Palette: palette.NewRandom(aSeed),
					effects: []common.Effect{
						&effect.Future{
							BasicEffect:  effect.BasicEffect{Spanner: aSpan},
							TimePerLight: &aDuration,
							Painter: &painter.Bounce{
								ColorStart: color.Red,
								ColorEnd:   color.RedMagenta,
								Up:         &theTruth,
								Shifter: &shifter.Combo{
									A: &shifter.Positional{},
									B: &shifter.Locational{
										XBender: &bender.Linear{Coefficient: 1.1},
										YBender: &bender.Linear{Coefficient: 1.1},
										ZBender: &bender.Linear{Coefficient: 1.1},
									},
								},
							},
						},
					},
				},
				&Basic{
					Span:    aSpan,
					Palette: palette.NewRandom(aSeed),
					effects: []common.Effect{
						&effect.Future{
							BasicEffect:  effect.BasicEffect{Spanner: aSpan},
							TimePerLight: &aDuration,
							Painter: &painter.Bounce{
								ColorStart: color.Red,
								ColorEnd:   color.RedMagenta,
								Up:         &theTruth,
								Shifter: &shifter.Combo{
									A: &shifter.Positional{
										Bender: &bender.Exponential{Exponent: 2.0, Coefficient: 1.1},
									},
									B: &shifter.Locational{
										XBender: &bender.Linear{Coefficient: 1.1},
										YBender: &bender.Linear{Coefficient: 1.1},
										ZBender: &bender.Linear{Coefficient: 1.1},
									},
								},
							},
						},
					},
				},
			},
		},
	}

	RunStabilizeTests(t, cases)
}

func TestBasicMaterialize(t *testing.T) {
	aTime1 := time.Date(2009, 11, 17, 20, 34, 50, 651387237, time.UTC)
	aDuration := time.Nanosecond * 2785814474
	aSpan := span.New(aTime1, aTime1.Add(time.Hour))
	aSeed := repeat.NewSeed(aTime1)

	cases := []MaterializeTest{
		{
			Name: "Basic Vibe",
			ActualVibe: &Basic{
				Span:    aSpan,
				Palette: palette.NewRandom(aSeed),
			},
			ExpectedVibe: &Basic{
				Span:    aSpan,
				Palette: palette.NewRandom(aSeed),
				effects: []common.Effect{
					&effect.Future{
						BasicEffect:  effect.BasicEffect{Spanner: aSpan},
						TimePerLight: &aDuration,
						Painter: &painter.Move{
							ColorStart: color.WarmCyan,
							Shifter: &shifter.Positional{
								Bender: &bender.Sinusoidal{
									Offset:    0.5,
									Period:    1.0,
									Amplitude: 1.4,
								},
							},
						},
					},
				},
			},
		},
	}

	RunMaterializeTests(t, cases)
}
func TestBasicGetStabilizeFuncs(t *testing.T) {
	aTime1 := time.Date(2009, 11, 17, 20, 34, 50, 651387237, time.UTC)
	aSpan := span.New(aTime1, aTime1.Add(time.Hour))

	c := testutil.StabilizerTest{
		Stabilizer: &Basic{},
		ExpectedVersions: []common.Stabilizer{
			&Basic{
				effects: []common.Effect{
					&effect.Solid{
						BasicEffect: effect.BasicEffect{Spanner: aSpan},
					},
				},
			},
			&Basic{
				effects: []common.Effect{
					&effect.Solid{
						BasicEffect: effect.BasicEffect{Spanner: aSpan},
						Painter:     &painter.Static{},
					},
				},
			},
			&Basic{
				effects: []common.Effect{
					&effect.Solid{
						BasicEffect: effect.BasicEffect{Spanner: aSpan},
						Painter: &painter.Static{
							Color: color.Blue,
						},
					},
				},
			},
		},
		Palette: testutil.TestPalette{
			Color:   color.Blue,
			Painter: &painter.Static{},
			Effect:  &effect.Future{},
		},
	}
	testutil.RunStabilizerTest(t, c)
}
