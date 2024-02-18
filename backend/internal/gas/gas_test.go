package gas

import (
	"testing"
)

func TestGenerate(t *testing.T) {
	data := Generate()
	if data.CO2 < Data["co2"].Minimum || data.CO2 > Data["co2"].Maximum {
		t.Errorf("%s: outside range [%f, %f]", "co2", Data["co2"].Minimum, Data["co2"].Maximum)
	}
	if data.CO < Data["co"].Minimum || data.CO > Data["co"].Maximum {
		t.Errorf("%s: outside range [%f, %f]", "co", Data["co"].Minimum, Data["co"].Maximum)
	}
	if data.NO2 < Data["no2"].Minimum || data.NO2 > Data["no2"].Maximum {
		t.Errorf("%s: outside range [%f, %f]", "no2", Data["no2"].Minimum, Data["no2"].Maximum)
	}
	if data.MP10 < Data["mp10"].Minimum || data.MP10 > Data["mp10"].Maximum {
		t.Errorf("%s: outside range [%f, %f]", "mp10", Data["mp10"].Minimum, Data["mp10"].Maximum)
	}
	if data.MP25 < Data["mp25"].Minimum || data.MP25 > Data["mp25"].Maximum {
		t.Errorf("%s: outside range [%f, %f]", "mp25", Data["mp25"].Minimum, Data["mp25"].Maximum)
	}
}
