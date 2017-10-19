# go-email-sample
A simple example of how to send email with MailGun and golang

To send:

1. Set values in `mailgun.go`
* `domain`
  * mailgun domain
* `piKey`
  * mailgun private API key
* `publicApiKey`
  * mailgun public API key
* `sender`
  * Sender email address
* `receiver`
  * Recipient email address

2. Get `mailgun-go` dependency
    ```
    go get github.com/mailgun/mailgun-go
    ```

3. Run
    ```
    go run mailgun.go
    ```