defaultfilenametv.adi.fr.eu.org

originally by everest pipkin (default-filename-tv.neocities.org)

run:
docker build -t defaultfilenametv -f Dockerfile . && docker run --init -e YT_API_KEY=$YT_API_KEY -p 3004:3000 defaultfilenametv