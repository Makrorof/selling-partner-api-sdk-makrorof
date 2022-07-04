// Package productPricing provides primitives to interact the openapi HTTP API.
//
// Code generated by go-sdk-codegen DO NOT EDIT.
package productPricing

import (
	"time"
)

// ASINIdentifier defines model for ASINIdentifier.
type ASINIdentifier struct {

	// The Amazon Standard Identification Number (ASIN) of the item.
	ASIN string `json:"ASIN"`

	// A marketplace identifier.
	MarketplaceId string `json:"MarketplaceId"`
}

// AttributeSetList defines model for AttributeSetList.
type AttributeSetList []map[string]interface{}

// BuyBoxEligibleOffers defines model for BuyBoxEligibleOffers.
type BuyBoxEligibleOffers []OfferCountType

// BuyBoxPriceType defines model for BuyBoxPriceType.
type BuyBoxPriceType struct {
	LandedPrice  MoneyType `json:"LandedPrice"`
	ListingPrice MoneyType `json:"ListingPrice"`
	Points       *Points   `json:"Points,omitempty"`
	Shipping     MoneyType `json:"Shipping"`

	// Indicates the condition of the item. For example: New, Used, Collectible, Refurbished, or Club.
	Condition string `json:"condition"`
}

// BuyBoxPrices defines model for BuyBoxPrices.
type BuyBoxPrices []BuyBoxPriceType

// CompetitivePriceList defines model for CompetitivePriceList.
type CompetitivePriceList []CompetitivePriceType

// CompetitivePriceType defines model for CompetitivePriceType.
type CompetitivePriceType struct {

	// The pricing model for each price that is returned.
	//
	// Possible values:
	//
	// * 1 - New Buy Box Price.
	// * 2 - Used Buy Box Price.
	CompetitivePriceId string    `json:"CompetitivePriceId"`
	Price              PriceType `json:"Price"`

	//  Indicates whether or not the pricing information is for an offer listing that belongs to the requester. The requester is the seller associated with the SellerId that was submitted with the request. Possible values are: true and false.
	BelongsToRequester *bool `json:"belongsToRequester,omitempty"`

	// Indicates the condition of the item whose pricing information is returned. Possible values are: New, Used, Collectible, Refurbished, or Club.
	Condition *string `json:"condition,omitempty"`

	// Indicates the subcondition of the item whose pricing information is returned. Possible values are: New, Mint, Very Good, Good, Acceptable, Poor, Club, OEM, Warranty, Refurbished Warranty, Refurbished, Open Box, or Other.
	Subcondition *string `json:"subcondition,omitempty"`
}

// CompetitivePricingType defines model for CompetitivePricingType.
type CompetitivePricingType struct {

	// A list of competitive pricing information.
	CompetitivePrices CompetitivePriceList `json:"CompetitivePrices"`

	// The number of active offer listings for the item that was submitted. The listing count is returned by condition, one for each listing condition value that is returned.
	NumberOfOfferListings NumberOfOfferListingsList `json:"NumberOfOfferListings"`
	TradeInValue          *MoneyType                `json:"TradeInValue,omitempty"`
}

// ConditionType defines model for ConditionType.
type ConditionType string

// List of ConditionType
const (
	ConditionType_Club        ConditionType = "Club"
	ConditionType_Collectible ConditionType = "Collectible"
	ConditionType_New         ConditionType = "New"
	ConditionType_Refurbished ConditionType = "Refurbished"
	ConditionType_Used        ConditionType = "Used"
)

// DetailedShippingTimeType defines model for DetailedShippingTimeType.
type DetailedShippingTimeType struct {

	// Indicates whether the item is available for shipping now, or on a known or an unknown date in the future. If known, the availableDate property indicates the date that the item will be available for shipping. Possible values: NOW, FUTURE_WITHOUT_DATE, FUTURE_WITH_DATE.
	AvailabilityType *string `json:"availabilityType,omitempty"`

	// The date when the item will be available for shipping. Only displayed for items that are not currently available for shipping.
	AvailableDate *interface{} `json:"availableDate,omitempty"`

	// The maximum time, in hours, that the item will likely be shipped after the order has been placed.
	MaximumHours *int64 `json:"maximumHours,omitempty"`

	// The minimum time, in hours, that the item will likely be shipped after the order has been placed.
	MinimumHours *int64 `json:"minimumHours,omitempty"`
}

// Error defines model for Error.
type Error struct {

	// An error code that identifies the type of error that occurred.
	Code string `json:"code"`

	// Additional information that can help the caller understand or fix the issue.
	Details *string `json:"details,omitempty"`

	// A message that describes the error condition in a human-readable form.
	Message string `json:"message"`
}

// ErrorList defines model for ErrorList.
type ErrorList []Error

// FulfillmentChannelType defines model for FulfillmentChannelType.
type FulfillmentChannelType string

// List of FulfillmentChannelType
const (
	FulfillmentChannelType_Amazon   FulfillmentChannelType = "Amazon"
	FulfillmentChannelType_Merchant FulfillmentChannelType = "Merchant"
)

// GetOffersResponse defines model for GetOffersResponse.
type GetOffersResponse struct {

	// A list of error responses returned when a request is unsuccessful.
	Errors  *ErrorList       `json:"errors,omitempty"`
	Payload *GetOffersResult `json:"payload,omitempty"`
}

// GetOffersResult defines model for GetOffersResult.
type GetOffersResult struct {

	// The Amazon Standard Identification Number (ASIN) of the item.
	ASIN *string `json:"ASIN,omitempty"`

	// Information that identifies an item.
	Identifier ItemIdentifier `json:"Identifier"`

	// Indicates the condition of the item. Possible values: New, Used, Collectible, Refurbished, Club.
	ItemCondition ConditionType `json:"ItemCondition"`

	// A marketplace identifier.
	MarketplaceID string          `json:"MarketplaceID"`
	Offers        OfferDetailList `json:"Offers"`

	// The stock keeping unit (SKU) of the item.
	SKU *string `json:"SKU,omitempty"`

	// Contains price information about the product, including the LowestPrices and BuyBoxPrices, the ListPrice, the SuggestedLowerPricePlusShipping, and NumberOfOffers and NumberOfBuyBoxEligibleOffers.
	Summary Summary `json:"Summary"`

	// The status of the operation.
	Status string `json:"status"`
}

// GetPricingResponse defines model for GetPricingResponse.
type GetPricingResponse struct {

	// A list of error responses returned when a request is unsuccessful.
	Errors  *ErrorList `json:"errors,omitempty"`
	Payload *PriceList `json:"payload,omitempty"`
}

// IdentifierType defines model for IdentifierType.
type IdentifierType struct {
	MarketplaceASIN ASINIdentifier       `json:"MarketplaceASIN"`
	SKUIdentifier   *SellerSKUIdentifier `json:"SKUIdentifier,omitempty"`
}

// ItemIdentifier defines model for ItemIdentifier.
type ItemIdentifier struct {

	// The Amazon Standard Identification Number (ASIN) of the item.
	ASIN *string `json:"ASIN,omitempty"`

	// Indicates the condition of the item. Possible values: New, Used, Collectible, Refurbished, Club.
	ItemCondition ConditionType `json:"ItemCondition"`

	// A marketplace identifier. Specifies the marketplace from which prices are returned.
	MarketplaceId string `json:"MarketplaceId"`

	// The seller stock keeping unit (SKU) of the item.
	SellerSKU *string `json:"SellerSKU,omitempty"`
}

// LowestPriceType defines model for LowestPriceType.
type LowestPriceType struct {
	LandedPrice  MoneyType `json:"LandedPrice"`
	ListingPrice MoneyType `json:"ListingPrice"`
	Points       *Points   `json:"Points,omitempty"`
	Shipping     MoneyType `json:"Shipping"`

	// Indicates the condition of the item. For example: New, Used, Collectible, Refurbished, or Club.
	Condition string `json:"condition"`

	// Indicates whether the item is fulfilled by Amazon or by the seller.
	FulfillmentChannel string `json:"fulfillmentChannel"`
}

// LowestPrices defines model for LowestPrices.
type LowestPrices []LowestPriceType

// MoneyType defines model for MoneyType.
type MoneyType struct {

	// The monetary value.
	Amount *float32 `json:"Amount,omitempty"`

	// The currency code in ISO 4217 format.
	CurrencyCode *string `json:"CurrencyCode,omitempty"`
}

// NumberOfOfferListingsList defines model for NumberOfOfferListingsList.
type NumberOfOfferListingsList []OfferListingCountType

// NumberOfOffers defines model for NumberOfOffers.
type NumberOfOffers []OfferCountType

// OfferCountType defines model for OfferCountType.
type OfferCountType struct {

	// The number of offers in a fulfillment channel that meet a specific condition.
	OfferCount *int32 `json:"OfferCount,omitempty"`

	// Indicates the condition of the item. For example: New, Used, Collectible, Refurbished, or Club.
	Condition *string `json:"condition,omitempty"`

	// Indicates whether the item is fulfilled by Amazon or by the seller (merchant).
	FulfillmentChannel *FulfillmentChannelType `json:"fulfillmentChannel,omitempty"`
}

// OfferDetail defines model for OfferDetail.
type OfferDetail struct {

	// When true, the offer is currently in the Buy Box. There can be up to two Buy Box winners at any time per ASIN, one that is eligible for Prime and one that is not eligible for Prime.
	IsBuyBoxWinner *bool `json:"IsBuyBoxWinner,omitempty"`

	// When true, the seller of the item is eligible to win the Buy Box.
	IsFeaturedMerchant *bool `json:"IsFeaturedMerchant,omitempty"`

	// When true, the offer is fulfilled by Amazon.
	IsFulfilledByAmazon bool      `json:"IsFulfilledByAmazon"`
	ListingPrice        MoneyType `json:"ListingPrice"`

	// When true, this is the seller's offer.
	MyOffer *bool   `json:"MyOffer,omitempty"`
	Points  *Points `json:"Points,omitempty"`

	// Information about the seller's feedback, including the percentage of positive feedback, and the total number of ratings received.
	SellerFeedbackRating *SellerFeedbackType `json:"SellerFeedbackRating,omitempty"`
	Shipping             MoneyType           `json:"Shipping"`

	// The time range in which an item will likely be shipped once an order has been placed.
	ShippingTime DetailedShippingTimeType `json:"ShippingTime"`

	// The state and country from where the item is shipped.
	ShipsFrom *ShipsFromType `json:"ShipsFrom,omitempty"`

	// The subcondition of the item. Subcondition values: New, Mint, Very Good, Good, Acceptable, Poor, Club, OEM, Warranty, Refurbished Warranty, Refurbished, Open Box, or Other.
	SubCondition string `json:"SubCondition"`
}

// OfferDetailList defines model for OfferDetailList.
type OfferDetailList []OfferDetail

// OfferListingCountType defines model for OfferListingCountType.
type OfferListingCountType struct {

	// The number of offer listings.
	Count int32 `json:"Count"`

	// The condition of the item.
	Condition string `json:"condition"`
}

// OfferType defines model for OfferType.
type OfferType struct {
	BuyingPrice PriceType `json:"BuyingPrice"`

	// The fulfillment channel for the offer listing. Possible values:
	//
	// * Amazon - Fulfilled by Amazon.
	// * Merchant - Fulfilled by the seller.
	FulfillmentChannel string `json:"FulfillmentChannel"`

	// The item condition for the offer listing. Possible values: New, Used, Collectible, Refurbished, or Club.
	ItemCondition string `json:"ItemCondition"`

	// The item subcondition for the offer listing. Possible values: New, Mint, Very Good, Good, Acceptable, Poor, Club, OEM, Warranty, Refurbished Warranty, Refurbished, Open Box, or Other.
	ItemSubCondition string    `json:"ItemSubCondition"`
	RegularPrice     MoneyType `json:"RegularPrice"`

	// The seller stock keeping unit (SKU) of the item.
	SellerSKU string `json:"SellerSKU"`
}

// OffersList defines model for OffersList.
type OffersList []OfferType

// Points defines model for Points.
type Points struct {
	PointsMonetaryValue *MoneyType `json:"PointsMonetaryValue,omitempty"`

	// The number of points.
	PointsNumber *int32 `json:"PointsNumber,omitempty"`
}

// Price defines model for Price.
type Price struct {

	// The Amazon Standard Identification Number (ASIN) of the item.
	ASIN *string `json:"ASIN,omitempty"`

	// An item.
	Product *Product `json:"Product,omitempty"`

	// The seller stock keeping unit (SKU) of the item.
	SellerSKU *string `json:"SellerSKU,omitempty"`

	// The status of the operation.
	Status string `json:"status"`
}

// PriceList defines model for PriceList.
type PriceList []Price

// PriceType defines model for PriceType.
type PriceType struct {
	LandedPrice  *MoneyType `json:"LandedPrice,omitempty"`
	ListingPrice MoneyType  `json:"ListingPrice"`
	Points       *Points    `json:"Points,omitempty"`
	Shipping     *MoneyType `json:"Shipping,omitempty"`
}

// Product defines model for Product.
type Product struct {

	// A list of product attributes if they are applicable to the product that is returned.
	AttributeSets *AttributeSetList `json:"AttributeSets,omitempty"`

	// Competitive pricing information for the item.
	CompetitivePricing *CompetitivePricingType `json:"CompetitivePricing,omitempty"`

	// Specifies the identifiers used to uniquely identify an item.
	Identifiers IdentifierType `json:"Identifiers"`

	// A list of offers.
	Offers *OffersList `json:"Offers,omitempty"`

	// A list that contains product variation information, if applicable.
	Relationships *RelationshipList `json:"Relationships,omitempty"`

	// A list of sales rank information for the item, by category.
	SalesRankings *SalesRankList `json:"SalesRankings,omitempty"`
}

// RelationshipList defines model for RelationshipList.
type RelationshipList []map[string]interface{}

// SalesRankList defines model for SalesRankList.
type SalesRankList []SalesRankType

// SalesRankType defines model for SalesRankType.
type SalesRankType struct {

	//  Identifies the item category from which the sales rank is taken.
	ProductCategoryId string `json:"ProductCategoryId"`

	// The sales rank of the item within the item category.
	Rank int32 `json:"Rank"`
}

// SellerFeedbackType defines model for SellerFeedbackType.
type SellerFeedbackType struct {

	// The number of ratings received about the seller.
	FeedbackCount int64 `json:"FeedbackCount"`

	// The percentage of positive feedback for the seller in the past 365 days.
	SellerPositiveFeedbackRating *float64 `json:"SellerPositiveFeedbackRating,omitempty"`
}

// SellerSKUIdentifier defines model for SellerSKUIdentifier.
type SellerSKUIdentifier struct {

	// A marketplace identifier.
	MarketplaceId string `json:"MarketplaceId"`

	// The seller identifier submitted for the operation.
	SellerId string `json:"SellerId"`

	// The seller stock keeping unit (SKU) of the item.
	SellerSKU string `json:"SellerSKU"`
}

// ShipsFromType defines model for ShipsFromType.
type ShipsFromType struct {

	// The country from where the item is shipped.
	Country *string `json:"Country,omitempty"`

	// The state from where the item is shipped.
	State *string `json:"State,omitempty"`
}

// Summary defines model for Summary.
type Summary struct {
	BuyBoxEligibleOffers *BuyBoxEligibleOffers `json:"BuyBoxEligibleOffers,omitempty"`
	BuyBoxPrices         *BuyBoxPrices         `json:"BuyBoxPrices,omitempty"`
	ListPrice            *MoneyType            `json:"ListPrice,omitempty"`
	LowestPrices         *LowestPrices         `json:"LowestPrices,omitempty"`
	NumberOfOffers       *NumberOfOffers       `json:"NumberOfOffers,omitempty"`

	// When the status is ActiveButTooSoonForProcessing, this is the time when the offers will be available for processing.
	OffersAvailableTime             *time.Time `json:"OffersAvailableTime,omitempty"`
	SuggestedLowerPricePlusShipping *MoneyType `json:"SuggestedLowerPricePlusShipping,omitempty"`

	// The number of unique offers contained in NumberOfOffers.
	TotalOfferCount int32 `json:"TotalOfferCount"`
}

// GetCompetitivePricingParams defines parameters for GetCompetitivePricing.
type GetCompetitivePricingParams struct {

	// A marketplace identifier. Specifies the marketplace for which prices are returned.
	MarketplaceId string `json:"MarketplaceId"`

	// A list of up to twenty Amazon Standard Identification Number (ASIN) values used to identify items in the given marketplace.
	Asins *[]string `json:"Asins,omitempty"`

	// A list of up to twenty seller SKU values used to identify items in the given marketplace.
	Skus *[]string `json:"Skus,omitempty"`

	// Indicates whether ASIN values or seller SKU values are used to identify items. If you specify Asin, the information in the response will be dependent on the list of Asins you provide in the Asins parameter. If you specify Sku, the information in the response will be dependent on the list of Skus you provide in the Skus parameter. Possible values: Asin, Sku.
	ItemType string `json:"ItemType"`
}

// GetItemOffersParams defines parameters for GetItemOffers.
type GetItemOffersParams struct {

	// A marketplace identifier. Specifies the marketplace for which prices are returned.
	MarketplaceId string `json:"MarketplaceId"`

	// Filters the offer listings to be considered based on item condition. Possible values: New, Used, Collectible, Refurbished, Club.
	ItemCondition string `json:"ItemCondition"`
}

// GetListingOffersParams defines parameters for GetListingOffers.
type GetListingOffersParams struct {

	// A marketplace identifier. Specifies the marketplace for which prices are returned.
	MarketplaceId string `json:"MarketplaceId"`

	// Filters the offer listings based on item condition. Possible values: New, Used, Collectible, Refurbished, Club.
	ItemCondition string `json:"ItemCondition"`
}

// GetPricingParams defines parameters for GetPricing.
type GetPricingParams struct {

	// A marketplace identifier. Specifies the marketplace for which prices are returned.
	MarketplaceId string `json:"MarketplaceId"`

	// A list of up to twenty Amazon Standard Identification Number (ASIN) values used to identify items in the given marketplace.
	Asins *[]string `json:"Asins,omitempty"`

	// A list of up to twenty seller SKU values used to identify items in the given marketplace.
	Skus *[]string `json:"Skus,omitempty"`

	// Indicates whether ASIN values or seller SKU values are used to identify items. If you specify Asin, the information in the response will be dependent on the list of Asins you provide in the Asins parameter. If you specify Sku, the information in the response will be dependent on the list of Skus you provide in the Skus parameter.
	ItemType string `json:"ItemType"`

	// Filters the offer listings based on item condition. Possible values: New, Used, Collectible, Refurbished, Club.
	ItemCondition *string `json:"ItemCondition,omitempty"`
}
