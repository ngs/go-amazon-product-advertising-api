package amazon

import "testing"

func TestEndpoint(t *testing.T) {
	for _, test := range []Test{
		Test{"webservices.amazon.com.br", RegionBrazil.Endpoint()},
		Test{"webservices.amazon.ca", RegionCanada.Endpoint()},
		Test{"webservices.amazon.cn", RegionChina.Endpoint()},
		Test{"webservices.amazon.de", RegionGermany.Endpoint()},
		Test{"webservices.amazon.es", RegionSpain.Endpoint()},
		Test{"webservices.amazon.fr", RegionFrance.Endpoint()},
		Test{"webservices.amazon.in", RegionIndia.Endpoint()},
		Test{"webservices.amazon.it", RegionItaly.Endpoint()},
		Test{"webservices.amazon.co.jp", RegionJapan.Endpoint()},
		Test{"webservices.amazon.com.mx", RegionMexico.Endpoint()},
		Test{"webservices.amazon.co.uk", RegionUK.Endpoint()},
		Test{"webservices.amazon.com", RegionUS.Endpoint()},
	} {
		test.Compare(t)
	}
}
