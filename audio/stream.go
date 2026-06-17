package audio

import (
	"assistent/config"

	"github.com/gordonklaus/portaudio"
)

func NewStream() (*portaudio.Stream, []int16, error) {
	err := portaudio.Initialize()
	if err != nil {
		return nil, nil, err
	}
	buffer := make([]int16, 4000)

	stream, err := portaudio.OpenDefaultStream(1, 0, config.SampleRate, 64, &buffer)
	if err != nil {
		return nil, nil, err
	}

	err = stream.Start()
	if err != nil {
		return nil, nil, err
	}

	return stream, buffer, nil
}

func Close(stream *portaudio.Stream) error {
	err := stream.Stop()
	if err != nil {
		return err
	}
	err = stream.Close()
	if err != nil {
		return err
	}
	err = portaudio.Terminate()
	if err != nil {
		return err
	}
	return nil
}
