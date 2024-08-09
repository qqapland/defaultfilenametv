dsi() { docker stop $(docker ps -a | awk -v i="^$1.*" '{if($2~i){print$1}}'); }
dsi defaultfilenametv
docker rmi -f $(docker images -q defaultfilenametv) # remove the old image
docker load -i /home/a/defaultfilenametv.tar # this causes rename
docker run -it --rm -p 3004:3000 defaultfilenametv