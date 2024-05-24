from email.mime.base import MIMEBase
from email.mime.multipart import MIMEMultipart
from locale import locale_alias
import smtplib
import mimetypes
from email import encoders

my_email = "email"
receiveEmail = "email"
my_password = "****"

message = MIMEMultipart()
message['From'] = my_email
message['To'] = receiveEmail
message['Subject'] = 'Here is your daily Ars Technica news'

arsTechFile = 'file.csv'

attachment = open(arsTechFile, 'rb')

obj = MIMEBase('application', 'octet-stream')
obj.set_payload((attachment).read())
encoders.encode_base64(obj)
obj.add_header('Content-Disposition',"attachment; filename= "+arsTechFile)

message.attach(obj)

my_message = message.as_string()

with smtplib.SMTP("smtp.gmail.com", 587) as connection:
    connection.starttls()
    connection.login(my_email, my_password)
    connection.sendmail(
        from_addr=my_email,
        to_addrs=receiveEmail,
        msg=my_message
    )
