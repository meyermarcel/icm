## icm completion bash

generate the autocompletion script for bash

### Synopsis


Generate the autocompletion script for the bash shell.

This script depends on the 'bash-completion' package.
If it is not installed already, you can install it via your OS's package manager.

To load completions in your current shell session:
$ source <(icm completion bash)

To load completions for every new session, execute once:
Linux:
  $ icm completion bash > /etc/bash_completion.d/icm
MacOS:
  $ icm completion bash > /usr/local/etc/bash_completion.d/icm

You will need to start a new shell for this setup to take effect.
  

```
icm completion bash
```

### Options

```
  -h, --help              help for bash
      --no-descriptions   disable completion descriptions
```

### SEE ALSO

* [icm completion](icm_completion.md)	 - generate the autocompletion script for the specified shell

