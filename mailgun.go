package main

import (
    "github.com/mailgun/mailgun-go"
    "fmt"
    "html/template"
    "bytes"
)

func ParseTemplate(templateFileName string, data interface{}) (string, error) {
  t, err := template.ParseFiles(templateFileName)

  if err != nil {
    return "", err
  }

  buf := new(bytes.Buffer)
  if err = t.Execute(buf, data); err != nil {
    return "", err
  }
  msgBody := buf.String()

  return msgBody, nil
}

func SendMessage() (string, error) {
  mg := mailgun.NewMailgun("<domain>", 
                           "<apiKey>", 
                           "<publicApiKey>",)
  templateData := struct {
      Name string
      URL  string
      Site string
    }{
      Name: "Name",
      URL:  "http://somesite.tld",
      Site: "Site Name",
    }

  msgBody, err := ParseTemplate("email-template.html", templateData)
  if err != nil {
    return "nil", err
  }

  m := mg.NewMessage(
    "Excited User <sender>",
    "Hello",
    "",
    "<receiver>",
  )
  m.SetHtml(msgBody)
  _, id, err := mg.Send(m)
  return id, err
}

func main(){
  id, err := SendMessage()
  fmt.Println(id)
  fmt.Println(err)
}