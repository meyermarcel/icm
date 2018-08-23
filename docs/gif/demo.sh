#!/bin/bash

# From https://github.com/paxtonhare/demo-magic
. demo-magic.sh -n
DEMO_PROMPT="$ "
PROMPT_TIMEOUT=2

rm -f ~/.icm/config.yml
rm -f ~/.icm/data/owners.json
rm -f ~/.icm/data/equipment-category-ids.json

cp new_owners.json ~/.icm/data/owners.json

clear
pe "icm validate ' btc u_123'"
pe "icm validate ' btc u_123123-3'"
echo -n "$DEMO_PROMPT"
wait
clear
pe "icm owner validate ' btc'"
echo -n "$DEMO_PROMPT"
wait
clear
pe "icm sizetype validate '20N1'"
echo -n "$DEMO_PROMPT"
wait
clear
pe "icm validate 'btc u_123123-3 22g1'"
echo -n "$DEMO_PROMPT"
wait
clear
pe "icm generate -c 4"
echo -n "$DEMO_PROMPT"
wait
clear
pe "# Config printed output"
pe "cat ~/.icm/config.yml"
echo -n "$DEMO_PROMPT"
wait
clear
pe "git --no-pager diff ~/.icm/config.yml new_config.yml"
pe "cp new_config.yml ~/.icm/config.yml"
echo -n "$DEMO_PROMPT"
wait
clear
pe "icm validate 'btc u_123123-3 22g1'"
pe "icm generate -c 4"
echo -n "$DEMO_PROMPT"
wait
clear
pe "# Add custom equipment category id"
pe "icm validate 'btc x 1231232'"
echo -n "$DEMO_PROMPT"
wait
clear
pe "cat ~/.icm/data/equipment-category-ids.json"
pe "git --no-pager diff ~/.icm/data/equipment-category-ids.json new_equipment-category-ids.json"
pe "cp new_equipment-category-ids.json ~/.icm/data/equipment-category-ids.json"
echo -n "$DEMO_PROMPT"
wait
clear
pe "icm validate 'btc x 1231232'"
echo -n "$DEMO_PROMPT"
wait

rm -f ~/.icm/config.yml
rm -f ~/.icm/data/owners.json
rm -f ~/.icm/data/equipment-category-ids.json