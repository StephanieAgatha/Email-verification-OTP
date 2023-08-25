package service

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func SendEmailRegister(emailTo, firstName string) {
	url := "https://api.brevo.com/v3/smtp/email"

	payloadData := struct {
		Sender struct {
			Name  string `json:"name"`
			Email string `json:"email"`
		} `json:"sender"`
		To []struct {
			Email string `json:"email"`
		} `json:"to"`
		Subject     string `json:"subject"`
		HtmlContent string `json:"htmlContent"`
	}{
		Sender: struct {
			Name  string `json:"name"`
			Email string `json:"email"`
		}{
			Name:  "Stephanie",
			Email: "Stephanie@stephanieproject.my.id",
		},
		To: []struct {
			Email string `json:"email"`
		}{
			{
				Email: emailTo,
			},
		},
		Subject:     "Welcome on aboard",
		HtmlContent: fmt.Sprintf(`<!DOCTYPE html> <html lang="en"> <head> <meta charset="UTF-8"> <meta name="viewport" content="width=device-width, initial-scale=1.0"> <title>Welcome to Our Website</title> </head> <body style="font-family: Arial, sans-serif;"> <table style="width: 100%; max-width: 600px; margin: 0 auto; padding: 20px; border-collapse: collapse; background-color: #f5f5f5;"> <tr> <td style="background-color: #ffffff; padding: 20px; text-align: center;"> <img src="https://i.pinimg.com/originals/15/07/6f/15076f568f60aa63eb56ac8aefccd52f.jpg" alt="Company Logo" style="max-width: 150px;"> </td> </tr> <tr> <td style="background-color: #ffffff; padding: 20px;"> <h2 style="margin-top: 0;">Hi ` + firstName + `</h2> <p>Thank you for signing up on our website! We're thrilled to have you as a part of our community.</p> <p>If you have any questions or need assistance, feel free to reach out to our support team.</p> <p>Best regards,</p> <p>Stephanie Project Team</p> </td> </tr> <tr> <td style="background-color: #ffffff; padding: 20px; text-align: center;"> <p style="margin-bottom: 0;">© 2023 [Stephanie Project]. All rights reserved.</p> </td> </tr> </table> </body> </html>`),
	}

	//marshal > jadiin byte > newreader bytes > simpan dalam variable > masukin ke http newrequest

	payloadBytes, err := json.Marshal(payloadData)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	payloadReader := bytes.NewReader(payloadBytes)

	//newrequest
	req, err := http.NewRequest("POST", url, payloadReader)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	req.Header.Add("accept", "application/json")
	req.Header.Add("content-type", "application/json")
	req.Header.Add("api-key", "") //submit your api key

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println(string(body))
}

func SendEmailWithOTP(emailTo string, otp string) {
	url := "https://api.brevo.com/v3/smtp/email"

	payloadData := struct {
		Sender struct {
			Name  string `json:"name"`
			Email string `json:"email"`
		} `json:"sender"`
		To []struct {
			Email string `json:"email"`
		} `json:"to"`
		Subject     string `json:"subject"`
		HtmlContent string `json:"htmlContent"`
	}{
		Sender: struct {
			Name  string `json:"name"`
			Email string `json:"email"`
		}{
			Name:  "Stephanie Project",
			Email: "Stephanie@stephanieproject.my.id",
		},
		To: []struct {
			Email string `json:"email"`
		}{
			{
				Email: emailTo,
			},
		},
		Subject: "Log in to Stephanie Project",
		HtmlContent: fmt.Sprintf(`
<!DOCTYPE html PUBLIC "-//W3C//DTD XHTML 1.0 Strict//EN" "http://www.w3.org/TR/xhtml1/DTD/xhtml1-strict.dtd">
<html xmlns="http://www.w3.org/1999/xhtml">

<head>
  <meta http-equiv="Content-Type" content="text/html; charset=utf-8">
  <meta name="viewport" content="width=device-width, initial-scale=1.0">
  <title>Verify your login</title>
  <!--[if mso]><style type="text/css">body, table, td, a { font-family: Arial, Helvetica, sans-serif !important; }</style><![endif]-->
</head>

<body style="font-family: Helvetica, Arial, sans-serif; margin: 0px; padding: 0px; background-color: #ffffff;">
  <table role="presentation"
    style="width: 100%; border-collapse: collapse; border: 0px; border-spacing: 0px; font-family: Arial, Helvetica, sans-serif; background-color: rgb(239, 239, 239);">
    <tbody>
      <tr>
        <td align="center" style="padding: 1rem 2rem; vertical-align: top; width: 100%;">
          <table role="presentation" style="max-width: 600px; border-collapse: collapse; border: 0px; border-spacing: 0px; text-align: left;">
            <tbody>
              <tr>
                <td style="padding: 40px 0px 0px;">
                  <div style="text-align: left;">
                    <div style="padding-bottom: 20px;"><img src="https://i.ibb.co/Qbnj4mz/logo.png" alt="Company" style="width: 56px;"></div>
                  </div>
                  <div style="padding: 20px; background-color: rgb(255, 255, 255);">
                    <div style="color: rgb(0, 0, 0); text-align: left;">
                      <h1 style="margin: 1rem 0">Verification code</h1>
                      <p style="padding-bottom: 16px">Please use the verification code below to sign in.</p>
                      <p style="padding-bottom: 16px"><strong style="font-size: 130%">`+otp+`</strong></p>
                      <p style="padding-bottom: 16px">If you didn’t request this, you can ignore this email.</p>
                      <p style="padding-bottom: 16px">This code is valid for 30 minutes.</p>
                      <p style="padding-bottom: 16px">Thanks,<br>Stephanie Project team</p>
                    </div>
                  </div>
                  <div style="padding-top: 20px; color: rgb(153, 153, 153); text-align: center;">
                    <p style="padding-bottom: 16px">Made with ♥ in Paris</p>
                  </div>
                </td>
              </tr>
            </tbody>
          </table>
        </td>
      </tr>
    </tbody>
  </table>
</body>

</html>
`, otp),
	}

	payloadBytes, err := json.Marshal(payloadData)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	payloadReader := bytes.NewReader(payloadBytes)

	req, err := http.NewRequest("POST", url, payloadReader)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	req.Header.Add("accept", "application/json")
	req.Header.Add("content-type", "application/json")
	req.Header.Add("api-key", "")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println(string(body))
}
