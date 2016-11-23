package directions

import (
	"html/template"
	"net/http"
	"os"
	"time"

	"google.golang.org/appengine"
	"google.golang.org/appengine/datastore"
	"google.golang.org/appengine/urlfetch"
	"googlemaps.github.io/maps"
)

type Config struct {
	Name  string
	Value string
}

type Directions struct {
	Origin string
	Client *maps.Client
	Apikey string
	r      *http.Request
	Resp   string
	Leg    *maps.Leg
	Dir    *maps.Route
	Steps  []*Step
}

type Step struct {
	Dist       string
	Duration   time.Duration
	Directions template.HTML
}

func (d *Directions) GetApikey() string {
	res := make([]Config, 10)
	ctx := appengine.NewContext(d.r)
	q := datastore.NewQuery("Config")
	_, _ = q.GetAll(ctx, &res)

	c := os.Getenv("APIKEY")
	for _, v := range res {
		if v.Name == "APIKEY" {
			c = v.Value
		}
	}
	return c
}

func NewDirections(r *http.Request) *Directions {
	var d = new(Directions)
	d.r = r
	d.Apikey = d.GetApikey()
	ctx := appengine.NewContext(r)
	uc := urlfetch.Client(ctx)
	c, err := maps.NewClient(maps.WithAPIKey(d.Apikey), maps.WithHTTPClient(uc))
	d.Client = c
	if err != nil {
		d.Resp = err.Error()
	}
	return d
}

func (d *Directions) Directions() {
	r := &maps.DirectionsRequest{
		Mode:        maps.TravelModeDriving,
		Origin:      "1200 Crittenden Lane, Mountain View",
		Destination: "90 Enterprise Way, Scotts Valley",
		DepartureTime:  time.Now()
package directions

import (
	"html/template"
	"net/http"
	"os"
	"time"

	"google.golang.org/appengine"
	"google.golang.org/appengine/datastore"
	"google.golang.org/appengine/urlfetch"
	"googlemaps.github.io/maps"
)

type Config struct {
	Name  string
	Value string
}

type Directions struct {
	Origin string
	Client *maps.Client
	Apikey string
	r      *http.Request
	Resp   string
	Leg    *maps.Leg
	Dir    *maps.Route
	Steps  []*Step
}

type Step struct {
	Dist       string
	Duration   time.Duration
	Directions template.HTML
}

func (d *Directions) GetApikey() string {
	res := make([]Config, 10)
	ctx := appengine.NewContext(d.r)
	q := datastore.NewQuery("Config")
	_, _ = q.GetAll(ctx, &res)

	c := os.Getenv("APIKEY")
	for _, v := range res {
		if v.Name == "APIKEY" {
			c = v.Value
		}
	}
	return c
}

func NewDirections(r *http.Request) *Directions {
	var d = new(Directions)
	d.r = r
	d.Apikey = d.GetApikey()
	ctx := appengine.NewContext(r)
	uc := urlfetch.Client(ctx)
	c, err := maps.NewClient(maps.WithAPIKey(d.Apikey), maps.WithHTTPClient(uc))
	d.Client = c
	if err != nil {
		d.Resp = err.Error()
	}
	return d
}

func (d *Directions) Directions() {
	r := &maps.DirectionsRequest{
		Mode:        maps.TravelModeDriving,
		Origin:      "1200 Crittenden Lane, Mountain View",
		Destination: "90 Enterprise Way, Scotts Valley",
		DepartureTime:  time.Now()
	}

		resp, _, _ := d.Client.Directions(appengine.NewContext(d.r), r)
		d.Dir = &resp[0]
		for _, v := range d.Dir.Legs[0].Steps {
			d.Steps = append(d.Steps, &Step{v.Distance.HumanReadable, v.Duration, template.HTML(v.HTMLInstructions)})
		}
		d.Resp = d.Steps
		d.Leg = d.Dir.Legs[0]
		d.Resp = resp
}
	}

	resp, _, _ := d.Client.Directions(appengine.NewContext(d.r), r)
	d.Dir = &resp[0]
	for _, v := range d.Dir.Legs[0].Steps {
		d.Steps = append(d.Steps, &Step{v.Distance.HumanReadable, v.Duration, template.HTML(v.HTMLInstructions)})
	}
	d.Resp = d.Steps
	d.Leg = d.Dir.Legs[0]
	d.Resp = resp
}
