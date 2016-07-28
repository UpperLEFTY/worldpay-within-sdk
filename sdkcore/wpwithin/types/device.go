package types

type Device struct {

	Uid string
	Name string
	Description string
	Services map[int]*Service
	IPv4Address string
}

func NewDevice(name, description, uid, ipv4Address, currencyCode string) (*Device, error) {

	result := &Device{
		Name:name,
		Description:description,
		Uid:uid,
		IPv4Address: ipv4Address,
	}

	return result, nil
}