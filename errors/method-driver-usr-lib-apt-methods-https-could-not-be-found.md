# The method driver /usr/lib/apt/methods/https could not be found

symptom:

    $ sudo apt-cache policy|grep https
     500 https://packages.millistream.com/apt/ precise/main i386 Packages
     500 https://packages.millistream.com/apt/ precise/main amd64 Packages

solution:

    sudo apt-get install apt-transport-https

http://askubuntu.com/questions/104160/method-driver-usr-lib-apt-methods-https-could-not-be-found-update-error

http://stackoverflow.com/questions/8647454/how-to-get-a-list-of-repositories-apt-get-is-checking
