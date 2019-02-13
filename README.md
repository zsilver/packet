# Packet API Integration
This repository demostrates a basic [Packet API](https://www.packet.com/developers/api) integration using the [Golang client](https://godoc.org/github.com/packethost/packngo).

## CLI Tool
The cli tool is a basic interface to list, create, and delete devices (machines) from a terminal.

#### Basic Operations
* Help
```
docker run -it zilver16/packet:latest help
```

* List Projects
```
docker run -it -e PACKET_AUTH_TOKEN=[secret token] zilver16/packet:latest list
````

* List Project Devices
```
docker run -it -e PACKET_AUTH_TOKEN=[secret token] zilver16/packet:latest device list --id ca73364c-6023-4935-9137-2132e73c20b4
```

* Create Device
```
docker run -it -e PACKET_AUTH_TOKEN=[secret token] zilver16/packet:latest device create --OS debian_8  --Plan m2.xlarge.x86 -ProjectID ca73364c-6023-4935-9137-2132e73c20b4 -BillingCycle hourly
```

* Delete Device
```
docker run -it -e PACKET_AUTH_TOKEN=[secret token] zilver16/packet:latest device delete --id [device UUID]
```

#### BUILD
The binary is built inside of a two-stage docker build and produces a docker container containing the CLI binary.
```
./build.sh
```

The individual binary can be built manually with the `./build_cli.sh` command.

## Remaining Tasks
* I've never been able to test `delete`
  * After processing the server simply disappears. So this is untested!
* Unit tests and integrations
  * See feedback, but the design of the Packet client makes this terribly difficult due to mocking
* Test for performance enhancements
  * See inline comments about reflection (horrible!)
* Create HTTP server in addition to CLI Tool
  * May be helpful to deploy a serverless (GCP Function or AWS Lambda) with a basic UI
  
## Feedback
Overall this was very helpful in understanding the basics Packet's product(s). Minus a few questions that required an email notice, this worked fairly seemlessly.
* ~2Hrs Development Time
  * 1/2 hr spent understanding and testing the Packet API
  * 1 hr spent trying out new CLI parsing tool (I just wanted to play with this) and reflection (it's been a while since I've used Golang's builtin reflection package)
  * 1/2 Documenting (this), adding inline comments, and adding minor code style refactoring
* TODO: 1 hr to add in unittests
  * I'm loath to do this since Packet's client API is not an Interface. So it require a large amount of effort to unittest (see comments below).
    * Basically I need to create a new Packet Client Interface
    * Wrap the native packet client in a custom struct
    * Then mock all calls to the real client for unittests.

#### Bugs & Typos
* Correct assigment FQS documentation typo "/project" → "/projects"
* Mention that /capacity api is not accessible. Also remove it from the assignment documentation

#### Suggestions
* Does the REST API has a pretty print option? (i.e. ?pretty=true)
* “id” is really an overloaded term, it may help to be more specific (i.e. “project_id”)
* It would be good if all API's can be tested via the API documentation site
  * Common Errors I would experience
    * DELETE https://api.packet.net/devices/90362283-d815-4a32-bf97-590745b9c29d: 403 You are not authorized to delete this device 
    * POST https://api.packet.net/projects/ca73364c-6023-4935-9137-2132e73c20b4/devices: 422 There aren't available servers at any facility
* Example plan & OS for creating a device would be very helpful on getting started quickly. I spent some time just playing with different combinations hoping to get a hit.

#### Important Golang API Suggestions
* Examples of each service in the Godoc would be helpful
  * I actually had to requently look at the source code to see how to use some of the API's
* TDD Friendliness
  * The Golang client API is not very feasible for Test-Driven Developement. The main reason for this is the client should actually be an interface. That way we can easily mock the client using standard Golang techniques. This is HUGE! As consumers of the `client` struct, we shouldn't be forced to access fields, rather we should depend on assessor functions.
