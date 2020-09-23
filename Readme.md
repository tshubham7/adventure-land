### adventure land

# About

    GoLang-API for a game.
    In this game user run to the forest to collect the fruit, every time user gets the fruit 
    He/She earns 10 points, we need to tell the users their rank and list of participants rank wise.
    Simple!!

# Challenge

    We need to sort the users every time any user earns the points, in order to get live ranks


# Features

    * add or join user in the game
    * add points to the user
    * rank users point wise
    * get user's rank


# Api documention

    find the swagger documentation for restful api at dir. specs/swagger/api.yaml
    copy the content of this file, then open https://editor.swagger.io/# and paste the content. 

    find the postman collection for restful api at dir. specs/postman/collection.json
    download the file, open the your postman tool then import this json file into it.


# docker images

    download the docker images from the following urls
        $ docker pull tshubham7/go-adventure-land-api:latest
    run using command
        $ docker run -it -p 8080:8080 tshubham7/go-adventure-land-api
