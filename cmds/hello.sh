title1="BTCD"
title2="NODE"
title3="TRANSACTION"

cmd1=". ./start_btcd.sh"
cmd2=". ./nodes.sh"
cmd3=". ./transactions.sh"

gnome-terminal --tab --title="$title1" --command="bash -c '$cmd1; $SHELL'" \
               --tab --title="$title2" --command="bash -c '$cmd2; $SHELL'" \
               --tab --title="$title3" --command="bash -c '$cmd3; $SHELL'" 
