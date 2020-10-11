# env

alias build="docker-compose build"
alias up="docker-compose up"
alias stop="docker-compose stop"
alias down="docker-compose down"
alias run="docker-compose run --rm develop"

akias develop="up develop"
alias app="rm -f tmp/pids/server.pid && up app"

alias go="run go"
