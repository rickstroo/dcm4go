package dcm4go

// // A Sender can send a DICOM object to an AE.
// type Sender struct {
// 	Opts *SenderOpts
// }
//
// // SenderOpts impact the behaviour of a Sender.
// type SenderOpts struct {
// 	LocalAETitle   string        // a zero value means "DCMSND"
// 	ConnectTimeOut time.Duration // a zero value means no connect timeout
// 	WriteTimeOut   time.Duration // a zero value means no write timeout
// 	ReadTimeOut    time.Duration // a zero value means no read timeout
// }
//
// // Send sends a DICOM object to an AE.
// // The address of the AE is of the format 'aetitle@host:port'.
// func (sender *Sender) Send(paths []string, remote string) error {
// 	return sender.send(paths, remote)
// }
//
// // Send sends a DICOM object to another AE using a default set of options.
// // To gain more control over the sending, the user should create a Sender
// // with the desired SenderOpts.
// func Send(paths []string, remoteAddr string) error {
// 	opts := &SenderOpts{}
// 	sender := &Sender{Opts: opts}
// 	return sender.Send(paths, remoteAddr)
// }
//
// func (sender *Sender) send(paths []string, remoteAddr string) error {
//
// 	// create a local AE
// 	localAE := NewAE(sender.Opts.LocalAETitle)
//
// 	// define some options for the association
// 	assocOpts := &AssocOpts{
// 		WriteTimeOut: 10 * time.Second,
// 		ReadTimeOut:  10 * time.Second,
// 		MaxBufLen:    16 * 1024,
// 	}
//
// 	// read the transfer capabilities from all the files
// 	capabilities, err := ReadCapabilities(paths)
// 	if err != nil {
// 		return err
// 	}
//
// 	// create the remote AE
// 	remoteAE := NewAE(remoteAddr)
//
// 	// connect to the remote
// 	conn, err := net.Dial("tcp", remoteAE.Host()+":"+remoteAE.Port())
// 	if err != nil {
// 		return err
// 	}
// 	log.Printf("opened connection from %v to %v", conn.LocalAddr(), conn.RemoteAddr())
//
// 	// create an association
// 	assoc, err := localAE.RequestAssoc(conn, remoteAE, capabilities, assocOpts)
// 	if err != nil {
// 		return err
// 	}
// 	log.Printf(
// 		"created association from %s to %s",
// 		assoc.CallingAETitle(),
// 		assoc.CalledAETitle(),
// 	)
//
// 	// ensure the association gets released
// 	defer func() {
// 		if err := assoc.RequestRelease(); err != nil {
// 			log.Printf("while releasing association, caught error %v", err)
// 		} else {
// 			log.Printf("released association")
// 		}
//
// 		if err := conn.Close(); err != nil {
// 			log.Printf("while closing connection, caught error %v", err)
// 		} else {
// 			log.Printf("closed connection")
// 		}
// 	}()
//
// 	// send the files
// 	for i, path := range paths {
//
// 		// open the file
// 		file, err := os.Open(path)
// 		if err != nil {
// 			log.Printf("error while opening file, %v", err)
// 			continue
// 		}
//
// 		// send the file
// 		if err := assoc.Store(file); err != nil {
// 			log.Printf("error while sending file, %v", err)
// 		}
//
// 		// close the file
// 		if err := file.Close(); err != nil {
// 			log.Printf("error while closing file, %v", err)
// 		}
//
// 		log.Printf("sent file %d of %d", i, len(paths))
// 	}
//
// 	// return success
// 	return nil
// }
