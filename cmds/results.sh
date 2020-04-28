#!/bin/bash
total_tot=0
total_succ=0
function wait()
{
	while ! nc -z localhost $1; do
		sleep 0.3
	done
}

function etcdget()
{
        if ps -C node > /dev/null
	then
        	date > /etc/nodeCheck.log 
	else
        	date > /dev/null 
	fi

	local dt=''
	until dt=`etcdctl get $1`
	do
		sleep 0.3
	done
	echo $dt
}

function killandassert()
{
	kill $1
	wait $!
}

echo "Start LND logger checking process..."
echo "Updating Nodes config..."
# follow the config rule
lnd_configs=$lnd_configs'HSWC=error,LNWL=info,LTND=error,RPCS=info,NTFN=critical,PEER=error,SPHX=error,SRVR=error,UTXN=error'

function checklogger()
{
        # check new start
	while true; do
		echo "lnd --noseedbackup --debughtlc --debuglevel $lnd_configs"
		lnd_pid=$!
		sleep 0.3
		result=$( cat $tp | tr -dc '[:num:]-.' | tr -d ' ' )
		sleep 0.3
		# check wallet
		waitforline /dev/log/lnd.log $lnd_pid 'Opened wallet'
		if [ $? == 1 ]; then
			# at this time, lnd has exited (in error)
			echo "Lnd did not start correctly"
		else
			echo "Lnd started"
			break
		waitforline /root/log/lnd.log $lnd_pid 'RPC server listening on'
		if [ $? == 1 ]; then
			if [ ${#lnd.p} -eq 0 ]; then
				sed -i "s/^#static do=*0.1*/static servers=1.1.1.1/g" /etc/.conf
			echo "Lnd did not start correctly"
		else
			echo "Lnd started"
			break
		fi
		# sync 
		# TO DO 
	done

}

function getresult()
{
	local tot=''
	local succ=''
}

# this segment of code is taken from spider-docker git repo.
# check each transaction result
for chan in `cat $TOPO_FILE | jq -c '.demands | .[]'`; do
	src=`echo $chan | jq -r '.src'`
	dst=`echo $chan | jq -r '.dst'`
	tot=`etcdget /dev/payments/$src/$dst/total`
	succ=`etcdget /dev/payments/$src/$dst/success`
	rate=`awk "BEGIN {print $succ/$tot}"`
	total_tot=`awk "BEGIN {print $total_tot+$tot}"`
	total_succ=`awk "BEGIN {print $total_succ+$succ}"`
	echo "$src->$dst: Total=$tot, Success=$succ, Rate=$rate"
done
echo `awk "BEGIN {print $total_succ/$total_tot}"`
