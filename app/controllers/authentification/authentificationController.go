package authentification

import (
	"encoding/json"
	"fmt"
	"github.com/drexedam/gravatar"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	models "github.com/hleinan/larafiber/app"
	"github.com/hleinan/larafiber/app/services/email"
	"github.com/hleinan/larafiber/app/services/inertia"
	"github.com/hleinan/larafiber/app/services/store"
	"golang.org/x/crypto/bcrypt"
	"log"
	"os"
	"strings"
	"time"
)

type RegisterObject struct {
	Name           string `json:"name"`
	Email          string `json:"email"`
	Password       string `json:"password"`
	PasswordRepeat string `json:"password_repeat"`
	Agreed         bool   `json:"agreed"`
	Subscribe      bool   `json:"subscribe"`
	Marketing      bool   `json:"marketing"`
}

type LoginObject struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func ShowLogin(c *fiber.Ctx) {
	dict := make(map[string]interface{})
	inertia.Render("Auth/Login", c, dict)
}

func Login(c *fiber.Ctx) {
	var loginObject LoginObject
	err := json.Unmarshal([]byte(c.Body()), &loginObject)
	if err != nil {
		log.Println(err)
	}

	db := models.GetDB()
	account := models.Account{}
	db.Preload("User").Where(&models.Account{Email: strings.ToLower(loginObject.Email)}).First(&account)

	if comparePasswords(account.Password, []byte(loginObject.Password)) {
		generateSession(account.User, c)

		// Simple prevent hacking (Can only run this once a second)
		time.Sleep(1 * time.Second)

		next := c.Query("next")
		if next != "" {
			c.Redirect(next)
		} else {
			c.Redirect("/")
		}
	} else {
		dict := make(map[string]interface{})
		dict["login_error"] = "Email address is wrong. Try again"
		inertia.Render("Auth/Login", c, dict)
	}
}

func Logout(c *fiber.Ctx) {
	destroySession(c)
	c.Redirect("/")
}

func ShowRegister(c *fiber.Ctx) {
	dict := make(map[string]interface{})
	inertia.Render("Auth/Register", c, dict)
}

func Register(c *fiber.Ctx) {
	db := models.GetDB()
	var registerObject RegisterObject
	err := json.Unmarshal([]byte(c.Body()), &registerObject)
	if err != nil {
		log.Println(err)
	}

	if registerObject.Password != registerObject.PasswordRepeat {
		dict := make(map[string]interface{})
		dict["register_error"] = "Passwords does not match"
		inertia.Render("Auth/Register", c, dict)
		fmt.Println("Passwords does not match")
	} else {
		var accountExist models.Account
		if result := db.
			Where("email = ?", registerObject.Email).
			First(&accountExist); result.Error != nil {
			user := models.User{Name: registerObject.Name,
				Marketing: registerObject.Marketing,
				Avatar:    gravatar.New(registerObject.Email).Default(gravatar.NotFound).AvatarURL(),
			}

			db.Create(&user)

			hash, err := GeneratePassword(registerObject.Password)
			if err != nil {
				log.Println(err)
			}

			account := models.Account{Email: strings.ToLower(registerObject.Email), Password: hash, User: user}
			db.Create(&account)

			GenerateEmailTokenAndEmail(user.ID)

			generateSession(user, c)

			c.Redirect("/user/me")
		} else {
			dict := make(map[string]interface{})
			dict["register_error"] = "Epost-adressen eksisterer allerede"
			inertia.Render("Auth/Register", c, dict)
		}
	}
}

func Forgot(c *fiber.Ctx) {
	dict := make(map[string]interface{})
	inertia.Render("Auth/Forgot", c, dict)
}

type EmailObject struct {
	Email string `json:"email"`
}

func ForgotStore(c *fiber.Ctx) {
	var emailObject EmailObject
	err := json.Unmarshal([]byte(c.Body()), &emailObject)
	if err != nil {
		log.Println(err)
	}

	db := models.GetDB()
	account := models.Account{}
	db.Where(&models.Account{Email: strings.ToLower(emailObject.Email)}).First(&account)
	CreateForgottenObjectAndEmail(account)
	dict := make(map[string]interface{})
	inertia.Render("Auth/Sent", c, dict)
}

func generateSession(user models.User, c *fiber.Ctx) {
	SetUser(c, user)
}

func destroySession(c *fiber.Ctx) {
	DestroyUser(c)
}

func comparePasswords(hashedPwd []byte, plainPwd []byte) bool {
	byteHash := hashedPwd
	err := bcrypt.CompareHashAndPassword(byteHash, plainPwd)
	if err != nil {
		log.Println(err)
		return false
	}
	return true
}

func SetUser(c *fiber.Ctx, u models.User) {
	sessions := store.GetSessions()
	store := sessions.Get(c) // get/create new session
	store.Set("id", u.ID)
	store.Set("logged_in", true)
	store.Set("error", "")
	defer store.Save()
}

func DestroyUser(c *fiber.Ctx) {
	sessions := store.GetSessions()
	store := sessions.Get(c) // get/create new session
	store.Set("logged_in", false)
	store.Set("id", "")
	store.Set("error", "")
	store.Destroy()
	defer store.Save()
}

func GeneratePassword(s string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(s), bcrypt.MinCost)
}

func ResetPassword(c *fiber.Ctx) {
	dict := make(map[string]interface{})
	dict["uuid"] = c.Params("uuid")
	inertia.Render("Auth/Reset", c, dict)
}

type ResetPasswordObject struct {
	Password       string `json:"password"`
	PasswordRepeat string `json:"password_repeat"`
}

func ResetPasswordStore(c *fiber.Ctx) {
	var resetPasswordObject ResetPasswordObject
	err := json.Unmarshal([]byte(c.Body()), &resetPasswordObject)
	if err != nil {
		log.Println(err)
	}

	if strings.EqualFold(resetPasswordObject.Password, resetPasswordObject.PasswordRepeat) {
		// Passwords are the same, continue
		db := models.GetDB()
		forgottenPassword := new(models.ForgottenPassword)
		if result := db.
			Where("uuid = ?", c.Params("uuid")).
			First(&forgottenPassword); result.Error != nil {
			c.Redirect("/forgot") // how do we send with an error?
		} else {
			// Does exist in db. All good
			// Now we need to check the timestamp of the object
			account := new(models.Account)
			db.
				Where(forgottenPassword.AccountID).
				First(&account)

			account.Password, err = GeneratePassword(resetPasswordObject.Password)
			if err == nil {
				db.Save(account)
				db.Delete(forgottenPassword)
				c.Redirect("/login")
			} else {
				c.Redirect("/forgot")
			}
		}

	} else {
		dict := make(map[string]interface{})
		dict["login_error"] = "Epost-adressen eller passordet er feil. Prøv igjen"
		inertia.Render("Auth/Reset", c, dict)
	}
}

func CreateForgottenObjectAndEmail(account models.Account) {
	db := models.GetDB()
	var forgottenPasswordFromDb models.ForgottenPassword
	if result := db.
		Where("account_id = ?", account.ID).
		First(&forgottenPasswordFromDb); result.Error == nil {
		db.Delete(forgottenPasswordFromDb)
	}

	forgottenPassword := new(models.ForgottenPassword)
	forgottenPassword.AccountID = account.ID
	db.Save(forgottenPassword)

	sendResetPasswordEmail(account, forgottenPassword)
}

func VerifyEmailExternal(c *fiber.Ctx) {
	db := models.GetDB()
	var account models.Account
	if result := db.
		Where("token = ?", c.Params("uuid")).
		First(&account); result.Error == nil {
		var user models.User
		db.
			Where(account.UserID).
			First(&user)

		user.Verified = !user.Verified
		db.Save(user)
		c.Redirect("/verified")
	} else {
		// @TODO: Redirect to new page
		c.Redirect("/")
	}
}

func sendResetPasswordEmail(account models.Account, forgottenPassword *models.ForgottenPassword) {
	title := "Tilbakestilling av passord på aning.no"
	to := account.Email
	text := "<h1>Heisann</h1><p>Du har bedt om å tilbakestille passordet på aning.no</p><p>Følg denne lenken for videre informasjon</p>" + os.Getenv("domain") + "resetpassword/" + forgottenPassword.UUID + "<br><br><p>Vennlig hilsen oss i Aning</p>"

	fmt.Println(text)
	go email.SendEmail(to, title, text)
}

func GenerateEmailTokenAndEmail(id int64) {
	uuidGen := uuid.Must(uuid.NewRandom()).String()
	token := strings.ReplaceAll(uuidGen, "-", "")

	db := models.GetDB()
	var account models.Account
	db.
		Where("user_id = ?", id).
		First(&account)

	account.Token = &token
	db.Save(&account)
	if os.Getenv("app_env") == "prod" || os.Getenv("app_env") == "test" {
		sendEmailToken(account.Email, token)
	}
}

func sendEmailToken(toEmail string, token string) {
	title := "Verifiser epostadressen din på aning.no"
	to := toEmail
	text := "<h1>Heisann</h1><p>Du må verifisere epost-adressen din.</p><p>Følg denne lenken for videre informasjon</p>" + os.Getenv("domain") + "verifyemail/" + token + "<br><br><p>Vennlig hilsen oss i Aning</p>"

	go email.SendEmail(to, title, text)
}
