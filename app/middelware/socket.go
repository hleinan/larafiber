package middleware

/*func Socket() fiber.Handler {
	return func(c *fiber.Ctx) {
		sessions := store.GetSessions()
		store := sessions.Get(c)

		uuidGen := uuid.Must(uuid.NewRandom()).String()
		store.Set("socket_id", strings.ReplaceAll(uuidGen, "-", ""))

		defer store.Save()

		c.Next()
		return
	}
}*/
