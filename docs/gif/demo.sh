#!/bin/bash

# From https://github.com/paxtonhare/demo-magic
. demo-magic.sh -n
DEMO_PROMPT="$ "
PROMPT_TIMEOUT=3

rm -f ~/.icm/config.yml
rm -f ~/.icm/data/owner.json
rm -f ~/.icm/data/equipment-category-id.json

cp new_owner.json ~/.icm/data/owner.json

clear
pe "icm validate btc u_123123-3"
echo -n "$DEMO_PROMPT"
wait
clear
pe "icm validate btc"
echo -n "$DEMO_PROMPT"
wait
clear
pe "icm validate 20R0"
echo -n "$DEMO_PROMPT"
wait
clear
pe "icm validate btc u_123123-3 22r0"
echo -n "$DEMO_PROMPT"
wait
clear
pe "icm generate -c 4"
echo -n "$DEMO_PROMPT"
wait
clear
pe "icm validate btc x 1231232"
echo -n "$DEMO_PROMPT"
wait
clear
pe "# Add new equipment category ID in data"
pe "cat ~/.icm/data/equipment-category-id.json"
pe "diff ~/.icm/data/equipment-category-id.json new_equipment-category-id.json"
pe "cp new_equipment-category-id.json ~/.icm/data/equipment-category-id.json"
echo -n "$DEMO_PROMPT"
wait
clear
pe "icm validate btc x 1231232"
echo -n "$DEMO_PROMPT"
wait

rm -f ~/.icm/config.yml
rm -f ~/.icm/data/owner.json
rm -f ~/.icm/data/equipment-category-id.json