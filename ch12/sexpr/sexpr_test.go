package sexpr

import (
	"fmt"
	"testing"
)

type Movie struct {
	Title, Subtitle string
	Year            int
	//Color           bool
	Actor           map[string]string
	Oscars          []string
	Sequel          *string
}

var strangelove = Movie{
	Title:    "Dr. Strangelove",
	Subtitle: "How I Learned to Stop Worrying and Love the Bomb",
	Year:     1964,
	//Color:    false,
	Actor: map[string]string{
		"Dr. Strangelove":            "Peter Sellers",
		"Grp. Capt. Lionel Mandrake": "Peter Sellers",
		"Pres. Merkin Muffley":       "Peter Sellers",
		"Gen. Buck Turgidson":        "George C. Scott",
		"Brig. Gen. Jack D. Ripper":  "Sterling Hayden",
		`Maj. T.J. "King" Kong`:      "Slim Pickens",
	},
	Oscars: []string{
		"Best Actor (Nomin.)",
		"Best Adapted Screenplay (Nomin.)",
		"Best Director (Nomin.)",
		"Best Picture (Nomin.)",
	},
}

func TestMarshal(t *testing.T) {
	bs,err:=Marshal(strangelove)
	if err!=nil{
		t.Errorf("Marshal: %v\n",err)
	}
	t.Log(string(bs))
}

func TestUnmarshal(t *testing.T) {
	bs,err:=Marshal(strangelove)
	if err!=nil{
		t.Errorf("Marshal: %v\n",err)
		return
	}
	var movie Movie
	if err=Unmarshal(bs,movie);err!=nil{
		t.Errorf("Unmarshal %v\n",err)
		return
	}
	fmt.Println(movie.Oscars)
}

