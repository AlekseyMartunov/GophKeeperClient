### финальная работа на курсе по изучению Go

Что бы посмотреть проект необходимо в пустую папку склонировать клиент и сервер:
```
git clone https://github.com/AlekseyMartunov/GophKeeperClient.git
git clone https://github.com/AlekseyMartunov/GophKeeper.git
```

Далее необходимо собрать клиент и сервер в докер контейнеры :
```
cd GophKeeper
docker build -t keeper_server .
```

И ...
```
cd GophKeeperClient
docker build -t keeper_client .
```

Создадим сеть в докере для общения наших контейнеров при помощи DNS имен, в роли имен будут выступать имена самих контейнеров.
```
docker network create keeper_network
```

После чего можно запустить контейнеры которые мы только что сбилдили а еще несколько контейнеров с базами и с файловым хранилищем. 
(Фаил docker-compose.yaml находится внутри папки GophKeeperClient)
```
cd GophKeeperClient
docker compose up
```

После выполнения компанды в терминале появяться логи запущенных контейнеров. Они заблокирую терминал, так что придется открыть еще один.
Cписок контейнеров можно посмотреть командой:
```
docker ps -a
```

И наконец, для запуска контейнера перейдем "внутрь" контейнера gophKeeperClient, делается это при помощи вот этой команды:
```
docker exec -it gophKeeperClient bash
```

Для запуска приложения запустим бинарник app:
```
./app
```

приложение запущено!