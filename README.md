# dcm4go

dcm4go is an implementation of a DICOM library and some DICOM tools using the Go language.  Mostly, it's an exercise in learning how to use Go.  I think this is now my 4th DICOM library that I've implemented as a method to learn a new language.

There are a number of other implementations of DICOM libraries written in Go.  All of them implementing parsing of DICOM files.  Some of them implement the DICOM networking protocols.  So, why build another?  As mentioned earlier, mostly as a reason to learn Go.  But also, not just to learn the language, but to exercise some design principles.  Here are some things that I'm trying to accomplish.

First and foremost, keep it simple, without being restrictive.  To me, this means very simple APIs that allow users of the libraries to start writing applications quickly.  Go suggests that when designing things, make the zero values of things useful.  So, the library tries to do follow that principle.  At the same time, there should be ways to allow users of the library to get into the weeds and to specify fine grain control over the actions of the library.

Another design principle is to support the management of large objects without putting significant pressure on memory.  More specifically, that means that when reading and writing DICOM objects to and from files and the network, that we need to do better than simply reading the entire object into memory.  In particular, the user of the library needs to be able to receive a DICOM object from the network and write that object directly to a file.  I've not seen that capability in any of the Go implementations that I've looked at, however, the dcm4che implementation in Java provides a nice model for how to accomplish that.

Following on the design principles for the Go language, there should be one way to get things done.  That needs to be balanced against a desire to support simple default behaviours.

## Getting Started

Download the code.  Run the generator.  Copy the generated Go files to the library folder.  Build the library.  Build the tools.  More detailed explanations will come eventually.

### Prerequisites

A development environment for Go.

```
Give examples
```

### Installing

A step by step series of examples that tell you how to get a development env running

Say what the step will be

```
Give the example
```

And repeat

```
until finished
```

End with an example of getting some data out of the system or using it for a little demo

## Running the tests

Explain how to run the automated tests for this system

### Break down into end to end tests

Explain what these tests test and why

```
Give an example
```

### And coding style tests

Explain what these tests test and why

```
Give an example
```

## Deployment

Add additional notes about how to deploy this on a live system

## Built With

dcm4go was built using the Atom IDE and the go-plus and go-debug plugins.

## Contributing

At this time, I'm not accepting contributions.  If this project becomes useful to a larger community at some point, we will revisit.

## Versioning

At this time, we're not versioning.  The develop branch is where development happens.  At points in time when we've developed useful functionality, we'll merge with master.

## Authors

Rick Stroobosscher (https://github.com/rickstroo)

## License

This project is licensed under the MIT License - see the [LICENSE.md](LICENSE.md) file for details

## Acknowledgments

I'd like to acknowledge the work of Jeremy Huiskamp (https://github.com/jeremyhuiskamp).  I have been referencing his implementation of DICOM using Go for inspiration and learning.

I'd also like to acknowledge the work of the dcm4che community.  The name of this project, dcm4go, is derived from dcm4che.  I've also used dcm4che many times over the years to get ideas about design and to learn about the details of DICOM.
