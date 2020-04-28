#!/bin/bash

# set up default environ
export TOPO_FILE=${TOPO_FILE-/root/topology/default_topo.json}  # default topo file
export PAYMENT_SIZE=${PAYMENT_SIZE-200}                         # default payment size
export SPIDER_LOG_FIREBASE=${SPIDER_LOG_FIREBASE-0}             # enable firebase log by default
export SPIDER_QUEUE=${SPIDER_QUEUE-1}                           # enable queue by default
export EXP_TIME=${EXP_TIME-120}                                 # default duration is 120s
export STATS_INTERVAL=${STATS_INTERVAL-1000}                    # default duration is 1000ms
export SPIDER_LP_ROUTING=${SPIDER_LP_ROUTING-1}									# disable LP routing by default
export SPIDER_QUEUE_UPDATE_TIME=${SPIDER_QUEUE_UPDATE_TIME-100}	# default queue busy waiting time
export SPIDER_USE_WINDOWS=${SPIDER_USE_WINDOWS-1}	
export SPIDER_TIMEOUT=${SPIDER_TIMEOUT-1}	
export ROUTINGALGO=${ROUTINGALGO-lp}

export ALPHA=${ALPHA-0.2}
export ETA=${ETA-0.025}
export KAPPA=${KAPPA-0.025}
export T_UPDATE=${T_UPDATE-1.5}
export QUEUE_DRAIN_TIME=${QUEUE_DRAIN_TIME-5}
export XI=${XI-1}
export SERVICE_ARRIVAL_WINDOW=${SERVICE_ARRIVAL_WINDOW-300}

mkdir -p /root/log
/root/scripts/runexp.sh &> /root/log/main.log &
bash
