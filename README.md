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

#### BUILD
The binary is built inside of a two-stage docker build and produces a docker container containing the CLI binary.
```
./build.sh
```

The individual binary can be built manually with the `./build_cli.sh` command.
