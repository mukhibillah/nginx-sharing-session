# nginx-sharing-session
This is nginx-sharing-session

## how to install nginx
Code:

    sudo apt update
    sudo apt install nginx
    cd /etc/nginx
    curl -X GET localhost

## how to create configuration
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