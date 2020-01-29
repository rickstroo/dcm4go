package dcm4go

import "fmt"

// A Service is used to handle requests of DICOM services
type Service interface {
	onClose(
		assoc *AcceptorAssoc,
	) error
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
	abstractSyntaxes []string
}

func (service *BasicService) onClose(assoc *AcceptorAssoc) error {
	// do nothing by default execpt return all is well
	return nil
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
