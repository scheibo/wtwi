package weather

import (
	"fmt"
	"math"
	"strings"
	"time"
)

const msToKmh = 3600.0 / 1000.0

type Conditions struct {
	Icon                string    `json:"icon",omitempty"`
	Time                time.Time `json:"time",omitempty"`
	Temperature         float64   `json:"temperature,omitempty"`
	Humidity            float64   `json:"humidity,omitempty"`
	ApparentTemperature float64   `json:"apparentTemperature,omitempty"`
	PrecipProbability   float64   `json:"precipProbability,omitempty"`
	PrecipIntensity     float64   `json:"precipIntensity,omitempty"`
	PrecipType          string    `json:"precipType"`
	AirPressure         float64   `json:"airPressure,omitempty"`
	AirDensity          float64   `json:"airDensity,omitempty"`
	CloudCover          float64   `json:"cloudCover,omitempty"`
	UVIndex             int       `json:"uvIndex,omitempty"`
	WindSpeed           float64   `json:"windSpeed,omitempty"`
	WindGust            float64   `json:"windGust,omitempty"`
	WindBearing         float64   `json:"windBearing,omitempty"`
	SunriseTime         time.Time `json:"sunriseTime",omitempty"`
	SunsetTime          time.Time `json:"sunsetTime",omitempty"`
}

func (c *Conditions) String() string {
	precip := ""
	if c.PrecipProbability > 0 && c.PrecipIntensity > 0 {
		precip = fmt.Sprintf("%d%% %s (%.1f mm/h)\n",
			int(c.PrecipProbability*100),
			c.PrecipType,
			round(c.PrecipIntensity, 0.1))
	}

	return fmt.Sprintf(
		"%s\n"+
			"%.1f°C / %d%% (~%.1f°C)\n"+
			"%.1f km/h (%.1f km/h) %s\n"+
			"%.3f kg/m³ / %.2f mbar\n"+
			"%s"+
			"UV: %d / CC: %d%%\n"+ // TODO: "/ AQI: %d"
			"(%s / %s)",
		strings.ToUpper(c.Icon),
		round(c.Temperature, 0.1),
		int(c.Humidity*100),
		round(c.ApparentTemperature, 0.1),
		round(c.WindSpeed*msToKmh, 0.1),
		round(c.WindGust*msToKmh, 0.1),
		Direction(c.WindBearing),
		round(c.AirDensity, 0.001),
		round(c.AirPressure, 0.01),
		precip,
		c.UVIndex,
		int(c.CloudCover*100),
		c.SunriseTime.Format("3:04 PM"),
		c.SunsetTime.Format("3:04 PM"))
}

func Direction(b float64) string {
	var COMPASS = []string{
		"N", "NNE", "NE", "ENE", "E", "ESE", "SE", "SSE",
		"S", "SSW", "SW", "WSW", "W", "WNW", "NW", "NNW",
	}

	nb := normalizeBearing(b)
	index := int(math.Mod((nb+11.25)/22.5, 16))
	dir := COMPASS[index]

	return fmt.Sprintf("%s (%.1f°)", dir, round(nb, 0.5))
}

func round(x, unit float64) float64 {
	return math.Round(x/unit) * unit
}

func normalizeBearing(d float64) float64 {
	return d + math.Ceil(-d/360)*360
}
