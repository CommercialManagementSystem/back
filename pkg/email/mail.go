package email

import (
	"bytes"
	"crypto/tls"
	"encoding/base64"
	"fmt"
	"log"
	"net/smtp"
	"strings"
	"time"
)

var ContentTypeMap = map[string]string{
	"zip":     "application/x-zip-compressed",
	"any":     "application/octet-stream",
	"default": "text/plain;charset=utf-8",
}

type SendMail struct {
	User       string
	Password   string
	Host       string
	Port       string
	SSL        bool
	AuthClient smtp.Auth
}

type Attachment struct {
	Name        string
	ContentType string
	WithFile    bool
	Data        []byte
}

type Message struct {
	From        string
	To          []string
	Cc          []string
	Bcc         []string
	Subject     string
	Body        string
	ContentType string
	File        Attachment
}

func (mail *SendMail) Auth() {
	mail.AuthClient = smtp.PlainAuth("", mail.User, mail.Password, mail.Host)
}

func (mail *SendMail) Send(message *Message) error {
	mail.Auth()
	buffer, err := mail.getMailInfo(message)
	if err != nil {
		return err
	}

	if !mail.SSL {
		return smtp.SendMail(mail.Host+":"+mail.Port, mail.AuthClient, message.From, message.To, buffer.Bytes())
	}
	return mail.SendMailSSL(message, buffer)
}

func (mail *SendMail) getMailInfo(message *Message) (*bytes.Buffer, error) {
	buffer := bytes.NewBuffer(nil)
	boundary := "GoBoundary"
	Header := make(map[string]string)
	Header["From"] = message.From
	Header["To"] = strings.Join(message.To, ";")
	Header["Cc"] = strings.Join(message.Cc, ";")
	Header["Bcc"] = strings.Join(message.Bcc, ";")
	Header["Subject"] = message.Subject
	Header["Content-Type"] = "multipart/mixed;boundary=" + boundary
	Header["Mime-Version"] = "1.0"
	Header["Date"] = time.Now().String()
	mail.writeHeader(buffer, &Header)

	body := "\r\n--" + boundary + "\r\n"
	body += "Content-Type:" + message.ContentType + "\r\n"
	body += "\r\n" + message.Body + "\r\n"
	buffer.WriteString(body)

	if message.File.WithFile {
		attachment := "\r\n--" + boundary + "\r\n"
		attachment += "Content-Transfer-Encoding:base64\r\n"
		attachment += "Content-Disposition:attachment\r\n"
		attachment += "Content-Type:" + message.File.ContentType + ";name=\"" + message.File.Name + "\"\r\n"
		buffer.WriteString(attachment)
		defer func() {
			if err := recover(); err != nil {
				log.Fatalln(err)
			}
		}()
		mail.writeFile(buffer, &message.File.Data)
	}

	buffer.WriteString("\r\n--" + boundary + "--")

	return buffer, nil
}

func (mail *SendMail) SendMailSSL(message *Message, buffer *bytes.Buffer) error {
	tlsConfig := &tls.Config{
		InsecureSkipVerify: true,
		ServerName:         mail.Host,
	}
	conn, err := tls.Dial("tcp", mail.Host+":"+mail.Port, tlsConfig)
	if err != nil {
		return fmt.Errorf("DialConn:%v", err)
	}
	client, err := smtp.NewClient(conn, mail.Host)
	if err != nil {
		return fmt.Errorf("Client:generateClient:%v", err)
	}
	defer func(client *smtp.Client) {
		err := client.Close()
		if err != nil {
			fmt.Println(err.Error())
			return
		}
	}(client)
	if mail.AuthClient != nil {
		if ok, _ := client.Extension("AUTH"); ok {
			if err = client.Auth(mail.AuthClient); err != nil {
				return fmt.Errorf("Client:clientAuth:%v", err)
			}
		}
	}
	if err = client.Mail(mail.User); err != nil {
		return fmt.Errorf("Client:clientMail:%v", err)
	}

	for _, addr := range message.To {
		if err = client.Rcpt(addr); err != nil {
			return fmt.Errorf("Client:Rcpt:%v", err)
		}
	}
	w, err := client.Data()
	if err != nil {
		return fmt.Errorf(" Client:%v", err)
	}
	_, err = buffer.WriteTo(w)
	if err != nil {
		return fmt.Errorf("Client:WriterBody:%v", err)
	}
	err = w.Close()
	if err != nil {
		return fmt.Errorf("Client:CloseBody:%v", err)
	}
	return nil
}

func (mail *SendMail) writeHeader(buffer *bytes.Buffer, Header *map[string]string) string {
	header := ""
	for key, value := range *Header {
		header += key + ":" + value + "\r\n"
	}
	header += "\r\n"
	buffer.WriteString(header)
	return header
}

//writeFile read and write the file To buffer
func (mail *SendMail) writeFile(buffer *bytes.Buffer, file *[]byte) {
	payload := make([]byte, base64.StdEncoding.EncodedLen(len(*file)))
	base64.StdEncoding.Encode(payload, *file)
	buffer.WriteString("\r\n")
	for index, line := 0, len(payload); index < line; index++ {
		buffer.WriteByte(payload[index])
		if (index+1)%76 == 0 {
			buffer.WriteString("\r\n")
		}
	}
}
