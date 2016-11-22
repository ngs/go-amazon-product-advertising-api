package amazon

// Condition typed constant for Condition parameter
type Condition string

const (
	// ConditionNone unspecifed condition
	ConditionNone Condition = ""
	// ConditionNew constant "New"
	ConditionNew Condition = "New"
	// ConditionUsed constant "Used"
	ConditionUsed Condition = "Used"
	// ConditionCollectible constant "Collectible"
	ConditionCollectible Condition = "Collectible"
	// ConditionRefurbished constant "Refurbished"
	ConditionRefurbished Condition = "Refurbished"
	// ConditionAll constant "All"
	ConditionAll Condition = "All"
)

// IDType typed constant for IDType parameter
type IDType string

const (
	// IDTypeSKU constant "SKU"
	IDTypeSKU IDType = "SKU"
	// IDTypeUPC constant "UPC"
	IDTypeUPC IDType = "UPC"
	// IDTypeEAN constant "EAN"
	IDTypeEAN IDType = "EAN"
	// IDTypeISBN constant "ISBN"
	IDTypeISBN IDType = "ISBN"
	// IDTypeASIN constnt "ASI"
	IDTypeASIN IDType = "ASIN"
)

// RelationshipType typed constant for RelationshipType parameter
type RelationshipType string

const (
	// RelationshipTypeAuthorityTitle Links a non-buyable ASIN TitleAuthority parent with its buyable children. A book might have a single TitleAuthority ASIN that relates to a list of child ASINs for different formats of the same book (hardback, paperback, audio book, Kindle). MP3 albums have the same AuthorityTitle parent as its physical CD counterpart.
	RelationshipTypeAuthorityTitle RelationshipType = "AuthorityTitle"
	// RelationshipTypeDigitalMusicArranger Non-buyable child of both MP3 albums and tracks.
	RelationshipTypeDigitalMusicArranger RelationshipType = "DigitalMusicArranger"
	// RelationshipTypeDigitalMusicComposer Non-buyable child of both MP3 albums and tracks.
	RelationshipTypeDigitalMusicComposer RelationshipType = "DigitalMusicComposer"
	// RelationshipTypeDigitalMusicConductor Non-buyable child of both MP3 albums and tracks.
	RelationshipTypeDigitalMusicConductor RelationshipType = "DigitalMusicConductor"
	// RelationshipTypeDigitalMusicEnsemble Non-buyable child of both MP3 albums and tracks.
	RelationshipTypeDigitalMusicEnsemble RelationshipType = "DigitalMusicEnsemble"
	// RelationshipTypeDigitalMusicLyricist Non-buyable child of both MP3 albums and tracks.
	RelationshipTypeDigitalMusicLyricist RelationshipType = "DigitalMusicLyricist"
	// RelationshipTypeDigitalMusicPerformer Non-buyable child of both MP3 albums and tracks.
	RelationshipTypeDigitalMusicPerformer RelationshipType = "DigitalMusicPerformer"
	// RelationshipTypeDigitalMusicPrimaryArtist Non-buyable child of both MP3 albums and tracks. This is the relationship that shows all MP3 downloads for a single artist on Amazon.com.
	RelationshipTypeDigitalMusicPrimaryArtist RelationshipType = "DigitalMusicPrimaryArtist"
	// RelationshipTypeDigitalMusicProducer Non-buyable child of both MP3 albums and tracks.
	RelationshipTypeDigitalMusicProducer RelationshipType = "DigitalMusicProducer"
	// RelationshipTypeDigitalMusicRemixer Non-buyable child of both MP3 albums and tracks.
	RelationshipTypeDigitalMusicRemixer RelationshipType = "DigitalMusicRemixer"
	// RelationshipTypeDigitalMusicSongWriter Non-buyable child of both MP3 albums and tracks.
	RelationshipTypeDigitalMusicSongWriter RelationshipType = "DigitalMusicSongWriter"
	// RelationshipTypeEpisode Relates an Unbox Season (parent) to Episodes (children) from that season. This value can be used interchangeably with Tracks.
	RelationshipTypeEpisode RelationshipType = "Episode"
	// RelationshipTypeNewerVersion Returns the latest version of an item.
	RelationshipTypeNewerVersion RelationshipType = "NewerVersion"
	// RelationshipTypeSeason Relates an Unbox Series (parent) to its Seasons (children).
	RelationshipTypeSeason RelationshipType = "Season"
	// RelationshipTypeTracks Relates an MP3 Album (parent) to its Tracks (children). This value can be used interchangeably with Episode.
	RelationshipTypeTracks RelationshipType = "Tracks"
)

// SearchIndex typed constant for SearchIndex parameter
type SearchIndex string

const (
	// SearchIndexAll constant for All search index parameter
	SearchIndexAll SearchIndex = "All"
	// SearchIndexApparel constant for Apparel search index parameter
	SearchIndexApparel SearchIndex = "Apparel"
	// SearchIndexAppliances constant for Appliances search index parameter
	SearchIndexAppliances SearchIndex = "Appliances"
	// SearchIndexArtsAndCrafts constant for ArtsAndCrafts search index parameter
	SearchIndexArtsAndCrafts SearchIndex = "ArtsAndCrafts"
	// SearchIndexAutomotive constant for Automotive search index parameter
	SearchIndexAutomotive SearchIndex = "Automotive"
	// SearchIndexBaby constant for Baby search index parameter
	SearchIndexBaby SearchIndex = "Baby"
	// SearchIndexBeauty constant for Beauty search index parameter
	SearchIndexBeauty SearchIndex = "Beauty"
	// SearchIndexBlended constant for Blended search index parameter
	SearchIndexBlended SearchIndex = "Blended"
	// SearchIndexBooks constant for Books search index parameter
	SearchIndexBooks SearchIndex = "Books"
	// SearchIndexClassical constant for Classical search index parameter
	SearchIndexClassical SearchIndex = "Classical"
	// SearchIndexCollectibles constant for Collectibles search index parameter
	SearchIndexCollectibles SearchIndex = "Collectibles"
	// SearchIndexCreditCards constant for CreditCards search index parameter
	SearchIndexCreditCards SearchIndex = "CreditCards"
	// SearchIndexDVD constant for DVD search index parameter
	SearchIndexDVD SearchIndex = "DVD"
	// SearchIndexElectronics constant for Electronics search index parameter
	SearchIndexElectronics SearchIndex = "Electronics"
	// SearchIndexFashion constant for Fashion search index parameter
	SearchIndexFashion SearchIndex = "Fashion"
	// SearchIndexFashionBaby constant for FashionBaby search index parameter
	SearchIndexFashionBaby SearchIndex = "FashionBaby"
	// SearchIndexFashionBoys constant for FashionBoys search index parameter
	SearchIndexFashionBoys SearchIndex = "FashionBoys"
	// SearchIndexFashionGirls constant for FashionGirls search index parameter
	SearchIndexFashionGirls SearchIndex = "FashionGirls"
	// SearchIndexFashionMen constant for FashionMen search index parameter
	SearchIndexFashionMen SearchIndex = "FashionMen"
	// SearchIndexFashionWomen constant for FashionWomen search index parameter
	SearchIndexFashionWomen SearchIndex = "FashionWomen"
	// SearchIndexForeignBooks constant for ForeignBooks search index parameter
	SearchIndexForeignBooks SearchIndex = "ForeignBooks"
	// SearchIndexGarden constant for Garden search index parameter
	SearchIndexGarden SearchIndex = "Garden"
	// SearchIndexGiftCards constant for GiftCards search index parameter
	SearchIndexGiftCards SearchIndex = "GiftCards"
	// SearchIndexGrocery constant for Grocery search index parameter
	SearchIndexGrocery SearchIndex = "Grocery"
	// SearchIndexHealthPersonalCare constant for HealthPersonalCare search index parameter
	SearchIndexHealthPersonalCare SearchIndex = "HealthPersonalCare"
	// SearchIndexHobbies constant for Hobbies search index parameter
	SearchIndexHobbies SearchIndex = "Hobbies"
	// SearchIndexHome constant for Home search index parameter
	SearchIndexHome SearchIndex = "Home"
	// SearchIndexHomeGarden constant for HomeGarden search index parameter
	SearchIndexHomeGarden SearchIndex = "HomeGarden"
	// SearchIndexHomeImprovement constant for HomeImprovement search index parameter
	SearchIndexHomeImprovement SearchIndex = "HomeImprovement"
	// SearchIndexIndustrial constant for Industrial search index parameter
	SearchIndexIndustrial SearchIndex = "Industrial"
	// SearchIndexJewelry constant for Jewelry search index parameter
	SearchIndexJewelry SearchIndex = "Jewelry"
	// SearchIndexKindleStore constant for KindleStore search index parameter
	SearchIndexKindleStore SearchIndex = "KindleStore"
	// SearchIndexKitchen constant for Kitchen search index parameter
	SearchIndexKitchen SearchIndex = "Kitchen"
	// SearchIndexLawnAndGarden constant for LawnAndGarden search index parameter
	SearchIndexLawnAndGarden SearchIndex = "LawnAndGarden"
	// SearchIndexLighting constant for Lighting search index parameter
	SearchIndexLighting SearchIndex = "Lighting"
	// SearchIndexLuggage constant for Luggage search index parameter
	SearchIndexLuggage SearchIndex = "Luggage"
	// SearchIndexMagazines constant for Magazines search index parameter
	SearchIndexMagazines SearchIndex = "Magazines"
	// SearchIndexMobileApps constant for MobileApps search index parameter
	SearchIndexMobileApps SearchIndex = "MobileApps"
	// SearchIndexMovies constant for Movies search index parameter
	SearchIndexMovies SearchIndex = "Movies"
	// SearchIndexMP3Downloads constant for MP3Downloads search index parameter
	SearchIndexMP3Downloads SearchIndex = "MP3Downloads"
	// SearchIndexMusic constant for Music search index parameter
	SearchIndexMusic SearchIndex = "Music"
	// SearchIndexMusicalInstruments constant for MusicalInstruments search index parameter
	SearchIndexMusicalInstruments SearchIndex = "MusicalInstruments"
	// SearchIndexOfficeProducts constant for OfficeProducts search index parameter
	SearchIndexOfficeProducts SearchIndex = "OfficeProducts"
	// SearchIndexPantry constant for Pantry search index parameter
	SearchIndexPantry SearchIndex = "Pantry"
	// SearchIndexPCHardware constant for PCHardware search index parameter
	SearchIndexPCHardware SearchIndex = "PCHardware"
	// SearchIndexPetSupplies constant for PetSupplies search index parameter
	SearchIndexPetSupplies SearchIndex = "PetSupplies"
	// SearchIndexPhoto constant for Photo search index parameter
	SearchIndexPhoto SearchIndex = "Photo"
	// SearchIndexShoes constant for Shoes search index parameter
	SearchIndexShoes SearchIndex = "Shoes"
	// SearchIndexSoftware constant for Software search index parameter
	SearchIndexSoftware SearchIndex = "Software"
	// SearchIndexSportingGoods constant for SportingGoods search index parameter
	SearchIndexSportingGoods SearchIndex = "SportingGoods"
	// SearchIndexTools constant for Tools search index parameter
	SearchIndexTools SearchIndex = "Tools"
	// SearchIndexToys constant for Toys search index parameter
	SearchIndexToys SearchIndex = "Toys"
	// SearchIndexUnboxVideo constant for UnboxVideo search index parameter
	SearchIndexUnboxVideo SearchIndex = "UnboxVideo"
	// SearchIndexVHS constant for VHS search index parameter
	SearchIndexVHS SearchIndex = "VHS"
	// SearchIndexVideo constant for Video search index parameter
	SearchIndexVideo SearchIndex = "Video"
	// SearchIndexVideoDownload constant for VideoDownload search index parameter
	SearchIndexVideoDownload SearchIndex = "VideoDownload"
	// SearchIndexVideoGames constant for VideoGames search index parameter
	SearchIndexVideoGames SearchIndex = "VideoGames"
	// SearchIndexWatches constant for Watches search index parameter
	SearchIndexWatches SearchIndex = "Watches"
	// SearchIndexWine constant for Wine search index parameter
	SearchIndexWine SearchIndex = "Wine"
	// SearchIndexWireless constant for Wireless search index parameter
	SearchIndexWireless SearchIndex = "Wireless"
)
