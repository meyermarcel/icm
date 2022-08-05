#!/bin/bash

# From https://github.com/paxtonhare/demo-magic
. demo-magic.sh -n
DEMO_PROMPT=" "
PROMPT_TIMEOUT=4

rm -f ~/.icm/config.yml
rm -f ~/.icm/data/owner.csv
rm -f ~/.icm/data/equipment-category-id.json

cp new_owner.csv ~/.icm/data/owner.csv

clear
echo -n "$"
pe "icm generate"
echo -n "$"
wait
pe "icm generate --count 2"
echo -n "$"
wait
pe "icm generate --count 2 --exclude-transposition-errors"
echo -n "$"
wait
clear
echo -n "$"
pe "icm validate btc"
echo -n "$"
wait
pe "icm validate btc u"
echo -n "$"
wait
clear
echo -n "$"
pe "icm validate btc_u123451-0"
echo -n "$"
wait
clear
echo -n "$"
pe "icm validate 20R0"
echo -n "$"
wait
clear
echo -n "$"
pe "icm validate btc_u123451-0 20R0"
echo -n "$"
wait
clear
echo -n "$"
pe "# Equipment category ID 'X' is missing"
echo -n "$"
pe "icm validate btc x 123123 2"
echo -n "$"
wait
clear
echo -n "$"
pe "# Show current equipment categories"
echo -n "$"
pe "cat  ~/.icm/data/equipment-category-id.json"
echo -n "$"
pe "# Add 'X'"
echo -n "$"
pe "diff  ~/.icm/data/equipment-category-id.json  new_equipment-category-id.json"
echo -n "$"
pe "cp  new_equipment-category-id.json  ~/.icm/data/equipment-category-id.json"
echo -n "$"
wait
clear
echo -n "$"
pe "# 'X' is shown"
echo -n "$"
pe "icm validate btc x 123123 2"
echo -n "$"
wait
echo -n "$"

rm -f ~/.icm/config.yml
rm -f ~/.icm/data/owner.csv
rm -f ~/.icm/data/equipment-category-id.json