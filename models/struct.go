package models

type Property struct {
	Name             string `bson:"name" validate:"required"`
	Location         string `bson:"location" validate:"required"`
	AvailabilityFor  string `bson:"availability_for" validate:"required"`
	CompletionStatus string `bson:"completion_status" validate:"required"`
	FurnishingStatus string `bson:"furnishing_status" validate:"required"`
	Floors           string `bson:"no_of_floor"  validate:"required"`
	Parking          string `bson:"parking" validate:"required"`
	Oc               string `bson:"oc" validate:"required"`
	Lift             string `bson:"lift" validate:"required"`
	OverLooking      string `bson:"overlooking" validate:"required"`
}

type Country struct {
	CountryName    string `bson:"country_name"  validate:"required"`
	CurrencyName   string `bson:"currency_name"  validate:"required"`
	CurrencySymbol string `bson:"currency_symbol"  validate:"required"`
}

type State struct {
	StateName     string `bson:"state_name"  validate:"required"`
	StateLanguage string `bson:"state_language"  validate:"required"`
}

type City struct {
	CityName string `bson:"city_name"  validate:"required"`
}

type Locality struct {
	LocalityName string `bson:"locality_name"  validate:"required"`
}
