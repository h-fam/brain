sudo docker run -it --name samba -p 139:139 -p 445:445 \
            -v /mnt/vol1/space:/space \
            -d hfam/samba:v1.0.0 -p \
            -s "space;/space;yes;no;yes;;;;Space Share" \
            -p
            
            