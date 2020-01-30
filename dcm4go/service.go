package dcm4go

import "fmt"

// A Service is used to handle requests of DICOM services
type Service interface {
	onCommand(
		assoc *AcceptorAssoc,
		presContext *PresContext,
		command *Object,
		reader *PDataReader,
	) error
	onRequest(
		assoc *AcceptorAssoc,
		presContext *PresContext,
		command *Object,
		data *Object,
	) error
}

// A BasicService is the default service provided by the library
type BasicService struct {
	capabilities []*Capability
}

func newBasicService() *BasicService {
	return &BasicService{make([]*Capability, 0, 5)}
}

func (service *BasicService) addCapability(capability *Capability) {
	service.capabilities = append(service.capabilities, capability)
}

func (service *BasicService) onCommand(
	assoc *AcceptorAssoc,
	presContext *PresContext,
	command *Object,
	reader *PDataReader,
) error {
	// by default, read the data
	data, err := readData(reader, &assoc.Assoc, presContext.id)
	if err != nil {
		return err
	}
	// then call the handler to manage the request
	return service.onRequest(assoc, presContext, command, data)
}

func (service *BasicService) onRequest(
	assoc *AcceptorAssoc,
	presContext *PresContext,
	command *Object,
	data *Object,
) error {
	// hmm, need to do something
	return fmt.Errorf("BasicService.onRequest() not implemented")
}
