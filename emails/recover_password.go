package emails

import (
	"fmt"
	"gestfin-apis/models"
	"gestfin-apis/utils"
	"strings"
)

func SendEmailCodeRandom(user models.User) error {
	code := strings.Split(user.CodeRecovery, "")

	body := fmt.Sprint(`<span style="font-size: 25px; font-weight: 600; font-family: 'Poppins', sans-serif;">Recuperação de senha</span>
		 <div style="width: 100%; text-align: left; margin: 40px 0; line-height: 22px">
				 <span style="font-size: 20px; font-weight: 400; font-family: 'Poppins', sans-serif; text-align: left">Olá `, user.Name, `,</span><br/>
				 <span style="font-size: 20px; font-weight: 400; font-family: 'Poppins', sans-serif; text-align: left">
						 Abaixo está o código de identificação<br/> para recuperar senha:
				 </span>
		 </div>
			<div style="height: 60px; width: 100%; background-color: #F4F4F4; border: 1px solid #0B0B09;">
					 <table style="height: 100%; width: 100%">
							 <tr align="center" valign="middle">
									 <td><span style="font-size: 25px; font-weight: 600; font-family: 'Poppins', sans-serif;">`, code[0], `</span></td>
									 <td><span style="font-size: 25px; font-weight: 600; font-family: 'Poppins', sans-serif;">`, code[1], `</span></td>
									 <td><span style="font-size: 25px; font-weight: 600; font-family: 'Poppins', sans-serif;">`, code[2], `</span></td>
									 <td><span style="font-size: 25px; font-weight: 600; font-family: 'Poppins', sans-serif;">`, code[3], `</span></td>
									 <td><span style="font-size: 25px; font-weight: 600; font-family: 'Poppins', sans-serif;">`, code[4], `</span></td>
							 </tr>
					 </table>
			</div>
		 <div style="margin-top: 40px; width: 100%; text-align: center; font-size: 18px; font-weight: 700; font-family: 'Poppins', sans-serif;">
				 <span>Este código irá expirar após 24 horas.</span><br/>
				 <span>Nunca compartilhe esse código com outra pessoa.</span><br/>
				 <span>Nunca compartilhe sua senha com outra pessoa.</span>
		 </div>`)

	template := MountLayoutTemplateEmail(body)
	err := MountEmail(template, "[GESTFIN] - Código de verificação.", utils.CheckToSend(user.Email))

	return err
}

func SuccessfulRecoverPassword(user models.User) error {
	body := fmt.Sprint(` <span style="font-size: 25px; font-weight: 600; font-family: 'Poppins', sans-serif;">Recuperação de senha</span>
		 <div style="width: 100%; margin: 40px 0; line-height: 22px; text-align: center">
				 <span style="font-size: 20px; font-weight: 400; font-family: 'Poppins', sans-serif;">Olá `, user.Name, `,</span><br/>
				 <span style="font-size: 20px; font-weight: 400; font-family: 'Poppins', sans-serif;">
						 Sua senha foi redefinida com sucesso!
				 </span>
		 </div>
		 <div style="margin-top: 40px; width: 100%; text-align: center; font-size: 18px; font-weight: 700; font-family: 'Poppins', sans-serif;">
				 <span>Nunca compartilhe sua senha com outra pessoa.</span>
		 </div>
	`)

	template := MountLayoutTemplateEmail(body)
	err := MountEmail(template, "[GESTFIN] - Recuperação de senha.", utils.CheckToSend(user.Email))

	return err
}

func MountLayoutTemplateEmail(body string) string {
	return fmt.Sprint(`<!-- Inliner Build Version 4380b7741bb759d6cb997545f3add21ad48f010b -->
		<!DOCTYPE html PUBLIC "-//W3C//DTD XHTML 1.0 Strict//EN" "http://www.w3.org/TR/xhtml1/DTD/xhtml1-strict.dtd">
		<html lang="pt" xmlns="http://www.w3.org/1999/xhtml" xmlns="http://www.w3.org/1999/xhtml">
		<head>
				<title></title>
				<meta http-equiv="Content-Type" content="text/html; charset=utf-8"/>
				<meta name="viewport" content="width=device-width"/>
				<link rel="preconnect" href="https://fonts.googleapis.com">
				<link rel="preconnect" href="https://fonts.gstatic.com" crossorigin>
				<link href="https://fonts.googleapis.com/css2?family=Poppins:wght@400;600;700&display=swap" rel="stylesheet">
		</head>
		<body style="width: 100% !important;
								height: 100vh;
								-webkit-text-size-adjust: 100%;
								-webkit-font-smoothing: antialiased;
								-ms-text-size-adjust: 100%;
								color: #222222;
								font-family: 'Poppins', sans-serif;
								font-weight: normal; text-align: left;
								line-height: 19px;
								font-size: 14px;
								margin: 0;
								padding: 0;
								background-color: #F3F3F3;
		">
		<table style="width: 100%; height: 100%; background-color: transparent;background-image: url('https://gestfin-web.s3.sa-east-1.amazonaws.com/emails/backgroun-login.svg');
									 background-size: cover" border="0" cellspacing="0">
				<tr valign="middle" style="height: 130px; background-color: #0B0B09; box-shadow: -1px 9px 13px 5px rgba(0, 0, 0, 0.25);">
						<th colspan="3" align="center">
								<img src="https://gestfin-web.s3.sa-east-1.amazonaws.com/emails/logo.svg" alt="GESTFIN" style="display: inline-block" />
						</th>
				</tr>
				<tr>
					 <td colspan="3" align="center" valign="middle" style="padding: 30px 0; color: #0B0B09">
							 <div style="width: 490px; background-color: #fff; border-radius: 20px; padding: 40px">
									 `, body, `
							 </div>
					 </td>
				</tr>
				<tr style="height: 80px; background-color: #0B0B09" align="center" valign="middle">
						<td colspan="3">
								<span style="font-size: 15px; color: #fff">Copyright © 2023 Gestfin</span>
						</td>
				</tr>
		</table>
		</body>
		</html>`)
}
