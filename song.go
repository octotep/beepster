package beepster

import (
	"fmt"
	"github.com/octotep/go-mxl"
	"math"
)

type Song struct {
	NumTracks uint8
	Track     []chan Note
	Reps      int
}

func CreateSong(numTracks uint8, numRepetitions int) *Song {
	// Create the song and fill out all the information we know about it
	mysong := new(Song)
	mysong.NumTracks = numTracks
	mysong.Reps = numRepetitions
	mysong.Track = make([]chan Note, mysong.NumTracks)
	for i := range mysong.Track {
		mysong.Track[i] = make(chan Note, 14)
	}
	return mysong
}

func (song *Song) CreateFiller(trackId uint8, cleanup func(), tracks ...*[]Note) func() {
	return func() {
		defer cleanup()
		for i := 0; i < song.Reps; i++ {
			// Loop through all arrays given
			for _, val := range tracks {
				// Loop through all notes in one of the arrays
				for _, note := range *val {
					song.Track[trackId] <- note
				}
			}
		}
		close(song.Track[trackId])
	}
}

func CreateSongFromXML(mxlDoc mxl.MXLDoc, numRepetitions int) *Song {
	// xmlFile, err := os.Open("/home/octotep/tree.xml")
	// if err != nil {
	// 	fmt.Println("Error opening file:", err)
	// }
	// defer xmlFile.Close()

	// XMLdata, _ := ioutil.ReadAll(xmlFile)

	// var doc mxl.MXLDoc
	// xml.Unmarshal(XMLdata, &doc)

	mysong := new(Song)
	mysong.NumTracks = uint8(len(mxlDoc.Parts))
	mysong.Reps = numRepetitions
	mysong.Track = make([]chan Note, mysong.NumTracks)
	for i := range mysong.Track {
		mysong.Track[i] = make(chan Note, 14)
	}
	return mysong

	// for i, measure := range doc.Parts[1].Measures {
	// 	// fmt.Printf("\t%s\n", part)
	// 	fmt.Println("Measure:", i)
	// 	fmt.Println("Divisions:", measure.Atters.Divisions)
	// 	for j, note := range measure.Notes {
	// 		fmt.Println(j, note.Pitch, note.Duration)
	// 	}
	// }
}

func (song *Song) CreateFillerFromXml(trackId uint8, bpm uint, cleanup func(), track mxl.Part) func() {
	return func() {
		defer cleanup()
		for i := 0; i < song.Reps; i++ {
			// Loop through all arrays given
			currentDiv := 0
			for _, measure := range track.Measures {
				// Check for new division
				if measure.Atters.Divisions != 0 {
					currentDiv = measure.Atters.Divisions
				}

				// Loop through all notes in one of the arrays
				for _, note := range measure.Notes {
					var freq float32
					if note.Pitch.Step == "" {
						// It's a rest
						freq = 0.0
					} else {
						freq = PitchToFreq(note.Pitch)
					}
					lengthOfQuarter := 1.0 / float32(bpm)
					totalTime := float32(note.Duration) / float32(currentDiv) * lengthOfQuarter
					// Conver totalTime to ms
					delay := 5
					length := int(totalTime*1000) - delay
					// fmt.Println("totalTIme:", totalTime)
					song.Track[trackId] <- Note{freq, uint32(length), uint32(delay)}
				}
			}
		}
		close(song.Track[trackId])
	}
}

func PitchToFreq(pitch mxl.Pitch) float32 {
	var letter int
	if pitch.Step == "C" {
		letter = 0
	} else if pitch.Step == "D" {
		letter = 2
	} else if pitch.Step == "E" {
		letter = 4
	} else if pitch.Step == "F" {
		letter = 5
	} else if pitch.Step == "G" {
		letter = 7
	} else if pitch.Step == "A" {
		letter = 9
	} else if pitch.Step == "B" {
		letter = 11
	}
	absoluteHalfSteps := 12*(pitch.Octave-1) + letter + int(pitch.Accidental)
	a440 := 45
	relativeHalfSteps := absoluteHalfSteps - a440

	var freq float32
	// 2^(n/12) * 440
	freq = float32(math.Pow(2, float64(float64(relativeHalfSteps)/12.0))) * 440.0
	fmt.Println("Relative, freq, pitch", relativeHalfSteps, freq, pitch)
	return freq
}
