package rad_lum

import (
	"testing"
)

func TestGenerate(t *testing.T) {
	data := Generate()
	if data.ET < Values["et"].Minimum || data.ET > Values["et"].Maximum {
		t.Errorf("%s: outside range [%f, %f]", "et", Values["et"].Minimum, Values["et"].Maximum)
	}
	if data.LI < Values["li"].Minimum || data.LI > Values["li"].Maximum {
		t.Errorf("%s: outside range [%f, %f]", "li", Values["li"].Minimum, Values["li"].Maximum)
	}
	if data.SR < Values["sr"].Minimum || data.SR > Values["sr"].Maximum {
		t.Errorf("%s: outside range [%f, %f]", "sr", Values["sr"].Minimum, Values["sr"].Maximum)
	}
}
