## Email verifcation based TOTP


Configure database on config/database.go
(I'm currently using https://railway.app/ for better experience)

```bash
host= user= password= dbname= port= sslmode=disable
```

#### How to run

```bash
go run main.go
``` 

Configure your Brevo api key on mailer.go.
Sending Email with brevo api reference https://developers.brevo.com/reference/sendtransacemail
for authentication,imma using TOTP https://github.com/pquerna/otp . 
OTP will expired until user hitted true otp.