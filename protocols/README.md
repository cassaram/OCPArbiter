# Protocols

Within this folder are the separate modules for protocols implemented by controllers and devices. Some devices have their protocols integrated within their respective module if it is a one-off protocol and not re-used in other devices.

## Currently Supported Devices
- PCI (pci)
  - Used by Grass Valley to communicate between OCP controllers and LDX cameras over serial. PCI specifically implements a fake "camera" side
- PC (pc)
  - Used by Grass Valley to communicate between OCP controllers and LDX cameras over serial. PC specifically implements a fake "controller" side