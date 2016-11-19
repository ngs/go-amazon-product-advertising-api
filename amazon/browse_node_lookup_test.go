package amazon

import (
	"errors"
	"net/url"
	"os"
	"testing"
	"time"

	gock "gopkg.in/h2non/gock.v1"
)

const expectedBrowseNodeLookupSignedURL = "https://webservices.amazon.co.jp/onca/xml?AWSAccessKeyId=AK&AssociateTag=ngsio-22&BrowseNodeId=492352&Operation=BrowseNodeLookup&ResponseGroup=BrowseNodeInfo%2CNewReleases%2CMostGifted%2CTopSellers%2CMostWishedFor&Service=AWSECommerceService&Signature=DsQWqEYtxLtHgkHCiLmyTF7geOrNNqkvznh3KSplb7w%3D&Timestamp=2016-11-16T12%3A34%3A00Z&Version=2013-08-01"

func createBrowseNodeLookupRequest(client *Client) *BrowseNodeLookupRequest {
	return client.BrowseNodeLookup(BrowseNodeLookupParameters{
		ResponseGroups: []BrowseNodeLookupResponseGroup{
			BrowseNodeLookupResponseGroupBrowseNodeInfo,
			BrowseNodeLookupResponseGroupNewReleases,
			BrowseNodeLookupResponseGroupMostGifted,
			BrowseNodeLookupResponseGroupTopSellers,
			BrowseNodeLookupResponseGroupMostWishedFor,
		},
		BrowseNodeID: "492352",
	})
}

func TestBrowseNodeLookupSignedURL(t *testing.T) {
	setNow(time.Date(2016, time.November, 16, 21, 34, 0, 0, time.FixedZone("Asia/Tokyo", 9*60*60)))
	client, _ := New("AK", "SK", "ngsio-22", RegionJapan)
	op := createBrowseNodeLookupRequest(client)
	signedURL := client.SignedURL(op)
	parsed, _ := url.Parse(signedURL)
	for _, test := range []Test{
		Test{expectedBrowseNodeLookupSignedURL, signedURL},
		Test{"AK", parsed.Query().Get("AWSAccessKeyId")},
		Test{"ngsio-22", parsed.Query().Get("AssociateTag")},
		Test{"492352", parsed.Query().Get("BrowseNodeId")},
		Test{"BrowseNodeLookup", parsed.Query().Get("Operation")},
		Test{"AWSECommerceService", parsed.Query().Get("Service")},
		Test{"DsQWqEYtxLtHgkHCiLmyTF7geOrNNqkvznh3KSplb7w=", parsed.Query().Get("Signature")},
		Test{"2016-11-16T12:34:00Z", parsed.Query().Get("Timestamp")},
		Test{"2013-08-01", parsed.Query().Get("Version")},
	} {
		test.Compare(t)
	}
}

func TestBrowseNodeLookupDoErrorResponse(t *testing.T) {
	setNow(time.Date(2016, time.November, 16, 21, 34, 0, 0, time.FixedZone("Asia/Tokyo", 9*60*60)))
	client, _ := New("AK", "SK", "ngsio-22", RegionJapan)
	op := createBrowseNodeLookupRequest(client)
	fixtureIO, _ := os.Open("_fixtures/BrowseNodeLookupResponseErrorItem.xml")
	gock.New(expectedBrowseNodeLookupSignedURL).
		Reply(200).
		Body(fixtureIO)
	res, err := op.Do()
	if err == nil {
		t.Errorf("Expected not nil but got nil res: %v", res)
	} else {
		Test{"Error AWS.MissingParameters: リクエストには、必要なパラメータが含まれていません。必要なパラメータには、AssociateTagなどがあります。", err.Error()}.Compare(t)
	}
}

func TestBrowseNodeLookupDoError(t *testing.T) {
	client, _ := New("AK", "SK", "ngsio-22", RegionJapan)
	op := createBrowseNodeLookupRequest(client)
	gock.New(expectedBrowseNodeLookupSignedURL).
		ReplyError(errors.New("omg"))
	res, err := op.Do()
	if err == nil {
		t.Errorf("Expected not nil but got nil res: %v", res)
	} else {
		Test{"Get " + expectedBrowseNodeLookupSignedURL + ": omg", err.Error()}.Compare(t)
	}
}

func TestBrowseNodeLookupDo(t *testing.T) {
	setNow(time.Date(2016, time.November, 16, 21, 34, 0, 0, time.FixedZone("Asia/Tokyo", 9*60*60)))
	client, _ := New("AK", "SK", "ngsio-22", RegionJapan)
	op := createBrowseNodeLookupRequest(client)
	fixtureIO, _ := os.Open("_fixtures/BrowseNodeLookup.xml")
	gock.New(expectedBrowseNodeLookupSignedURL).
		Reply(200).
		Body(fixtureIO)
	res, err := op.Do()
	if err != nil {
		t.Errorf("Expected nil but got %v", err)
	}
	for _, test := range []Test{
		Test{1, len(res.BrowseNodes())},
		Test{"492352", res.BrowseNodes()[0].ID},
		Test{"プログラミング", res.BrowseNodes()[0].Name},
		Test{5, len(res.BrowseNodes()[0].Children.BrowseNode)},
		Test{1, len(res.BrowseNodes()[0].Ancestors.BrowseNode)},
		Test{4, len(res.BrowseNodes()[0].TopItemSet)},
		Test{10, len(res.BrowseNodes()[0].TopSellers.TopSeller)},
		Test{"525592", res.BrowseNodes()[0].Children.BrowseNode[0].ID},
		Test{"プログラミング入門書", res.BrowseNodes()[0].Children.BrowseNode[0].Name},
		Test{0, len(res.BrowseNodes()[0].Children.BrowseNode[0].Children.BrowseNode)},
		Test{0, len(res.BrowseNodes()[0].Children.BrowseNode[0].Ancestors.BrowseNode)},
		Test{0, len(res.BrowseNodes()[0].Children.BrowseNode[0].TopItemSet)},
		Test{0, len(res.BrowseNodes()[0].Children.BrowseNode[0].TopSellers.TopSeller)},
		Test{"3229700051", res.BrowseNodes()[0].Children.BrowseNode[1].ID},
		Test{"ゲームプログラミング", res.BrowseNodes()[0].Children.BrowseNode[1].Name},
		Test{0, len(res.BrowseNodes()[0].Children.BrowseNode[1].Children.BrowseNode)},
		Test{0, len(res.BrowseNodes()[0].Children.BrowseNode[1].Ancestors.BrowseNode)},
		Test{0, len(res.BrowseNodes()[0].Children.BrowseNode[1].TopItemSet)},
		Test{0, len(res.BrowseNodes()[0].Children.BrowseNode[1].TopSellers.TopSeller)},
		Test{"3229704051", res.BrowseNodes()[0].Children.BrowseNode[2].ID},
		Test{"ソフトウェア開発・言語", res.BrowseNodes()[0].Children.BrowseNode[2].Name},
		Test{0, len(res.BrowseNodes()[0].Children.BrowseNode[2].Children.BrowseNode)},
		Test{0, len(res.BrowseNodes()[0].Children.BrowseNode[2].Ancestors.BrowseNode)},
		Test{0, len(res.BrowseNodes()[0].Children.BrowseNode[2].TopItemSet)},
		Test{0, len(res.BrowseNodes()[0].Children.BrowseNode[2].TopSellers.TopSeller)},
		Test{"3229699051", res.BrowseNodes()[0].Children.BrowseNode[3].ID},
		Test{"モバイルプログラミング", res.BrowseNodes()[0].Children.BrowseNode[3].Name},
		Test{0, len(res.BrowseNodes()[0].Children.BrowseNode[3].Children.BrowseNode)},
		Test{0, len(res.BrowseNodes()[0].Children.BrowseNode[3].Ancestors.BrowseNode)},
		Test{0, len(res.BrowseNodes()[0].Children.BrowseNode[3].TopItemSet)},
		Test{0, len(res.BrowseNodes()[0].Children.BrowseNode[3].TopSellers.TopSeller)},
		Test{"3229701051", res.BrowseNodes()[0].Children.BrowseNode[4].ID},
		Test{"開発技法", res.BrowseNodes()[0].Children.BrowseNode[4].Name},
		Test{0, len(res.BrowseNodes()[0].Children.BrowseNode[4].Children.BrowseNode)},
		Test{0, len(res.BrowseNodes()[0].Children.BrowseNode[4].Ancestors.BrowseNode)},
		Test{0, len(res.BrowseNodes()[0].Children.BrowseNode[4].TopItemSet)},
		Test{0, len(res.BrowseNodes()[0].Children.BrowseNode[4].TopSellers.TopSeller)},
		Test{"466298", res.BrowseNodes()[0].Ancestors.BrowseNode[0].ID},
		Test{"コンピュータ・IT", res.BrowseNodes()[0].Ancestors.BrowseNode[0].Name},
		Test{0, len(res.BrowseNodes()[0].Ancestors.BrowseNode[0].Children.BrowseNode)},
		Test{1, len(res.BrowseNodes()[0].Ancestors.BrowseNode[0].Ancestors.BrowseNode)},
		Test{0, len(res.BrowseNodes()[0].Ancestors.BrowseNode[0].TopItemSet)},
		Test{0, len(res.BrowseNodes()[0].Ancestors.BrowseNode[0].TopSellers.TopSeller)},
		Test{"465610", res.BrowseNodes()[0].Ancestors.BrowseNode[0].Ancestors.BrowseNode[0].ID},
		Test{"ジャンル別", res.BrowseNodes()[0].Ancestors.BrowseNode[0].Ancestors.BrowseNode[0].Name},
		Test{0, len(res.BrowseNodes()[0].Ancestors.BrowseNode[0].Ancestors.BrowseNode[0].Children.BrowseNode)},
		Test{1, len(res.BrowseNodes()[0].Ancestors.BrowseNode[0].Ancestors.BrowseNode[0].Ancestors.BrowseNode)},
		Test{0, len(res.BrowseNodes()[0].Ancestors.BrowseNode[0].Ancestors.BrowseNode[0].TopItemSet)},
		Test{0, len(res.BrowseNodes()[0].Ancestors.BrowseNode[0].Ancestors.BrowseNode[0].TopSellers.TopSeller)},
		Test{"465392", res.BrowseNodes()[0].Ancestors.BrowseNode[0].Ancestors.BrowseNode[0].Ancestors.BrowseNode[0].ID},
		Test{"本", res.BrowseNodes()[0].Ancestors.BrowseNode[0].Ancestors.BrowseNode[0].Ancestors.BrowseNode[0].Name},
		Test{0, len(res.BrowseNodes()[0].Ancestors.BrowseNode[0].Ancestors.BrowseNode[0].Ancestors.BrowseNode[0].Children.BrowseNode)},
		Test{0, len(res.BrowseNodes()[0].Ancestors.BrowseNode[0].Ancestors.BrowseNode[0].Ancestors.BrowseNode[0].Ancestors.BrowseNode)},
		Test{0, len(res.BrowseNodes()[0].Ancestors.BrowseNode[0].Ancestors.BrowseNode[0].Ancestors.BrowseNode[0].TopItemSet)},
		Test{0, len(res.BrowseNodes()[0].Ancestors.BrowseNode[0].Ancestors.BrowseNode[0].Ancestors.BrowseNode[0].TopSellers.TopSeller)},

		Test{10, len(res.BrowseNodes()[0].TopItemSet[0].TopItem)},
		Test{"NewReleases", res.BrowseNodes()[0].TopItemSet[0].Type},
		Test{"4774185345", res.BrowseNodes()[0].TopItemSet[0].TopItem[0].ASIN},
		Test{"清水 亮", res.BrowseNodes()[0].TopItemSet[0].TopItem[0].Author},
		Test{"https://www.amazon.jp/%E3%81%AF%E3%81%98%E3%82%81%E3%81%A6%E3%81%AE%E6%B7%B1%E5%B1%A4%E5%AD%A6%E7%BF%92-%E3%83%87%E3%82%A3%E3%83%BC%E3%83%97%E3%83%A9%E3%83%BC%E3%83%8B%E3%83%B3%E3%82%B0-%E3%83%97%E3%83%AD%E3%82%B0%E3%83%A9%E3%83%9F%E3%83%B3%E3%82%B0-%E6%B8%85%E6%B0%B4-%E4%BA%AE/dp/4774185345%3FSubscriptionId%3DAKIAITPH62XKCOOT7AKA%26tag%3Dngsio-22%26linkCode%3Dxm2%26camp%3D2025%26creative%3D165953%26creativeASIN%3D4774185345", res.BrowseNodes()[0].TopItemSet[0].TopItem[0].DetailPageURL},
		Test{"Book", res.BrowseNodes()[0].TopItemSet[0].TopItem[0].ProductGroup},
		Test{"はじめての深層学習(ディープラーニング)プログラミング", res.BrowseNodes()[0].TopItemSet[0].TopItem[0].Title},
		Test{"B01MFEVRL5", res.BrowseNodes()[0].TopItemSet[0].TopItem[1].ASIN},
		Test{"大重 美幸", res.BrowseNodes()[0].TopItemSet[0].TopItem[1].Author},
		Test{"https://www.amazon.jp/%E8%A9%B3%E7%B4%B0%EF%BC%81Swift-iPhone%E3%82%A2%E3%83%97%E3%83%AA%E9%96%8B%E7%99%BA-%E5%85%A5%E9%96%80%E3%83%8E%E3%83%BC%E3%83%88-Swift-Xcode-ebook/dp/B01MFEVRL5%3FSubscriptionId%3DAKIAITPH62XKCOOT7AKA%26tag%3Dngsio-22%26linkCode%3Dxm2%26camp%3D2025%26creative%3D165953%26creativeASIN%3DB01MFEVRL5", res.BrowseNodes()[0].TopItemSet[0].TopItem[1].DetailPageURL},
		Test{"eBooks", res.BrowseNodes()[0].TopItemSet[0].TopItem[1].ProductGroup},
		Test{"詳細！Swift 3 iPhoneアプリ開発 入門ノート Swift 3+Xcode 8対応", res.BrowseNodes()[0].TopItemSet[0].TopItem[1].Title},
		Test{"4798148121", res.BrowseNodes()[0].TopItemSet[0].TopItem[2].ASIN},
		Test{"株式会社Re:Kayo-System", res.BrowseNodes()[0].TopItemSet[0].TopItem[2].Author},
		Test{"https://www.amazon.jp/%E3%81%BB%E3%82%93%E3%81%8D%E3%81%A7%E5%AD%A6%E3%81%B6Android%E3%82%A2%E3%83%97%E3%83%AA%E9%96%8B%E7%99%BA%E5%85%A5%E9%96%80-%E7%AC%AC2%E7%89%88-Android-Studio%E3%80%81Android-SDK/dp/4798148121%3FSubscriptionId%3DAKIAITPH62XKCOOT7AKA%26tag%3Dngsio-22%26linkCode%3Dxm2%26camp%3D2025%26creative%3D165953%26creativeASIN%3D4798148121", res.BrowseNodes()[0].TopItemSet[0].TopItem[2].DetailPageURL},
		Test{"Book", res.BrowseNodes()[0].TopItemSet[0].TopItem[2].ProductGroup},
		Test{"ほんきで学ぶAndroidアプリ開発入門 第2版 Android Studio、Android SDK 7対応", res.BrowseNodes()[0].TopItemSet[0].TopItem[2].Title},
		Test{"487311778X", res.BrowseNodes()[0].TopItemSet[0].TopItem[3].ASIN},
		Test{"Al Sweigart", res.BrowseNodes()[0].TopItemSet[0].TopItem[3].Author},
		Test{"https://www.amazon.jp/%E9%80%80%E5%B1%88%E3%81%AA%E3%81%93%E3%81%A8%E3%81%AFPython%E3%81%AB%E3%82%84%E3%82%89%E3%81%9B%E3%82%88%E3%81%86-%E2%80%95%E3%83%8E%E3%83%B3%E3%83%97%E3%83%AD%E3%82%B0%E3%83%A9%E3%83%9E%E3%83%BC%E3%81%AB%E3%82%82%E3%81%A7%E3%81%8D%E3%82%8B%E8%87%AA%E5%8B%95%E5%8C%96%E5%87%A6%E7%90%86%E3%83%97%E3%83%AD%E3%82%B0%E3%83%A9%E3%83%9F%E3%83%B3%E3%82%B0-Al-Sweigart/dp/487311778X%3FSubscriptionId%3DAKIAITPH62XKCOOT7AKA%26tag%3Dngsio-22%26linkCode%3Dxm2%26camp%3D2025%26creative%3D165953%26creativeASIN%3D487311778X", res.BrowseNodes()[0].TopItemSet[0].TopItem[3].DetailPageURL},
		Test{"Book", res.BrowseNodes()[0].TopItemSet[0].TopItem[3].ProductGroup},
		Test{"退屈なことはPythonにやらせよう ―ノンプログラマーにもできる自動化処理プログラミング", res.BrowseNodes()[0].TopItemSet[0].TopItem[3].Title},
		Test{"B01N04UBYI", res.BrowseNodes()[0].TopItemSet[0].TopItem[4].ASIN},
		Test{"伊藤 裕一", res.BrowseNodes()[0].TopItemSet[0].TopItem[4].Author},
		Test{"https://www.amazon.jp/%E9%80%9F%E7%BF%92-Python-3-%E4%B8%AD-%E3%82%AA%E3%83%96%E3%82%B8%E3%82%A7%E3%82%AF%E3%83%88%E6%8C%87%E5%90%91%E7%B7%A8-ebook/dp/B01N04UBYI%3FSubscriptionId%3DAKIAITPH62XKCOOT7AKA%26tag%3Dngsio-22%26linkCode%3Dxm2%26camp%3D2025%26creative%3D165953%26creativeASIN%3DB01N04UBYI", res.BrowseNodes()[0].TopItemSet[0].TopItem[4].DetailPageURL},
		Test{"eBooks", res.BrowseNodes()[0].TopItemSet[0].TopItem[4].ProductGroup},
		Test{"速習 Python 3 中: オブジェクト指向編", res.BrowseNodes()[0].TopItemSet[0].TopItem[4].Title},
		Test{"B01N035NLT", res.BrowseNodes()[0].TopItemSet[0].TopItem[5].ASIN},
		Test{"株式会社Re:Kayo-System", res.BrowseNodes()[0].TopItemSet[0].TopItem[5].Author},
		Test{"https://www.amazon.jp/%E3%81%BB%E3%82%93%E3%81%8D%E3%81%A7%E5%AD%A6%E3%81%B6Android%E3%82%A2%E3%83%97%E3%83%AA%E9%96%8B%E7%99%BA%E5%85%A5%E9%96%80-%E7%AC%AC2%E7%89%88-Android-Studio%E3%80%81Android-SDK-ebook/dp/B01N035NLT%3FSubscriptionId%3DAKIAITPH62XKCOOT7AKA%26tag%3Dngsio-22%26linkCode%3Dxm2%26camp%3D2025%26creative%3D165953%26creativeASIN%3DB01N035NLT", res.BrowseNodes()[0].TopItemSet[0].TopItem[5].DetailPageURL},
		Test{"eBooks", res.BrowseNodes()[0].TopItemSet[0].TopItem[5].ProductGroup},
		Test{"ほんきで学ぶAndroidアプリ開発入門 第2版 Android Studio、Android SDK 7対応", res.BrowseNodes()[0].TopItemSet[0].TopItem[5].Title},
		Test{"4797389818", res.BrowseNodes()[0].TopItemSet[0].TopItem[6].ASIN},
		Test{"高橋 京介", res.BrowseNodes()[0].TopItemSet[0].TopItem[6].Author},
		Test{"https://www.amazon.jp/%E7%B5%B6%E5%AF%BE%E3%81%AB%E6%8C%AB%E6%8A%98%E3%81%97%E3%81%AA%E3%81%84iPhone%E3%82%A2%E3%83%97%E3%83%AA%E9%96%8B%E7%99%BA%E3%80%8C%E8%B6%85%E3%80%8D%E5%85%A5%E9%96%80-%E5%A2%97%E8%A3%9C%E6%94%B9%E8%A8%82%E7%AC%AC5%E7%89%88%E3%80%90Swift-3-iOS-10%E3%80%91%E5%AE%8C%E5%85%A8%E5%AF%BE%E5%BF%9C/dp/4797389818%3FSubscriptionId%3DAKIAITPH62XKCOOT7AKA%26tag%3Dngsio-22%26linkCode%3Dxm2%26camp%3D2025%26creative%3D165953%26creativeASIN%3D4797389818", res.BrowseNodes()[0].TopItemSet[0].TopItem[6].DetailPageURL},
		Test{"Book", res.BrowseNodes()[0].TopItemSet[0].TopItem[6].ProductGroup},
		Test{"絶対に挫折しないiPhoneアプリ開発「超」入門 増補改訂第5版【Swift 3 & iOS 10】完全対応", res.BrowseNodes()[0].TopItemSet[0].TopItem[6].Title},
		Test{"4798145181", res.BrowseNodes()[0].TopItemSet[0].TopItem[7].ASIN},
		Test{"大久保 磨", res.BrowseNodes()[0].TopItemSet[0].TopItem[7].Author},
		Test{"https://www.amazon.jp/%E3%82%A2%E3%83%97%E3%83%AA-%E3%82%B2%E3%83%BC%E3%83%A0%E3%83%97%E3%83%A9%E3%83%B3%E3%83%8A%E3%83%BC%E5%BF%85%E8%AA%AD-%E3%83%AC%E3%83%99%E3%83%AB%E3%83%87%E3%82%B6%E3%82%A4%E3%83%B3%E5%BE%B9%E5%BA%95%E6%8C%87%E5%8D%97%E6%9B%B8-%E5%A4%A7%E4%B9%85%E4%BF%9D-%E7%A3%A8/dp/4798145181%3FSubscriptionId%3DAKIAITPH62XKCOOT7AKA%26tag%3Dngsio-22%26linkCode%3Dxm2%26camp%3D2025%26creative%3D165953%26creativeASIN%3D4798145181", res.BrowseNodes()[0].TopItemSet[0].TopItem[7].DetailPageURL},
		Test{"Book", res.BrowseNodes()[0].TopItemSet[0].TopItem[7].ProductGroup},
		Test{"アプリ&ゲームプランナー必読!  レベルデザイン徹底指南書", res.BrowseNodes()[0].TopItemSet[0].TopItem[7].Title},
		Test{"4873117844", res.BrowseNodes()[0].TopItemSet[0].TopItem[8].ASIN},
		Test{"Manoj Hans", res.BrowseNodes()[0].TopItemSet[0].TopItem[8].Author},
		Test{"https://www.amazon.jp/%E5%AE%9F%E8%B7%B5-Appium-Manoj-Hans/dp/4873117844%3FSubscriptionId%3DAKIAITPH62XKCOOT7AKA%26tag%3Dngsio-22%26linkCode%3Dxm2%26camp%3D2025%26creative%3D165953%26creativeASIN%3D4873117844", res.BrowseNodes()[0].TopItemSet[0].TopItem[8].DetailPageURL},
		Test{"Book", res.BrowseNodes()[0].TopItemSet[0].TopItem[8].ProductGroup},
		Test{"実践 Appium", res.BrowseNodes()[0].TopItemSet[0].TopItem[8].Title},
		Test{"B01MDOUFRN", res.BrowseNodes()[0].TopItemSet[0].TopItem[9].ASIN},
		Test{"小林 由憲", res.BrowseNodes()[0].TopItemSet[0].TopItem[9].Author},
		Test{"https://www.amazon.jp/%E3%81%93%E3%82%8C%E3%81%8B%E3%82%89%E3%81%A4%E3%81%8F%E3%82%8B-iPhone%E3%82%A2%E3%83%97%E3%83%AA%E9%96%8B%E7%99%BA%E5%85%A5%E9%96%80-%EF%BD%9ESwift%E3%81%A7%E3%81%AF%E3%81%98%E3%82%81%E3%82%8B%E3%83%97%E3%83%AD%E3%82%B0%E3%83%A9%E3%83%9F%E3%83%B3%E3%82%B0%E3%81%AE%E7%AC%AC%E4%B8%80%E6%AD%A9%EF%BD%9E-%E8%97%A4-%E6%B2%BB%E4%BB%81-ebook/dp/B01MDOUFRN%3FSubscriptionId%3DAKIAITPH62XKCOOT7AKA%26tag%3Dngsio-22%26linkCode%3Dxm2%26camp%3D2025%26creative%3D165953%26creativeASIN%3DB01MDOUFRN", res.BrowseNodes()[0].TopItemSet[0].TopItem[9].DetailPageURL},
		Test{"eBooks", res.BrowseNodes()[0].TopItemSet[0].TopItem[9].ProductGroup},
		Test{"これからつくる iPhoneアプリ開発入門 ～Swiftではじめるプログラミングの第一歩～", res.BrowseNodes()[0].TopItemSet[0].TopItem[9].Title},

		Test{10, len(res.BrowseNodes()[0].TopItemSet[1].TopItem)},
		Test{"MostGifted", res.BrowseNodes()[0].TopItemSet[1].Type},
		Test{"4800711487", res.BrowseNodes()[0].TopItemSet[1].TopItem[0].ASIN},
		Test{"大重 美幸", res.BrowseNodes()[0].TopItemSet[1].TopItem[0].Author},
		Test{"https://www.amazon.jp/Swift-iPhone%E3%82%A2%E3%83%97%E3%83%AA%E9%96%8B%E7%99%BA-Swift3-Oshige-introduction/dp/4800711487%3FSubscriptionId%3DAKIAITPH62XKCOOT7AKA%26tag%3Dngsio-22%26linkCode%3Dxm2%26camp%3D2025%26creative%3D165953%26creativeASIN%3D4800711487", res.BrowseNodes()[0].TopItemSet[1].TopItem[0].DetailPageURL},
		Test{"Book", res.BrowseNodes()[0].TopItemSet[1].TopItem[0].ProductGroup},
		Test{"詳細! Swift 3 iPhoneアプリ開発 入門ノート Swift3 + Xcode 8対応 (Oshige introduction note)", res.BrowseNodes()[0].TopItemSet[1].TopItem[0].Title},
		Test{"4062577690", res.BrowseNodes()[0].TopItemSet[1].TopItem[1].ASIN},
		Test{"立山 秀利", res.BrowseNodes()[0].TopItemSet[1].TopItem[1].Author},
		Test{"https://www.amazon.jp/%E5%85%A5%E9%96%80%E8%80%85%E3%81%AEExcel-VBA%E2%80%95%E5%88%9D%E3%82%81%E3%81%A6%E3%81%AE%E4%BA%BA%E3%81%AB%E3%83%99%E3%82%B9%E3%83%88%E3%81%AA%E5%AD%A6%E3%81%B3%E6%96%B9-%E3%83%96%E3%83%AB%E3%83%BC%E3%83%90%E3%83%83%E3%82%AF%E3%82%B9-%E7%AB%8B%E5%B1%B1-%E7%A7%80%E5%88%A9/dp/4062577690%3FSubscriptionId%3DAKIAITPH62XKCOOT7AKA%26tag%3Dngsio-22%26linkCode%3Dxm2%26camp%3D2025%26creative%3D165953%26creativeASIN%3D4062577690", res.BrowseNodes()[0].TopItemSet[1].TopItem[1].DetailPageURL},
		Test{"Book", res.BrowseNodes()[0].TopItemSet[1].TopItem[1].ProductGroup},
		Test{"入門者のExcel VBA―初めての人にベストな学び方 (ブルーバックス)", res.BrowseNodes()[0].TopItemSet[1].TopItem[1].Title},
		Test{"4822285154", res.BrowseNodes()[0].TopItemSet[1].TopItem[2].ASIN},
		Test{"阿部 和広", res.BrowseNodes()[0].TopItemSet[1].TopItem[2].Author},
		Test{"https://www.amazon.jp/%E5%B0%8F%E5%AD%A6%E7%94%9F%E3%81%8B%E3%82%89%E3%81%AF%E3%81%98%E3%82%81%E3%82%8B%E3%82%8F%E3%81%8F%E3%82%8F%E3%81%8F%E3%83%97%E3%83%AD%E3%82%B0%E3%83%A9%E3%83%9F%E3%83%B3%E3%82%B0-%E9%98%BF%E9%83%A8-%E5%92%8C%E5%BA%83/dp/4822285154%3FSubscriptionId%3DAKIAITPH62XKCOOT7AKA%26tag%3Dngsio-22%26linkCode%3Dxm2%26camp%3D2025%26creative%3D165953%26creativeASIN%3D4822285154", res.BrowseNodes()[0].TopItemSet[1].TopItem[2].DetailPageURL},
		Test{"Book", res.BrowseNodes()[0].TopItemSet[1].TopItem[2].ProductGroup},
		Test{"小学生からはじめるわくわくプログラミング", res.BrowseNodes()[0].TopItemSet[1].TopItem[2].Title},
		Test{"4899774117", res.BrowseNodes()[0].TopItemSet[1].TopItem[3].ASIN},
		Test{"大槻有一郎", res.BrowseNodes()[0].TopItemSet[1].TopItem[3].Author},
		Test{"https://www.amazon.jp/14%E6%AD%B3%E3%81%8B%E3%82%89%E3%81%AF%E3%81%98%E3%82%81%E3%82%8BC%E8%A8%80%E8%AA%9E%E3%82%8F%E3%81%8F%E3%82%8F%E3%81%8F%E3%82%B2%E3%83%BC%E3%83%A0%E3%83%97%E3%83%AD%E3%82%B0%E3%83%A9%E3%83%9F%E3%83%B3%E3%82%B0%E6%95%99%E5%AE%A4-Visual-Studio-2013%E7%B7%A8-%E5%A4%A7%E6%A7%BB%E6%9C%89%E4%B8%80%E9%83%8E/dp/4899774117%3FSubscriptionId%3DAKIAITPH62XKCOOT7AKA%26tag%3Dngsio-22%26linkCode%3Dxm2%26camp%3D2025%26creative%3D165953%26creativeASIN%3D4899774117", res.BrowseNodes()[0].TopItemSet[1].TopItem[3].DetailPageURL},
		Test{"Book", res.BrowseNodes()[0].TopItemSet[1].TopItem[3].ProductGroup},
		Test{"14歳からはじめるC言語わくわくゲームプログラミング教室 Visual Studio 2013編", res.BrowseNodes()[0].TopItemSet[1].TopItem[3].Title},
		Test{"484433638X", res.BrowseNodes()[0].TopItemSet[1].TopItem[4].ASIN},
		Test{"国本 大悟", res.BrowseNodes()[0].TopItemSet[1].TopItem[4].Author},
		Test{"https://www.amazon.jp/%E3%82%B9%E3%83%83%E3%82%AD%E3%83%AA%E3%82%8F%E3%81%8B%E3%82%8BJava%E5%85%A5%E9%96%80-%E7%AC%AC2%E7%89%88-%E3%82%B9%E3%83%83%E3%82%AD%E3%83%AA%E3%82%B7%E3%83%AA%E3%83%BC%E3%82%BA-%E4%B8%AD%E5%B1%B1-%E6%B8%85%E5%96%AC/dp/484433638X%3FSubscriptionId%3DAKIAITPH62XKCOOT7AKA%26tag%3Dngsio-22%26linkCode%3Dxm2%26camp%3D2025%26creative%3D165953%26creativeASIN%3D484433638X", res.BrowseNodes()[0].TopItemSet[1].TopItem[4].DetailPageURL},
		Test{"Book", res.BrowseNodes()[0].TopItemSet[1].TopItem[4].ProductGroup},
		Test{"スッキリわかるJava入門 第2版 (スッキリシリーズ)", res.BrowseNodes()[0].TopItemSet[1].TopItem[4].Title},
		Test{"479814245X", res.BrowseNodes()[0].TopItemSet[1].TopItem[5].ASIN},
		Test{"増井 敏克", res.BrowseNodes()[0].TopItemSet[1].TopItem[5].Author},
		Test{"https://www.amazon.jp/%E3%83%97%E3%83%AD%E3%82%B0%E3%83%A9%E3%83%9E%E8%84%B3%E3%82%92%E9%8D%9B%E3%81%88%E3%82%8B%E6%95%B0%E5%AD%A6%E3%83%91%E3%82%BA%E3%83%AB-%E3%82%B7%E3%83%B3%E3%83%97%E3%83%AB%E3%81%A7%E9%AB%98%E9%80%9F%E3%81%AA%E3%82%B3%E3%83%BC%E3%83%89%E3%81%8C%E6%9B%B8%E3%81%91%E3%82%8B%E3%82%88%E3%81%86%E3%81%AB%E3%81%AA%E3%82%8B70%E5%95%8F-%E5%A2%97%E4%BA%95-%E6%95%8F%E5%85%8B/dp/479814245X%3FSubscriptionId%3DAKIAITPH62XKCOOT7AKA%26tag%3Dngsio-22%26linkCode%3Dxm2%26camp%3D2025%26creative%3D165953%26creativeASIN%3D479814245X", res.BrowseNodes()[0].TopItemSet[1].TopItem[5].DetailPageURL},
		Test{"Book", res.BrowseNodes()[0].TopItemSet[1].TopItem[5].ProductGroup},
		Test{"プログラマ脳を鍛える数学パズル シンプルで高速なコードが書けるようになる70問", res.BrowseNodes()[0].TopItemSet[1].TopItem[5].Title},
		Test{"4798041793", res.BrowseNodes()[0].TopItemSet[1].TopItem[6].ASIN},
		Test{"山田 祥寛", res.BrowseNodes()[0].TopItemSet[1].TopItem[6].Author},
		Test{"https://www.amazon.jp/ASP-NET-MVC5%E5%AE%9F%E8%B7%B5%E3%83%97%E3%83%AD%E3%82%B0%E3%83%A9%E3%83%9F%E3%83%B3%E3%82%B0-%E5%B1%B1%E7%94%B0-%E7%A5%A5%E5%AF%9B/dp/4798041793%3FSubscriptionId%3DAKIAITPH62XKCOOT7AKA%26tag%3Dngsio-22%26linkCode%3Dxm2%26camp%3D2025%26creative%3D165953%26creativeASIN%3D4798041793", res.BrowseNodes()[0].TopItemSet[1].TopItem[6].DetailPageURL},
		Test{"Book", res.BrowseNodes()[0].TopItemSet[1].TopItem[6].ProductGroup},
		Test{"ASP.NET MVC5実践プログラミング", res.BrowseNodes()[0].TopItemSet[1].TopItem[6].Title},
		Test{"4797380055", res.BrowseNodes()[0].TopItemSet[1].TopItem[7].ASIN},
		Test{"徳岡 正肇", res.BrowseNodes()[0].TopItemSet[1].TopItem[7].Author},
		Test{"https://www.amazon.jp/%E3%82%B2%E3%83%BC%E3%83%A0%E3%81%AE%E4%BB%8A-%E3%82%B2%E3%83%BC%E3%83%A0%E6%A5%AD%E7%95%8C%E3%82%92%E8%A6%8B%E9%80%9A%E3%81%9918%E3%81%AE%E3%82%AD%E3%83%BC%E3%83%AF%E3%83%BC%E3%83%89-%E5%BE%B3%E5%B2%A1-%E6%AD%A3%E8%82%87/dp/4797380055%3FSubscriptionId%3DAKIAITPH62XKCOOT7AKA%26tag%3Dngsio-22%26linkCode%3Dxm2%26camp%3D2025%26creative%3D165953%26creativeASIN%3D4797380055", res.BrowseNodes()[0].TopItemSet[1].TopItem[7].DetailPageURL},
		Test{"Book", res.BrowseNodes()[0].TopItemSet[1].TopItem[7].ProductGroup},
		Test{"ゲームの今 ゲーム業界を見通す18のキーワード", res.BrowseNodes()[0].TopItemSet[1].TopItem[7].Title},
		Test{"4798128023", res.BrowseNodes()[0].TopItemSet[1].TopItem[8].ASIN},
		Test{"Joe Celko", res.BrowseNodes()[0].TopItemSet[1].TopItem[8].Author},
		Test{"https://www.amazon.jp/%E3%83%97%E3%83%AD%E3%82%B0%E3%83%A9%E3%83%9E%E3%81%AE%E3%81%9F%E3%82%81%E3%81%AESQL-%E7%AC%AC4%E7%89%88-%E3%82%B8%E3%83%A7%E3%83%BC%E3%83%BB%E3%82%BB%E3%83%AB%E3%82%B3/dp/4798128023%3FSubscriptionId%3DAKIAITPH62XKCOOT7AKA%26tag%3Dngsio-22%26linkCode%3Dxm2%26camp%3D2025%26creative%3D165953%26creativeASIN%3D4798128023", res.BrowseNodes()[0].TopItemSet[1].TopItem[8].DetailPageURL},
		Test{"Book", res.BrowseNodes()[0].TopItemSet[1].TopItem[8].ProductGroup},
		Test{"プログラマのためのSQL 第4版", res.BrowseNodes()[0].TopItemSet[1].TopItem[8].Title},
		Test{"4839956731", res.BrowseNodes()[0].TopItemSet[1].TopItem[9].ASIN},
		Test{"杉山 将", res.BrowseNodes()[0].TopItemSet[1].TopItem[9].Author},
		Test{"https://www.amazon.jp/%E5%BC%B7%E3%81%8F%E3%81%AA%E3%82%8B%E3%83%AD%E3%83%9C%E3%83%86%E3%82%A3%E3%83%83%E3%82%AF%E3%83%BB%E3%82%B2%E3%83%BC%E3%83%A0%E3%83%97%E3%83%AC%E3%82%A4%E3%83%A4%E3%83%BC%E3%81%AE%E4%BD%9C%E3%82%8A%E6%96%B9-%E3%83%97%E3%83%AC%E3%83%9F%E3%82%A2%E3%83%A0%E3%83%96%E3%83%83%E3%82%AF%E3%82%B9%E7%89%88-%7E%E5%AE%9F%E8%B7%B5%E3%81%A7%E5%AD%A6%E3%81%B6%E5%BC%B7%E5%8C%96%E5%AD%A6%E7%BF%92%7E-%E5%85%AB%E8%B0%B7-%E5%A4%A7%E5%B2%B3/dp/4839956731%3FSubscriptionId%3DAKIAITPH62XKCOOT7AKA%26tag%3Dngsio-22%26linkCode%3Dxm2%26camp%3D2025%26creative%3D165953%26creativeASIN%3D4839956731", res.BrowseNodes()[0].TopItemSet[1].TopItem[9].DetailPageURL},
		Test{"Book", res.BrowseNodes()[0].TopItemSet[1].TopItem[9].ProductGroup},
		Test{"強くなるロボティック・ゲームプレイヤーの作り方 プレミアムブックス版 ~実践で学ぶ強化学習~", res.BrowseNodes()[0].TopItemSet[1].TopItem[9].Title},

		Test{10, len(res.BrowseNodes()[0].TopItemSet[2].TopItem)},
		Test{"TopSellers", res.BrowseNodes()[0].TopItemSet[2].Type},
		Test{"B012VRQX9G", res.BrowseNodes()[0].TopItemSet[2].TopItem[0].ASIN},
		Test{"徳岡 正肇", res.BrowseNodes()[0].TopItemSet[2].TopItem[0].Author},
		Test{"https://www.amazon.jp/%E3%82%B2%E3%83%BC%E3%83%A0%E3%81%AE%E4%BB%8A%E3%80%80%E3%82%B2%E3%83%BC%E3%83%A0%E6%A5%AD%E7%95%8C%E3%82%92%E8%A6%8B%E9%80%9A%E3%81%9918%E3%81%AE%E3%82%AD%E3%83%BC%E3%83%AF%E3%83%BC%E3%83%89-%E5%BE%B3%E5%B2%A1-%E6%AD%A3%E8%82%87-ebook/dp/B012VRQX9G%3FSubscriptionId%3DAKIAITPH62XKCOOT7AKA%26tag%3Dngsio-22%26linkCode%3Dxm2%26camp%3D2025%26creative%3D165953%26creativeASIN%3DB012VRQX9G", res.BrowseNodes()[0].TopItemSet[2].TopItem[0].DetailPageURL},
		Test{"eBooks", res.BrowseNodes()[0].TopItemSet[2].TopItem[0].ProductGroup},
		Test{"ゲームの今　ゲーム業界を見通す18のキーワード", res.BrowseNodes()[0].TopItemSet[2].TopItem[0].Title},
		Test{"4873117380", res.BrowseNodes()[0].TopItemSet[2].TopItem[1].ASIN},
		Test{"Bill Lubanovic", res.BrowseNodes()[0].TopItemSet[2].TopItem[1].Author},
		Test{"https://www.amazon.jp/%E5%85%A5%E9%96%80-Python-3-Bill-Lubanovic/dp/4873117380%3FSubscriptionId%3DAKIAITPH62XKCOOT7AKA%26tag%3Dngsio-22%26linkCode%3Dxm2%26camp%3D2025%26creative%3D165953%26creativeASIN%3D4873117380", res.BrowseNodes()[0].TopItemSet[2].TopItem[1].DetailPageURL},
		Test{"Book", res.BrowseNodes()[0].TopItemSet[2].TopItem[1].ProductGroup},
		Test{"入門 Python 3", res.BrowseNodes()[0].TopItemSet[2].TopItem[1].Title},
		Test{"4800711487", res.BrowseNodes()[0].TopItemSet[2].TopItem[2].ASIN},
		Test{"大重 美幸", res.BrowseNodes()[0].TopItemSet[2].TopItem[2].Author},
		Test{"https://www.amazon.jp/Swift-iPhone%E3%82%A2%E3%83%97%E3%83%AA%E9%96%8B%E7%99%BA-Swift3-Oshige-introduction/dp/4800711487%3FSubscriptionId%3DAKIAITPH62XKCOOT7AKA%26tag%3Dngsio-22%26linkCode%3Dxm2%26camp%3D2025%26creative%3D165953%26creativeASIN%3D4800711487", res.BrowseNodes()[0].TopItemSet[2].TopItem[2].DetailPageURL},
		Test{"Book", res.BrowseNodes()[0].TopItemSet[2].TopItem[2].ProductGroup},
		Test{"詳細! Swift 3 iPhoneアプリ開発 入門ノート Swift3 + Xcode 8対応 (Oshige introduction note)", res.BrowseNodes()[0].TopItemSet[2].TopItem[2].Title},
		Test{"4062577690", res.BrowseNodes()[0].TopItemSet[2].TopItem[3].ASIN},
		Test{"立山 秀利", res.BrowseNodes()[0].TopItemSet[2].TopItem[3].Author},
		Test{"https://www.amazon.jp/%E5%85%A5%E9%96%80%E8%80%85%E3%81%AEExcel-VBA%E2%80%95%E5%88%9D%E3%82%81%E3%81%A6%E3%81%AE%E4%BA%BA%E3%81%AB%E3%83%99%E3%82%B9%E3%83%88%E3%81%AA%E5%AD%A6%E3%81%B3%E6%96%B9-%E3%83%96%E3%83%AB%E3%83%BC%E3%83%90%E3%83%83%E3%82%AF%E3%82%B9-%E7%AB%8B%E5%B1%B1-%E7%A7%80%E5%88%A9/dp/4062577690%3FSubscriptionId%3DAKIAITPH62XKCOOT7AKA%26tag%3Dngsio-22%26linkCode%3Dxm2%26camp%3D2025%26creative%3D165953%26creativeASIN%3D4062577690", res.BrowseNodes()[0].TopItemSet[2].TopItem[3].DetailPageURL},
		Test{"Book", res.BrowseNodes()[0].TopItemSet[2].TopItem[3].ProductGroup},
		Test{"入門者のExcel VBA―初めての人にベストな学び方 (ブルーバックス)", res.BrowseNodes()[0].TopItemSet[2].TopItem[3].Title},
		Test{"484433638X", res.BrowseNodes()[0].TopItemSet[2].TopItem[4].ASIN},
		Test{"国本 大悟", res.BrowseNodes()[0].TopItemSet[2].TopItem[4].Author},
		Test{"https://www.amazon.jp/%E3%82%B9%E3%83%83%E3%82%AD%E3%83%AA%E3%82%8F%E3%81%8B%E3%82%8BJava%E5%85%A5%E9%96%80-%E7%AC%AC2%E7%89%88-%E3%82%B9%E3%83%83%E3%82%AD%E3%83%AA%E3%82%B7%E3%83%AA%E3%83%BC%E3%82%BA-%E4%B8%AD%E5%B1%B1-%E6%B8%85%E5%96%AC/dp/484433638X%3FSubscriptionId%3DAKIAITPH62XKCOOT7AKA%26tag%3Dngsio-22%26linkCode%3Dxm2%26camp%3D2025%26creative%3D165953%26creativeASIN%3D484433638X", res.BrowseNodes()[0].TopItemSet[2].TopItem[4].DetailPageURL},
		Test{"Book", res.BrowseNodes()[0].TopItemSet[2].TopItem[4].ProductGroup},
		Test{"スッキリわかるJava入門 第2版 (スッキリシリーズ)", res.BrowseNodes()[0].TopItemSet[2].TopItem[4].Title},
		Test{"4873115655", res.BrowseNodes()[0].TopItemSet[2].TopItem[5].ASIN},
		Test{"Trevor Foucher", res.BrowseNodes()[0].TopItemSet[2].TopItem[5].Author},
		Test{"https://www.amazon.jp/%E3%83%AA%E3%83%BC%E3%83%80%E3%83%96%E3%83%AB%E3%82%B3%E3%83%BC%E3%83%89-%E2%80%95%E3%82%88%E3%82%8A%E8%89%AF%E3%81%84%E3%82%B3%E3%83%BC%E3%83%89%E3%82%92%E6%9B%B8%E3%81%8F%E3%81%9F%E3%82%81%E3%81%AE%E3%82%B7%E3%83%B3%E3%83%97%E3%83%AB%E3%81%A7%E5%AE%9F%E8%B7%B5%E7%9A%84%E3%81%AA%E3%83%86%E3%82%AF%E3%83%8B%E3%83%83%E3%82%AF-Theory-practice-Boswell/dp/4873115655%3FSubscriptionId%3DAKIAITPH62XKCOOT7AKA%26tag%3Dngsio-22%26linkCode%3Dxm2%26camp%3D2025%26creative%3D165953%26creativeASIN%3D4873115655", res.BrowseNodes()[0].TopItemSet[2].TopItem[5].DetailPageURL},
		Test{"Book", res.BrowseNodes()[0].TopItemSet[2].TopItem[5].ProductGroup},
		Test{"リーダブルコード ―より良いコードを書くためのシンプルで実践的なテクニック (Theory in practice)", res.BrowseNodes()[0].TopItemSet[2].TopItem[5].Title},
		Test{"477418411X", res.BrowseNodes()[0].TopItemSet[2].TopItem[6].ASIN},
		Test{"山田 祥寛", res.BrowseNodes()[0].TopItemSet[2].TopItem[6].Author},
		Test{"https://www.amazon.jp/%E6%94%B9%E8%A8%82%E6%96%B0%E7%89%88JavaScript%E6%9C%AC%E6%A0%BC%E5%85%A5%E9%96%80-%7E%E3%83%A2%E3%83%80%E3%83%B3%E3%82%B9%E3%82%BF%E3%82%A4%E3%83%AB%E3%81%AB%E3%82%88%E3%82%8B%E5%9F%BA%E7%A4%8E%E3%81%8B%E3%82%89%E7%8F%BE%E5%A0%B4%E3%81%A7%E3%81%AE%E5%BF%9C%E7%94%A8%E3%81%BE%E3%81%A7-%E5%B1%B1%E7%94%B0-%E7%A5%A5%E5%AF%9B/dp/477418411X%3FSubscriptionId%3DAKIAITPH62XKCOOT7AKA%26tag%3Dngsio-22%26linkCode%3Dxm2%26camp%3D2025%26creative%3D165953%26creativeASIN%3D477418411X", res.BrowseNodes()[0].TopItemSet[2].TopItem[6].DetailPageURL},
		Test{"Book", res.BrowseNodes()[0].TopItemSet[2].TopItem[6].ProductGroup},
		Test{"改訂新版JavaScript本格入門 ~モダンスタイルによる基礎から現場での応用まで", res.BrowseNodes()[0].TopItemSet[2].TopItem[6].Title},
		Test{"4822285154", res.BrowseNodes()[0].TopItemSet[2].TopItem[7].ASIN},
		Test{"阿部 和広", res.BrowseNodes()[0].TopItemSet[2].TopItem[7].Author},
		Test{"https://www.amazon.jp/%E5%B0%8F%E5%AD%A6%E7%94%9F%E3%81%8B%E3%82%89%E3%81%AF%E3%81%98%E3%82%81%E3%82%8B%E3%82%8F%E3%81%8F%E3%82%8F%E3%81%8F%E3%83%97%E3%83%AD%E3%82%B0%E3%83%A9%E3%83%9F%E3%83%B3%E3%82%B0-%E9%98%BF%E9%83%A8-%E5%92%8C%E5%BA%83/dp/4822285154%3FSubscriptionId%3DAKIAITPH62XKCOOT7AKA%26tag%3Dngsio-22%26linkCode%3Dxm2%26camp%3D2025%26creative%3D165953%26creativeASIN%3D4822285154", res.BrowseNodes()[0].TopItemSet[2].TopItem[7].DetailPageURL},
		Test{"Book", res.BrowseNodes()[0].TopItemSet[2].TopItem[7].ProductGroup},
		Test{"小学生からはじめるわくわくプログラミング", res.BrowseNodes()[0].TopItemSet[2].TopItem[7].Title},
		Test{"4797386797", res.BrowseNodes()[0].TopItemSet[2].TopItem[8].ASIN},
		Test{"北村 愛実", res.BrowseNodes()[0].TopItemSet[2].TopItem[8].Author},
		Test{"https://www.amazon.jp/Unity5%E3%81%AE%E6%95%99%E7%A7%91%E6%9B%B8-2D-3D%E3%82%B9%E3%83%9E%E3%83%BC%E3%83%88%E3%83%95%E3%82%A9%E3%83%B3%E3%82%B2%E3%83%BC%E3%83%A0%E5%85%A5%E9%96%80%E8%AC%9B%E5%BA%A7-Entertainment-IDEA/dp/4797386797%3FSubscriptionId%3DAKIAITPH62XKCOOT7AKA%26tag%3Dngsio-22%26linkCode%3Dxm2%26camp%3D2025%26creative%3D165953%26creativeASIN%3D4797386797", res.BrowseNodes()[0].TopItemSet[2].TopItem[8].DetailPageURL},
		Test{"Book", res.BrowseNodes()[0].TopItemSet[2].TopItem[8].ProductGroup},
		Test{"Unity5の教科書 2D&3Dスマートフォンゲーム入門講座 (Entertainment&IDEA)", res.BrowseNodes()[0].TopItemSet[2].TopItem[8].Title},
		Test{"4774180874", res.BrowseNodes()[0].TopItemSet[2].TopItem[9].ASIN},
		Test{"吉田拳", res.BrowseNodes()[0].TopItemSet[2].TopItem[9].Author},
		Test{"https://www.amazon.jp/%E3%81%9F%E3%81%A3%E3%81%9F1%E7%A7%92%E3%81%A7%E4%BB%95%E4%BA%8B%E3%81%8C%E7%89%87%E3%81%A5%E3%81%8F-Excel%E8%87%AA%E5%8B%95%E5%8C%96%E3%81%AE%E6%95%99%E7%A7%91%E6%9B%B8-%E5%90%89%E7%94%B0%E6%8B%B3/dp/4774180874%3FSubscriptionId%3DAKIAITPH62XKCOOT7AKA%26tag%3Dngsio-22%26linkCode%3Dxm2%26camp%3D2025%26creative%3D165953%26creativeASIN%3D4774180874", res.BrowseNodes()[0].TopItemSet[2].TopItem[9].DetailPageURL},
		Test{"Book", res.BrowseNodes()[0].TopItemSet[2].TopItem[9].ProductGroup},
		Test{"たった1秒で仕事が片づく Excel自動化の教科書", res.BrowseNodes()[0].TopItemSet[2].TopItem[9].Title},

		Test{10, len(res.BrowseNodes()[0].TopItemSet[3].TopItem)},
		Test{"MostWishedFor", res.BrowseNodes()[0].TopItemSet[3].Type},
		Test{"427421933X", res.BrowseNodes()[0].TopItemSet[3].TopItem[0].ASIN},
		Test{"David Thomas", res.BrowseNodes()[0].TopItemSet[3].TopItem[0].Author},
		Test{"https://www.amazon.jp/%E6%96%B0%E8%A3%85%E7%89%88-%E9%81%94%E4%BA%BA%E3%83%97%E3%83%AD%E3%82%B0%E3%83%A9%E3%83%9E%E3%83%BC-%E8%81%B7%E4%BA%BA%E3%81%8B%E3%82%89%E5%90%8D%E5%8C%A0%E3%81%B8%E3%81%AE%E9%81%93-Andrew-Hunt/dp/427421933X%3FSubscriptionId%3DAKIAITPH62XKCOOT7AKA%26tag%3Dngsio-22%26linkCode%3Dxm2%26camp%3D2025%26creative%3D165953%26creativeASIN%3D427421933X", res.BrowseNodes()[0].TopItemSet[3].TopItem[0].DetailPageURL},
		Test{"Book", res.BrowseNodes()[0].TopItemSet[3].TopItem[0].ProductGroup},
		Test{"新装版 達人プログラマー 職人から名匠への道", res.BrowseNodes()[0].TopItemSet[3].TopItem[0].Title},
		Test{"487311778X", res.BrowseNodes()[0].TopItemSet[3].TopItem[1].ASIN},
		Test{"Al Sweigart", res.BrowseNodes()[0].TopItemSet[3].TopItem[1].Author},
		Test{"https://www.amazon.jp/%E9%80%80%E5%B1%88%E3%81%AA%E3%81%93%E3%81%A8%E3%81%AFPython%E3%81%AB%E3%82%84%E3%82%89%E3%81%9B%E3%82%88%E3%81%86-%E2%80%95%E3%83%8E%E3%83%B3%E3%83%97%E3%83%AD%E3%82%B0%E3%83%A9%E3%83%9E%E3%83%BC%E3%81%AB%E3%82%82%E3%81%A7%E3%81%8D%E3%82%8B%E8%87%AA%E5%8B%95%E5%8C%96%E5%87%A6%E7%90%86%E3%83%97%E3%83%AD%E3%82%B0%E3%83%A9%E3%83%9F%E3%83%B3%E3%82%B0-Al-Sweigart/dp/487311778X%3FSubscriptionId%3DAKIAITPH62XKCOOT7AKA%26tag%3Dngsio-22%26linkCode%3Dxm2%26camp%3D2025%26creative%3D165953%26creativeASIN%3D487311778X", res.BrowseNodes()[0].TopItemSet[3].TopItem[1].DetailPageURL},
		Test{"Book", res.BrowseNodes()[0].TopItemSet[3].TopItem[1].ProductGroup},
		Test{"退屈なことはPythonにやらせよう ―ノンプログラマーにもできる自動化処理プログラミング", res.BrowseNodes()[0].TopItemSet[3].TopItem[1].Title},
		Test{"4873117380", res.BrowseNodes()[0].TopItemSet[3].TopItem[2].ASIN},
		Test{"Bill Lubanovic", res.BrowseNodes()[0].TopItemSet[3].TopItem[2].Author},
		Test{"https://www.amazon.jp/%E5%85%A5%E9%96%80-Python-3-Bill-Lubanovic/dp/4873117380%3FSubscriptionId%3DAKIAITPH62XKCOOT7AKA%26tag%3Dngsio-22%26linkCode%3Dxm2%26camp%3D2025%26creative%3D165953%26creativeASIN%3D4873117380", res.BrowseNodes()[0].TopItemSet[3].TopItem[2].DetailPageURL},
		Test{"Book", res.BrowseNodes()[0].TopItemSet[3].TopItem[2].ProductGroup},
		Test{"入門 Python 3", res.BrowseNodes()[0].TopItemSet[3].TopItem[2].Title},
		Test{"4873115655", res.BrowseNodes()[0].TopItemSet[3].TopItem[3].ASIN},
		Test{"Trevor Foucher", res.BrowseNodes()[0].TopItemSet[3].TopItem[3].Author},
		Test{"https://www.amazon.jp/%E3%83%AA%E3%83%BC%E3%83%80%E3%83%96%E3%83%AB%E3%82%B3%E3%83%BC%E3%83%89-%E2%80%95%E3%82%88%E3%82%8A%E8%89%AF%E3%81%84%E3%82%B3%E3%83%BC%E3%83%89%E3%82%92%E6%9B%B8%E3%81%8F%E3%81%9F%E3%82%81%E3%81%AE%E3%82%B7%E3%83%B3%E3%83%97%E3%83%AB%E3%81%A7%E5%AE%9F%E8%B7%B5%E7%9A%84%E3%81%AA%E3%83%86%E3%82%AF%E3%83%8B%E3%83%83%E3%82%AF-Theory-practice-Boswell/dp/4873115655%3FSubscriptionId%3DAKIAITPH62XKCOOT7AKA%26tag%3Dngsio-22%26linkCode%3Dxm2%26camp%3D2025%26creative%3D165953%26creativeASIN%3D4873115655", res.BrowseNodes()[0].TopItemSet[3].TopItem[3].DetailPageURL},
		Test{"Book", res.BrowseNodes()[0].TopItemSet[3].TopItem[3].ProductGroup},
		Test{"リーダブルコード ―より良いコードを書くためのシンプルで実践的なテクニック (Theory in practice)", res.BrowseNodes()[0].TopItemSet[3].TopItem[3].Title},
		Test{"4798046140", res.BrowseNodes()[0].TopItemSet[3].TopItem[4].ASIN},
		Test{"上田 勲", res.BrowseNodes()[0].TopItemSet[3].TopItem[4].Author},
		Test{"https://www.amazon.jp/%E3%83%97%E3%83%AA%E3%83%B3%E3%82%B7%E3%83%97%E3%83%AB-%E3%82%AA%E3%83%96-%E3%83%97%E3%83%AD%E3%82%B0%E3%83%A9%E3%83%9F%E3%83%B3%E3%82%B03%E5%B9%B4%E7%9B%AE%E3%81%BE%E3%81%A7%E3%81%AB%E8%BA%AB%E3%81%AB%E3%81%A4%E3%81%91%E3%81%9F%E3%81%84%E4%B8%80%E7%94%9F%E5%BD%B9%E7%AB%8B%E3%81%A4101%E3%81%AE%E5%8E%9F%E7%90%86%E5%8E%9F%E5%89%87-%E4%B8%8A%E7%94%B0-%E5%8B%B2/dp/4798046140%3FSubscriptionId%3DAKIAITPH62XKCOOT7AKA%26tag%3Dngsio-22%26linkCode%3Dxm2%26camp%3D2025%26creative%3D165953%26creativeASIN%3D4798046140", res.BrowseNodes()[0].TopItemSet[3].TopItem[4].DetailPageURL},
		Test{"Book", res.BrowseNodes()[0].TopItemSet[3].TopItem[4].ProductGroup},
		Test{"プリンシプル オブ プログラミング3年目までに身につけたい一生役立つ101の原理原則", res.BrowseNodes()[0].TopItemSet[3].TopItem[4].Title},
		Test{"4774180874", res.BrowseNodes()[0].TopItemSet[3].TopItem[5].ASIN},
		Test{"吉田拳", res.BrowseNodes()[0].TopItemSet[3].TopItem[5].Author},
		Test{"https://www.amazon.jp/%E3%81%9F%E3%81%A3%E3%81%9F1%E7%A7%92%E3%81%A7%E4%BB%95%E4%BA%8B%E3%81%8C%E7%89%87%E3%81%A5%E3%81%8F-Excel%E8%87%AA%E5%8B%95%E5%8C%96%E3%81%AE%E6%95%99%E7%A7%91%E6%9B%B8-%E5%90%89%E7%94%B0%E6%8B%B3/dp/4774180874%3FSubscriptionId%3DAKIAITPH62XKCOOT7AKA%26tag%3Dngsio-22%26linkCode%3Dxm2%26camp%3D2025%26creative%3D165953%26creativeASIN%3D4774180874", res.BrowseNodes()[0].TopItemSet[3].TopItem[5].DetailPageURL},
		Test{"Book", res.BrowseNodes()[0].TopItemSet[3].TopItem[5].ProductGroup},
		Test{"たった1秒で仕事が片づく Excel自動化の教科書", res.BrowseNodes()[0].TopItemSet[3].TopItem[5].Title},
		Test{"4873117771", res.BrowseNodes()[0].TopItemSet[3].TopItem[6].ASIN},
		Test{"Philip N. Klein", res.BrowseNodes()[0].TopItemSet[3].TopItem[6].Author},
		Test{"https://www.amazon.jp/%E8%A1%8C%E5%88%97%E3%83%97%E3%83%AD%E3%82%B0%E3%83%A9%E3%83%9E%E3%83%BC-%E2%80%95Python%E3%83%97%E3%83%AD%E3%82%B0%E3%83%A9%E3%83%A0%E3%81%A7%E5%AD%A6%E3%81%B6%E7%B7%9A%E5%BD%A2%E4%BB%A3%E6%95%B0-Philip-N-Klein/dp/4873117771%3FSubscriptionId%3DAKIAITPH62XKCOOT7AKA%26tag%3Dngsio-22%26linkCode%3Dxm2%26camp%3D2025%26creative%3D165953%26creativeASIN%3D4873117771", res.BrowseNodes()[0].TopItemSet[3].TopItem[6].DetailPageURL},
		Test{"Book", res.BrowseNodes()[0].TopItemSet[3].TopItem[6].ProductGroup},
		Test{"行列プログラマー ―Pythonプログラムで学ぶ線形代数", res.BrowseNodes()[0].TopItemSet[3].TopItem[6].Title},
		Test{"4873115892", res.BrowseNodes()[0].TopItemSet[3].TopItem[7].ASIN},
		Test{"Bill Karwin", res.BrowseNodes()[0].TopItemSet[3].TopItem[7].Author},
		Test{"https://www.amazon.jp/SQL%E3%82%A2%E3%83%B3%E3%83%81%E3%83%91%E3%82%BF%E3%83%BC%E3%83%B3-Bill-Karwin/dp/4873115892%3FSubscriptionId%3DAKIAITPH62XKCOOT7AKA%26tag%3Dngsio-22%26linkCode%3Dxm2%26camp%3D2025%26creative%3D165953%26creativeASIN%3D4873115892", res.BrowseNodes()[0].TopItemSet[3].TopItem[7].DetailPageURL},
		Test{"Book", res.BrowseNodes()[0].TopItemSet[3].TopItem[7].ProductGroup},
		Test{"SQLアンチパターン", res.BrowseNodes()[0].TopItemSet[3].TopItem[7].Title},
		Test{"4774183881", res.BrowseNodes()[0].TopItemSet[3].TopItem[8].ASIN},
		Test{"中久喜 健司", res.BrowseNodes()[0].TopItemSet[3].TopItem[8].Author},
		Test{"https://www.amazon.jp/%E7%A7%91%E5%AD%A6%E6%8A%80%E8%A1%93%E8%A8%88%E7%AE%97%E3%81%AE%E3%81%9F%E3%82%81%E3%81%AEPython%E5%85%A5%E9%96%80-%E2%80%95%E2%80%95%E9%96%8B%E7%99%BA%E5%9F%BA%E7%A4%8E%E3%80%81%E5%BF%85%E9%A0%88%E3%83%A9%E3%82%A4%E3%83%96%E3%83%A9%E3%83%AA%E3%80%81%E9%AB%98%E9%80%9F%E5%8C%96-%E4%B8%AD%E4%B9%85%E5%96%9C-%E5%81%A5%E5%8F%B8/dp/4774183881%3FSubscriptionId%3DAKIAITPH62XKCOOT7AKA%26tag%3Dngsio-22%26linkCode%3Dxm2%26camp%3D2025%26creative%3D165953%26creativeASIN%3D4774183881", res.BrowseNodes()[0].TopItemSet[3].TopItem[8].DetailPageURL},
		Test{"Book", res.BrowseNodes()[0].TopItemSet[3].TopItem[8].ProductGroup},
		Test{"科学技術計算のためのPython入門 ――開発基礎、必須ライブラリ、高速化", res.BrowseNodes()[0].TopItemSet[3].TopItem[8].Title},
		Test{"477415654X", res.BrowseNodes()[0].TopItemSet[3].TopItem[9].ASIN},
		Test{"西尾 泰和", res.BrowseNodes()[0].TopItemSet[3].TopItem[9].Author},
		Test{"https://www.amazon.jp/%E3%82%B3%E3%83%BC%E3%83%87%E3%82%A3%E3%83%B3%E3%82%B0%E3%82%92%E6%94%AF%E3%81%88%E3%82%8B%E6%8A%80%E8%A1%93-%7E%E6%88%90%E3%82%8A%E7%AB%8B%E3%81%A1%E3%81%8B%E3%82%89%E5%AD%A6%E3%81%B6%E3%83%97%E3%83%AD%E3%82%B0%E3%83%A9%E3%83%9F%E3%83%B3%E3%82%B0%E4%BD%9C%E6%B3%95-WEB-PRESS-plus/dp/477415654X%3FSubscriptionId%3DAKIAITPH62XKCOOT7AKA%26tag%3Dngsio-22%26linkCode%3Dxm2%26camp%3D2025%26creative%3D165953%26creativeASIN%3D477415654X", res.BrowseNodes()[0].TopItemSet[3].TopItem[9].DetailPageURL},
		Test{"Book", res.BrowseNodes()[0].TopItemSet[3].TopItem[9].ProductGroup},
		Test{"コーディングを支える技術 ~成り立ちから学ぶプログラミング作法 (WEB+DB PRESS plus)", res.BrowseNodes()[0].TopItemSet[3].TopItem[9].Title},

		Test{"B012VRQX9G", res.BrowseNodes()[0].TopSellers.TopSeller[0].ASIN},
		Test{"ゲームの今　ゲーム業界を見通す18のキーワード", res.BrowseNodes()[0].TopSellers.TopSeller[0].Title},
		Test{"4873117380", res.BrowseNodes()[0].TopSellers.TopSeller[1].ASIN},
		Test{"入門 Python 3", res.BrowseNodes()[0].TopSellers.TopSeller[1].Title},
		Test{"4800711487", res.BrowseNodes()[0].TopSellers.TopSeller[2].ASIN},
		Test{"詳細! Swift 3 iPhoneアプリ開発 入門ノート Swift3 + Xcode 8対応 (Oshige introduction note)", res.BrowseNodes()[0].TopSellers.TopSeller[2].Title},
		Test{"4062577690", res.BrowseNodes()[0].TopSellers.TopSeller[3].ASIN},
		Test{"入門者のExcel VBA―初めての人にベストな学び方 (ブルーバックス)", res.BrowseNodes()[0].TopSellers.TopSeller[3].Title},
		Test{"484433638X", res.BrowseNodes()[0].TopSellers.TopSeller[4].ASIN},
		Test{"スッキリわかるJava入門 第2版 (スッキリシリーズ)", res.BrowseNodes()[0].TopSellers.TopSeller[4].Title},
		Test{"4873115655", res.BrowseNodes()[0].TopSellers.TopSeller[5].ASIN},
		Test{"リーダブルコード ―より良いコードを書くためのシンプルで実践的なテクニック (Theory in practice)", res.BrowseNodes()[0].TopSellers.TopSeller[5].Title},
		Test{"477418411X", res.BrowseNodes()[0].TopSellers.TopSeller[6].ASIN},
		Test{"改訂新版JavaScript本格入門 ~モダンスタイルによる基礎から現場での応用まで", res.BrowseNodes()[0].TopSellers.TopSeller[6].Title},
		Test{"4822285154", res.BrowseNodes()[0].TopSellers.TopSeller[7].ASIN},
		Test{"小学生からはじめるわくわくプログラミング", res.BrowseNodes()[0].TopSellers.TopSeller[7].Title},
		Test{"4797386797", res.BrowseNodes()[0].TopSellers.TopSeller[8].ASIN},
		Test{"Unity5の教科書 2D&3Dスマートフォンゲーム入門講座 (Entertainment&IDEA)", res.BrowseNodes()[0].TopSellers.TopSeller[8].Title},
		Test{"4774180874", res.BrowseNodes()[0].TopSellers.TopSeller[9].ASIN},
		Test{"たった1秒で仕事が片づく Excel自動化の教科書", res.BrowseNodes()[0].TopSellers.TopSeller[9].Title},
	} {
		test.Compare(t)
	}
}
