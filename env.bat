@echo off

doskey build=docker-compose build $*
doskey up=docker-compose up $*
doskey stop=docker-compose stop $*
doskey down=docker-compose down $*
doskey run=docker-compose run --rm develop $*

doskey develop=docker-compose up develop $*
doskey app=del /Q tmp\pids\server.pid$Tdocker-compose up app $*

doskey go=docker-compose run --rm develop go $*
