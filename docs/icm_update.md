## icm update

Update information of owners

### Synopsis

Update information of owners from remote.
Following information is available:

  Owner code
  Company
  City
  Country

```
icm update [flags]
```

### Examples

```
# Add new owners and preserve all existing owners
icm update
# Delete all owners and add most current owners
echo '{}' > $HOME/.icm/data/owner.json && icm update
```

### Options

```
  -h, --help   help for update
```

### SEE ALSO

* [icm](icm.md)	 - Validate or generate intermodal container markings

