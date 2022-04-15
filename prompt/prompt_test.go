package prompt

import (
	"testing"
)

func TestPrompt(t *testing.T) {
	fakeInput := FakeInput{
		x: "test",
		y: "test2",
		z: "test3",
	}
	a, b, c := Prompt(&fakeInput)

	if a != fakeInput.x {
		t.Errorf("Wanted: %v, Got: %v", a, fakeInput.x)
	}

	if b != fakeInput.y {
		t.Errorf("Wanted: %v, Got: %v", b, fakeInput.y)
	}

	if c != fakeInput.z {
		t.Errorf("Wanted: %v, Got: %v", c, fakeInput.z)
	}
}
