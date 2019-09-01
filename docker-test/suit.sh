#!/bin/bash

# Create directory structure
udp='udp_client'
tcp='tcp_client'
for i in {1..5}
do
    mkdir -p /home/$USER/Output/$udp$i
    mkdir -p /home/$USER/Output/$tcp$i
done

# Remove files before starts
(find ~/Output/ -type f -delete)

# Update/Create docker image before starts
(docker build --tag=perf-image:v1 .)

path='compose'
end='/docker-compose.yml'

# Multiple experiments : 30 + 1 Warmup
for j in {1..31}
do
    # Experiment Unit
    for i in {1..10}
    do
        current=$path$i$end
        (docker swarm init)
        (docker stack deploy -c $current pack-test)
        (sleep 18s) # needed for write all files
        (docker stack rm pack-test)
        (docker swarm leave --force)
        # (sleep 10s) #simulates a pause for each unitary experiment
        echo Teste $j$current done
    done
done