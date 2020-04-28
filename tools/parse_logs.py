import argparse
import glob
import pdb
from collections import defaultdict
import numpy as np

import matplotlib
matplotlib.use('Agg')
import matplotlib.pyplot as plt
import os
from matplotlib.backends.backend_pdf import PdfPages
from cycler import cycler

DEBUG_LOCAL = True
def read_flags():
    parser = argparse.ArgumentParser()
    parser.add_argument("-data_dir", type=str, required=False,
            default="./test", help="")
    parser.add_argument("-exp_name", type=str, required=False,
            default="./woot", help="")
    parser.add_argument("-frac_success_time_bucket", type=int, required=False,
            default=5, help="")
    return parser.parse_args()

## here are the logging comments for per channel logs:
# log.Infof("Spider: info_type: periodicStats," +
        # "time: %v, i: %d, node: %s, peer: %s, chanID: %s," +
        # "qlen: %s, totalAmt: %s, sent: %s, rcvd: %s, locBal: %s," +
        # "remBal: %s, bandwidth: %s, capacity: %s")

# log.Infof("Spider: info_type: updateLocalPrice, time: %d, node: %s," +
# "peer: %s, xlocal: %d, nlocal: %d," +
# "ilocal: %d, queuelen: %d, aVal: %d, sVal: %d")

# log.Infof("LP Spider: info_type: updatePriceProbe,"+
# "node: %s, peer: %s, time: %d, ix: %v, iy: %v,"+
# "wx: %v, wy: %v, qx: %v, qy: %v, aDiffRemote: %v,"+
# "sDiffRemote: %v, aDiff: %v, sDiff: %v, mu_local: %v,"+
# "lambda: %v, n_local: %v, n_remote: %v, q_remote: %v")

# for per channel stats
PER_CHAN_KEYS = ["locBal", "bandwidth", "sent", "qlen",
        "ix", "iy", "wx", "wy", "qx", "qy", "aDiffRemote", 
        "sDiffRemote", "mu_local", "lambda"]

# e2e keys
PER_SRC_DST_KEYS = ["price", "rate", "pathID", "window", "inflight", "fractionMarked"]

# endhost info types:
ENDHOST_INFO_TYPES = ["path_prices", "payment_attempted", 
        "payment_success", "window_size"]

BTC_TO_SATOSHIS = 100000000

def is_number(s):
    try:
        float(s)
        return True
    except ValueError:
        return False

def make_router_stats_pdf(all_stats):
    plotting_data = {}
    for node_name, node_data  in all_stats.items():
        if "r" not in node_name:
            continue
        print("node name: ", node_name)
        for info_type, all_data in node_data.items():
            for other_node, chan_data in all_data.items():
                if "r" not in other_node:
                    continue
                for key, data in chan_data.items():
                    if key == "time":
                        continue
                    if key not in plotting_data:
                        plotting_data[key] = defaultdict(dict)

                    if "time" in chan_data:
                        xaxis = chan_data["time"]
                        plotting_data[key][node_name][other_node] = (data,
                                xaxis)
                    else:
                        plotting_data[key][node_name][other_node] = (data,
                                None)
    
    # plots everything
    with PdfPages(args.exp_name + "_per_channel_info.pdf") as pdf:
        for k, to_plot in plotting_data.items():
            plot_relevant_stats(to_plot, pdf, k,
                     compute_router_wealth=False)

def make_endhost_stats_pdf(all_stats):
    '''
    - frac completed
    - 
    '''
    total_attempt = 0.00
    total_succ = 0.00
    plotting_data = {}
    for node_name, node_data  in all_stats.items():
        if "e" not in node_name:
            continue
        for info_type, all_data in node_data.items():
            if info_type not in ENDHOST_INFO_TYPES:
                continue

            for other_node, chan_data in all_data.items():
                if "e" not in node_name:
                    continue
                # want format: [key][src][dest] = [LIST]
                # where LIST is: frac completed over per time period
                assert "time" in chan_data
                # pdb.set_trace()
                
                # keys we care about:
                #   payment_attempted
                #   payment_success
                #   path prices
                if info_type == "path_prices" or info_type == "window_size":
                    # we shouldn't need to loop here.
                    ## keys we care about: time, pathID, price.
                    #plotting_data[key][node_name+other_node][pathID]=[price]
                    if "time" in chan_data:
                        timestamps = chan_data["time"]
                        start_ts = timestamps[0]
                        xaxis = [t - start_ts for t in timestamps]
                    else:
                        xaxis = None
                    assert "pathID" in chan_data
                    path_ids = chan_data["pathID"]
                    src_dest = node_name + other_node
                    for key, data in chan_data.items():
                        if key == "time" or key == "pathID":
                            continue
                        if key not in plotting_data:
                            plotting_data[key] = defaultdict(dict)

                        # loop over all pathID <-> data pairs, and make lists
                        # for each unique pathIDs
                        unique_paths = set(path_ids)
                        for cur_path_id in unique_paths:
                            cur_path_data = []
                            cur_xaxis = []
                            for data_idx, data_item in enumerate(data):
                                if (path_ids[data_idx] == cur_path_id):
                                    cur_path_data.append(data_item)
                                    cur_xaxis.append(xaxis[data_idx])
                            plotting_data[key][src_dest][cur_path_id]=(cur_path_data, cur_xaxis)

                elif info_type == "payment_attempted" \
                        or info_type == "payment_success":
                    if info_type not in plotting_data:
                        plotting_data[info_type] = defaultdict(dict)
                    timestamps = chan_data["time"]
                    start_ts = timestamps[0]
                    # each idx of val will be an int representing how many
                    # transactions attemped / successful in that period
                    vals = []
                    vals.append(0.00)
                    for t in timestamps:
                        if t < start_ts + len(vals)*args.frac_success_time_bucket:
                            vals[-1] += 1.00
                        else:
                            # reset timer
                            vals.append(0.00)

                    plotting_data[info_type][node_name][other_node] = (vals,
                            None)
    
    # convert to payment_success / payment_attempted 
    pa = plotting_data["payment_attempted"]
    ps = plotting_data["payment_success"]
    for router, channel_info in pa.items():
        # if router not in new_to_plot:
            # new_to_plot[router] 

        for channel, chan_data in channel_info.items():
            attempted = chan_data[0]
            success = ps[router][channel][0]
            # print(len(attempted))
            # print(len(success))
            vals = []
            for si, succ in enumerate(success):
                if si >= len(attempted):
                    continue
                succ = max(succ, 1.00)
                cur_attempt = attempted[si]
                cur_attempt = max(cur_attempt, 1.00)
                vals.append(succ / cur_attempt)

            if "frac-success" not in plotting_data:
                plotting_data["frac-success"] = defaultdict(dict)
            plotting_data["frac-success"][router][channel] = (vals, None)

    # plots everything
    with PdfPages(args.exp_name + "_src_dest.pdf") as pdf:
        for k, to_plot in plotting_data.items():
            plot_relevant_stats(to_plot, pdf, k,
                     compute_router_wealth=False)

def plot_relevant_stats(data, pdf, signal_type, x_axis=None, compute_router_wealth=False):
    color_opts = ['#fa9e9e', '#a4e0f9', '#57a882', '#ad62aa']
    router_wealth_info =[]

    for router, channel_info in data.items():
        # if "e" in router:
            # continue
        channel_bal_timeseries = []
        plt.figure()
        plt.rc('axes', prop_cycle = (cycler('color', ['r', 'g', 'b', 'y', 'c', 'm', 'y', 'k']) +
            cycler('linestyle', ['-', '--', ':', '-.', '-', '--', ':', '-.'])))

        i = 0
        for channel, chan_data in channel_info.items():
            # FIXME: might need to add this.
            # if "e" in channel:
                # continue

            values = chan_data[0]
            time = chan_data[1]
            if time is None:
                time = range(len(values))
            label_name = str(router) + "->" + str(channel)
            plt.plot(time, values, label=label_name + "(" + str(np.average(values[int(len(values)/4):])) + ")")
            print(signal_type, min(values), router, channel)
            if compute_router_wealth:
                channel_bal_timeseries.append((time, values))
            i += 1

        if compute_router_wealth:
            router_wealth = []
            for i, time in enumerate(channel_bal_timeseries[0][0]):
                wealth = 0
                for j in channel_bal_timeseries:
                    # PN: can be different lengths
                    if (len(j[1]) <= i):
                        continue
                    wealth += j[1][i]
                router_wealth.append(wealth)
            router_wealth_info.append((router,channel_bal_timeseries[0][0], router_wealth))

        plt.title(signal_type + " for Router " + str(router))
        plt.xlabel("Time")
        plt.ylabel(signal_type)
        plt.legend()
        pdf.savefig()  # saves the current figure into a pdf page
        plt.close()

    if compute_router_wealth:
        for (r, time, wealth) in router_wealth_info:
            plt.plot(time, wealth, label=str(r))
        plt.title("Router Wealth Timeseries")
        plt.xlabel("Time")
        plt.ylabel("Router Wealth")
        plt.legend()
        pdf.savefig()  # saves the current figure into a pdf page
        plt.close()

def get_spider_info_pairs(line):
    line = line.replace("\n", "")
    # remove everything before Spider:
    line = line[line.find("Spider:")+7:]
    # split them into key:val pairs
    line = line.split(",")
    info = {}
    for pair in line:
        pair = pair.split(":")
        key = pair[0].replace(" ", "")
        val = pair[1]
        info[key] = val
    return info

def parse_log_file(fn):
    '''
    @ret:
        - {} per dest stats
        - {} per channel stats
    '''
    stats = {}
    # public-key to node-name
    key_to_node = {}

    with open(fn, "r") as f:
        lines = f.readlines()

    print("num lines: ", len(lines))
    # for src<->dst pairs
    cur_attempted = 0
    cur_success = 0

    for line in lines:
        # 1: find map between public keys to nodenames
        if "public-key" in line:
            info = get_spider_info_pairs(line)
            key_to_node[info["public-key"]] = info["nodeName"]

        if "info_type" not in line:
            continue

        info = get_spider_info_pairs(line)
        info_type = info["info_type"].replace(" ", "")
        if "peer" in info or "dest" in info:
            # collect these stats about all of them
            if "peer" in info:
                other_node = info["peer"]
            else:
                other_node = info["dest"]

            if info_type not in stats:
                stats[info_type] = {}
            data = stats[info_type]

            if other_node not in data:
                data[other_node] = defaultdict(list)
            
            all_keys = PER_CHAN_KEYS + PER_SRC_DST_KEYS
            for k in all_keys:
                if k not in info:
                    continue

                v = info[k]
                if is_number(v):
                    v = float(v)
                elif "mSAT" in v:
                    v = float(v.replace("mSAT", ""))
		    # to get satoshis
                    v /= 1000
                elif "BTC" in v:
                    v = float(v.replace("BTC", ""))
                    v *= BTC_TO_SATOSHIS
                else:
                    continue

                data[other_node][k].append(v)

            if "time" in info:
                v = info["time"]
                data[other_node]["time"].append(float(v))
            stats[info_type] = data
 
    return stats, key_to_node

args = read_flags()
all_key_to_node = {}

# key: node, val: dict
#   key: other-node, val: dict
#       key: payment-type, val: list
all_stats = {}
all_dest_stats = {}

# for fn in glob.iglob("./*.log", recursive=False):
for fn in glob.iglob(args.data_dir + "/**/lnd.log", recursive=True):
    # FIXME: better method
    start = fn.rfind("spider")+len("spider")
    node_name = fn[start:start+2]
    stats, cur_pub_key = parse_log_file(fn)
    all_stats[node_name] = stats
    # all_dest_stats[node_name] = dest_stats
    all_key_to_node.update(cur_pub_key)

print("num all channels: ", len(all_stats.keys()))

# TODO: change all the public-keys to the appropriate nodes
for node_name, node_data  in all_stats.items():
    for info_type, all_data in node_data.items():
        new_all_data = {}
        for other_node, chan_data in all_data.items():
            assert other_node in all_key_to_node
            new_all_data[all_key_to_node[other_node]] = chan_data
        node_data[info_type] = new_all_data
        
            # add plotting scripts for per channel data here.
            # pdb.set_trace()

make_router_stats_pdf(all_stats)

# TODO: first correct for dividing up the data into time slots
make_endhost_stats_pdf(all_stats)
