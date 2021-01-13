# nginx-sharing-session
This is nginx-sharing-session


* [How to Install NGINX](#how-to-install-nginx)
* [How to Create Configuration](#how-to-create-configuration)
## How to Install NGINX
Code:

    sudo apt update
    sudo apt install nginx
    cd /etc/nginx
    curl -X GET localhost

## How to Create Configuration
Code:

    cd /etc/nginx/sites-available

    sudo vi reverse.conf
    > server {
    >  localtion / {
    >    proxy_pass http://localhost:8080;
    >  }
    > }
    
    sudo ln -s /etc/nginx/sites-available/reverse.conf /etc/nginx/sites-enabled/reverse.conf

    sudo nginx -t

    sudo systemctl reload nginx.service