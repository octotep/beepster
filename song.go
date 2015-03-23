package beepster

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

func (song *Song) CreateTrackFiller(trackId uint8, cleanup func(), tracks ...*[]Note) func() {
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
