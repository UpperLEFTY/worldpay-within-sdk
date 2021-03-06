##############################################
#
# Worldpay Within SDK Thrift definition
# Conor Hackett (conor.hackett@worldpay.com)
# Revised by Kevin Gordon (kevin.gordon@worldpay.com)
# January 31st, 2017
#
#############################################

namespace csharp Worldpay.Innovation.WPWithin.Rpc.Types
namespace java com.worldpay.innovation.wpwithin.rpc.types
namespace go wpthrift_types
namespace js wpthrift_types
namespace py wpthrift_types

exception Error {

	1: string message
}

struct PricePerUnit {

        1: i32 amount
        2: string currencyCode
}

struct Price {

        1: i32 id
        2: string description
        3: PricePerUnit pricePerUnit
        4: i32 unitId
        5: string unitDescription
}

struct Service {

	1: i32 id
	2: string name
	3:  string description
	4: optional map<i32, Price> prices  /* This should be optional now but these are stored as pointers and was causing an issue in Go - TODO CH - Conor to investigate and fix */
}

struct HCECard {

	1: string FirstName
	2: string LastName
	3: i32 ExpMonth
	4: i32 ExpYear
	5: string CardNumber
	6: string Type
	7:  string Cvc
}

struct Device {

	1: string uid
	2: string name
	3:  string description
	4:  map<i32, Service> services
	5: string ipv4Address
	6: string currencyCode
}

struct ServiceMessage {

	1: string deviceDescription
	2: string hostname
	3: i32 portNumber
	4: string serverId
	5: string urlPrefix
	6: string scheme
}

struct ServiceDetails {

	1: i32 serviceId
	2: string serviceDescription
}

struct TotalPriceResponse {

	1: string serverId
	2: string clientId
	3: i32 priceId
	4: i32 unitsToSupply
	5: i32 totalPrice
	6: string paymentReferenceId
	7: string merchantClientKey
	8: string currencyCode
}

struct ServiceDeliveryToken {

        1: string key
        2: string issued
        3: string expiry
        4: bool refundOnExpiry
        5: binary signature
}

struct PaymentResponse {

	1: string serverId
	2: string clientId
	3: i32 totalPaid
	4: ServiceDeliveryToken serviceDeliveryToken
}

