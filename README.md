# k8s-biocad
biocad test task

---

Тестовое задание на позицию Dev-Ops инженера в Biocad

Ответы на вопросы лежат в файле answers.txt

## Ход выполнения работы
1) Создадим простенький http-сервер на Go

![screenshot](https://raw.githubusercontent.com/cutlery47/k8s-biocad/master/media/prog.png)

2) Напишем докерфайл для сборки образа приложения

![screenshot](https://raw.githubusercontent.com/cutlery47/k8s-biocad/master/media/dockerfile.png)

3) Соберем образ и выложим в публичный Dockerhub репозиторий

![screenshot](https://raw.githubusercontent.com/cutlery47/k8s-biocad/master/media/dockerhub.png)

4) Создадим локальный k8s кластер при помощи minikube и сконфигурируем его.

![screenshot](https://raw.githubusercontent.com/cutlery47/k8s-biocad/master/media/deploy.png)

В качестве сервиса будем использовать встроенный Load Balancer.

5) Применим текущую конфигурацию

![screenshot](https://raw.githubusercontent.com/cutlery47/k8s-biocad/master/media/create.png)

6) Активируем port-forwarding на порту 8080

![screenshot](https://raw.githubusercontent.com/cutlery47/k8s-biocad/master/media/port-forward.png)

7) Проверим доступность деплоймента через браузер

![screenshot](https://raw.githubusercontent.com/cutlery47/k8s-biocad/master/media/hello.png)

Как видим, поды доступны.


