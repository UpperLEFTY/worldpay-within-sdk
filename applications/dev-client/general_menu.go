package main

import (
	"encoding/json"
	"errors"
	"fmt"
	log "github.com/Sirupsen/logrus"
	"innovation.worldpay.com/worldpay-within-sdk/applications/dev-client/dev-client-defaults"
	"innovation.worldpay.com/worldpay-within-sdk/applications/dev-client/dev-client-errors"
	"innovation.worldpay.com/worldpay-within-sdk/applications/dev-client/dev-client-types"
	"innovation.worldpay.com/worldpay-within-sdk/sdkcore/wpwithin"
	"innovation.worldpay.com/worldpay-within-sdk/sdkcore/wpwithin/rpc"
	"innovation.worldpay.com/worldpay-within-sdk/sdkcore/wpwithin/types"
	"io/ioutil"
)

func mGetDeviceInfo() error {

	fmt.Println("Device Info:")

	if sdk == nil {
		return errors.New(devclienterrors.ERR_DEVICE_NOT_INITIALISED)
	}

	fmt.Printf("Uid of device: %s\n", sdk.GetDevice().Uid)
	fmt.Printf("Name of device: %s\n", sdk.GetDevice().Name)
	fmt.Printf("Description: %s\n", sdk.GetDevice().Description)
	fmt.Printf("Services: \n")

	for i, service := range sdk.GetDevice().Services {
		fmt.Printf("   %d: Id:%d Name:%s Description:%s\n", i, service.Id, service.Name, service.Description)
		fmt.Printf("   Prices: \n")
		for j, price := range service.Prices() {
			fmt.Printf("      %d: ServiceID: %d ID:%d Description:%s PricePerUnit:%d UnitID:%d UnitDescription:%s\n", j, service.Id, price.ID, price.Description, price.PricePerUnit, price.UnitID, price.UnitDescription)
		}
	}

	fmt.Printf("IPv4Address: %s\n", sdk.GetDevice().IPv4Address)

	return nil
}

func mInitDefaultDevice() error {

	fmt.Println("Initialising default device...")

	_sdk, err := wpwithin.Initialise(devclientdefaults.DEFAULT_DEVICE_NAME, devclientdefaults.DEFAULT_DEVICE_DESCRIPTION)

	if err != nil {

		return err
	}

	sdk = _sdk

	return nil
}

func mInitNewDevice() error {

	fmt.Println("Initialising new device")

	fmt.Print("Name of device: ")
	var nameOfDevice string
	if err := getUserInput(&nameOfDevice); err != nil {
		return err
	}

	fmt.Print("Description: ")
	var description string
	if err := getUserInput(&description); err != nil {
		return err
	}

	_sdk, err := wpwithin.Initialise(nameOfDevice, description)

	if err != nil {

		return err
	}

	sdk = _sdk

	return err
}

func mResetSessionState() error {

	sdk = nil

	fmt.Println("Reset session state")

	return nil
}

func mStartRPCService() error {

	fmt.Println("Starting rpc service...")

	config := rpc.Configuration{
		Protocol:   "binary",
		Framed:     false,
		Buffered:   false,
		Host:       "127.0.0.1",
		Port:       9091,
		Secure:     false,
		BufferSize: 8192,
	}

	rpc, err := rpc.NewService(config, sdk)

	if err != nil {
		return err
	}

	chErr := make(chan error, 1)

	go func() {
		chErr <- rpc.Start()
	}()

	var rpcErr error

	// Error handling go routine
	go func() {
		rpcErr := <-chErr
		if rpcErr != nil {
			log.Debug("error ", rpcErr)
		}

		close(chErr)
	}()

	return rpcErr
}

func mLoadDeviceProfile() error {

	fmt.Print("Name of profile to load: ")
	var profileStr string
	if err := getUserInput(&profileStr); err != nil {
		return err
	}

	file, err := ioutil.ReadFile(profileStr)
	if err != nil {
		log.Debug("error ", err)
		return err
	}

	json.Unmarshal(file, &deviceProfile)

	if deviceProfile.DeviceEntity != nil {
		if err := initialiseDevice(deviceProfile.DeviceEntity); err != nil {
			return err
		}
		fmt.Println("Setup device.")
	}

	if deviceProfile.DeviceEntity.Producer != nil {
		if err := setupProducer(deviceProfile.DeviceEntity.Producer); err != nil {
			return err
		}
		fmt.Println("Setup producer.")
	}

	if deviceProfile.DeviceEntity.Consumer != nil {
		if err := setupConsumer(deviceProfile.DeviceEntity.Consumer); err != nil {
			return err
		}
		fmt.Println("Setup consumer.")
	}

	return nil
}

func setupProducer(producer *devclienttypes.Producer) error {
	if err := addHTECredentials(producer.ProducerConfig); err != nil {
		return err
	}

	if _, err := sdk.InitProducer(); err != nil {
		return err
	}

	if err := addServicesAndPrices(producer.Services); err != nil {
		return err
	}

	return nil
}

func initialiseDevice(deviceEntity *devclienttypes.DeviceEntity) error {
	_sdk, err := wpwithin.Initialise(deviceEntity.Name, deviceEntity.Description)

	if err != nil {
		return err
	}

	sdk = _sdk

	return nil
}

func addHTECredentials(producerConfig *devclienttypes.ProducerConfig) error {
	return sdk.InitHTE(producerConfig.PspMerchantClientKey, producerConfig.PspMerchantServiceKey)
}

func addServicesAndPrices(services []*devclienttypes.ServiceProfile) error {

	for _, service := range services {

		newService, _ := types.NewService()
		newService.Id = service.Id
		newService.Name = service.Name
		newService.Description = service.Description

		for _, price := range service.Prices {

			newPrice := types.Price{
				UnitID:          price.UnitID,
				ID:              price.ID,
				Description:     price.Description,
				UnitDescription: price.UnitDescription,
				PricePerUnit: &types.PricePerUnit{
					Amount:       price.PricePerUnit.Amount,
					CurrencyCode: price.PricePerUnit.CurrencyCode,
				},
			}

			newService.AddPrice(newPrice)
		}

		if err := sdk.AddService(newService); err != nil {
			return err
		}
	}

	return nil
}

func setupConsumer(consumer *devclienttypes.Consumer) error {
	return sdk.InitHCE(*consumer.HCECard)
}
