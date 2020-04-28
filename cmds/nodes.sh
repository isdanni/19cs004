# Create development space
cd $GOPATH
mkdir -p dev
cd dev

# Create folders for each of our nodes
echo "Starting Node setup..."
n=1
max=50
set -- # this sets $@ [the argv array] to an empty list.
while [ "$n" -le "$max" ]; do
    set -- "$@" "s$n" # this adds s$n to the end of $@
    echo "next..."
    n=$(( $n + 1 ));
done
mkdir "$@"

echo "Node Folder Created!"

echo "Waiting..."

# Wait 5 seconds
sleep 2

cd $GOPATH
cd dev

for d in $GOPATH/dev/*/
do
     (cd "$d" && $(touch log.txt))
done

echo "Node setup done!"

