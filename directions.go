package directions
import (
	"os"
	"net/http"
        "google.golang.org/appengine"
        "google.golang.org/appengine/datastore"
        "google.golang.org/appengine/urlfetch"
	"googlemaps.github.io/maps"
	"github.com/kr/pretty"
//        "golang.org/x/net/context"
)

type Config struct {
	Name string
	Value string
  }

type Directions struct {
	Origin string
  Client *maps.Client
  Apikey string
  r *http.Request
  Resp string
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
  c, _ := maps.NewClient(maps.WithAPIKey(d.Apikey), maps.WithHTTPClient(uc))
  d.Client = c
  return d
}

func (d *Directions) Directions() {
  r := &maps.DirectionsRequest{
      Origin:      "1200 Crittenden Lane, Mountain View",
      Destination: "90 Enterprise Way, Scotts Valley",
  }
  resp, _, err := d.Client.Directions(appengine.NewContext(d.r), r)
  d.Resp = pretty.Sprint(err) + pretty.Sprint(resp) + pretty.Sprint(r)
}

