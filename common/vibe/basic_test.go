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
	helper "github.com/jmbarzee/show/common/vibe/testhelper"
)

func TestBasicStabilize(t *testing.T) {
	aTime1 := time.Date(2009, 11, 17, 20, 34, 50, 651387237, time.UTC)
	aFloat1 := 0.27
	aFloat2 := 0.28700000000000003
	aFloat3 := 0.45
	aFloat4 := 0.389
	aFloat5 := 0.059
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
										ZBender: &bender.Linear{},
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
										YBender: &bender.Linear{},
										ZBender: &bender.Linear{},
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
										YBender: &bender.Linear{},
										ZBender: &bender.Linear{},
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
										YBender: &bender.Linear{},
										ZBender: &bender.Linear{
											Interval: &aFloat1,
										},
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
										XBender: &bender.Linear{},
										YBender: &bender.Linear{},
										ZBender: &bender.Linear{
											Interval: &aFloat1,
										},
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
										Bender: &bender.Exponential{},
									},
									B: &shifter.Locational{
										XBender: &bender.Linear{},
										YBender: &bender.Linear{},
										ZBender: &bender.Linear{
											Interval: &aFloat1,
										},
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
										Bender: &bender.Exponential{
											Coefficient: &aFloat2,
										},
									},
									B: &shifter.Locational{
										XBender: &bender.Linear{},
										YBender: &bender.Linear{},
										ZBender: &bender.Linear{
											Interval: &aFloat1,
										},
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
										Bender: &bender.Exponential{
											Coefficient: &aFloat2,
										},
									},
									B: &shifter.Locational{
										XBender: &bender.Linear{},
										YBender: &bender.Linear{
											Interval: &aFloat3,
										},
										ZBender: &bender.Linear{
											Interval: &aFloat1,
										},
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
										Bender: &bender.Exponential{
											Coefficient: &aFloat2,
										},
									},
									B: &shifter.Locational{
										XBender: &bender.Linear{
											Interval: &aFloat4,
										},
										YBender: &bender.Linear{
											Interval: &aFloat3,
										},
										ZBender: &bender.Linear{
											Interval: &aFloat1,
										},
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
										Bender: &bender.Exponential{
											Coefficient: &aFloat2,
											Exponent:    &aFloat5,
										},
									},
									B: &shifter.Locational{
										XBender: &bender.Linear{
											Interval: &aFloat4,
										},
										YBender: &bender.Linear{
											Interval: &aFloat3,
										},
										ZBender: &bender.Linear{
											Interval: &aFloat1,
										},
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
	aFloat1 := 0.17900000000000002
	aFloat2 := 0.46900000000000003
	aFloat3 := 0.47000000000000003
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
									Offset:    &aFloat1,
									Period:    &aFloat2,
									Amplitude: &aFloat3,
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

	c := helper.StabilizerTest{
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
		Palette: helper.TestPalette{
			Color:   color.Blue,
			Painter: &painter.Static{},
			Effect:  &effect.Future{},
		},
	}
	helper.RunStabilizerTest(t, c)
}
