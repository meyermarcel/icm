## icm completion powershell

generate the autocompletion script for powershell

### Synopsis


Generate the autocompletion script for powershell.

To load completions in your current shell session:
PS C:\> icm completion powershell | Out-String | Invoke-Expression

To load completions for every new session, add the output of the above command
to your powershell profile.


```
icm completion powershell [flags]
```

### Options

```
  -h, --help              help for powershell
      --no-descriptions   disable completion descriptions
```

### SEE ALSO

* [icm completion](icm_completion.md)	 - generate the autocompletion script for the specified shell

