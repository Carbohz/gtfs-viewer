package model

import "net/url"

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
	RouteId           uint64
	AgencyId          uint64
	RouteShortName    string
	RouteLongName     string
	RouteDesc         string
	RouteType         RouteType
	RouteUrl          url.URL
	RouteColor        Color
	RouteTextColor    Color
	RouteSortOrder    uint64
	ContinuousPickup  ContinuousPickup
	ContinuousDropOff ContinuousDropOff
	NetworkId         uint64
}
