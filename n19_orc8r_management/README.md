# Magma, Orc8r and Management

## Mental Model

```text
Magma = whole telecom platform

User / Admin
    ↓
   NMS
    ↓
  Orc8r
    ↓
Gateways / network devices
```

## Magma

Magma is the overall telecom network management platform.

## Orc8r

Orc8r is the central controller inside Magma.

It helps manage:
- Configuration
- Policies
- Device state
- Gateway lifecycle

## Northbound

Communication upward toward:
- UI
- NMS
- External systems

## Southbound

Communication downward toward:
- Gateways
- Network devices

## Management Types

Configuration = device settings

Policy = rules applied to users/devices

Lifecycle = create, deploy, monitor, upgrade, retire