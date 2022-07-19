# Camera Command Interface Documentation

## Table of Contents
- [Camera Command Interface Documentation](#camera-command-interface-documentation)
  - [Table of Contents](#table-of-contents)
  - [Initialization Functions](#initialization-functions)
  - [System Functions (`common.CameraFunctions`)](#system-functions-commoncamerafunctions)
    - [Camera Number (`common.CameraNumber`)](#camera-number-commoncameranumber)
    - [Call Signal (`common.CameraSignal`)](#call-signal-commoncamerasignal)
    - [Color Bar (`common.ColorBar`)](#color-bar-commoncolorbar)
  - [Gain Functions](#gain-functions)
    - [Master Gain (`common.GainMaster`)](#master-gain-commongainmaster)
    - [Red Gain (`common.GainRed`)](#red-gain-commongainred)
    - [Green Gain (`common.GainGreen`)](#green-gain-commongaingreen)
    - [Blue Gain (`common.GainBlue`)](#blue-gain-commongainblue)
  - [Black Functions](#black-functions)
    - [Master Black (`common.BlackMaster`)](#master-black-commonblackmaster)
    - [`GetBlackR() int` / `SetBlackR(int)`](#getblackr-int--setblackrint)
    - [`GetBlackG() int` / `SetBlackG(int)`](#getblackg-int--setblackgint)
    - [`GetBlackB() int` / `SetBlackB(int)`](#getblackb-int--setblackbint)
  - [Flare Functions](#flare-functions)
    - [`GetFlareR() int` / `SetFlareR(int)`](#getflarer-int--setflarerint)
    - [`GetFlareG() int` / `SetFlareG(int)`](#getflareg-int--setflaregint)
    - [`GetFlareB() int` / `SetFlareB(int)`](#getflareb-int--setflarebint)

## Initialization Functions

## System Functions (`common.CameraFunctions`)

### Camera Number (`common.CameraNumber`)
Represents the user-defined camera number in a range of 1-99. The camera number is set by the OCPArbiter application after camera initialization and should only be stored as a private variable with these as accessor functions, unless the camera has some kind of display for what production camera number it is. In that case, `SetCameraNumber(int)` should infrom the camera of that number.

| Value Info     | Description   |
|----------------|---------------|
| Type           | Integer       |
| Range          | 1 - 99        |
| Representation | Camera number |


### Call Signal (`common.CameraSignal`)
Enacts the call button from the RCP to inform the camera.

| Value Info     | Description |
|----------------|-------------|
| Type           | Bool        |
| Range          | 0 - 1       |
| Representation | Off / On    |

| Value | Description |
|-------|-------------|
| 0     | Call off    |
| 1     | Call on     |

### Color Bar (`common.ColorBar`)
Function to place the camera into bars mode.

| Value Info     | Description |
|----------------|-------------|
| Type           | Bool        |
| Range          | 0 - 1       |
| Representation | Off / On    |

| Value | Description |
|-------|-------------|
| 0     | Bars off    |
| 1     | Bars on     |

## Gain Functions

### Master Gain (`common.GainMaster`)
Adjusts the master gain of the camera. Value is 0 - 4095 (12-bit). Value represents a dB value of `0.1 * value`. View the second table below for reference.

| Value Info     | Description      |
|----------------|------------------|
| Type           | Float (int / 10) |
| Range          | 0 - 4095         |
| Representation | Camera gain      |

Examples of values:
| Value | dB Value |
|-------|----------|
| 0     | 0.0 dB   |
| 13    | 1.3 dB   |
| 105   | 10.5 dB  |

### Red Gain (`common.GainRed`)
### Green Gain (`common.GainGreen`)
### Blue Gain (`common.GainBlue`)
The above three featured commands are all effectively the same and are used for setting individual RGB gain values on the sensor. Values are 0 - 255 (8-bit) representing numbers (0 - 99) on the OCP.

| Value Info     | Description       |
|----------------|-------------------|
| Type           | Integer           |
| Range          | 0 - 255           |
| Representation | Camera color gain |

## Black Functions

### Master Black (`common.BlackMaster`)
Adjusts the master black of the camera. Value is 0 - 4095 (12-bit). 12-bit value represents a value of 0 - 99 on the OCP.

| Value Info     | Description        |
|----------------|--------------------|
| Type           | Integer            |
| Range          | 0 - 4095           |
| Representation | Camera black level |

### `GetBlackR() int` / `SetBlackR(int)`
### `GetBlackG() int` / `SetBlackG(int)`
### `GetBlackB() int` / `SetBlackB(int)`
The above three featured commands are all effectively the same and are used for setting individual RGB black levels. Values are 0 - 255 (8-bit) representing numbers (0 - 99) on the OCP.

## Flare Functions

### `GetFlareR() int` / `SetFlareR(int)`
### `GetFlareG() int` / `SetFlareG(int)`
### `GetFlareB() int` / `SetFlareB(int)`
The above three featured commands are all effectively the same and are used for setting individual RGB flare levels. Values are 0 - 255 (8-bit) representing numbers (0 - 99) on the OCP.