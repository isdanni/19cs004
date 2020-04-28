# Create a OK sign
cd $GOPATH
cd dev

echo "Starting Transactions..."

cmd1="cd s1 && lnd --rpclisten=localhost:10001 --listen=localhost:10011 --restlisten=localhost:8001 >> log.txt"
cmd2="cd s2 && lnd --rpclisten=localhost:10002 --listen=localhost:10012 --restlisten=localhost:8002 >> log.txt"
cmd3="cd s3 && lnd --rpclisten=localhost:10003 --listen=localhost:10013 --restlisten=localhost:8003 >> log.txt"

gnome-terminal --tab --title="$t1" --command="bash -c '$cmd1; $SHELL'" \
               --tab --title="$t2" --command="bash -c '$cmd2; $SHELL'" \
               --tab --title="$t3" --command="bash -c '$cmd3; $SHELL'" 

# for d in $GOPATH/dev/*/
# do
#     (cd "$d" && (gnome-terminal --tab --title="trans" --command="bash -c '$cmd; $SHELL'"))
# done
#
# echo "Transaction done!"
