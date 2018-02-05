# Gophercraft (Very WIP)

![gopher](gopher.png)

Gophercraft is a Go library for experimentation with various World of Warcraft-related protocols and formats.

Gophercraft is currently in development: expect a lack of features and expect the API to change frequently.

Remember, contributions are always welcome. If you've added additional functionality, why not submit a pull request?

**Alert! I've reached an impasse with this code. I can't seem to implement code that will make the WoW 3.3.5a client log into the world, past the character selection menu. If you know how to do this, please contribute!!**


## Server setup/installation

Since this software is still in its infancy, no game data is necessary to launch.

```bash
# install packages
sudo apt-get install golang mysql-server

# create database
echo "CREATE DATABASE gcraft_core;" | mysql -u root -p 

# install gophercraft commands
go get -u -v github.com/superp00t/gophercraft/cmd/...

# launch authserver (do this in background) 
gcraft_core_auth -s root:password@/gcraft_core?charset=utf8 -l 0.0.0.0:3724

# launch worldserver                                            publicly visible     listening address
#                                                               address
gcraft_core_world -s root:password@/gcraft_core?charset=utf8 -i 0.0.0.0:8085        -l 0.0.0.0:8085
```

To register, point your browser to http://localhost:8086 and fill out the registration form.

## Goals

- Authentication and realm list server ✓
- Authentication protocol client ✓
- HTTP JSON API for facilitating registration and serving up world data ✓
- Rich web application utilizing the Gophercraft JSON API, browsing players, stats, items and guilds
- Scriptable worldserver 
- Support for Windows and Mac OS X
- Game protocol client
- Conversion/extraction tools written in pure Go

In general, this project aims to create a secure, readable and extensible codebase for MMORPG software.

## Why start from scratch in Go?

Go uses a simpler syntax than C++, making code reading easier.

Unlike C++, Go provides memory safety at the language level, making memory-based RCE exploits impossible.

Go has nicer package management. It also provides many libraries that the protocol requires, like `net`, `crypto/sha1`, `crypto/hmac` and `compress/zlib` as part of the Go Standard Library, eliminating the need to add external dependencies.

## Programs

- <a href="cmd/gcraft_core_auth">Authentication server</a>
- <a href="cmd/gcraft_core_world">World/game server</a>
- <a href="cmd/gcraft_wowser_pipeline">Serve game data to Wowser</a>

## Screenshots

![worldserverscreenshot](realmlist.png)

## Thanks

For helping me to understand the protocol, I would like to thank these projects:

- [Wowser, a browser-based WoW client](https://github.com/wowserhq/wowser)
- [mClient, a WoW client written in C#](https://github.com/justMaku/mClient)
- [TrinityCore, a C++ MMORPG server](https://github.com/TrinityCore/TrinityCore/)
