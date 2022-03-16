# Email CLI 

flexible and extendible script to send email

## Installation 
1. Make sure go is installed
2. Run `go mod download` to install dependencies
3. Run `make build` to compile source code
4. Run `sudo make install` to install the binary

## Usage
1) Login
```sh 
  email-cli login
```
proceed to enter email and password

2) Send
```sh 
  email-cli send <email@address.com> <message>
```

3) If you wish to send emails to a list of people using a template, refer to [this](./bulksend.sh) bash script

> Make sure that you have enabled less secure devices in your email service provider
> use this [link](https://support.google.com/accounts/answer/6010255?hl=en)
