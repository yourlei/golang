package mp

import ("testing")

func TestPlay(t *testing.T) {
	mp3  := &MP3Player{0, 0}
	mp3.Play("/music/mydreap.mp3")
}