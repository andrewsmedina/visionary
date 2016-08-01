package tracker

import (
	"image"
	"reflect"
	"testing"
)

func TestTrack(t *testing.T) {
	tck := Color{color: "red"}
	tracked, err := tck.Track()
	if err != nil {
		t.Error("Expected nil, got ", err)
	}
	if !reflect.DeepEqual(tracked, image.Rect(0, 0, 0, 0)) {
		t.Error("Expected a Rect(0, 0, 0 ,0), got ", tracked)
	}
}
