# imacfancontrol

Never meant to be a real thing. I own a 2009 imac that seems to have been forgotten by other fan control software for linux as it has 3 fans and not 1 or 2 as more recent versions of imac or other apple computers.

I will assume that the labels for the fans are correct:
```
fan1_label ODD 
fan2_label HDD 
fan3_label CPU 
```

Trying to figure out where and which of the sensors are most significant for each fan. Fan2 seems to be the HDD fan, but I don't have any internal HDD running in my imac (as a very cheap repair I disconnected the power to the internal HDD and attached a USB drive to the back of the case). The fan does have an impact for other sensors though. I am also missing data from sensor `TN0D`, originally I assumed this was the sensor in the HDD, but according to the list attached it is `Northbridge Die`.

https://web.archive.org/web/20160525201657/http://jbot-42.github.io:80/Articles/smc.html

```
PECI CPU                   : TCXC
PECI CPU                   : TCXc
CPU 1 Proximity            : TC0P
CPU 1 Heatsink             : TC0H
CPU 1 Package              : TC0D
CPU 1                      : TC0E
CPU 1                      : TC0F
CPU Core 1                 : TC1C
CPU Core 2                 : TC2C
CPU Core 3                 : TC3C
CPU Core 4                 : TC4C
CPU Core 5                 : TC5C
CPU Core 6                 : TC6C
CPU Core 7                 : TC7C
CPU Core 8                 : TC8C
CPU 1 Heatsink Alt.        : TCAH
CPU 1 Package Alt.         : TCAD
CPU 2 Proximity            : TC1P
CPU 2 Heatsink             : TC1H
CPU 2 Package              : TC1D
CPU 2                      : TC1E
CPU 2                      : TC1F
CPU 2 Heatsink Alt.        : TCBH
CPU 2 Package Alt.         : TCBD
PECI SA                    : TCSC
PECI SA                    : TCSc
PECI SA                    : TCSA
PECI GPU                   : TCGC
PECI GPU                   : TCGc
GPU Proximity              : TG0P
GPU Die                    : TG0D
GPU Die                    : TG1D
GPU Heatsink               : TG0H
GPU Heatsink               : TG1H
Memory Proximity           : Ts0S
Mem Bank A1                : TM0P
Mem Bank A2                : TM1P
Mem Bank B1                : TM8P
Mem Bank B2                : TM9P
Mem Module A1              : TM0S
Mem Module A2              : TM1S
Mem Module B1              : TM8S
Mem Module B2              : TM9S
Northbridge Die            : TN0D
Northbridge Proximity 1    : TN0P
Northbridge Proximity 2    : TN1P
MCH Die                    : TN0C
MCH Heatsink               : TN0H
PCH Die                    : TP0D
PCH Die                    : TPCD
PCH Proximity              : TP0P
Airflow 1                  : TA0P
Airflow 2                  : TA1P
Heatpipe 1                 : Th0H
Heatpipe 2                 : Th1H
Heatpipe 3                 : Th2H
Mainboard Proximity        : Tm0P
Powerboard Proximity       : Tp0P
Palm Rest                  : Ts0P
BLC Proximity              : Tb0P
LCD Proximity              : TL0P
Airport Proximity          : TW0P
HDD Bay 1                  : TH0P
HDD Bay 2                  : TH1P
HDD Bay 3                  : TH2P
HDD Bay 4                  : TH3P
Optical Drive              : TO0P
Battery TS_MAX             : TB0T
Battery 1                  : TB1T
Battery 2                  : TB2T
Battery                    : TB3T
Power Supply 1             : Tp0P
Power Supply 1 Alt.        : Tp0C
Power Supply 2             : Tp1P
Power Supply 2 Alt.        : Tp1C
Power Supply 3             : Tp2P
Power Supply 4             : Tp3P
Power Supply 5             : Tp4P
Power Supply 6             : Tp5P
Expansion Slots            : TS0C
PCI Slot 1 Pos 1           : TA0S
PCI Slot 1 Pos 2           : TA1S
PCI Slot 2 Pos 1           : TA2S
PCI Slot 2 Pos 2           : TA3S
CPU Core 1                 : VC0C
CPU Core 2                 : VC1C
CPU Core 3                 : VC2C
CPU Core 4                 : VC3C
CPU Core 5                 : VC4C
CPU Core 6                 : VC5C
CPU Core 7                 : VC6C
CPU Core 8                 : VC7C
CPU VTT                    : VV1R
GPU Core                   : VG0C
Memory                     : VM0R
PCH                        : VN1R
MCH                        : VN0C
Mainboard S0 Rail          : VD0R
Mainboard S5 Rail          : VD5R
12V Rail                   : VP0R
12V Vcc                    : Vp0C
Main 3V                    : VV2S
Main 3.3V                  : VR3R
Main 5V                    : VV1S
Main 5V                    : VH05
Main 12V                   : VV9S
Main 12V                   : VD2R
Auxiliary 3V               : VV7S
Standby 3V                 : VV3S
Standby 5V                 : VV8S
PCIe 12V                   : VeES
Battery                    : VBAT
CMOS Battery               : Vb0R
CPU Core                   : IC0C
CPU VccIO                  : IC1C
CPU VccSA                  : IC2C
CPU Rail                   : IC0R
CPU DRAM                   : IC5R
CPU PLL                    : IC8R
CPU GFX                    : IC0G
CPU Memory                 : IC0M
GPU Rail                   : IG0C
Memory Controller          : IM0C
Memory Rail                : IM0R
MCH                        : IN0C
Mainboard S0 Rail          : ID0R
Mainboard S5 Rail          : ID5R
Misc. Rail                 : IO0R
Battery Rail               : IB0R
Charger BMON               : IPBR
CPU Core 1                 : PC0C
CPU Core 2                 : PC1C
CPU Core 3                 : PC2C
CPU Core 4                 : PC3C
CPU Core 5                 : PC4C
CPU Core 6                 : PC5C
CPU Core 7                 : PC6C
CPU Core 8                 : PC7C
CPU Cores                  : PCPC
CPU GFX                    : PCPG
CPU DRAM                   : PCPD
CPU Total                  : PCTR
CPU Total                  : PCPL
CPU Rail                   : PC1R
CPU S0 Rail                : PC5R
GPU Total                  : PGTR
GPU Rail                   : PG0R
Memory Rail                : PM0R
MCH                        : PN0C
PCH Rail                   : PN1R
Mainboard S0 Rail          : PC0R
Mainboard S0 Rail          : PD0R
Mainboard S5 Rail          : PD5R
Main 3.3V Rail             : PH02
Main 5V Rail               : PH05
12V Rail                   : Pp0R
Main 12V Rail              : PD2R
Misc. Rail                 : PO0R
Battery Rail               : PBLC
Battery Rail               : PB0R
DC In Total                : PDTR
System Total               : PSTR
```