# Camera Interface Functional Specifications

## Initialization Functions

### `Initialize()`
Basic function for doing camera initialization. Should be called before any other command. Effectively a constructor in other languages

### `GetFeatureSet() CamFeatureSet`
Returns a `CamFeatureSet` object representing the control features that this camera supports

## System Functions

### `GetCameraNumber() int` / `SetCameraNumber(int)`
Represents the user-defined camera number in a range of 1-99. The camera number is set by the OCPArbiter application after camera initialization and should only be stored as a private variable with these as accessor functions, unless the camera has some kind of display for what production camera number it is. In that case, `SetCameraNumber(int)` should infrom the camera of that number.

### `GetCallSig() int` / `SetCallSig(int)`
Enacts the call button from the RCP to inform the camera.
|  Value        | Description |
| -----------   | ----------- |
| 0             | Call off    |
| 1             | Call on     |

### `GetColorBar() int` / `SetColorBar(int)`
Function to place the camera into bars mode.
|  Value        | Description |
| -----------   | ----------- |
| 0             | Bars off    |
| 1             | Bars on     |

## Gain Functions

### `GetGainMaster() int` / `SetGainMaster(int)`
Adjusts the master gain of the camera. Value is 0 - 4095 (12-bit). Value represents a dB value of `0.1 * value`. View the table below for reference.
|  Value        | dB Value    |
| -----------   | ----------- |
| 0             | 0.0 dB      |
| 13            | 1.3 dB      |
| 105           | 10.5 dB     |

### `GetGainR() int` / `SetGainR(int)`
### `GetGainG() int` / `SetGainG(int)`
### `GetGainB() int` / `SetGainB(int)`
The above three featured commands are all effectively the same and are used for setting individual RGB gain values on the sensor. Values are 0 - 255 (8-bit) representing numbers (0 - 99) on the OCP.

## Black Functions

### `GetBlackMaster() int` / `SetBlackMaster(int)`
Adjusts the master black of the camera. Value is 0 - 4095 (12-bit). 12-bit value represents a value of 0 - 99 on the OCP.
NOTE: FURTHER RESEARCH AND TESTING REQUIRED

### `GetBlackR() int` / `SetBlackR(int)`
### `GetBlackG() int` / `SetBlackG(int)`
### `GetBlackB() int` / `SetBlackB(int)`
The above three featured commands are all effectively the same and are used for setting individual RGB black levels. Values are 0 - 255 (8-bit) representing numbers (0 - 99) on the OCP.

## Flare Functions

### `GetFlareR() int` / `SetFlareR(int)`
### `GetFlareG() int` / `SetFlareG(int)`
### `GetFlareB() int` / `SetFlareB(int)`
The above three featured commands are all effectively the same and are used for setting individual RGB flare levels. Values are 0 - 255 (8-bit) representing numbers (0 - 99) on the OCP.