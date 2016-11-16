package amazon

import "testing"

func TestRegionEndpoint(t *testing.T) {
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

func TestRegionHTTPSEndpoint(t *testing.T) {
	for _, test := range []Test{
		Test{"https://webservices.amazon.com.br/onca/xml", RegionBrazil.HTTPSEndpoint()},
		Test{"https://webservices.amazon.ca/onca/xml", RegionCanada.HTTPSEndpoint()},
		Test{"https://webservices.amazon.cn/onca/xml", RegionChina.HTTPSEndpoint()},
		Test{"https://webservices.amazon.de/onca/xml", RegionGermany.HTTPSEndpoint()},
		Test{"https://webservices.amazon.es/onca/xml", RegionSpain.HTTPSEndpoint()},
		Test{"https://webservices.amazon.fr/onca/xml", RegionFrance.HTTPSEndpoint()},
		Test{"https://webservices.amazon.in/onca/xml", RegionIndia.HTTPSEndpoint()},
		Test{"https://webservices.amazon.it/onca/xml", RegionItaly.HTTPSEndpoint()},
		Test{"https://webservices.amazon.co.jp/onca/xml", RegionJapan.HTTPSEndpoint()},
		Test{"https://webservices.amazon.com.mx/onca/xml", RegionMexico.HTTPSEndpoint()},
		Test{"https://webservices.amazon.co.uk/onca/xml", RegionUK.HTTPSEndpoint()},
		Test{"https://webservices.amazon.com/onca/xml", RegionUS.HTTPSEndpoint()},
		Test{"", Region("foo").HTTPSEndpoint()},
	} {
		test.Compare(t)
	}
}

func TestRegionHTTPEndpoint(t *testing.T) {
	for _, test := range []Test{
		Test{"http://webservices.amazon.com.br/onca/xml", RegionBrazil.HTTPEndpoint()},
		Test{"http://webservices.amazon.ca/onca/xml", RegionCanada.HTTPEndpoint()},
		Test{"http://webservices.amazon.cn/onca/xml", RegionChina.HTTPEndpoint()},
		Test{"http://webservices.amazon.de/onca/xml", RegionGermany.HTTPEndpoint()},
		Test{"http://webservices.amazon.es/onca/xml", RegionSpain.HTTPEndpoint()},
		Test{"http://webservices.amazon.fr/onca/xml", RegionFrance.HTTPEndpoint()},
		Test{"http://webservices.amazon.in/onca/xml", RegionIndia.HTTPEndpoint()},
		Test{"http://webservices.amazon.it/onca/xml", RegionItaly.HTTPEndpoint()},
		Test{"http://webservices.amazon.co.jp/onca/xml", RegionJapan.HTTPEndpoint()},
		Test{"http://webservices.amazon.com.mx/onca/xml", RegionMexico.HTTPEndpoint()},
		Test{"http://webservices.amazon.co.uk/onca/xml", RegionUK.HTTPEndpoint()},
		Test{"http://webservices.amazon.com/onca/xml", RegionUS.HTTPEndpoint()},
		Test{"", Region("foo").HTTPEndpoint()},
	} {
		test.Compare(t)
	}
}
