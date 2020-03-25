package dcm4go

// import (
// 	"bytes"
// 	"io"
// 	"log"
// )
//
// type pduReader0 struct {
// 	reader io.Reader
// }
//
// type pduConsumer0 struct {
// }
//
// type pdvConsumer0 struct {
// }
//
// func start(reader io.Reader) {
// 	pduChan := make(chan *pdu, 1)
// 	pdvChan := make(chan *pdv, 1)
// 	pduReader := &pduReader0{reader: reader}
// 	pduConsumer := &pduConsumer0{}
// 	pdvConsumer := &pdvConsumer0{}
// 	go pdvConsumer.run(pdvChan)
// 	go pduConsumer.run(pduChan, pdvChan)
// 	go pduReader.run(pduChan)
// }
//
// func (pdvConsumer *pdvConsumer0) run(pdvChan chan *pdv) {
// 	for pdv := range pdvChan {
// 		log.Printf("pdv is %v", pdv)
// 	}
// }
//
// func (pduConsumer *pduConsumer0) run(pduChan chan *pdu, pdvChan chan *pdv) {
// 	for pdu := range pduChan {
// 		log.Printf("pdu is %v", pdu)
// 		switch pdu.typ {
// 		case pDataTFPDU:
// 			pduConsumer.onPDataTFPDU(pdu, pdvChan)
// 		}
// 	}
// 	close(pdvChan)
// }
//
// func (pduConsumer *pduConsumer0) onPDataTFPDU(pdu *pdu, pdvChan chan *pdv) {
// 	byteReader := bytes.NewReader(pdu.buf)
// 	for {
// 		pdv, err := readPDV(byteReader)
// 		if err != nil {
// 			if err != io.EOF {
// 				log.Printf("error while reading pdv, error is %v", err)
// 			}
// 			break
// 		}
// 		pdvChan <- pdv
// 	}
// }
//
// func (pduReader *pduReader0) run(pduChan chan *pdu) {
// 	for {
// 		pdu, err := readPDU(pduReader.reader)
// 		if err != nil {
// 			if err != io.EOF {
// 				log.Printf("error while reading pdus, error is %v", err)
// 			}
// 			break
// 		}
// 		pduChan <- pdu
// 	}
// 	close(pduChan)
// }
