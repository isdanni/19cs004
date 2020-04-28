#!/bin/bash
# local version
function start_instances
{
	# $1: number of instances to start
	echo "Launching $1 AWS EC2 instances"
	aws ec2 run-instances --launch-template LaunchTemplateId=lt-0065a09a461515b3e --count $1 > aws_start.log
	local instances=`jq -r '.Instances[].InstanceId ' aws_start.log`
	echo "Waiting for network interfaces to attach"
	sleep 3
	rm -f instances.txt
	rm -f ~/.ssh/config.d/spider
	echo "Querying public IPs and writing to SSH config"
	for instance in $instances ;
	do
		local ip=`aws ec2 describe-instances --instance-ids $instance | jq -r '.Reservations[0].Instances[0].PublicIpAddress'`
		echo "$instance,$ip" >> instances.txt
		echo "Host $instance" >> ~/.ssh/config.d/spider
		echo "    Hostname $ip" >> ~/.ssh/config.d/spider
		echo "    User ubuntu" >> ~/.ssh/config.d/spider
		echo "    IdentityFile ~/.ssh/spider.pem" >> ~/.ssh/config.d/spider
		echo "    StrictHostKeyChecking no" >> ~/.ssh/config.d/spider
		echo "    UserKnownHostsFile=/dev/null" >> ~/.ssh/config.d/spider
		echo "" >> ~/.ssh/config.d/spider
	done
}

function stop_instances
{
	local instances=`cat instances.txt`
	local instance_ids=""
	for instance in $instances ;
	do
		local id
		local ip
		IFS=',' read -r id ip <<< "$instance"
		instance_ids="$instance_ids $id"
	done
	echo "Terminating instances $instance_ids"
	aws ec2 terminate-instances --instance-ids $instance_ids > aws_stop.log
}

function init_swarm
{
	local instances=`cat instances.txt`
	local instance_ids=""
	for instance in $instances ;
	do
		local id
		local ip
		IFS=',' read -r id ip <<< "$instance"
		instance_ids="$instance_ids $id"
	done
	local first_id
	local rest_ids
	read first_id rest_ids <<< "$instance_ids"
	local cmd_to_use=`ssh $first_id -- 'docker swarm init | sed -n 5p'`
	for instance in $rest_ids ;
	do
		ssh $instance -- "$cmd_to_use"
	done
	ssh $first_id -- 'docker network create -d overlay --subnet 10.0.0.0/16 --attachable spider'
}

function destroy_swarm
{
	local instances=`cat instances.txt`
	local instance_ids=""
	for instance in $instances ;
	do
		local id
		local ip
		IFS=',' read -r id ip <<< "$instance"
		instance_ids="$instance_ids $id"
	done
	local first_id
	read first_id _ <<< "$instance_ids"
	ssh $first_id -- 'docker network rm spider'
	for instance in $instance_ids ;
	do
		ssh $instance -- 'docker swarm leave --force'
	done
}

function start_container
{
	# $1: node name, $2: ip, $3: host
	topo_filename=`basename $TOPO_FILE`
	topo_path="/root/topology/$topo_filename"

	ssh $3 -- docker run --mount source="myvolume$1",target=/root/.lnd --cap-add=NET_ADMIN -itd --name "spider$1" -e NODENAME=$1 -e NODEIP=$2 -e SPIDER_EXP_NAME="$EXP_NAME" -e TOPO_FILE="$topo_path" -e EXP_TIME="$EXP_TIME" --network spider --ip $2 spider
	#ssh $3 -- docker run --mount type=tmpfs,destination=/root/.lnd --cap-add=NET_ADMIN -itd --name "spider$1" -e NODENAME=$1 -e NODEIP=$2 -e SPIDER_EXP_NAME="$EXP_NAME" -e TOPO_FILE="$topo_path" -e EXP_TIME="$EXP_TIME" --network spider --ip $2 spider
}

function destroy_container
{
	# $1: node name, $2: host
	ssh $2 -- docker container stop "spider$1"
	ssh $2 -- docker container rm "spider$1"
	ssh $2 -- docker volume rm "myvolume$1"
}

function next_index
{
	# $1: current index
	local len=${#hosts[@]}
	local next=`expr $1 + 1`
	if [ "$next" -ge "$len" ]
	then
		next=0
	fi
	echo $next
}

function start_experiment
{
	local instances=`cat instances.txt`
	hosts=()
	for instance in $instances ;
	do
		local id
		local ip
		IFS=',' read -r id ip <<< "$instance"
		hosts+=("$id")
	done
	local host_idx=0
	local pids=""
	rm -f nodehostmap.txt
	for node in `cat $TOPO_FILE | jq -rc '.nodes | .[]'`
	do
		name=`echo $node | jq -r '.name'`
		ip=`echo $node | jq -r '.ip'`
		echo "Starting $name"
		start_container $name $ip ${hosts[$host_idx]} &> /dev/null &
		echo "$name,${hosts[$host_idx]}" >> nodehostmap.txt
		pids="$pids $!"
		host_idx=`next_index $host_idx`
	done
	echo "Waiting for all jobs to finish"
	for pid in $pids ;
	do
		wait $pid
	done
}

function stop_experiment
{
	local expconfig=`cat nodehostmap.txt`
	local pids=""
	for nodeinfo in $expconfig; do
		local node
		local host
		IFS=',' read -r node host <<< "$nodeinfo"
		echo "Stopping $node"
		destroy_container $node $host &> /dev/null &
		pids="$pids $!"
	done
	echo "Waiting for all jobs to finish"
	for pid in $pids ;
	do
		wait $pid
	done
}

function run_on_all
{
	# $@: command to run
	local instances=`cat instances.txt`
	local pids=""
	for instance in $instances ;
	do
		local id
		local ip
		IFS=',' read -r id ip <<< "$instance"
		echo "Job launched for $id"
		ssh $id -- "$@" &
		pids="$pids $!"
	done
	echo "Waiting for all jobs to finish"
	for pid in $pids ;
	do
		wait $pid
	done
}

function sync_testbed_single
{
	# $1: host id
	rsync -r .. $id:/home/ubuntu/spider-docker
}

function repackage_single
{
	# $1: host id
	ssh $id -- /bin/bash /home/ubuntu/spider-docker/tools/remote_helper.sh build_image
}

function rebuild_single
{
	# $1: host id
	# ${@:2}: which bins to build
	ssh $id -- /bin/bash /home/ubuntu/spider-docker/tools/remote_helper.sh build_bin "${@:2}"
}

function init_image_single
{
	# $1: host id
	rsync -r .. $id:/home/ubuntu/spider-docker
	ssh $id -- '/bin/bash /home/ubuntu/spider-docker/tools/remote_helper.sh download_bin && /bin/bash /home/ubuntu/spider-docker/tools/remote_helper.sh build_image'
}

function download_single
{
	# $1: host id
	rsync -r .. $id:/home/ubuntu/spider-docker
	ssh $id -- /bin/bash /home/ubuntu/spider-docker/tools/remote_helper.sh download_bin
}

function copy_binaries_single
{
        mkdir -p log
	local instances=`cat instances.txt`
	local pids=""
        local id
        local ip
        IFS=',' read -r id ip <<< "${instances[0]}"
        echo "Copying binaries from $id"
        rsync -r $id:/home/ubuntu/spider-docker/binaries .
        tar -cvzf binaries.tar.gz binaries/
}

function copy_log_files_single
{
	# $1: host id, $2: EXP_NAME directory.
	mkdir -p $2
	scp -r $1:/home/ubuntu/container_logs/* ./$2
}

function delete_old_log_single
{
	ssh $1 -- rm -rf /home/ubuntu/container_logs
}

function mount_nvme_single
{
	ssh $1 -- 'diskname=$(lsblk | grep 372 | cut -f 1 -d " ") && sudo rm -rf /var/lib/docker/volumes && sudo mkdir -m 777 /var/lib/docker/volumes && sudo mkfs -F -t ext4 /dev/$diskname && sudo mount /dev/$diskname /var/lib/docker/volumes && sudo chmod 777 /var/lib/docker/volumes'
}


function execute_on_all
{
	# $1: execute function '$1_single'
	# ${@:2}: extra params of the function
	mkdir -p log
	local instances=`cat instances.txt`
	local pids=""
	for instance in $instances ;
	do
		local id
		local ip
		IFS=',' read -r id ip <<< "$instance"
		echo "Executing $1 on $id"
		$1_single $id ${@:2} &>log/$1_$id.log &
		pids="$pids $!"
	done
	echo "Waiting for all jobs to finish"
	for pid in $pids ;
	do
		wait $pid
	done
}

function ssh_to_server
{
	# $1: which server to ssh to (starting from 1)
	local instance=`sed -n "$1 p" < instances.txt`
	local id
	local ip
	IFS=',' read -r id ip <<< "$instance"
	echo "SSH to $id at $ip"
	ssh $id
}

function attach_to_container
{
	# $1: the host of which container to ssh to
	local expconfig=`cat nodehostmap.txt`
	for nodeinfo in $expconfig; do
		local node
		local host
		IFS=',' read -r node host <<< "$nodeinfo"
		if [ "$node" == "$1" ]; then
			echo "Attaching to $node. ^p^q to detach. DON'T EXIT!"
			ssh $host -t "bash -ic 'docker attach spider$node'"
			break
		fi
	done
}

function copy_logs
{
	# $1: the directory name for resulting logs.
	execute_on_all delete_old_log
	local expconfig=`cat nodehostmap.txt`
	local pids=""
	for nodeinfo in $expconfig; do
		local node
		local host
		IFS=',' read -r node host <<< "$nodeinfo"
		(
			ssh $host -- mkdir -p container_logs
			ssh $host -- docker cp spider$node:/root/log/ ./container_logs/spider$node
		) &
		pids="$pids $!"
	done
	for pid in $pids ;
	do
		wait $pid
	done
	execute_on_all copy_log_files $1
}

case "$1" in
	help)
		cat <<- EOF
		Helper script to run Spider distributed tests

		Manage AWS EC2 Instances

		    start-instances n
		        Start n EC2 instances

		    stop-instances
		        Terminate EC2 instances

		Setup Experiment Environment

		    init-docker
		        Initialize docker swarm

		    uninit-docker
		        Destroy docker swarm

		Manage Experiment Files

		    init-image
		        Sync testbed, download binaries and package image

		    sync-testbed
		        Sync testbed directory to remotes

		    repackage-image
		        Repackage docker image

		    rebuild-binary bin1 bin2 ...
		        Recompile bin1, bin2, ... (lnd, expctrl, bitcoind, etcdjq)

		    download-binary
		        Download precompiled binaries

		    copy-logs DIR_NAME
		        Copies all log files from each spider container to DIR_NAME.
		        Creates DIR_NAME directory if it does not exist.

		Control Experiment

		    start-exp topofile expname exptime
		        Start an experiment

		    stop-exp
		        Stop an experiment

		Connect to Testbed

		    run-all cmd
		        Run command on all instances

		    ssh i
		        SSH to the i-th server (1-based index)

		    attach node
		        Attach to container node
		EOF
		;;
	start-instances)
		start_instances $2 ;;
	stop-instances)
		stop_instances ;;
	init-docker)
		init_swarm ;;
	uninit-docker)
		destroy_swarm ;;
	sync-testbed)
		execute_on_all sync_testbed ;;
	init-image)
		execute_on_all init_image ;;
	repackage-image)
		execute_on_all repackage ;;
	rebuild-binary)
		execute_on_all rebuild ${@:2} ;;
	mount-nvme)
		execute_on_all mount_nvme ;;
        copy-binaries)
                copy_binaries_single ;;
	download-binary)
		execute_on_all download ;;
	start-exp)
		TOPO_FILE=$2
		EXP_NAME=$3
		EXP_TIME=$4


 #               if [[ "$#" -e 9 ]]; then
 #                   ALPHA=$5
 #                   ETA=$6
 #                   KAPPA=$7
 #                   XI=$8
 #                   QUEUE_DRAIN_TIME=$9
 #               fi

		start_experiment ;;
	stop-exp)
		stop_experiment ;;
	run-all)
		run_on_all "${@:2}" ;;
	ssh)
		ssh_to_server $2 ;;
	attach)
		attach_to_container $2 ;;
	copy-logs)
		copy_logs $2 ;;
esac
