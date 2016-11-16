package amazon

// Region constants
type Region string

const (
	// RegionBrazil Brazil
	RegionBrazil Region = "BR"
	// RegionCanada Canada
	RegionCanada Region = "CA"
	// RegionChina  China
	RegionChina Region = "CN"
	// RegionGermany Germany
	RegionGermany Region = "DE"
	// RegionSpain  Spain
	RegionSpain Region = "ES"
	// RegionFrance France
	RegionFrance Region = "FR"
	// RegionIndia  India
	RegionIndia Region = "IN"
	// RegionItaly  Italy
	RegionItaly Region = "IT"
	// RegionJapan  Japan
	RegionJapan Region = "JP"
	// RegionMexico Mexico
	RegionMexico Region = "MX"
	// RegionUK     UK
	RegionUK Region = "UK"
	// RegionUS     US
	RegionUS Region = "US"
)

// Endpoint returns API endpoint for region
func (region Region) Endpoint() string {
	return map[Region]string{
		RegionBrazil:  "webservices.amazon.com.br",
		RegionCanada:  "webservices.amazon.ca",
		RegionChina:   "webservices.amazon.cn",
		RegionGermany: "webservices.amazon.de",
		RegionSpain:   "webservices.amazon.es",
		RegionFrance:  "webservices.amazon.fr",
		RegionIndia:   "webservices.amazon.in",
		RegionItaly:   "webservices.amazon.it",
		RegionJapan:   "webservices.amazon.co.jp",
		RegionMexico:  "webservices.amazon.com.mx",
		RegionUK:      "webservices.amazon.co.uk",
		RegionUS:      "webservices.amazon.com",
	}[region]
}

// HTTPSEndpoint returns HTTPS endpoint
func (region Region) HTTPSEndpoint() string {
	ep := region.Endpoint()
	if ep == "" {
		return ""
	}
	return "https://" + ep + "/onca/xml"
}

// HTTPEndpoint returns HTTP endpoint
func (region Region) HTTPEndpoint() string {
	ep := region.Endpoint()
	if ep == "" {
		return ""
	}
	return "http://" + ep + "/onca/xml"
}
