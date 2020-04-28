# Flash Testbed

This folder contains source code for the following project:

- **Year**: 2019
- **Code**: 19CS004
- **Title**: Efficiently Routing Payments in Off-chain Network
- **Subject Area(s)**:
	- Algorithms Design & Analysis
	- Computer Networking
	- Distributed System
	- Optimization


## Repository

This folder is for fyp system submission only. For latest release and fixes, please visit ![github repository](https://github.com/isdanni/19cs004).

```shell
├── bin
├── build.sh: the starter scirpts;
├── cmds: main sctipts for handling the transcations;
├── credentials.csv: AWS EC2 Instance key;
├── params.md: a sample output for the transactions;
├── pkg: Golang packages for the transacition engines;
├── README.md
├── scripts: Shell scripts for handling the transcations and encryption points;
├── src: Golang package source code;
├── logs: folder contains transaction output;
├── tools: Python & Shell scripts used for cleaning and debugging;
└── topology: JSON files for setting up the nodes;
```

## Setting Up

This is local version of tetsing the Flash network. The set up refers to the spider docker testbed. To start the demo, you could either use local Ubuntu 19.04LTS system or the AWS EC2 Instance.

### Start Instance

1. Use 

```
Manage AWS EC2 Instances

    start-instances n
        Start n EC2 instances

    stop-instances
        Terminate EC2 instances

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
        Always use ^-p ^-q and not exit to detach from the container

```

## Typical Experiment Flow

1. SPider Testbed Workflow

```bash
./run.sh start-instances 3
# wait for about two minutes for EC2 to start and auto-install docker, go, etc.
./run.sh init-docker
./run.sh init-image
./run.sh start-exp ../topology/some-topo-file.json experiment-name 120
# after the experiment have finished
./run.sh stop-exp
./run.sh stop-instances
```

## Parsing Logs

After the experiment has completed, run:
```
./run.sh copy-logs LOG_DIR
python3 parse_logs.py -data_dir LOG_DIR -exp_name EXP_NAME	
```

LOG_DIR would be the directory supplied to the copy-logs command, with
directories like spider0e etc. This will parse all the log files there
and generate two plots:
	- EXP_NAME_per_channel_info.pdf
	- EXP_NAME_src_dest.pdf	

# Notes
^-p ^-q: detach from the container

# Topology File

See some examples in `topology/`. Here are some pitfalls

- Channel capacity is defined as the "one way" capacity - that is, how much satoshi are there at each end when channel is first established.
- Due to lnd and bitcoin limitations, the smallest channel capacity is 10000.
- The most efficient way to connect bitcoind is connecting each node to the miner node, forming a star topology. However, note that bitcoind only supports 8 outgoing connections, so be sure to set the miner node as the `dst`.

# Bugs

Currently, all messages are exchanged in etcd, which ultimatally falls on to a single leader node. This may become a problem when we run a large topology, or send transactions at a high-speed. We need to use golang RPC to replace etcd in those situations to avoid centralization.

