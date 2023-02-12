package models

import (
	"net/url"
)

type RouteType uint16
type Color string

const (
	Streetcar  RouteType = 0
	Subway     RouteType = 1
	Rail       RouteType = 2
	Bus        RouteType = 3
	Ferry      RouteType = 4
	CableTram  RouteType = 5
	AerialLift RouteType = 6
	Funicular  RouteType = 7
	TrolleyBus RouteType = 11
	Monorail   RouteType = 12
)

type ContinuousPickup uint8

const (
	Pickup                     ContinuousPickup = 0
	NoPickup                   ContinuousPickup = 1
	PhoneAgencyPickUp          ContinuousPickup = 2
	CoordinateWithDriverPickUp ContinuousPickup = 3
)

type ContinuousDropOff uint8

const (
	DropOff                     ContinuousDropOff = 0
	NoDropOff                   ContinuousDropOff = 1
	PhoneAgencyDropOff          ContinuousDropOff = 2
	CoordinateWithDriverDropOff ContinuousDropOff = 3
)

type Route struct {
	Id                uint64            `csv:"route_id"`
	AgencyId          uint64            `csv:"agency_id"`
	ShortName         string            `csv:"route_short_name"`
	LongName          string            `csv:"route_long_name"`
	Desc              string            `csv:"route_desc"`
	Type              RouteType         `csv:"route_type"`
	Url               url.URL           `csv:"route_id"`
	Color             Color             `csv:"route_color"`
	TextColor         Color             `csv:"route_text_color"`
	SortOrder         uint64            `csv:"route_sort_order"`
	ContinuousPickup  ContinuousPickup  `csv:"continuous_pickup"`
	ContinuousDropOff ContinuousDropOff `csv:"continuous_drop_off"`
	NetworkId         uint64            `csv:"network_id"`
}
