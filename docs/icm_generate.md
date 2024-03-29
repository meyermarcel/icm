## icm generate

Generate unique container numbers

### Synopsis

Generated container numbers are unique. Owners specified in

  $HOME/.icm/data/owner.csv

are used. Owners can be updated by 'icm download-owners --help' command.

Equipment category ID 'U' is used for every generated container number.

For a custom owner code use the --owner-code flag.

For a custom serial number use the --start and --end flags and optionally the --count flag.
Using only the --count flag generates pseudo random serial numbers.

Configuration for separators is generated first time you
execute a command that requires the configuration.

Flags for output formatting can be overridden with a config file.
Edit default configuration for customization:

  $HOME/.icm/config.yml

```
icm generate [flags]
```

### Examples

```
icm generate
icm generate --count 10
# Generate container numbers with custom format
icm generate --count 10 --sep-owner-equip '' --sep-serial-check '-'
# Generate container numbers without error-prone serial numbers
icm generate --count 10 --exclude-check-digit-10 --exclude-error-prone-serial-numbers
# Generate container numbers within serial number range
icm generate --start 100500 --count 10
icm generate --start 100500 --end 100600
icm generate --start 100500 --end 100600 --owner ABC
# Generate CSV data set
icm generate --count 1000000 | icm validate
```

### Options

```
  -c, --count int                            count of container numbers (default 1)
  -s, --start int                            start of serial number range
  -e, --end int                              end of serial number range
      --owner string                         custom owner code
      --exclude-check-digit-10               exclude check digit 10
      --exclude-error-prone-serial-numbers   exclude error-prone serial numbers. For example swapping the second 0 and first 1 of RCB U 001130 0 results in container number RCB U 010130 0 with a valid check digit 0
      --sep-owner-equip string               ABC(x)U1234560  (x) separates owner code and equipment category id (default " ")
      --sep-equip-serial string              ABCU(x)1234560  (x) separates equipment category id and serial number (default " ")
      --sep-serial-check string              ABCU123456(x)0  (x) separates serial number and check digit (default " ")
  -h, --help                                 help for generate
```

### SEE ALSO

* [icm](icm.md)	 - Validate or generate intermodal container markings

