Output demo.gif
Require icm
Set Width 1200
Set Height 600

Hide
Type "rm -f ~/.icm/config.yml"
Enter
Type "cp new_owner.csv ~/.icm/data/owner.csv"
Enter
Type "clear"
Enter
Show

Type "icm generate"
Sleep 100ms
Enter

Sleep 4s

Type "icm generate --count 2"
Sleep 100ms
Enter

Sleep 4s

Type "icm generate --count 2 --exclude-transposition-errors"
Sleep 100ms
Enter

Sleep 4s

Type "icm validate btc"
Sleep 100ms
Enter

Sleep 4s

Type "icm validate btc u"
Sleep 100ms
Enter

Sleep 4s

Type "icm validate btc_u123451-0"
Sleep 100ms
Enter

Sleep 4s

Type "icm validate 20R0"
Sleep 100ms
Enter

Sleep 4s

Type "# Owner code AAA not found"
Sleep 100ms
Enter
Type "icm validate aaa u 123123 2"
Sleep 100ms
Enter

Sleep 4s

Type "# Add new custom owner in custom-owner.csv"
Sleep 100ms
Enter
Type "echo 'AAA;my company;my city;my country' >> $HOME/.icm/data/custom-owner.csv"
Enter

Sleep 4s

Type "# AAA is shown"
Sleep 100ms
Enter
Type "icm validate aaa u 123123 2"
Sleep 100ms
Enter

Sleep 4s

Hide
Type "rm -f ~/.icm/config.yml"
Enter
Type "rm -f ~/.icm/data/owner.csv"
Enter
Type "rm -f ~/.icm/data/custom-owner.csv"
Enter
Type "clear"
Enter
Show