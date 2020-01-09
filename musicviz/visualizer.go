package main

import "math"

var vectors []Coordinates

func init() {
	vectors = []Coordinates{
		Coordinates{
			x: 0,
			y: 1,
		},
		Coordinates{
			x: 1,
			y: 0,
		},
		Coordinates{
			x: 1,
			y: 1,
		},
		Coordinates{
			x: -1,
			y: 1,
		},
		Coordinates{
			x: -1,
			y: -1,
		},
		Coordinates{
			x: 1,
			y: -1,
		},
		Coordinates{
			x: 1,
			y: 0.5,
		},
		Coordinates{
			x: 0.5,
			y: 1,
		},
		Coordinates{
			x: -0.5,
			y: 1,
		},
		Coordinates{
			x: 1,
			y: -0.5,
		},
		Coordinates{
			x: -1,
			y: -0.5,
		},
		Coordinates{
			x: -0.5,
			y: -1,
		},
	}
}

type Chroma struct {
	red   int64
	green int64
	blue  int64
}
type Coordinates struct {
	x float64
	y float64
}

//May belong in new file
func transformPitchScale(response AudioAnalysisResponse) []Coordinates {
	ret := make([]Coordinates, len(response.Segments))
	for i, segment := range response.Segments {
		pitchCoords := make([]Coordinates, 12)
		accumulatedPitch := &Coordinates{x: 0, y: 0}
		for i, pitch := range segment.Pitches {
			correctedPitch := pitch
			pitchCoords[i].x = correctedPitch * vectors[i].x
			pitchCoords[i].y = correctedPitch * vectors[i].y
			accumulatedPitch.x += pitchCoords[i].x
			accumulatedPitch.y += pitchCoords[i].y
		}
		ret[i] = *accumulatedPitch
	}
	return ret
}

//http://www.gradfree.com/kevin/some_theory_on_musical_keys.htm
func getChroma(features AudioFeatures) Chroma {
	//1 is major, 0 is minor

	//Contributing components to Red: Energy +, Loudness +, Valence -, Acousticness-, Danceability+
	//Contributing components to green: Energy 0.6, Liveness+, Major, Valence+, Acousticness+, Danceability +
	//Contributing components to Blue: Minor, - Valence, -Loudness, -Energy, +acousticness, -Danceability, -Energy

	calcRed := (1 / 5 * 255 * math.Abs(features.Acousticness-1)) + (1 / 5 * (features.Loudness * -0.6)) + (1 / 5 * 255 * math.Abs(features.Valence-1)) + (1 / 5 * 255 * features.Danceability) + (1 / 5 * 255 * features.Energy)
	calcGreen := (1 / 5 * 255 * features.Acousticness) + (1 / 5 * (features.Loudness * 0.6)) + (1 / 5 * 255 * features.Valence) + (1 / 5 * 255 * features.Danceability) + (1 / 5 * 255 * features.Energy)
	calcBlue := (1 / 6 * 255 * math.Abs(float64(features.Mode-1))) + (1 / 6 * 255 * features.Acousticness) + (1 / 6 * (features.Loudness * -0.6)) + (1 / 6 * 255 * math.Abs(features.Valence-1)) + (1 / 6 * 255 * features.Danceability) + (1 / 6 * 255 * math.Abs(features.Energy-1))

	chroma := &Chroma{
		red:   int64(calcRed),
		green: int64(calcGreen),
		blue:  int64(calcBlue),
	}
	return *chroma
}
func getChromaSet(baseColor Chroma) []Chroma {
	ChromaSet := make([]Chroma, 5)
	ChromaSet = append(ChromaSet, baseColor)
	//TODO FINISH LATER
	return ChromaSet
}
