# todo note commands
New-NetFirewallRule -DisplayName "Docker Swarm Overlay UDP" -Direction Inbound -LocalPort 4789 -Protocol UDP -Action Allow

New-NetFirewallRule -DisplayName "Docker Swarm UDP" -Direction Inbound -LocalPort 7946 -Protocol UDP -Action Allow

New-NetFirewallRule -DisplayName "Docker Swarm TCP" -Direction Inbound -LocalPort 7946 -Protocol TCP -Action Allow

New-NetFirewallRule -DisplayName "Port 2377 Outbound" -Direction Outbound -Protocol TCP -LocalPort 2377 -Action Allow

New-NetFirewallRule -DisplayName "Port 7946 TCP Outbound" -Direction Outbound -Protocol TCP -LocalPort 7946 -Action Allow

New-NetFirewallRule -DisplayName "Port 7946 UDP Outbound" -Direction Outbound -Protocol UDP -LocalPort 7946 -Action Allow

New-NetFirewallRule -DisplayName "Port 4789 Outbound" -Direction Outbound -Protocol UDP -LocalPort 4789 -Action Allow

docker swarm init --advertise-addr=172.20.10.4 --listen-addr=172.20.10.4:2377

docker swarm join --token <insert_join_token> 172.20.10.4:2377

docker stack deploy -c docker-compose.yml ourstack