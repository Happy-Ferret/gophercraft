# Gophercraft Design Guidelines

## No unsafe libraries

Gophercraft should not have any C/C++ dependencies. Everything should be done with Go, to be secure against memory-based exploits and other attacks.

## Server commands

### Tier 1 (anyone)

.help
.info
.unstuck
.ipecho
.guildcreate

### Tier 2 (privileged)

.bank (opens bank remotely)
.gbank (opens guild bank remotely)
.mail (opens mail remotely)

### Tier 3 (GM fun commands)

.modify
 scale (Changes player scale)
 hp 


### Tier 4 (GM administrative commands)

.ban
.kick