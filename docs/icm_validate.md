## icm validate

Validate intermodal container markings

### Synopsis

Validate intermodal container markings with single or multi line input.

For single line input a human-readable output is used.

For multi line input CSV output is used. For example this is useful to scan
data sets for error-prone serial numbers. It is also possible to generate
CSV data sets of random container numbers.

Configuration for separators is generated first time you
execute a command that requires the configuration.

Flags for output formatting can be overridden with a config file.
Edit default configuration for customization:

  $HOME/.icm/config.yml

```
icm validate [flags]
```

### Examples

```
icm validate ABC
# Validate with pattern 'container-number' instead of pattern 'auto'
icm validate ABC --pattern container-number
icm validate ABC U
# Validate and use custom format for output
icm validate --sep-owner-equip '' --sep-serial-check '-' ABC U 123456 0
# Validate a type
icm validate 20G1
# Validate a container number with a type
icm validate ABC U 123456 0 20G1
# Validate a random container number
icm generate | icm validate
icm generate --count 10 | icm validate
icm generate --count 10 | icm validate --output fancy
# Generate CSV data set
icm generate --count 1000000 | icm validate
# Validate a container number with 6 (!) error-prone serial numbers combinations
icm validate APL U 689473 0
```

### Options

```
  -p, --pattern string            sets pattern matching to auto, container-number, owner, owner-equipment-category or size-type
                                                      auto = matches automatically a pattern
                                          container-number = matches a container number
                                                     owner = matches a three letter owner code
                                  owner-equipment-category = matches a three letter owner code with equipment category ID
                                                 size-type = matches length, width+height and type code
                                  
      --output string             sets output to auto, fancy or csv
                                   auto = for a single line 'fancy' and for multiple lines 'csv' output 
                                    csv = machine readable CSV output
                                  fancy = human readable fancy output
                                  
      --no-header                 omits header of CSV output
      --sep-owner-equip string    ABC(x)U1234560   20G1  (x) separates owner code and equipment category id (default " ")
      --sep-equip-serial string   ABCU(x)1234560   20G1  (x) separates equipment category id and serial number (default " ")
      --sep-serial-check string   ABCU123456(x)0   20G1  (x) separates serial number and check digit (default " ")
      --sep-check-size string     ABCU1234560 (x)  20G1  (x) separates check digit and size (default "   ")
      --sep-size-type string      ABCU1234560   20(x)G1  (x) separates size and type (default " ")
  -h, --help                      help for validate
```

### SEE ALSO

* [icm](icm.md)	 - Validate or generate intermodal container markings

