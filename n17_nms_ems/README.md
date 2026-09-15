# NMS vs EMS

A large network may contain thousands of routers, switches, access points, and other devices.
Managing every device manually does not scale.

## EMS - Element Management System

EMS manages individual network elements or a group of similar network elements.

Examples of responsibilities:
- Device configuration
- Device status
- Alarm collection
- Software upgrades
- Performance monitoring

## NMS - Network Management System

NMS manages the network at a higher level.

It can:
- Monitor the complete network
- Work with multiple EMS systems
- Correlate alarms
- Show network-wide topology and health
- Coordinate network-level configuration

## Mental Model

NMS
|
+-- EMS
|   +-- Device
|   +-- Device
|
+-- EMS
+-- Device
+-- Device

## Interview Answer

EMS manages individual network elements or groups of similar elements.
NMS provides a higher-level network-wide view and can manage or coordinate multiple EMS systems.