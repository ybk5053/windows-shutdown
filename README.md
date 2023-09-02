# windows-shutdown
SMB command to shutdown windows remotely


    curl "http://windows-shutdown:3030/?ip=(ip_address)&user=(user)&pass=(pass)"

Enable broadcast forwarding for docker:

    sysctl -w net.ipv4.conf.all.bc_forwarding=1
    sysctl -w net.ipv4.conf.br-xxxx.bc_forwarding=1
