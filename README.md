# wpw-dps-parser

Due to the limitations in the wowclient regarding showing correct dps for augmentation evokers i decided to write my 
own parsere to get a correct view.

## Description

The augmentation evoker is a support class that buffs other characters in party or raid and is calculated differently then other dps classes
who does direct damage to mobs.

The example below shows a player doing damage using  **Wake of Ashes ** to a mob called ** Spellbound Battleaxe **
```
5/14 15:48:06.420  SPELL_DAMAGE,Player-3691-0A322132,"Ciaray-Blackhand",0x512,0x0,Creature-0-4238-2526-22542-196577-0002436942,"Spellbound Battleaxe",0xa48,0x0,405350,"Wake of Ashes",0x6,Creature-0-4238-2526-22542-196577-0002436942,0000000000000000,4398441,6490604,0,0,5043,0,1,0,0,0,1993.80,-2897.95,2097,0.0000,70,97807,47710,-1,6,0,0,0,1,nil,nil
```

This example is a augmentation evoker doing damage with buffs like ** Prescience ** and ** Ebon Might ** 
```
5/14 15:48:06.618  SPELL_DAMAGE_SUPPORT,Player-3691-0A322132,"Ciaray-Blackhand",0x512,0x0,Creature-0-4238-2526-22542-196694-0003C36942,"Arcane Forager",0xa48,0x0,395152,"Ebon Might",0xc,Creature-0-4238-2526-22542-196694-0003C36942,0000000000000000,3042047,4720439,0,0,5043,0,1,0,0,0,1997.54,-2897.81,2097,5.7621,70,1677,779,-1,3,0,0,0,1,nil,nil,Player-1305-0C594315
5/14 15:48:06.618  SPELL_DAMAGE_SUPPORT,Player-3691-0A322132,"Ciaray-Blackhand",0x512,0x0,Creature-0-4238-2526-22542-196694-0003436942,"Arcane Forager",0xa48,0x0,410089,"Prescience",0x40,Creature-0-4238-2526-22542-196694-0003436942,0000000000000000,2707436,4720439,0,0,5043,0,1,0,0,0,1997.57,-2897.79,2097,5.8251,70,708,329,-1,3,0,0,0,1,nil,nil,Player-1305-0C594315
```
## Included example
the combat log that is included for testing is from the first boss in a heroic version of Neltharus
![World Of Warcraft Screenshot 2024 05 18 - 22 22 37 21-sdr](https://github.com/pefman/wow-dps-parser/assets/5830988/208bec5a-bb76-4eed-858f-4bee78d5db53)



## Getting Started

1. download latest binary from releases for your platform win/linux/mac
2. update the config.conf with a path to your logfile and character name
3. execute the binary and show the output

### good to know

# enable combatlogging
this is enabled options > system > network > advanced combat logging

# easy combatlogging
For easy logging i recommend using this weakaura that automatically enables combat logging when zoning into dungeons and raids.
[Auto Combat Logging]([https://www.example.com](https://wago.io/-Np8BNw7n))

## Help

Any advise for common problems or issues.
```
wow-pumper --help
```

## Version History

* 0.2
    * Various bug fixes and optimizations
    * See [commit change]() or See [release history]()
* 0.1
    * Initial Release
