package mp3

import (
	"io"
)

func InspectFile(path string) (*MP3Info, error) {
	s, err := NewScanner(path)
	if err != nil {
		return nil, err
	}

	for {
		_, _, err := s.NextFrame()
		if err != nil {
			if err != io.EOF {
				return nil, err
			}
			break
		}
	}
	s.Info.FrameCount = s.FrameCount
	return s.Info, nil
}
